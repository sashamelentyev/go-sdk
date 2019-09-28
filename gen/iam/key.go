// Code generated by sdkgen. DO NOT EDIT.

//nolint
package iam

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"

	iam "github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
)

//revive:disable

// KeyServiceClient is a iam.KeyServiceClient with
// lazy GRPC connection initialization.
type KeyServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

var _ iam.KeyServiceClient = &KeyServiceClient{}

// Create implements iam.KeyServiceClient
func (c *KeyServiceClient) Create(ctx context.Context, in *iam.CreateKeyRequest, opts ...grpc.CallOption) (*iam.CreateKeyResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewKeyServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements iam.KeyServiceClient
func (c *KeyServiceClient) Delete(ctx context.Context, in *iam.DeleteKeyRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewKeyServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements iam.KeyServiceClient
func (c *KeyServiceClient) Get(ctx context.Context, in *iam.GetKeyRequest, opts ...grpc.CallOption) (*iam.Key, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewKeyServiceClient(conn).Get(ctx, in, opts...)
}

// List implements iam.KeyServiceClient
func (c *KeyServiceClient) List(ctx context.Context, in *iam.ListKeysRequest, opts ...grpc.CallOption) (*iam.ListKeysResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewKeyServiceClient(conn).List(ctx, in, opts...)
}
