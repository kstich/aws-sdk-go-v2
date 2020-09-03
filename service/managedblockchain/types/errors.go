// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/ptr"
)

// You do not have sufficient access to perform this action.
type AccessDeniedException struct {
	Message *string
}

func (e *AccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessDeniedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessDeniedException) ErrorCode() string             { return "AccessDeniedException" }
func (e *AccessDeniedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

//
type IllegalActionException struct {
	Message *string
}

func (e *IllegalActionException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IllegalActionException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IllegalActionException) ErrorCode() string             { return "IllegalActionException" }
func (e *IllegalActionException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *IllegalActionException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *IllegalActionException) HasMessage() bool {
	return e.Message != nil
}

// The request processing has failed because of an unknown error, exception or
// failure.
type InternalServiceErrorException struct {
	Message *string
}

func (e *InternalServiceErrorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServiceErrorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServiceErrorException) ErrorCode() string             { return "InternalServiceErrorException" }
func (e *InternalServiceErrorException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The action or operation requested is invalid. Verify that the action is typed
// correctly.
type InvalidRequestException struct {
	Message *string
}

func (e *InvalidRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRequestException) ErrorCode() string             { return "InvalidRequestException" }
func (e *InvalidRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidRequestException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidRequestException) HasMessage() bool {
	return e.Message != nil
}

// A resource request is issued for a resource that already exists.
type ResourceAlreadyExistsException struct {
	Message *string
}

func (e *ResourceAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceAlreadyExistsException) ErrorCode() string             { return "ResourceAlreadyExistsException" }
func (e *ResourceAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ResourceAlreadyExistsException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceAlreadyExistsException) HasMessage() bool {
	return e.Message != nil
}

// The maximum number of resources of that type already exist. Ensure the resources
// requested are within the boundaries of the service edition and your account
// limits.
type ResourceLimitExceededException struct {
	Message *string
}

func (e *ResourceLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceLimitExceededException) ErrorCode() string             { return "ResourceLimitExceededException" }
func (e *ResourceLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ResourceLimitExceededException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceLimitExceededException) HasMessage() bool {
	return e.Message != nil
}

// A requested resource does not exist on the network. It may have been deleted or
// referenced inaccurately.
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
func (e *ResourceNotFoundException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceNotFoundException) HasMessage() bool {
	return e.Message != nil
}

// The requested resource exists but is not in a status that can complete the
// operation.
type ResourceNotReadyException struct {
	Message *string
}

func (e *ResourceNotReadyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotReadyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotReadyException) ErrorCode() string             { return "ResourceNotReadyException" }
func (e *ResourceNotReadyException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ResourceNotReadyException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceNotReadyException) HasMessage() bool {
	return e.Message != nil
}

// The request or operation could not be performed because a service is throttling
// requests. The most common source of throttling errors is launching EC2 instances
// such that your service limit for EC2 instances is exceeded. Request a limit
// increase or delete unused resources if possible.
type ThrottlingException struct {
	Message *string
}

func (e *ThrottlingException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ThrottlingException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ThrottlingException) ErrorCode() string             { return "ThrottlingException" }
func (e *ThrottlingException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
