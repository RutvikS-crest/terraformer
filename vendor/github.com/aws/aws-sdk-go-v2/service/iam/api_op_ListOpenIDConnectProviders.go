// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists information about the IAM OpenID Connect (OIDC) provider resource objects
// defined in the AWS account. IAM resource-listing operations return a subset of
// the available attributes for the resource. For example, this operation does not
// return tags, even though they are an attribute of the returned object. To view
// all of the information for an OIDC provider, see GetOpenIDConnectProvider.
func (c *Client) ListOpenIDConnectProviders(ctx context.Context, params *ListOpenIDConnectProvidersInput, optFns ...func(*Options)) (*ListOpenIDConnectProvidersOutput, error) {
	if params == nil {
		params = &ListOpenIDConnectProvidersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOpenIDConnectProviders", params, optFns, addOperationListOpenIDConnectProvidersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOpenIDConnectProvidersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOpenIDConnectProvidersInput struct {
}

// Contains the response to a successful ListOpenIDConnectProviders request.
type ListOpenIDConnectProvidersOutput struct {

	// The list of IAM OIDC provider resource objects defined in the AWS account.
	OpenIDConnectProviderList []types.OpenIDConnectProviderListEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListOpenIDConnectProvidersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListOpenIDConnectProviders{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListOpenIDConnectProviders{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOpenIDConnectProviders(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListOpenIDConnectProviders(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "ListOpenIDConnectProviders",
	}
}
