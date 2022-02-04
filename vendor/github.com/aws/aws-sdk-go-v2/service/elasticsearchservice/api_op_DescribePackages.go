// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes all packages available to Amazon ES. Includes options for filtering,
// limiting the number of results, and pagination.
func (c *Client) DescribePackages(ctx context.Context, params *DescribePackagesInput, optFns ...func(*Options)) (*DescribePackagesOutput, error) {
	if params == nil {
		params = &DescribePackagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePackages", params, optFns, addOperationDescribePackagesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePackagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for request parameters to DescribePackage operation.
type DescribePackagesInput struct {

	// Only returns packages that match the DescribePackagesFilterList values.
	Filters []types.DescribePackagesFilter

	// Limits results to a maximum number of packages.
	MaxResults int32

	// Used for pagination. Only necessary if a previous API call includes a non-null
	// NextToken value. If provided, returns results for the next page.
	NextToken *string
}

// Container for response returned by DescribePackages operation.
type DescribePackagesOutput struct {
	NextToken *string

	// List of PackageDetails objects.
	PackageDetailsList []types.PackageDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribePackagesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribePackages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribePackages{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePackages(options.Region), middleware.Before); err != nil {
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

// DescribePackagesAPIClient is a client that implements the DescribePackages
// operation.
type DescribePackagesAPIClient interface {
	DescribePackages(context.Context, *DescribePackagesInput, ...func(*Options)) (*DescribePackagesOutput, error)
}

var _ DescribePackagesAPIClient = (*Client)(nil)

// DescribePackagesPaginatorOptions is the paginator options for DescribePackages
type DescribePackagesPaginatorOptions struct {
	// Limits results to a maximum number of packages.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribePackagesPaginator is a paginator for DescribePackages
type DescribePackagesPaginator struct {
	options   DescribePackagesPaginatorOptions
	client    DescribePackagesAPIClient
	params    *DescribePackagesInput
	nextToken *string
	firstPage bool
}

// NewDescribePackagesPaginator returns a new DescribePackagesPaginator
func NewDescribePackagesPaginator(client DescribePackagesAPIClient, params *DescribePackagesInput, optFns ...func(*DescribePackagesPaginatorOptions)) *DescribePackagesPaginator {
	if params == nil {
		params = &DescribePackagesInput{}
	}

	options := DescribePackagesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribePackagesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribePackagesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribePackages page.
func (p *DescribePackagesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribePackagesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribePackages(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribePackages(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "DescribePackages",
	}
}
