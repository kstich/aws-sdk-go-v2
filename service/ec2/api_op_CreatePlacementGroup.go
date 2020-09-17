// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a placement group in which to launch instances. The strategy of the
// placement group determines how the instances are organized within the group. A
// cluster placement group is a logical grouping of instances within a single
// Availability Zone that benefit from low network latency, high network
// throughput. A spread placement group places instances on distinct hardware. A
// partition placement group places groups of instances in different partitions,
// where instances in one partition do not share the same hardware with instances
// in another partition. For more information, see Placement groups
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html) in
// the Amazon Elastic Compute Cloud User Guide.
func (c *Client) CreatePlacementGroup(ctx context.Context, params *CreatePlacementGroupInput, optFns ...func(*Options)) (*CreatePlacementGroupOutput, error) {
	stack := middleware.NewStack("CreatePlacementGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpCreatePlacementGroupMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePlacementGroup(options.Region), middleware.Before)
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
			OperationName: "CreatePlacementGroup",
			Err:           err,
		}
	}
	out := result.(*CreatePlacementGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePlacementGroupInput struct {
	// A name for the placement group. Must be unique within the scope of your account
	// for the Region. Constraints: Up to 255 ASCII characters
	GroupName *string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The tags to apply to the new placement group.
	TagSpecifications []*types.TagSpecification
	// The placement strategy.
	Strategy types.PlacementStrategy
	// The number of partitions. Valid only when Strategy is set to partition.
	PartitionCount *int32
}

type CreatePlacementGroupOutput struct {
	// Describes a placement group.
	PlacementGroup *types.PlacementGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpCreatePlacementGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpCreatePlacementGroup{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpCreatePlacementGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreatePlacementGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreatePlacementGroup",
	}
}
