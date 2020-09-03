// Code generated by smithy-go-codegen DO NOT EDIT.

package groundstation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a mission profile.
func (c *Client) DeleteMissionProfile(ctx context.Context, params *DeleteMissionProfileInput, optFns ...func(*Options)) (*DeleteMissionProfileOutput, error) {
	stack := middleware.NewStack("DeleteMissionProfile", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteMissionProfileMiddlewares(stack)
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
	addOpDeleteMissionProfileValidationMiddleware(stack)

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
			OperationName: "DeleteMissionProfile",
			Err:           err,
		}
	}
	out := result.(*DeleteMissionProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DeleteMissionProfileInput struct {
	// UUID of a mission profile.
	MissionProfileId *string
}

//
type DeleteMissionProfileOutput struct {
	// UUID of a mission profile.
	MissionProfileId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteMissionProfileMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteMissionProfile{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteMissionProfile{}, middleware.After)
}