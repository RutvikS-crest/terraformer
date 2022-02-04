// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a file system, permanently severing access to its contents. Upon return,
// the file system no longer exists and you can't access any contents of the
// deleted file system. You can't delete a file system that is in use. That is, if
// the file system has any mount targets, you must first delete them. For more
// information, see DescribeMountTargets and DeleteMountTarget. The
// DeleteFileSystem call returns while the file system state is still deleting. You
// can check the file system deletion status by calling the DescribeFileSystems
// operation, which returns a list of file systems in your account. If you pass
// file system ID or creation token for the deleted file system, the
// DescribeFileSystems returns a 404 FileSystemNotFound error. This operation
// requires permissions for the elasticfilesystem:DeleteFileSystem action.
func (c *Client) DeleteFileSystem(ctx context.Context, params *DeleteFileSystemInput, optFns ...func(*Options)) (*DeleteFileSystemOutput, error) {
	if params == nil {
		params = &DeleteFileSystemInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteFileSystem", params, optFns, addOperationDeleteFileSystemMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteFileSystemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DeleteFileSystemInput struct {

	// The ID of the file system you want to delete.
	//
	// This member is required.
	FileSystemId *string
}

type DeleteFileSystemOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteFileSystemMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteFileSystem{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteFileSystem{}, middleware.After)
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
	if err = addOpDeleteFileSystemValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteFileSystem(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteFileSystem(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticfilesystem",
		OperationName: "DeleteFileSystem",
	}
}
