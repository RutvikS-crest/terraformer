// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// For Redis engine version 6.x onwards: Deletes a user group. The user group must
// first be disassociated from the replication group before it can be deleted. For
// more information, see Using Role Based Access Control (RBAC)
// (http://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Clusters.RBAC.html).
func (c *Client) DeleteUserGroup(ctx context.Context, params *DeleteUserGroupInput, optFns ...func(*Options)) (*DeleteUserGroupOutput, error) {
	if params == nil {
		params = &DeleteUserGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteUserGroup", params, optFns, addOperationDeleteUserGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteUserGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteUserGroupInput struct {

	// The ID of the user group.
	//
	// This member is required.
	UserGroupId *string
}

type DeleteUserGroupOutput struct {

	// The Amazon Resource Name (ARN) of the user group.
	ARN *string

	// The current supported value is Redis.
	Engine *string

	// A list of updates being applied to the user groups.
	PendingChanges *types.UserGroupPendingChanges

	// A list of replication groups that the user group can access.
	ReplicationGroups []string

	// Indicates user group status. Can be "creating", "active", "modifying",
	// "deleting".
	Status *string

	// The ID of the user group.
	UserGroupId *string

	// The list of user IDs that belong to the user group.
	UserIds []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteUserGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteUserGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteUserGroup{}, middleware.After)
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
	if err = addOpDeleteUserGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteUserGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteUserGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "DeleteUserGroup",
	}
}
