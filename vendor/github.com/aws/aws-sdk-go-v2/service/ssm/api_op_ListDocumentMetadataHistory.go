// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Information about approval reviews for a version of an SSM document.
func (c *Client) ListDocumentMetadataHistory(ctx context.Context, params *ListDocumentMetadataHistoryInput, optFns ...func(*Options)) (*ListDocumentMetadataHistoryOutput, error) {
	if params == nil {
		params = &ListDocumentMetadataHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDocumentMetadataHistory", params, optFns, addOperationListDocumentMetadataHistoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDocumentMetadataHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDocumentMetadataHistoryInput struct {

	// The type of data for which details are being requested. Currently, the only
	// supported value is DocumentReviews.
	//
	// This member is required.
	Metadata types.DocumentMetadataEnum

	// The name of the document.
	//
	// This member is required.
	Name *string

	// The version of the document.
	DocumentVersion *string

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string
}

type ListDocumentMetadataHistoryOutput struct {

	// The user ID of the person in the organization who requested the document review.
	Author *string

	// The version of the document.
	DocumentVersion *string

	// Information about the response to the document approval request.
	Metadata *types.DocumentMetadataResponseInfo

	// The name of the document.
	Name *string

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDocumentMetadataHistoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDocumentMetadataHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDocumentMetadataHistory{}, middleware.After)
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
	if err = addOpListDocumentMetadataHistoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDocumentMetadataHistory(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListDocumentMetadataHistory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "ListDocumentMetadataHistory",
	}
}
