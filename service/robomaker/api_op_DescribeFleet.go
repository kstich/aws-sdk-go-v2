// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes a fleet.
func (c *Client) DescribeFleet(ctx context.Context, params *DescribeFleetInput, optFns ...func(*Options)) (*DescribeFleetOutput, error) {
	stack := middleware.NewStack("DescribeFleet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeFleetMiddlewares(stack)
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
	addOpDescribeFleetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFleet(options.Region), middleware.Before)
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
			OperationName: "DescribeFleet",
			Err:           err,
		}
	}
	out := result.(*DescribeFleetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFleetInput struct {
	// The Amazon Resource Name (ARN) of the fleet.
	Fleet *string
}

type DescribeFleetOutput struct {
	// The Amazon Resource Name (ARN) of the fleet.
	Arn *string
	// The status of the last deployment.
	LastDeploymentStatus types.DeploymentStatus
	// The time of the last deployment.
	LastDeploymentTime *time.Time
	// The list of all tags added to the specified fleet.
	Tags map[string]*string
	// A list of robots.
	Robots []*types.Robot
	// The Amazon Resource Name (ARN) of the last deployment job.
	LastDeploymentJob *string
	// The time, in milliseconds since the epoch, when the fleet was created.
	CreatedAt *time.Time
	// The name of the fleet.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeFleetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeFleet{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeFleet{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeFleet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "DescribeFleet",
	}
}
