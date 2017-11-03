/*
Package csi is CSI driver interface for OSD
Copyright 2017 Portworx

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package csi

import (
	"fmt"
	"net"
	"sync"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"go.pedge.io/dlog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// OsdCsiServerConfig provides the configuration to the
// the gRPC CSI server created by NewOsdCsiServer()
type OsdCsiServerConfig struct {
	Net     string
	Address string
}

// OsdCsiServer is a OSD CSI compliant server which
// proxies CSI requests for a single specific driver
type OsdCsiServer struct {
	listener net.Listener
	server   *grpc.Server
	wg       sync.WaitGroup
	running  bool
	lock     sync.Mutex
}

// NewOsdCsiServer creates a gRPC CSI complient server on the
// specified port and transport.
func NewOsdCsiServer(config *OsdCsiServerConfig) (*OsdCsiServer, error) {
	if nil == config {
		return nil, fmt.Errorf("Configuration must be provided")
	}
	if len(config.Address) == 0 {
		return nil, fmt.Errorf("Address must be provided")
	}
	if len(config.Net) == 0 {
		return nil, fmt.Errorf("Net must be provided")
	}

	l, err := net.Listen(config.Net, config.Address)
	if err != nil {
		return nil, fmt.Errorf("Unable to setup server: %s", err.Error())
	}

	return &OsdCsiServer{
		listener: l,
	}, nil
}

// Start is used to start the server.
// It will return an error if the server is already running.
func (s *OsdCsiServer) Start() error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.running {
		return fmt.Errorf("Server already running")
	}

	s.server = grpc.NewServer()

	csi.RegisterIdentityServer(s.server, s)
	reflection.Register(s.server)

	// Start listening for requests
	dlog.Infof("CSI Server ready on %s", s.Address())
	waitForServer := make(chan bool)
	s.goServe(waitForServer)
	<-waitForServer

	s.running = true
	return nil
}

// Stop is used to stop the gRPC CSI complient server.
// It can be called multiple times. It does nothing if the server
// has already been stopped.
func (s *OsdCsiServer) Stop() {
	s.lock.Lock()
	defer s.lock.Unlock()

	if !s.running {
		return
	}

	s.server.Stop()
	s.wg.Wait()
	s.running = false
}

// Address returns the address of the server which can be
// used by clients to connect.
func (s *OsdCsiServer) Address() string {
	return s.listener.Addr().String()
}

func (s *OsdCsiServer) goServe(started chan<- bool) {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		started <- true
		err := s.server.Serve(s.listener)
		if err != nil {
			dlog.Fatalf("ERROR: Unable to start gRPC server: %s\n", err.Error())
		}
	}()
}