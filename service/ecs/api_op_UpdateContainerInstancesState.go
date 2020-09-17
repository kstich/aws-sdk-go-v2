// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the status of an Amazon ECS container instance. Once a container
// instance has reached an ACTIVE state, you can change the status of a container
// instance to DRAINING to manually remove an instance from a cluster, for example
// to perform system updates, update the Docker daemon, or scale down the cluster
// size. A container instance cannot be changed to DRAINING until it has reached an
// ACTIVE status. If the instance is in any other status, an error will be
// received. When you set a container instance to DRAINING, Amazon ECS prevents new
// tasks from being scheduled for placement on the container instance and
// replacement service tasks are started on other container instances in the
// cluster if the resources are available. Service tasks on the container instance
// that are in the PENDING state are stopped immediately. Service tasks on the
// container instance that are in the RUNNING state are stopped and replaced
// according to the service's deployment configuration parameters,
// minimumHealthyPercent and maximumPercent. You can change the deployment
// configuration of your service using UpdateService ().
//
//     * If
// minimumHealthyPercent is below 100%, the scheduler can ignore desiredCount
// temporarily during task replacement. For example, desiredCount is four tasks, a
// minimum of 50% allows the scheduler to stop two existing tasks before starting
// two new tasks. If the minimum is 100%, the service scheduler can't remove
// existing tasks until the replacement tasks are considered healthy. Tasks for
// services that do not use a load balancer are considered healthy if they are in
// the RUNNING state. Tasks for services that use a load balancer are considered
// healthy if they are in the RUNNING state and the container instance they are
// hosted on is reported as healthy by the load balancer.
//
//     * The maximumPercent
// parameter represents an upper limit on the number of running tasks during task
// replacement, which enables you to define the replacement batch size. For
// example, if desiredCount is four tasks, a maximum of 200% starts four new tasks
// before stopping the four tasks to be drained, provided that the cluster
// resources required to do this are available. If the maximum is 100%, then
// replacement tasks can't start until the draining tasks have stopped.
//
// Any
// PENDING or RUNNING tasks that do not belong to a service are not affected. You
// must wait for them to finish or stop them manually. A container instance has
// completed draining when it has no more RUNNING tasks. You can verify this using
// ListTasks (). When a container instance has been drained, you can set a
// container instance to ACTIVE status and once it has reached that status the
// Amazon ECS scheduler can begin scheduling tasks on the instance again.
func (c *Client) UpdateContainerInstancesState(ctx context.Context, params *UpdateContainerInstancesStateInput, optFns ...func(*Options)) (*UpdateContainerInstancesStateOutput, error) {
	stack := middleware.NewStack("UpdateContainerInstancesState", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateContainerInstancesStateMiddlewares(stack)
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
	addOpUpdateContainerInstancesStateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateContainerInstancesState(options.Region), middleware.Before)
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
			OperationName: "UpdateContainerInstancesState",
			Err:           err,
		}
	}
	out := result.(*UpdateContainerInstancesStateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateContainerInstancesStateInput struct {
	// The container instance state with which to update the container instance. The
	// only valid values for this action are ACTIVE and DRAINING. A container instance
	// can only be updated to DRAINING status once it has reached an ACTIVE state. If a
	// container instance is in REGISTERING, DEREGISTERING, or REGISTRATION_FAILED
	// state you can describe the container instance but will be unable to update the
	// container instance state.
	Status types.ContainerInstanceStatus
	// A list of container instance IDs or full ARN entries.
	ContainerInstances []*string
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the
	// container instance to update. If you do not specify a cluster, the default
	// cluster is assumed.
	Cluster *string
}

type UpdateContainerInstancesStateOutput struct {
	// The list of container instances.
	ContainerInstances []*types.ContainerInstance
	// Any failures associated with the call.
	Failures []*types.Failure

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateContainerInstancesStateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateContainerInstancesState{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateContainerInstancesState{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateContainerInstancesState(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "UpdateContainerInstancesState",
	}
}
