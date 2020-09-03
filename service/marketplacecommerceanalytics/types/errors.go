// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/ptr"
)

// This exception is thrown when an internal service error occurs.
type MarketplaceCommerceAnalyticsException struct {
	Message *string
}

func (e *MarketplaceCommerceAnalyticsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MarketplaceCommerceAnalyticsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MarketplaceCommerceAnalyticsException) ErrorCode() string {
	return "MarketplaceCommerceAnalyticsException"
}
func (e *MarketplaceCommerceAnalyticsException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultServer
}
func (e *MarketplaceCommerceAnalyticsException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *MarketplaceCommerceAnalyticsException) HasMessage() bool {
	return e.Message != nil
}
