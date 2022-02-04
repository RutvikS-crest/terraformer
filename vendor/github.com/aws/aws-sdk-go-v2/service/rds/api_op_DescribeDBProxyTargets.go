// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about DBProxyTarget objects. This API supports pagination.
func (c *Client) DescribeDBProxyTargets(ctx context.Context, params *DescribeDBProxyTargetsInput, optFns ...func(*Options)) (*DescribeDBProxyTargetsOutput, error) {
	if params == nil {
		params = &DescribeDBProxyTargetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBProxyTargets", params, optFns, c.addOperationDescribeDBProxyTargetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBProxyTargetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDBProxyTargetsInput struct {

	// The identifier of the DBProxyTarget to describe.
	//
	// This member is required.
	DBProxyName *string

	// This parameter is not currently supported.
	Filters []types.Filter

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	// The identifier of the DBProxyTargetGroup to describe.
	TargetGroupName *string

	noSmithyDocumentSerde
}

type DescribeDBProxyTargetsOutput struct {

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// An arbitrary number of DBProxyTarget objects, containing details of the
	// corresponding targets.
	Targets []types.DBProxyTarget

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDBProxyTargetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBProxyTargets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBProxyTargets{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeDBProxyTargetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBProxyTargets(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeDBProxyTargetsAPIClient is a client that implements the
// DescribeDBProxyTargets operation.
type DescribeDBProxyTargetsAPIClient interface {
	DescribeDBProxyTargets(context.Context, *DescribeDBProxyTargetsInput, ...func(*Options)) (*DescribeDBProxyTargetsOutput, error)
}

var _ DescribeDBProxyTargetsAPIClient = (*Client)(nil)

// DescribeDBProxyTargetsPaginatorOptions is the paginator options for
// DescribeDBProxyTargets
type DescribeDBProxyTargetsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeDBProxyTargetsPaginator is a paginator for DescribeDBProxyTargets
type DescribeDBProxyTargetsPaginator struct {
	options   DescribeDBProxyTargetsPaginatorOptions
	client    DescribeDBProxyTargetsAPIClient
	params    *DescribeDBProxyTargetsInput
	nextToken *string
	firstPage bool
}

// NewDescribeDBProxyTargetsPaginator returns a new DescribeDBProxyTargetsPaginator
func NewDescribeDBProxyTargetsPaginator(client DescribeDBProxyTargetsAPIClient, params *DescribeDBProxyTargetsInput, optFns ...func(*DescribeDBProxyTargetsPaginatorOptions)) *DescribeDBProxyTargetsPaginator {
	if params == nil {
		params = &DescribeDBProxyTargetsInput{}
	}

	options := DescribeDBProxyTargetsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeDBProxyTargetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDBProxyTargetsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeDBProxyTargets page.
func (p *DescribeDBProxyTargetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDBProxyTargetsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeDBProxyTargets(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeDBProxyTargets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBProxyTargets",
	}
}
