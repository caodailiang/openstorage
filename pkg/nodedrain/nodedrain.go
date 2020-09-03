package nodedrain

import (
	"context"

	"github.com/libopenstorage/openstorage/api"
	"github.com/libopenstorage/openstorage/api/errors"
)

// Provider is a collection of APIs for performing different kinds of drain
// operations on a node
type Provider interface {
	// RemoveVolumeAttachments creates a task to drain volume attachments
	// from the provided node in the cluster.
	RemoveVolumeAttachments(ctx context.Context, in *api.SdkNodeRemoveVolumeAttachmentsRequest) (*api.SdkJobResponse, error)
}

// NewDefaultNodeDrainProvider does not any node drain related operations
func NewDefaultNodeDrainProvider() Provider {
	return &UnsupportedNodeDrainProvider{}
}

// UnsupportedNodeDrainProvider unsupported implementation of drain.
type UnsupportedNodeDrainProvider struct {
}

func (u *UnsupportedNodeDrainProvider) RemoveVolumeAttachments(ctx context.Context, in *api.SdkNodeRemoveVolumeAttachmentsRequest) (*api.SdkJobResponse, error) {
	return nil, &errors.ErrNotSupported{}
}
