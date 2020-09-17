// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Deletes the lifecycle policy associated with the specified repository.
func (c *Client) DeleteLifecyclePolicy(ctx context.Context, params *DeleteLifecyclePolicyInput, optFns ...func(*Options)) (*DeleteLifecyclePolicyOutput, error) {
	stack := middleware.NewStack("DeleteLifecyclePolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteLifecyclePolicyMiddlewares(stack)
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
	addOpDeleteLifecyclePolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteLifecyclePolicy(options.Region), middleware.Before)
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
			OperationName: "DeleteLifecyclePolicy",
			Err:           err,
		}
	}
	out := result.(*DeleteLifecyclePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteLifecyclePolicyInput struct {
	// The AWS account ID associated with the registry that contains the repository. If
	// you do not specify a registry, the default registry is assumed.
	RegistryId *string
	// The name of the repository.
	RepositoryName *string
}

type DeleteLifecyclePolicyOutput struct {
	// The JSON lifecycle policy text.
	LifecyclePolicyText *string
	// The repository name associated with the request.
	RepositoryName *string
	// The time stamp of the last time that the lifecycle policy was run.
	LastEvaluatedAt *time.Time
	// The registry ID associated with the request.
	RegistryId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteLifecyclePolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteLifecyclePolicy{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteLifecyclePolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteLifecyclePolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr",
		OperationName: "DeleteLifecyclePolicy",
	}
}
