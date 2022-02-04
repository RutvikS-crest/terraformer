// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// The provided iterator exceeds the maximum age allowed.
type ExpiredIteratorException struct {
	Message *string
}

func (e *ExpiredIteratorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ExpiredIteratorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ExpiredIteratorException) ErrorCode() string             { return "ExpiredIteratorException" }
func (e *ExpiredIteratorException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The pagination token passed to the operation is expired.
type ExpiredNextTokenException struct {
	Message *string
}

func (e *ExpiredNextTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ExpiredNextTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ExpiredNextTokenException) ErrorCode() string             { return "ExpiredNextTokenException" }
func (e *ExpiredNextTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A specified parameter exceeds its restrictions, is not supported, or can't be
// used. For more information, see the returned message.
type InvalidArgumentException struct {
	Message *string
}

func (e *InvalidArgumentException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidArgumentException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidArgumentException) ErrorCode() string             { return "InvalidArgumentException" }
func (e *InvalidArgumentException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The ciphertext references a key that doesn't exist or that you don't have access
// to.
type KMSAccessDeniedException struct {
	Message *string
}

func (e *KMSAccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KMSAccessDeniedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KMSAccessDeniedException) ErrorCode() string             { return "KMSAccessDeniedException" }
func (e *KMSAccessDeniedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because the specified customer master key (CMK) isn't
// enabled.
type KMSDisabledException struct {
	Message *string
}

func (e *KMSDisabledException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KMSDisabledException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KMSDisabledException) ErrorCode() string             { return "KMSDisabledException" }
func (e *KMSDisabledException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because the state of the specified resource isn't valid
// for this request. For more information, see How Key State Affects Use of a
// Customer Master Key
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
// AWS Key Management Service Developer Guide.
type KMSInvalidStateException struct {
	Message *string
}

func (e *KMSInvalidStateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KMSInvalidStateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KMSInvalidStateException) ErrorCode() string             { return "KMSInvalidStateException" }
func (e *KMSInvalidStateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because the specified entity or resource can't be
// found.
type KMSNotFoundException struct {
	Message *string
}

func (e *KMSNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KMSNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KMSNotFoundException) ErrorCode() string             { return "KMSNotFoundException" }
func (e *KMSNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The AWS access key ID needs a subscription for the service.
type KMSOptInRequired struct {
	Message *string
}

func (e *KMSOptInRequired) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KMSOptInRequired) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KMSOptInRequired) ErrorCode() string             { return "KMSOptInRequired" }
func (e *KMSOptInRequired) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was denied due to request throttling. For more information about
// throttling, see Limits
// (https://docs.aws.amazon.com/kms/latest/developerguide/limits.html#requests-per-second)
// in the AWS Key Management Service Developer Guide.
type KMSThrottlingException struct {
	Message *string
}

func (e *KMSThrottlingException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KMSThrottlingException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KMSThrottlingException) ErrorCode() string             { return "KMSThrottlingException" }
func (e *KMSThrottlingException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested resource exceeds the maximum number allowed, or the number of
// concurrent stream requests exceeds the maximum number allowed.
type LimitExceededException struct {
	Message *string
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string             { return "LimitExceededException" }
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request rate for the stream is too high, or the requested data is too large
// for the available throughput. Reduce the frequency or size of your requests. For
// more information, see Streams Limits
// (https://docs.aws.amazon.com/kinesis/latest/dev/service-sizes-and-limits.html)
// in the Amazon Kinesis Data Streams Developer Guide, and Error Retries and
// Exponential Backoff in AWS
// (https://docs.aws.amazon.com/general/latest/gr/api-retries.html) in the AWS
// General Reference.
type ProvisionedThroughputExceededException struct {
	Message *string
}

func (e *ProvisionedThroughputExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ProvisionedThroughputExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ProvisionedThroughputExceededException) ErrorCode() string {
	return "ProvisionedThroughputExceededException"
}
func (e *ProvisionedThroughputExceededException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The resource is not available for this operation. For successful operation, the
// resource must be in the ACTIVE state.
type ResourceInUseException struct {
	Message *string
}

func (e *ResourceInUseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceInUseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceInUseException) ErrorCode() string             { return "ResourceInUseException" }
func (e *ResourceInUseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested resource could not be found. The stream might not be specified
// correctly.
type ResourceNotFoundException struct {
	Message *string
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string             { return "ResourceNotFoundException" }
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
