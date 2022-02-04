// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves settings for one or more trails associated with the current region for
// your account.
func (c *Client) DescribeTrails(ctx context.Context, params *DescribeTrailsInput, optFns ...func(*Options)) (*DescribeTrailsOutput, error) {
	if params == nil {
		params = &DescribeTrailsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTrails", params, optFns, addOperationDescribeTrailsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTrailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Returns information about the trail.
type DescribeTrailsInput struct {

	// Specifies whether to include shadow trails in the response. A shadow trail is
	// the replication in a region of a trail that was created in a different region,
	// or in the case of an organization trail, the replication of an organization
	// trail in member accounts. If you do not include shadow trails, organization
	// trails in a member account and region replication trails will not be returned.
	// The default is true.
	IncludeShadowTrails *bool

	// Specifies a list of trail names, trail ARNs, or both, of the trails to describe.
	// The format of a trail ARN is:
	// arn:aws:cloudtrail:us-east-2:123456789012:trail/MyTrail If an empty list is
	// specified, information for the trail in the current region is returned.
	//
	// * If an
	// empty list is specified and IncludeShadowTrails is false, then information for
	// all trails in the current region is returned.
	//
	// * If an empty list is specified
	// and IncludeShadowTrails is null or true, then information for all trails in the
	// current region and any associated shadow trails in other regions is
	// returned.
	//
	// If one or more trail names are specified, information is returned
	// only if the names match the names of trails belonging only to the current
	// region. To return information about a trail in another region, you must specify
	// its trail ARN.
	TrailNameList []string
}

// Returns the objects or data listed below if successful. Otherwise, returns an
// error.
type DescribeTrailsOutput struct {

	// The list of trail objects. Trail objects with string values are only returned if
	// values for the objects exist in a trail's configuration. For example,
	// SNSTopicName and SNSTopicARN are only returned in results if a trail is
	// configured to send SNS notifications. Similarly, KMSKeyId only appears in
	// results if a trail's log files are encrypted with AWS KMS-managed keys.
	TrailList []types.Trail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeTrailsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTrails{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTrails{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTrails(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeTrails(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudtrail",
		OperationName: "DescribeTrails",
	}
}
