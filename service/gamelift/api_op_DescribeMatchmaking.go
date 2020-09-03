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

// Retrieves one or more matchmaking tickets. Use this operation to retrieve ticket
// information, including status and--once a successful match is made--acquire
// connection information for the resulting new game session. You can use this
// operation to track the progress of matchmaking requests (through polling) as an
// alternative to using event notifications. See more details on tracking
// matchmaking requests through polling or notifications in StartMatchmaking (). To
// request matchmaking tickets, provide a list of up to 10 ticket IDs. If the
// request is successful, a ticket object is returned for each requested ID that
// currently exists. Learn more  Add FlexMatch to a Game Client
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/match-client.html)
// Set Up FlexMatch Event Notification
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/match-notification.html)
// Related operations  <ul> <li> <p> <a>StartMatchmaking</a> </p> </li> <li> <p>
// <a>DescribeMatchmaking</a> </p> </li> <li> <p> <a>StopMatchmaking</a> </p> </li>
// <li> <p> <a>AcceptMatch</a> </p> </li> <li> <p> <a>StartMatchBackfill</a> </p>
// </li> </ul>
func (c *Client) DescribeMatchmaking(ctx context.Context, params *DescribeMatchmakingInput, optFns ...func(*Options)) (*DescribeMatchmakingOutput, error) {
	stack := middleware.NewStack("DescribeMatchmaking", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeMatchmakingMiddlewares(stack)
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
	addOpDescribeMatchmakingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMatchmaking(options.Region), middleware.Before)

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
			OperationName: "DescribeMatchmaking",
			Err:           err,
		}
	}
	out := result.(*DescribeMatchmakingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type DescribeMatchmakingInput struct {
	// A unique identifier for a matchmaking ticket. You can include up to 10 ID
	// values.
	TicketIds []*string
}

// Represents the returned data in response to a request action.
type DescribeMatchmakingOutput struct {
	// A collection of existing matchmaking ticket objects matching the request.
	TicketList []*types.MatchmakingTicket

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeMatchmakingMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMatchmaking{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMatchmaking{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeMatchmaking(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeMatchmaking",
	}
}
