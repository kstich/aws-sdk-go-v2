// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates termination protection for the specified stack. If a user attempts to
// delete a stack with termination protection enabled, the operation fails and the
// stack remains unchanged. For more information, see Protecting a Stack From Being
// Deleted
// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-protect-stacks.html)
// in the AWS CloudFormation User Guide. For nested stacks
// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-nested-stacks.html),
// termination protection is set on the root stack and cannot be changed directly
// on the nested stack.
func (c *Client) UpdateTerminationProtection(ctx context.Context, params *UpdateTerminationProtectionInput, optFns ...func(*Options)) (*UpdateTerminationProtectionOutput, error) {
	stack := middleware.NewStack("UpdateTerminationProtection", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpUpdateTerminationProtectionMiddlewares(stack)
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
	addOpUpdateTerminationProtectionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTerminationProtection(options.Region), middleware.Before)
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
			OperationName: "UpdateTerminationProtection",
			Err:           err,
		}
	}
	out := result.(*UpdateTerminationProtectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTerminationProtectionInput struct {
	// Whether to enable termination protection on the specified stack.
	EnableTerminationProtection *bool
	// The name or unique ID of the stack for which you want to set termination
	// protection.
	StackName *string
}

type UpdateTerminationProtectionOutput struct {
	// The unique ID of the stack.
	StackId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpUpdateTerminationProtectionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpUpdateTerminationProtection{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateTerminationProtection{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateTerminationProtection(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "UpdateTerminationProtection",
	}
}
