// Code generated by smithy-go-codegen DO NOT EDIT.

package amplify

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new backend environment for an Amplify app.
func (c *Client) CreateBackendEnvironment(ctx context.Context, params *CreateBackendEnvironmentInput, optFns ...func(*Options)) (*CreateBackendEnvironmentOutput, error) {
	stack := middleware.NewStack("CreateBackendEnvironment", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateBackendEnvironmentMiddlewares(stack)
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
	addOpCreateBackendEnvironmentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBackendEnvironment(options.Region), middleware.Before)
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
			OperationName: "CreateBackendEnvironment",
			Err:           err,
		}
	}
	out := result.(*CreateBackendEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request structure for the backend environment create request.
type CreateBackendEnvironmentInput struct {
	// The AWS CloudFormation stack name of a backend environment.
	StackName *string
	// The unique ID for an Amplify app.
	AppId *string
	// The name of deployment artifacts.
	DeploymentArtifacts *string
	// The name for the backend environment.
	EnvironmentName *string
}

// The result structure for the create backend environment request.
type CreateBackendEnvironmentOutput struct {
	// Describes the backend environment for an Amplify app.
	BackendEnvironment *types.BackendEnvironment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateBackendEnvironmentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateBackendEnvironment{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateBackendEnvironment{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateBackendEnvironment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplify",
		OperationName: "CreateBackendEnvironment",
	}
}
