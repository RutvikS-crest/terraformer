// Code generated by sdkgen. DO NOT EDIT.

//nolint
package ydb

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	ydb "github.com/yandex-cloud/go-genproto/yandex/cloud/ydb/v1"
)

//revive:disable

// BackupServiceClient is a ydb.BackupServiceClient with
// lazy GRPC connection initialization.
type BackupServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Delete implements ydb.BackupServiceClient
func (c *BackupServiceClient) Delete(ctx context.Context, in *ydb.DeleteBackupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return ydb.NewBackupServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements ydb.BackupServiceClient
func (c *BackupServiceClient) Get(ctx context.Context, in *ydb.GetBackupRequest, opts ...grpc.CallOption) (*ydb.Backup, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return ydb.NewBackupServiceClient(conn).Get(ctx, in, opts...)
}

// List implements ydb.BackupServiceClient
func (c *BackupServiceClient) List(ctx context.Context, in *ydb.ListBackupsRequest, opts ...grpc.CallOption) (*ydb.ListBackupsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return ydb.NewBackupServiceClient(conn).List(ctx, in, opts...)
}

type BackupIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *BackupServiceClient
	request *ydb.ListBackupsRequest

	items []*ydb.Backup
}

func (c *BackupServiceClient) BackupIterator(ctx context.Context, req *ydb.ListBackupsRequest, opts ...grpc.CallOption) *BackupIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &BackupIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *BackupIterator) Next() bool {
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

	it.items = response.Backups
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *BackupIterator) Take(size int64) ([]*ydb.Backup, error) {
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

	var result []*ydb.Backup

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *BackupIterator) TakeAll() ([]*ydb.Backup, error) {
	return it.Take(0)
}

func (it *BackupIterator) Value() *ydb.Backup {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *BackupIterator) Error() error {
	return it.err
}

// ListPaths implements ydb.BackupServiceClient
func (c *BackupServiceClient) ListPaths(ctx context.Context, in *ydb.ListPathsRequest, opts ...grpc.CallOption) (*ydb.ListPathsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return ydb.NewBackupServiceClient(conn).ListPaths(ctx, in, opts...)
}

type BackupPathsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *BackupServiceClient
	request *ydb.ListPathsRequest

	items []string
}

func (c *BackupServiceClient) BackupPathsIterator(ctx context.Context, req *ydb.ListPathsRequest, opts ...grpc.CallOption) *BackupPathsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &BackupPathsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *BackupPathsIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
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

	response, err := it.client.ListPaths(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Paths
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *BackupPathsIterator) Take(size int64) ([]string, error) {
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

	var result []string

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *BackupPathsIterator) TakeAll() ([]string, error) {
	return it.Take(0)
}

func (it *BackupPathsIterator) Value() string {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *BackupPathsIterator) Error() error {
	return it.err
}
