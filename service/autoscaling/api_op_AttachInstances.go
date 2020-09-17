// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Attaches one or more EC2 instances to the specified Auto Scaling group. When you
// attach instances, Amazon EC2 Auto Scaling increases the desired capacity of the
// group by the number of instances being attached. If the number of instances
// being attached plus the desired capacity of the group exceeds the maximum size
// of the group, the operation fails. If there is a Classic Load Balancer attached
// to your Auto Scaling group, the instances are also registered with the load
// balancer. If there are target groups attached to your Auto Scaling group, the
// instances are also registered with the target groups. For more information, see
// Attach EC2 Instances to Your Auto Scaling Group
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/attach-instance-asg.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) AttachInstances(ctx context.Context, params *AttachInstancesInput, optFns ...func(*Options)) (*AttachInstancesOutput, error) {
	stack := middleware.NewStack("AttachInstances", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpAttachInstancesMiddlewares(stack)
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
	addOpAttachInstancesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAttachInstances(options.Region), middleware.Before)
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
			OperationName: "AttachInstances",
			Err:           err,
		}
	}
	out := result.(*AttachInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AttachInstancesInput struct {
	// The name of the Auto Scaling group.
	AutoScalingGroupName *string
	// The IDs of the instances. You can specify up to 20 instances.
	InstanceIds []*string
}

type AttachInstancesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpAttachInstancesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpAttachInstances{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpAttachInstances{}, middleware.After)
}

func newServiceMetadataMiddleware_opAttachInstances(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "AttachInstances",
	}
}
