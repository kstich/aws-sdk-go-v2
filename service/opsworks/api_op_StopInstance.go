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

// Stops a specified instance. When you stop a standard instance, the data
// disappears and must be reinstalled when you restart the instance. You can stop
// an Amazon EBS-backed instance without losing data. For more information, see
// Starting, Stopping, and Rebooting Instances
// (https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-starting.html).
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) StopInstance(ctx context.Context, params *StopInstanceInput, optFns ...func(*Options)) (*StopInstanceOutput, error) {
	stack := middleware.NewStack("StopInstance", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStopInstanceMiddlewares(stack)
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
	addOpStopInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopInstance(options.Region), middleware.Before)
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
			OperationName: "StopInstance",
			Err:           err,
		}
	}
	out := result.(*StopInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopInstanceInput struct {
	// The instance ID.
	InstanceId *string
	// Specifies whether to force an instance to stop. If the instance's root device
	// type is ebs, or EBS-backed, adding the Force parameter to the StopInstances API
	// call disassociates the AWS OpsWorks Stacks instance from EC2, and forces
	// deletion of only the OpsWorks Stacks instance. You must also delete the
	// formerly-associated instance in EC2 after troubleshooting and replacing the AWS
	// OpsWorks Stacks instance with a new one.
	Force *bool
}

type StopInstanceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStopInstanceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStopInstance{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopInstance{}, middleware.After)
}

func newServiceMetadataMiddleware_opStopInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "StopInstance",
	}
}
