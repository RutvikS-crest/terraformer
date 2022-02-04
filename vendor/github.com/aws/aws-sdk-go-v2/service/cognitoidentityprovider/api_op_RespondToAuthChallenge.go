// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Responds to the authentication challenge.
func (c *Client) RespondToAuthChallenge(ctx context.Context, params *RespondToAuthChallengeInput, optFns ...func(*Options)) (*RespondToAuthChallengeOutput, error) {
	if params == nil {
		params = &RespondToAuthChallengeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RespondToAuthChallenge", params, optFns, addOperationRespondToAuthChallengeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RespondToAuthChallengeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to respond to an authentication challenge.
type RespondToAuthChallengeInput struct {

	// The challenge name. For more information, see InitiateAuth
	// (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_InitiateAuth.html).
	// ADMIN_NO_SRP_AUTH is not a valid value.
	//
	// This member is required.
	ChallengeName types.ChallengeNameType

	// The app client ID.
	//
	// This member is required.
	ClientId *string

	// The Amazon Pinpoint analytics metadata for collecting metrics for
	// RespondToAuthChallenge calls.
	AnalyticsMetadata *types.AnalyticsMetadataType

	// The challenge responses. These are inputs corresponding to the value of
	// ChallengeName, for example: SECRET_HASH (if app client is configured with client
	// secret) applies to all inputs below (including SOFTWARE_TOKEN_MFA).
	//
	// * SMS_MFA:
	// SMS_MFA_CODE, USERNAME.
	//
	// * PASSWORD_VERIFIER: PASSWORD_CLAIM_SIGNATURE,
	// PASSWORD_CLAIM_SECRET_BLOCK, TIMESTAMP, USERNAME.
	//
	// * NEW_PASSWORD_REQUIRED:
	// NEW_PASSWORD, any other required attributes, USERNAME.
	//
	// * SOFTWARE_TOKEN_MFA:
	// USERNAME and SOFTWARE_TOKEN_MFA_CODE are required attributes.
	//
	// * DEVICE_SRP_AUTH
	// requires USERNAME, DEVICE_KEY, SRP_A (and SECRET_HASH).
	//
	// *
	// DEVICE_PASSWORD_VERIFIER requires everything that PASSWORD_VERIFIER requires
	// plus DEVICE_KEY.
	ChallengeResponses map[string]string

	// A map of custom key-value pairs that you can provide as input for any custom
	// workflows that this action triggers. You create custom workflows by assigning
	// AWS Lambda functions to user pool triggers. When you use the
	// RespondToAuthChallenge API action, Amazon Cognito invokes any functions that are
	// assigned to the following triggers: post authentication, pre token generation,
	// define auth challenge, create auth challenge, and verify auth challenge. When
	// Amazon Cognito invokes any of these functions, it passes a JSON payload, which
	// the function receives as input. This payload contains a clientMetadata
	// attribute, which provides the data that you assigned to the ClientMetadata
	// parameter in your RespondToAuthChallenge request. In your function code in AWS
	// Lambda, you can process the clientMetadata value to enhance your workflow for
	// your specific needs. For more information, see Customizing User Pool Workflows
	// with Lambda Triggers
	// (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html)
	// in the Amazon Cognito Developer Guide. Take the following limitations into
	// consideration when you use the ClientMetadata parameter:
	//
	// * Amazon Cognito does
	// not store the ClientMetadata value. This data is available only to AWS Lambda
	// triggers that are assigned to a user pool to support custom workflows. If your
	// user pool configuration does not include triggers, the ClientMetadata parameter
	// serves no purpose.
	//
	// * Amazon Cognito does not validate the ClientMetadata
	// value.
	//
	// * Amazon Cognito does not encrypt the the ClientMetadata value, so don't
	// use it to provide sensitive information.
	ClientMetadata map[string]string

	// The session which should be passed both ways in challenge-response calls to the
	// service. If InitiateAuth or RespondToAuthChallenge API call determines that the
	// caller needs to go through another challenge, they return a session with other
	// challenge parameters. This session should be passed as it is to the next
	// RespondToAuthChallenge API call.
	Session *string

	// Contextual data such as the user's device fingerprint, IP address, or location
	// used for evaluating the risk of an unexpected event by Amazon Cognito advanced
	// security.
	UserContextData *types.UserContextDataType
}

// The response to respond to the authentication challenge.
type RespondToAuthChallengeOutput struct {

	// The result returned by the server in response to the request to respond to the
	// authentication challenge.
	AuthenticationResult *types.AuthenticationResultType

	// The challenge name. For more information, see InitiateAuth
	// (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_InitiateAuth.html).
	ChallengeName types.ChallengeNameType

	// The challenge parameters. For more information, see InitiateAuth
	// (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_InitiateAuth.html).
	ChallengeParameters map[string]string

	// The session which should be passed both ways in challenge-response calls to the
	// service. If the caller needs to go through another challenge, they return a
	// session with other challenge parameters. This session should be passed as it is
	// to the next RespondToAuthChallenge API call.
	Session *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRespondToAuthChallengeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRespondToAuthChallenge{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRespondToAuthChallenge{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = addOpRespondToAuthChallengeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRespondToAuthChallenge(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRespondToAuthChallenge(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RespondToAuthChallenge",
	}
}
