// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackage

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediapackage/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets details about an existing HarvestJob.
func (c *Client) DescribeHarvestJob(ctx context.Context, params *DescribeHarvestJobInput, optFns ...func(*Options)) (*DescribeHarvestJobOutput, error) {
	if params == nil {
		params = &DescribeHarvestJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeHarvestJob", params, optFns, c.addOperationDescribeHarvestJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeHarvestJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeHarvestJobInput struct {

	// The ID of the HarvestJob.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type DescribeHarvestJobOutput struct {

	// The Amazon Resource Name (ARN) assigned to the HarvestJob.
	Arn *string

	// The ID of the Channel that the HarvestJob will harvest from.
	ChannelId *string

	// The time the HarvestJob was submitted
	CreatedAt *string

	// The end of the time-window which will be harvested.
	EndTime *string

	// The ID of the HarvestJob. The ID must be unique within the region and it cannot
	// be changed after the HarvestJob is submitted.
	Id *string

	// The ID of the OriginEndpoint that the HarvestJob will harvest from. This cannot
	// be changed after the HarvestJob is submitted.
	OriginEndpointId *string

	// Configuration parameters for where in an S3 bucket to place the harvested
	// content
	S3Destination *types.S3Destination

	// The start of the time-window which will be harvested.
	StartTime *string

	// The current status of the HarvestJob. Consider setting up a CloudWatch Event to
	// listen for HarvestJobs as they succeed or fail. In the event of failure, the
	// CloudWatch Event will include an explanation of why the HarvestJob failed.
	Status types.Status

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeHarvestJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeHarvestJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeHarvestJob{}, middleware.After)
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
	if err = addOpDescribeHarvestJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeHarvestJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeHarvestJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediapackage",
		OperationName: "DescribeHarvestJob",
	}
}
