// Code generated by smithy-go-codegen DO NOT EDIT.

package synthetics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Use this operation to run a canary that has already been created. The frequency
// of the canary runs is determined by the value of the canary's Schedule. To see a
// canary's schedule, use GetCanary
// (https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_GetCanary.html).
func (c *Client) StartCanary(ctx context.Context, params *StartCanaryInput, optFns ...func(*Options)) (*StartCanaryOutput, error) {
	stack := middleware.NewStack("StartCanary", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpStartCanaryMiddlewares(stack)
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
	addOpStartCanaryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartCanary(options.Region), middleware.Before)
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
			OperationName: "StartCanary",
			Err:           err,
		}
	}
	out := result.(*StartCanaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartCanaryInput struct {
	// The name of the canary that you want to run. To find canary names, use
	// DescribeCanaries
	// (https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_DescribeCanaries.html).
	Name *string
}

type StartCanaryOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpStartCanaryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpStartCanary{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpStartCanary{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartCanary(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "synthetics",
		OperationName: "StartCanary",
	}
}
