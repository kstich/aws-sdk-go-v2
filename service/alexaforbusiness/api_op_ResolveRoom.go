// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Determines the details for the room from which a skill request was invoked. This
// operation is used by skill developers.
func (c *Client) ResolveRoom(ctx context.Context, params *ResolveRoomInput, optFns ...func(*Options)) (*ResolveRoomOutput, error) {
	stack := middleware.NewStack("ResolveRoom", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpResolveRoomMiddlewares(stack)
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
	addOpResolveRoomValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opResolveRoom(options.Region), middleware.Before)
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
			OperationName: "ResolveRoom",
			Err:           err,
		}
	}
	out := result.(*ResolveRoomOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResolveRoomInput struct {
	// The ARN of the skill that was requested. Required.
	SkillId *string
	// The ARN of the user. Required.
	UserId *string
}

type ResolveRoomOutput struct {
	// The name of the room from which the skill request was invoked.
	RoomName *string
	// The ARN of the room from which the skill request was invoked.
	RoomArn *string
	// Response to get the room profile request. Required.
	RoomSkillParameters []*types.RoomSkillParameter

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpResolveRoomMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpResolveRoom{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpResolveRoom{}, middleware.After)
}

func newServiceMetadataMiddleware_opResolveRoom(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "ResolveRoom",
	}
}
