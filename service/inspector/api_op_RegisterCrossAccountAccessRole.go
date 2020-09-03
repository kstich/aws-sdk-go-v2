// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Registers the IAM role that grants Amazon Inspector access to AWS Services
// needed to perform security assessments.
func (c *Client) RegisterCrossAccountAccessRole(ctx context.Context, params *RegisterCrossAccountAccessRoleInput, optFns ...func(*Options)) (*RegisterCrossAccountAccessRoleOutput, error) {
	stack := middleware.NewStack("RegisterCrossAccountAccessRole", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRegisterCrossAccountAccessRoleMiddlewares(stack)
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
	addOpRegisterCrossAccountAccessRoleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterCrossAccountAccessRole(options.Region), middleware.Before)

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
			OperationName: "RegisterCrossAccountAccessRole",
			Err:           err,
		}
	}
	out := result.(*RegisterCrossAccountAccessRoleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterCrossAccountAccessRoleInput struct {
	// The ARN of the IAM role that grants Amazon Inspector access to AWS Services
	// needed to perform security assessments.
	RoleArn *string
}

type RegisterCrossAccountAccessRoleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRegisterCrossAccountAccessRoleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterCrossAccountAccessRole{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterCrossAccountAccessRole{}, middleware.After)
}

func newServiceMetadataMiddleware_opRegisterCrossAccountAccessRole(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "RegisterCrossAccountAccessRole",
	}
}
