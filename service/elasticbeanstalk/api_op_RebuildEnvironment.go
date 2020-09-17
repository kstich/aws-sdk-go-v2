// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes and recreates all of the AWS resources (for example: the Auto Scaling
// group, load balancer, etc.) for a specified environment and forces a restart.
func (c *Client) RebuildEnvironment(ctx context.Context, params *RebuildEnvironmentInput, optFns ...func(*Options)) (*RebuildEnvironmentOutput, error) {
	stack := middleware.NewStack("RebuildEnvironment", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpRebuildEnvironmentMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opRebuildEnvironment(options.Region), middleware.Before)
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
			OperationName: "RebuildEnvironment",
			Err:           err,
		}
	}
	out := result.(*RebuildEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type RebuildEnvironmentInput struct {
	// The name of the environment to rebuild. Condition: You must specify either this
	// or an EnvironmentId, or both. If you do not specify either, AWS Elastic
	// Beanstalk returns MissingRequiredParameter error.
	EnvironmentName *string
	// The ID of the environment to rebuild. Condition: You must specify either this or
	// an EnvironmentName, or both. If you do not specify either, AWS Elastic Beanstalk
	// returns MissingRequiredParameter error.
	EnvironmentId *string
}

type RebuildEnvironmentOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpRebuildEnvironmentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpRebuildEnvironment{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpRebuildEnvironment{}, middleware.After)
}

func newServiceMetadataMiddleware_opRebuildEnvironment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "RebuildEnvironment",
	}
}
