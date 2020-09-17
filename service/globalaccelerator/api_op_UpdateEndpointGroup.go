// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Update an endpoint group. To see an AWS CLI example of updating an endpoint
// group, scroll down to Example.
func (c *Client) UpdateEndpointGroup(ctx context.Context, params *UpdateEndpointGroupInput, optFns ...func(*Options)) (*UpdateEndpointGroupOutput, error) {
	stack := middleware.NewStack("UpdateEndpointGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateEndpointGroupMiddlewares(stack)
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
	addOpUpdateEndpointGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEndpointGroup(options.Region), middleware.Before)
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
			OperationName: "UpdateEndpointGroup",
			Err:           err,
		}
	}
	out := result.(*UpdateEndpointGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEndpointGroupInput struct {
	// The list of endpoint objects.
	EndpointConfigurations []*types.EndpointConfiguration
	// The time—10 seconds or 30 seconds—between each health check for an endpoint. The
	// default value is 30.
	HealthCheckIntervalSeconds *int32
	// The number of consecutive health checks required to set the state of a healthy
	// endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default
	// value is 3.
	ThresholdCount *int32
	// The protocol that AWS Global Accelerator uses to check the health of endpoints
	// that are part of this endpoint group. The default value is TCP.
	HealthCheckProtocol types.HealthCheckProtocol
	// The Amazon Resource Name (ARN) of the endpoint group.
	EndpointGroupArn *string
	// The port that AWS Global Accelerator uses to check the health of endpoints that
	// are part of this endpoint group. The default port is the listener port that this
	// endpoint group is associated with. If the listener port is a list of ports,
	// Global Accelerator uses the first port in the list.
	HealthCheckPort *int32
	// The percentage of traffic to send to an AWS Region. Additional traffic is
	// distributed to other endpoint groups for this listener. Use this action to
	// increase (dial up) or decrease (dial down) traffic to a specific Region. The
	// percentage is applied to the traffic that would otherwise have been routed to
	// the Region based on optimal routing. The default value is 100.
	TrafficDialPercentage *float32
	// If the protocol is HTTP/S, then this specifies the path that is the destination
	// for health check targets. The default value is slash (/).
	HealthCheckPath *string
}

type UpdateEndpointGroupOutput struct {
	// The information about the endpoint group that was updated.
	EndpointGroup *types.EndpointGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateEndpointGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateEndpointGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateEndpointGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateEndpointGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "globalaccelerator",
		OperationName: "UpdateEndpointGroup",
	}
}
