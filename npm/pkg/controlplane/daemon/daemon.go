package daemon

import (
	"context"

	"github.com/Azure/azure-container-networking/npm/pkg/dataplane"
	"k8s.io/klog"
)

type Daemon struct {
	ctx    context.Context
	cancel context.CancelFunc
	nodeID string
	dp     dataplane.GenericDataplane
}

func NewDaemon(
	ctx context.Context,
	nodeID string,
	dp dataplane.GenericDataplane) *Daemon {

	klog.Infof("Creating daemon for node %s", nodeID)
	return &Daemon{
		ctx:    ctx,
		cancel: nil,
		nodeID: nodeID,
		dp:     dp,
	}
}
