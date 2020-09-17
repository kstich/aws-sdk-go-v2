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

// Defines and adds a quota to the service quota template. To add a quota to the
// template, you must provide the ServiceCode, QuotaCode, AwsRegion, and
// DesiredValue. Once you add a quota to the template, use
// ListServiceQuotaIncreaseRequestsInTemplate () to see the list of quotas in the
// template.
func (c *Client) PutServiceQuotaIncreaseRequestIntoTemplate(ctx context.Context, params *PutServiceQuotaIncreaseRequestIntoTemplateInput, optFns ...func(*Options)) (*PutServiceQuotaIncreaseRequestIntoTemplateOutput, error) {
	stack := middleware.NewStack("PutServiceQuotaIncreaseRequestIntoTemplate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutServiceQuotaIncreaseRequestIntoTemplateMiddlewares(stack)
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
	addOpPutServiceQuotaIncreaseRequestIntoTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutServiceQuotaIncreaseRequestIntoTemplate(options.Region), middleware.Before)
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
			OperationName: "PutServiceQuotaIncreaseRequestIntoTemplate",
			Err:           err,
		}
	}
	out := result.(*PutServiceQuotaIncreaseRequestIntoTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutServiceQuotaIncreaseRequestIntoTemplateInput struct {
	// Specifies the new, increased value for the quota.
	DesiredValue *float64
	// Specifies the AWS Region for the quota.
	AwsRegion *string
	// Specifies the service that you want to use.
	ServiceCode *string
	// Specifies the service quota that you want to use.
	QuotaCode *string
}

type PutServiceQuotaIncreaseRequestIntoTemplateOutput struct {
	// A structure that contains information about one service quota increase request.
	ServiceQuotaIncreaseRequestInTemplate *types.ServiceQuotaIncreaseRequestInTemplate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutServiceQuotaIncreaseRequestIntoTemplateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutServiceQuotaIncreaseRequestIntoTemplate{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutServiceQuotaIncreaseRequestIntoTemplate{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutServiceQuotaIncreaseRequestIntoTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicequotas",
		OperationName: "PutServiceQuotaIncreaseRequestIntoTemplate",
	}
}
