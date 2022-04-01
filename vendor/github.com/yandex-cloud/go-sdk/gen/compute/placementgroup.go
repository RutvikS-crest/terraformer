// Code generated by sdkgen. DO NOT EDIT.

//nolint
package compute

import (
	"context"

	"google.golang.org/grpc"

	compute "github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// PlacementGroupServiceClient is a compute.PlacementGroupServiceClient with
// lazy GRPC connection initialization.
type PlacementGroupServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements compute.PlacementGroupServiceClient
func (c *PlacementGroupServiceClient) Create(ctx context.Context, in *compute.CreatePlacementGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewPlacementGroupServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements compute.PlacementGroupServiceClient
func (c *PlacementGroupServiceClient) Delete(ctx context.Context, in *compute.DeletePlacementGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewPlacementGroupServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements compute.PlacementGroupServiceClient
func (c *PlacementGroupServiceClient) Get(ctx context.Context, in *compute.GetPlacementGroupRequest, opts ...grpc.CallOption) (*compute.PlacementGroup, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewPlacementGroupServiceClient(conn).Get(ctx, in, opts...)
}

// List implements compute.PlacementGroupServiceClient
func (c *PlacementGroupServiceClient) List(ctx context.Context, in *compute.ListPlacementGroupsRequest, opts ...grpc.CallOption) (*compute.ListPlacementGroupsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewPlacementGroupServiceClient(conn).List(ctx, in, opts...)
}

type PlacementGroupIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *PlacementGroupServiceClient
	request *compute.ListPlacementGroupsRequest

	items []*compute.PlacementGroup
}

func (c *PlacementGroupServiceClient) PlacementGroupIterator(ctx context.Context, req *compute.ListPlacementGroupsRequest, opts ...grpc.CallOption) *PlacementGroupIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &PlacementGroupIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *PlacementGroupIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.List(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.PlacementGroups
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *PlacementGroupIterator) Take(size int64) ([]*compute.PlacementGroup, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*compute.PlacementGroup

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *PlacementGroupIterator) TakeAll() ([]*compute.PlacementGroup, error) {
	return it.Take(0)
}

func (it *PlacementGroupIterator) Value() *compute.PlacementGroup {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *PlacementGroupIterator) Error() error {
	return it.err
}

// ListInstances implements compute.PlacementGroupServiceClient
func (c *PlacementGroupServiceClient) ListInstances(ctx context.Context, in *compute.ListPlacementGroupInstancesRequest, opts ...grpc.CallOption) (*compute.ListPlacementGroupInstancesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewPlacementGroupServiceClient(conn).ListInstances(ctx, in, opts...)
}

type PlacementGroupInstancesIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *PlacementGroupServiceClient
	request *compute.ListPlacementGroupInstancesRequest

	items []*compute.Instance
}

func (c *PlacementGroupServiceClient) PlacementGroupInstancesIterator(ctx context.Context, req *compute.ListPlacementGroupInstancesRequest, opts ...grpc.CallOption) *PlacementGroupInstancesIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &PlacementGroupInstancesIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *PlacementGroupInstancesIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.ListInstances(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Instances
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *PlacementGroupInstancesIterator) Take(size int64) ([]*compute.Instance, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*compute.Instance

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *PlacementGroupInstancesIterator) TakeAll() ([]*compute.Instance, error) {
	return it.Take(0)
}

func (it *PlacementGroupInstancesIterator) Value() *compute.Instance {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *PlacementGroupInstancesIterator) Error() error {
	return it.err
}

// ListOperations implements compute.PlacementGroupServiceClient
func (c *PlacementGroupServiceClient) ListOperations(ctx context.Context, in *compute.ListPlacementGroupOperationsRequest, opts ...grpc.CallOption) (*compute.ListPlacementGroupOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewPlacementGroupServiceClient(conn).ListOperations(ctx, in, opts...)
}

type PlacementGroupOperationsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *PlacementGroupServiceClient
	request *compute.ListPlacementGroupOperationsRequest

	items []*operation.Operation
}

func (c *PlacementGroupServiceClient) PlacementGroupOperationsIterator(ctx context.Context, req *compute.ListPlacementGroupOperationsRequest, opts ...grpc.CallOption) *PlacementGroupOperationsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &PlacementGroupOperationsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *PlacementGroupOperationsIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.ListOperations(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Operations
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *PlacementGroupOperationsIterator) Take(size int64) ([]*operation.Operation, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*operation.Operation

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *PlacementGroupOperationsIterator) TakeAll() ([]*operation.Operation, error) {
	return it.Take(0)
}

func (it *PlacementGroupOperationsIterator) Value() *operation.Operation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *PlacementGroupOperationsIterator) Error() error {
	return it.err
}

// Update implements compute.PlacementGroupServiceClient
func (c *PlacementGroupServiceClient) Update(ctx context.Context, in *compute.UpdatePlacementGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewPlacementGroupServiceClient(conn).Update(ctx, in, opts...)
}
