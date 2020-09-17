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

// Returns a list of the quota increase requests in the template.
func (c *Client) ListServiceQuotaIncreaseRequestsInTemplate(ctx context.Context, params *ListServiceQuotaIncreaseRequestsInTemplateInput, optFns ...func(*Options)) (*ListServiceQuotaIncreaseRequestsInTemplateOutput, error) {
	stack := middleware.NewStack("ListServiceQuotaIncreaseRequestsInTemplate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListServiceQuotaIncreaseRequestsInTemplateMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListServiceQuotaIncreaseRequestsInTemplate(options.Region), middleware.Before)
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
			OperationName: "ListServiceQuotaIncreaseRequestsInTemplate",
			Err:           err,
		}
	}
	out := result.(*ListServiceQuotaIncreaseRequestsInTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListServiceQuotaIncreaseRequestsInTemplateInput struct {
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
	// Specifies the AWS Region for the quota that you want to use.
	AwsRegion *string
	// The identifier for a service. When performing an operation, use the ServiceCode
	// to specify a particular service.
	ServiceCode *string
}

type ListServiceQuotaIncreaseRequestsInTemplateOutput struct {
	// If present in the response, this value indicates there's more output available
	// that what's included in the current response. This can occur even when the
	// response includes no values at all, such as when you ask for a filtered view of
	// a very long list. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to continue processing and get the next part of
	// the output. You should repeat this until the NextToken response element comes
	// back empty (as null).
	NextToken *string
	// Returns the list of values of the quota increase request in the template.
	ServiceQuotaIncreaseRequestInTemplateList []*types.ServiceQuotaIncreaseRequestInTemplate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListServiceQuotaIncreaseRequestsInTemplateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListServiceQuotaIncreaseRequestsInTemplate{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListServiceQuotaIncreaseRequestsInTemplate{}, middleware.After)
}

func newServiceMetadataMiddleware_opListServiceQuotaIncreaseRequestsInTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicequotas",
		OperationName: "ListServiceQuotaIncreaseRequestsInTemplate",
	}
}
