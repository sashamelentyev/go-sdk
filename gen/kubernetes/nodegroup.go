// Code generated by sdkgen. DO NOT EDIT.

//nolint
package k8s

import (
	"context"

	"google.golang.org/grpc"

	k8s "github.com/yandex-cloud/go-genproto/yandex/cloud/k8s/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// NodeGroupServiceClient is a k8s.NodeGroupServiceClient with
// lazy GRPC connection initialization.
type NodeGroupServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

var _ k8s.NodeGroupServiceClient = &NodeGroupServiceClient{}

// Create implements k8s.NodeGroupServiceClient
func (c *NodeGroupServiceClient) Create(ctx context.Context, in *k8s.CreateNodeGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return k8s.NewNodeGroupServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements k8s.NodeGroupServiceClient
func (c *NodeGroupServiceClient) Delete(ctx context.Context, in *k8s.DeleteNodeGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return k8s.NewNodeGroupServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements k8s.NodeGroupServiceClient
func (c *NodeGroupServiceClient) Get(ctx context.Context, in *k8s.GetNodeGroupRequest, opts ...grpc.CallOption) (*k8s.NodeGroup, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return k8s.NewNodeGroupServiceClient(conn).Get(ctx, in, opts...)
}

// List implements k8s.NodeGroupServiceClient
func (c *NodeGroupServiceClient) List(ctx context.Context, in *k8s.ListNodeGroupsRequest, opts ...grpc.CallOption) (*k8s.ListNodeGroupsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return k8s.NewNodeGroupServiceClient(conn).List(ctx, in, opts...)
}

// ListOperations implements k8s.NodeGroupServiceClient
func (c *NodeGroupServiceClient) ListOperations(ctx context.Context, in *k8s.ListNodeGroupOperationsRequest, opts ...grpc.CallOption) (*k8s.ListNodeGroupOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return k8s.NewNodeGroupServiceClient(conn).ListOperations(ctx, in, opts...)
}

// Update implements k8s.NodeGroupServiceClient
func (c *NodeGroupServiceClient) Update(ctx context.Context, in *k8s.UpdateNodeGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return k8s.NewNodeGroupServiceClient(conn).Update(ctx, in, opts...)
}
