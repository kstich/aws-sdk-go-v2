// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Assign a registered instance to a layer.
//
//     * You can assign registered
// on-premises instances to any layer type.
//
//     * You can assign registered Amazon
// EC2 instances only to custom layers.
//
//     * You cannot use this action with
// instances that were created with AWS OpsWorks Stacks.
//
// Required Permissions: To
// use this action, an AWS Identity and Access Management (IAM) user must have a
// Manage permissions level for the stack or an attached policy that explicitly
// grants permissions. For more information on user permissions, see Managing User
// Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) AssignInstance(ctx context.Context, params *AssignInstanceInput, optFns ...func(*Options)) (*AssignInstanceOutput, error) {
	stack := middleware.NewStack("AssignInstance", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAssignInstanceMiddlewares(stack)
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
	addOpAssignInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssignInstance(options.Region), middleware.Before)

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
			OperationName: "AssignInstance",
			Err:           err,
		}
	}
	out := result.(*AssignInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssignInstanceInput struct {
	// The instance ID.
	InstanceId *string
	// The layer ID, which must correspond to a custom layer. You cannot assign a
	// registered instance to a built-in layer.
	LayerIds []*string
}

type AssignInstanceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAssignInstanceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAssignInstance{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssignInstance{}, middleware.After)
}

func newServiceMetadataMiddleware_opAssignInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "AssignInstance",
	}
}
