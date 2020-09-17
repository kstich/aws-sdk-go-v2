// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates fleet properties, including name and description, for a fleet. To update
// metadata, specify the fleet ID and the property values that you want to change.
// If successful, the fleet ID for the updated fleet is returned. Learn more
// Setting up GameLift Fleets
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html)
// Related operations
//
//     * CreateFleet ()
//
//     * ListFleets ()
//
//     * DeleteFleet
// ()
//
//     * DescribeFleetAttributes ()
//
//     * Update fleets:
//
//         *
// UpdateFleetAttributes ()
//
//         * UpdateFleetCapacity ()
//
//         *
// UpdateFleetPortSettings ()
//
//         * UpdateRuntimeConfiguration ()
//
//     *
// StartFleetActions () or StopFleetActions ()
func (c *Client) UpdateFleetAttributes(ctx context.Context, params *UpdateFleetAttributesInput, optFns ...func(*Options)) (*UpdateFleetAttributesOutput, error) {
	stack := middleware.NewStack("UpdateFleetAttributes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateFleetAttributesMiddlewares(stack)
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
	addOpUpdateFleetAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFleetAttributes(options.Region), middleware.Before)
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
			OperationName: "UpdateFleetAttributes",
			Err:           err,
		}
	}
	out := result.(*UpdateFleetAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type UpdateFleetAttributesInput struct {
	// Names of metric groups to include this fleet in. Amazon CloudWatch uses a fleet
	// metric group is to aggregate metrics from multiple fleets. Use an existing
	// metric group name to add this fleet to the group. Or use a new name to create a
	// new metric group. A fleet can only be included in one metric group at a time.
	MetricGroups []*string
	// A descriptive label that is associated with a fleet. Fleet names do not need to
	// be unique.
	Name *string
	// Game session protection policy to apply to all new instances created in this
	// fleet. Instances that already exist are not affected. You can set protection for
	// individual instances using UpdateGameSession ().
	//
	//     * NoProtection -- The game
	// session can be terminated during a scale-down event.
	//
	//     * FullProtection -- If
	// the game session is in an ACTIVE status, it cannot be terminated during a
	// scale-down event.
	NewGameSessionProtectionPolicy types.ProtectionPolicy
	// Human-readable description of a fleet.
	Description *string
	// A unique identifier for a fleet to update attribute metadata for. You can use
	// either the fleet ID or ARN value.
	FleetId *string
	// Policy that limits the number of game sessions an individual player can create
	// over a span of time.
	ResourceCreationLimitPolicy *types.ResourceCreationLimitPolicy
}

// Represents the returned data in response to a request action.
type UpdateFleetAttributesOutput struct {
	// A unique identifier for a fleet that was updated. Use either the fleet ID or ARN
	// value.
	FleetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateFleetAttributesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateFleetAttributes{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateFleetAttributes{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateFleetAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "UpdateFleetAttributes",
	}
}
