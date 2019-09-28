// Code generated by sdkgen. DO NOT EDIT.

//nolint
package mysql

import (
	"context"

	"google.golang.org/grpc"

	mysql "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/mysql/v1"
)

//revive:disable

// ResourcePresetServiceClient is a mysql.ResourcePresetServiceClient with
// lazy GRPC connection initialization.
type ResourcePresetServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

var _ mysql.ResourcePresetServiceClient = &ResourcePresetServiceClient{}

// Get implements mysql.ResourcePresetServiceClient
func (c *ResourcePresetServiceClient) Get(ctx context.Context, in *mysql.GetResourcePresetRequest, opts ...grpc.CallOption) (*mysql.ResourcePreset, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewResourcePresetServiceClient(conn).Get(ctx, in, opts...)
}

// List implements mysql.ResourcePresetServiceClient
func (c *ResourcePresetServiceClient) List(ctx context.Context, in *mysql.ListResourcePresetsRequest, opts ...grpc.CallOption) (*mysql.ListResourcePresetsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewResourcePresetServiceClient(conn).List(ctx, in, opts...)
}
