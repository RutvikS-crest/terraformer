// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of the dashboards for your account. If you include
// DashboardNamePrefix, only those dashboards with names starting with the prefix
// are listed. Otherwise, all dashboards in your account are listed. ListDashboards
// returns up to 1000 results on one page. If there are more than 1000 dashboards,
// you can call ListDashboards again and include the value you received for
// NextToken in the first call, to receive the next 1000 results.
func (c *Client) ListDashboards(ctx context.Context, params *ListDashboardsInput, optFns ...func(*Options)) (*ListDashboardsOutput, error) {
	if params == nil {
		params = &ListDashboardsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDashboards", params, optFns, addOperationListDashboardsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDashboardsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDashboardsInput struct {

	// If you specify this parameter, only the dashboards with names starting with the
	// specified string are listed. The maximum length is 255, and valid characters are
	// A-Z, a-z, 0-9, ".", "-", and "_".
	DashboardNamePrefix *string

	// The token returned by a previous call to indicate that there is more data
	// available.
	NextToken *string
}

type ListDashboardsOutput struct {

	// The list of matching dashboards.
	DashboardEntries []types.DashboardEntry

	// The token that marks the start of the next batch of returned results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDashboardsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListDashboards{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListDashboards{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDashboards(options.Region), middleware.Before); err != nil {
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

// ListDashboardsAPIClient is a client that implements the ListDashboards
// operation.
type ListDashboardsAPIClient interface {
	ListDashboards(context.Context, *ListDashboardsInput, ...func(*Options)) (*ListDashboardsOutput, error)
}

var _ ListDashboardsAPIClient = (*Client)(nil)

// ListDashboardsPaginatorOptions is the paginator options for ListDashboards
type ListDashboardsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDashboardsPaginator is a paginator for ListDashboards
type ListDashboardsPaginator struct {
	options   ListDashboardsPaginatorOptions
	client    ListDashboardsAPIClient
	params    *ListDashboardsInput
	nextToken *string
	firstPage bool
}

// NewListDashboardsPaginator returns a new ListDashboardsPaginator
func NewListDashboardsPaginator(client ListDashboardsAPIClient, params *ListDashboardsInput, optFns ...func(*ListDashboardsPaginatorOptions)) *ListDashboardsPaginator {
	if params == nil {
		params = &ListDashboardsInput{}
	}

	options := ListDashboardsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDashboardsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDashboardsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListDashboards page.
func (p *ListDashboardsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDashboardsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListDashboards(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListDashboards(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "monitoring",
		OperationName: "ListDashboards",
	}
}
