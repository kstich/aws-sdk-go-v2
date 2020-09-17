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

// Updates game session properties. This includes the session name, maximum player
// count, protection policy, which controls whether or not an active game session
// can be terminated during a scale-down event, and the player session creation
// policy, which controls whether or not new players can join the session. To
// update a game session, specify the game session ID and the values you want to
// change. If successful, an updated GameSession () object is returned.
//
//     *
// CreateGameSession ()
//
//     * DescribeGameSessions ()
//
//     *
// DescribeGameSessionDetails ()
//
//     * SearchGameSessions ()
//
//     *
// UpdateGameSession ()
//
//     * GetGameSessionLogUrl ()
//
//     * Game session
// placements
//
//         * StartGameSessionPlacement ()
//
//         *
// DescribeGameSessionPlacement ()
//
//         * StopGameSessionPlacement ()
func (c *Client) UpdateGameSession(ctx context.Context, params *UpdateGameSessionInput, optFns ...func(*Options)) (*UpdateGameSessionOutput, error) {
	stack := middleware.NewStack("UpdateGameSession", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateGameSessionMiddlewares(stack)
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
	addOpUpdateGameSessionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGameSession(options.Region), middleware.Before)
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
			OperationName: "UpdateGameSession",
			Err:           err,
		}
	}
	out := result.(*UpdateGameSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type UpdateGameSessionInput struct {
	// Game session protection policy to apply to this game session only.
	//
	//     *
	// NoProtection -- The game session can be terminated during a scale-down event.
	//
	//
	// * FullProtection -- If the game session is in an ACTIVE status, it cannot be
	// terminated during a scale-down event.
	ProtectionPolicy types.ProtectionPolicy
	// Policy determining whether or not the game session accepts new players.
	PlayerSessionCreationPolicy types.PlayerSessionCreationPolicy
	// A descriptive label that is associated with a game session. Session names do not
	// need to be unique.
	Name *string
	// The maximum number of players that can be connected simultaneously to the game
	// session.
	MaximumPlayerSessionCount *int32
	// A unique identifier for the game session to update.
	GameSessionId *string
}

// Represents the returned data in response to a request action.
type UpdateGameSessionOutput struct {
	// The updated game session metadata.
	GameSession *types.GameSession

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateGameSessionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateGameSession{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateGameSession{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateGameSession(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "UpdateGameSession",
	}
}
