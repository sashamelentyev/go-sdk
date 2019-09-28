// Code generated by sdkgen. DO NOT EDIT.

//nolint
package iam

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	iam "github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// ServiceAccountServiceClient is a iam.ServiceAccountServiceClient with
// lazy GRPC connection initialization.
type ServiceAccountServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

var _ iam.ServiceAccountServiceClient = &ServiceAccountServiceClient{}

// Create implements iam.ServiceAccountServiceClient
func (c *ServiceAccountServiceClient) Create(ctx context.Context, in *iam.CreateServiceAccountRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewServiceAccountServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements iam.ServiceAccountServiceClient
func (c *ServiceAccountServiceClient) Delete(ctx context.Context, in *iam.DeleteServiceAccountRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewServiceAccountServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements iam.ServiceAccountServiceClient
func (c *ServiceAccountServiceClient) Get(ctx context.Context, in *iam.GetServiceAccountRequest, opts ...grpc.CallOption) (*iam.ServiceAccount, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewServiceAccountServiceClient(conn).Get(ctx, in, opts...)
}

// List implements iam.ServiceAccountServiceClient
func (c *ServiceAccountServiceClient) List(ctx context.Context, in *iam.ListServiceAccountsRequest, opts ...grpc.CallOption) (*iam.ListServiceAccountsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewServiceAccountServiceClient(conn).List(ctx, in, opts...)
}

// ListAccessBindings implements iam.ServiceAccountServiceClient
func (c *ServiceAccountServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewServiceAccountServiceClient(conn).ListAccessBindings(ctx, in, opts...)
}

// ListOperations implements iam.ServiceAccountServiceClient
func (c *ServiceAccountServiceClient) ListOperations(ctx context.Context, in *iam.ListServiceAccountOperationsRequest, opts ...grpc.CallOption) (*iam.ListServiceAccountOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewServiceAccountServiceClient(conn).ListOperations(ctx, in, opts...)
}

// SetAccessBindings implements iam.ServiceAccountServiceClient
func (c *ServiceAccountServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewServiceAccountServiceClient(conn).SetAccessBindings(ctx, in, opts...)
}

// Update implements iam.ServiceAccountServiceClient
func (c *ServiceAccountServiceClient) Update(ctx context.Context, in *iam.UpdateServiceAccountRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewServiceAccountServiceClient(conn).Update(ctx, in, opts...)
}

// UpdateAccessBindings implements iam.ServiceAccountServiceClient
func (c *ServiceAccountServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewServiceAccountServiceClient(conn).UpdateAccessBindings(ctx, in, opts...)
}
