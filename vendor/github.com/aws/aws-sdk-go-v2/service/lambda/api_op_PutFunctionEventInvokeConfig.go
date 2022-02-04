// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Configures options for asynchronous invocation
// (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html) on a
// function, version, or alias. If a configuration already exists for a function,
// version, or alias, this operation overwrites it. If you exclude any settings,
// they are removed. To set one option without affecting existing settings for
// other options, use UpdateFunctionEventInvokeConfig. By default, Lambda retries
// an asynchronous invocation twice if the function returns an error. It retains
// events in a queue for up to six hours. When an event fails all processing
// attempts or stays in the asynchronous invocation queue for too long, Lambda
// discards it. To retain discarded events, configure a dead-letter queue with
// UpdateFunctionConfiguration. To send an invocation record to a queue, topic,
// function, or event bus, specify a destination
// (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#invocation-async-destinations).
// You can configure separate destinations for successful invocations (on-success)
// and events that fail all processing attempts (on-failure). You can configure
// destinations in addition to or instead of a dead-letter queue.
func (c *Client) PutFunctionEventInvokeConfig(ctx context.Context, params *PutFunctionEventInvokeConfigInput, optFns ...func(*Options)) (*PutFunctionEventInvokeConfigOutput, error) {
	if params == nil {
		params = &PutFunctionEventInvokeConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutFunctionEventInvokeConfig", params, optFns, addOperationPutFunctionEventInvokeConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutFunctionEventInvokeConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutFunctionEventInvokeConfigInput struct {

	// The name of the Lambda function, version, or alias. Name formats
	//
	// * Function
	// name - my-function (name-only), my-function:v1 (with alias).
	//
	// * Function ARN -
	// arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	// * Partial ARN -
	// 123456789012:function:my-function.
	//
	// You can append a version number or alias to
	// any of the formats. The length constraint applies only to the full ARN. If you
	// specify only the function name, it is limited to 64 characters in length.
	//
	// This member is required.
	FunctionName *string

	// A destination for events after they have been sent to a function for processing.
	// Destinations
	//
	// * Function - The Amazon Resource Name (ARN) of a Lambda
	// function.
	//
	// * Queue - The ARN of an SQS queue.
	//
	// * Topic - The ARN of an SNS
	// topic.
	//
	// * Event Bus - The ARN of an Amazon EventBridge event bus.
	DestinationConfig *types.DestinationConfig

	// The maximum age of a request that Lambda sends to a function for processing.
	MaximumEventAgeInSeconds *int32

	// The maximum number of times to retry when the function returns an error.
	MaximumRetryAttempts *int32

	// A version number or alias name.
	Qualifier *string
}

type PutFunctionEventInvokeConfigOutput struct {

	// A destination for events after they have been sent to a function for processing.
	// Destinations
	//
	// * Function - The Amazon Resource Name (ARN) of a Lambda
	// function.
	//
	// * Queue - The ARN of an SQS queue.
	//
	// * Topic - The ARN of an SNS
	// topic.
	//
	// * Event Bus - The ARN of an Amazon EventBridge event bus.
	DestinationConfig *types.DestinationConfig

	// The Amazon Resource Name (ARN) of the function.
	FunctionArn *string

	// The date and time that the configuration was last updated.
	LastModified *time.Time

	// The maximum age of a request that Lambda sends to a function for processing.
	MaximumEventAgeInSeconds *int32

	// The maximum number of times to retry when the function returns an error.
	MaximumRetryAttempts *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutFunctionEventInvokeConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutFunctionEventInvokeConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutFunctionEventInvokeConfig{}, middleware.After)
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
	if err = addOpPutFunctionEventInvokeConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutFunctionEventInvokeConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutFunctionEventInvokeConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "PutFunctionEventInvokeConfig",
	}
}
