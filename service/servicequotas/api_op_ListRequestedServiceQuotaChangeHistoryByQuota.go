// Code generated by smithy-go-codegen DO NOT EDIT.

package servicequotas

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Requests a list of the changes to specific service quotas. This command provides
// additional granularity over the ListRequestedServiceQuotaChangeHistory command.
// Once a quota change request has reached CASE_CLOSED, APPROVED, or DENIED, the
// history has been kept for 90 days.
func (c *Client) ListRequestedServiceQuotaChangeHistoryByQuota(ctx context.Context, params *ListRequestedServiceQuotaChangeHistoryByQuotaInput, optFns ...func(*Options)) (*ListRequestedServiceQuotaChangeHistoryByQuotaOutput, error) {
	stack := middleware.NewStack("ListRequestedServiceQuotaChangeHistoryByQuota", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListRequestedServiceQuotaChangeHistoryByQuotaMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpListRequestedServiceQuotaChangeHistoryByQuotaValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListRequestedServiceQuotaChangeHistoryByQuota(options.Region), middleware.Before)
	addResponseErrorWrapper(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "ListRequestedServiceQuotaChangeHistoryByQuota",
			Err:           err,
		}
	}
	out := result.(*ListRequestedServiceQuotaChangeHistoryByQuotaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRequestedServiceQuotaChangeHistoryByQuotaInput struct {
	// Specifies the service that you want to use.
	ServiceCode *string
	// Specifies the status value of the quota increase request.
	Status types.RequestStatus
	// Specifies the service quota that you want to use
	QuotaCode *string
	// (Optional) Limits the number of results that you want to include in the
	// response. If you don't include this parameter, the response defaults to a value
	// that's specific to the operation. If additional items exist beyond the specified
	// maximum, the NextToken element is present and has a value (isn't null). Include
	// that value as the NextToken request parameter in the call to the operation to
	// get the next part of the results. You should check NextToken after every
	// operation to ensure that you receive all of the results.
	MaxResults *int32
	// (Optional) Use this parameter in a request if you receive a NextToken response
	// in a previous request that indicates that there's more output available. In a
	// subsequent call, set it to the value of the previous call's NextToken response
	// to indicate where the output should continue from.
	NextToken *string
}

type ListRequestedServiceQuotaChangeHistoryByQuotaOutput struct {
	// If present in the response, this value indicates there's more output available
	// that what's included in the current response. This can occur even when the
	// response includes no values at all, such as when you ask for a filtered view of
	// a very long list. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to continue processing and get the next part of
	// the output. You should repeat this until the NextToken response element comes
	// back empty (as null).
	NextToken *string
	// Returns a list of service quota requests.
	RequestedQuotas []*types.RequestedServiceQuotaChange

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListRequestedServiceQuotaChangeHistoryByQuotaMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListRequestedServiceQuotaChangeHistoryByQuota{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListRequestedServiceQuotaChangeHistoryByQuota{}, middleware.After)
}

func newServiceMetadataMiddleware_opListRequestedServiceQuotaChangeHistoryByQuota(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicequotas",
		OperationName: "ListRequestedServiceQuotaChangeHistoryByQuota",
	}
}
