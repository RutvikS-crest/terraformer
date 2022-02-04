// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of products or offerings that the user can manage through the
// API. Each offering record indicates the recurring price per unit and the
// frequency for that offering. The API returns a NotEligible error if the user is
// not permitted to invoke the operation. If you must be able to invoke this
// operation, contact aws-devicefarm-support@amazon.com
// (mailto:aws-devicefarm-support@amazon.com).
func (c *Client) ListOfferings(ctx context.Context, params *ListOfferingsInput, optFns ...func(*Options)) (*ListOfferingsOutput, error) {
	if params == nil {
		params = &ListOfferingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOfferings", params, optFns, addOperationListOfferingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOfferingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to list all offerings.
type ListOfferingsInput struct {

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string
}

// Represents the return values of the list of offerings.
type ListOfferingsOutput struct {

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string

	// A value that represents the list offering results.
	Offerings []types.Offering

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListOfferingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListOfferings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListOfferings{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOfferings(options.Region), middleware.Before); err != nil {
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

// ListOfferingsAPIClient is a client that implements the ListOfferings operation.
type ListOfferingsAPIClient interface {
	ListOfferings(context.Context, *ListOfferingsInput, ...func(*Options)) (*ListOfferingsOutput, error)
}

var _ ListOfferingsAPIClient = (*Client)(nil)

// ListOfferingsPaginatorOptions is the paginator options for ListOfferings
type ListOfferingsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListOfferingsPaginator is a paginator for ListOfferings
type ListOfferingsPaginator struct {
	options   ListOfferingsPaginatorOptions
	client    ListOfferingsAPIClient
	params    *ListOfferingsInput
	nextToken *string
	firstPage bool
}

// NewListOfferingsPaginator returns a new ListOfferingsPaginator
func NewListOfferingsPaginator(client ListOfferingsAPIClient, params *ListOfferingsInput, optFns ...func(*ListOfferingsPaginatorOptions)) *ListOfferingsPaginator {
	if params == nil {
		params = &ListOfferingsInput{}
	}

	options := ListOfferingsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListOfferingsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListOfferingsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListOfferings page.
func (p *ListOfferingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListOfferingsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListOfferings(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListOfferings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "ListOfferings",
	}
}
