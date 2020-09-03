// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/ptr"
)

// A second request to use or change an object was not allowed. This can result
// from retrying a request using a parameter that was not present in the original
// request.
type IdempotentParameterMismatchException struct {
	Message *string

	Code *int32
}

func (e *IdempotentParameterMismatchException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IdempotentParameterMismatchException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IdempotentParameterMismatchException) ErrorCode() string {
	return "IdempotentParameterMismatchException"
}
func (e *IdempotentParameterMismatchException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}
func (e *IdempotentParameterMismatchException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *IdempotentParameterMismatchException) HasMessage() bool {
	return e.Message != nil
}
func (e *IdempotentParameterMismatchException) GetCode() int32 {
	return ptr.ToInt32(e.Code)
}
func (e *IdempotentParameterMismatchException) HasCode() bool {
	return e.Code != nil
}

// An error on the server occurred when trying to process a request.
type InternalServerException struct {
	Message *string

	Code *int32
}

func (e *InternalServerException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerException) ErrorCode() string             { return "InternalServerException" }
func (e *InternalServerException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }
func (e *InternalServerException) GetCode() int32 {
	return ptr.ToInt32(e.Code)
}
func (e *InternalServerException) HasCode() bool {
	return e.Code != nil
}
func (e *InternalServerException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InternalServerException) HasMessage() bool {
	return e.Message != nil
}

// An error on the client occurred. Typically, the cause is an invalid input value.
type InvalidInputException struct {
	Message *string

	Code *int32
}

func (e *InvalidInputException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidInputException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidInputException) ErrorCode() string             { return "InvalidInputException" }
func (e *InvalidInputException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidInputException) GetCode() int32 {
	return ptr.ToInt32(e.Code)
}
func (e *InvalidInputException) HasCode() bool {
	return e.Code != nil
}
func (e *InvalidInputException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidInputException) HasMessage() bool {
	return e.Message != nil
}

// A submitted tag is invalid.
type InvalidTagException struct {
	Message *string
}

func (e *InvalidTagException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidTagException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidTagException) ErrorCode() string             { return "InvalidTagException" }
func (e *InvalidTagException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidTagException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidTagException) HasMessage() bool {
	return e.Message != nil
}

// The subscriber exceeded the maximum number of operations. This exception can
// occur when listing objects such as DataSource.
type LimitExceededException struct {
	Message *string

	Code *int32
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
func (e *LimitExceededException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *LimitExceededException) HasMessage() bool {
	return e.Message != nil
}
func (e *LimitExceededException) GetCode() int32 {
	return ptr.ToInt32(e.Code)
}
func (e *LimitExceededException) HasCode() bool {
	return e.Code != nil
}

// The exception is thrown when a predict request is made to an unmounted MLModel.
type PredictorNotMountedException struct {
	Message *string
}

func (e *PredictorNotMountedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PredictorNotMountedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PredictorNotMountedException) ErrorCode() string             { return "PredictorNotMountedException" }
func (e *PredictorNotMountedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *PredictorNotMountedException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *PredictorNotMountedException) HasMessage() bool {
	return e.Message != nil
}

// A specified resource cannot be located.
type ResourceNotFoundException struct {
	Message *string

	Code *int32
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
func (e *ResourceNotFoundException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceNotFoundException) HasMessage() bool {
	return e.Message != nil
}
func (e *ResourceNotFoundException) GetCode() int32 {
	return ptr.ToInt32(e.Code)
}
func (e *ResourceNotFoundException) HasCode() bool {
	return e.Code != nil
}

// The limit in the number of tags has been exceeded.
type TagLimitExceededException struct {
	Message *string
}

func (e *TagLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TagLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TagLimitExceededException) ErrorCode() string             { return "TagLimitExceededException" }
func (e *TagLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *TagLimitExceededException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *TagLimitExceededException) HasMessage() bool {
	return e.Message != nil
}
