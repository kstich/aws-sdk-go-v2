// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appconfig/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List deployment strategies.
func (c *Client) ListDeploymentStrategies(ctx context.Context, params *ListDeploymentStrategiesInput, optFns ...func(*Options)) (*ListDeploymentStrategiesOutput, error) {
	stack := middleware.NewStack("ListDeploymentStrategies", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListDeploymentStrategiesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListDeploymentStrategies(options.Region), middleware.Before)

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
			OperationName: "ListDeploymentStrategies",
			Err:           err,
		}
	}
	out := result.(*ListDeploymentStrategiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDeploymentStrategiesInput struct {
	// A token to start the list. Use this token to get the next set of results.
	NextToken *string
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32
}

type ListDeploymentStrategiesOutput struct {
	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string
	// The elements from this collection.
	Items []*types.DeploymentStrategy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListDeploymentStrategiesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListDeploymentStrategies{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListDeploymentStrategies{}, middleware.After)
}

func newServiceMetadataMiddleware_opListDeploymentStrategies(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appconfig",
		OperationName: "ListDeploymentStrategies",
	}
}
