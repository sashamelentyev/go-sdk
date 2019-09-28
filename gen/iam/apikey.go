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

// ApiKeyServiceClient is a iam.ApiKeyServiceClient with
// lazy GRPC connection initialization.
type ApiKeyServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

var _ iam.ApiKeyServiceClient = &ApiKeyServiceClient{}

// Create implements iam.ApiKeyServiceClient
func (c *ApiKeyServiceClient) Create(ctx context.Context, in *iam.CreateApiKeyRequest, opts ...grpc.CallOption) (*iam.CreateApiKeyResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewApiKeyServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements iam.ApiKeyServiceClient
func (c *ApiKeyServiceClient) Delete(ctx context.Context, in *iam.DeleteApiKeyRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewApiKeyServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements iam.ApiKeyServiceClient
func (c *ApiKeyServiceClient) Get(ctx context.Context, in *iam.GetApiKeyRequest, opts ...grpc.CallOption) (*iam.ApiKey, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewApiKeyServiceClient(conn).Get(ctx, in, opts...)
}

// List implements iam.ApiKeyServiceClient
func (c *ApiKeyServiceClient) List(ctx context.Context, in *iam.ListApiKeysRequest, opts ...grpc.CallOption) (*iam.ListApiKeysResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewApiKeyServiceClient(conn).List(ctx, in, opts...)
}
