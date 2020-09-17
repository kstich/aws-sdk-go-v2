// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about all the endpoints that are associated with a
// specific user ID.
func (c *Client) GetUserEndpoints(ctx context.Context, params *GetUserEndpointsInput, optFns ...func(*Options)) (*GetUserEndpointsOutput, error) {
	stack := middleware.NewStack("GetUserEndpoints", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetUserEndpointsMiddlewares(stack)
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
	addOpGetUserEndpointsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetUserEndpoints(options.Region), middleware.Before)
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
			OperationName: "GetUserEndpoints",
			Err:           err,
		}
	}
	out := result.(*GetUserEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetUserEndpointsInput struct {
	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	ApplicationId *string
	// The unique identifier for the user.
	UserId *string
}

type GetUserEndpointsOutput struct {
	// Provides information about all the endpoints that are associated with a user ID.
	EndpointsResponse *types.EndpointsResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetUserEndpointsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetUserEndpoints{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetUserEndpoints{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetUserEndpoints(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "GetUserEndpoints",
	}
}
