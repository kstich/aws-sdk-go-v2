// Code generated by smithy-go-codegen DO NOT EDIT.

package cloud9

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an environment member from an AWS Cloud9 development environment.
func (c *Client) DeleteEnvironmentMembership(ctx context.Context, params *DeleteEnvironmentMembershipInput, optFns ...func(*Options)) (*DeleteEnvironmentMembershipOutput, error) {
	stack := middleware.NewStack("DeleteEnvironmentMembership", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteEnvironmentMembershipMiddlewares(stack)
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
	addOpDeleteEnvironmentMembershipValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteEnvironmentMembership(options.Region), middleware.Before)
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
			OperationName: "DeleteEnvironmentMembership",
			Err:           err,
		}
	}
	out := result.(*DeleteEnvironmentMembershipOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteEnvironmentMembershipInput struct {
	// The Amazon Resource Name (ARN) of the environment member to delete from the
	// environment.
	UserArn *string
	// The ID of the environment to delete the environment member from.
	EnvironmentId *string
}

type DeleteEnvironmentMembershipOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteEnvironmentMembershipMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteEnvironmentMembership{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteEnvironmentMembership{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteEnvironmentMembership(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloud9",
		OperationName: "DeleteEnvironmentMembership",
	}
}
