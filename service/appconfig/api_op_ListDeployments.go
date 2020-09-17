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

// Lists the deployments for an environment.
func (c *Client) ListDeployments(ctx context.Context, params *ListDeploymentsInput, optFns ...func(*Options)) (*ListDeploymentsOutput, error) {
	stack := middleware.NewStack("ListDeployments", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListDeploymentsMiddlewares(stack)
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
	addOpListDeploymentsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListDeployments(options.Region), middleware.Before)
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
			OperationName: "ListDeployments",
			Err:           err,
		}
	}
	out := result.(*ListDeploymentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDeploymentsInput struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32
	// A token to start the list. Use this token to get the next set of results.
	NextToken *string
	// The application ID.
	ApplicationId *string
	// The environment ID.
	EnvironmentId *string
}

type ListDeploymentsOutput struct {
	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string
	// The elements from this collection.
	Items []*types.DeploymentSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListDeploymentsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListDeployments{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListDeployments{}, middleware.After)
}

func newServiceMetadataMiddleware_opListDeployments(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appconfig",
		OperationName: "ListDeployments",
	}
}
