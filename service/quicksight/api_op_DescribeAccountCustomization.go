// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the customizations associated with your AWS account.
func (c *Client) DescribeAccountCustomization(ctx context.Context, params *DescribeAccountCustomizationInput, optFns ...func(*Options)) (*DescribeAccountCustomizationOutput, error) {
	stack := middleware.NewStack("DescribeAccountCustomization", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeAccountCustomizationMiddlewares(stack)
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
	addOpDescribeAccountCustomizationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAccountCustomization(options.Region), middleware.Before)
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
			OperationName: "DescribeAccountCustomization",
			Err:           err,
		}
	}
	out := result.(*DescribeAccountCustomizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAccountCustomizationInput struct {
	// The ID for the AWS account that you want to describe QuickSight customizations
	// for.
	AwsAccountId *string
	// The status of the creation of the customization. This is an asynchronous
	// process. A status of CREATED means that your customization is ready to use.
	Resolved *bool
	// The namespace associated with the customization that you're describing.
	Namespace *string
}

type DescribeAccountCustomizationOutput struct {
	// The AWS request ID for this operation.
	RequestId *string
	// The namespace associated with the customization that you're describing.
	Namespace *string
	// The ID for the AWS account that you want to describe QuickSight customizations
	// for.
	AwsAccountId *string
	// The customizations associated with QuickSight.
	AccountCustomization *types.AccountCustomization

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeAccountCustomizationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAccountCustomization{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAccountCustomization{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAccountCustomization(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DescribeAccountCustomization",
	}
}
