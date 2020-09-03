// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Resets a group's deployments.
func (c *Client) ResetDeployments(ctx context.Context, params *ResetDeploymentsInput, optFns ...func(*Options)) (*ResetDeploymentsOutput, error) {
	stack := middleware.NewStack("ResetDeployments", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpResetDeploymentsMiddlewares(stack)
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
	addOpResetDeploymentsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opResetDeployments(options.Region), middleware.Before)

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
			OperationName: "ResetDeployments",
			Err:           err,
		}
	}
	out := result.(*ResetDeploymentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Information needed to reset deployments.
type ResetDeploymentsInput struct {
	// A client token used to correlate requests and responses.
	AmznClientToken *string
	// The ID of the Greengrass group.
	GroupId *string
	// If true, performs a best-effort only core reset.
	Force *bool
}

type ResetDeploymentsOutput struct {
	// The ID of the deployment.
	DeploymentId *string
	// The ARN of the deployment.
	DeploymentArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpResetDeploymentsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpResetDeployments{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpResetDeployments{}, middleware.After)
}

func newServiceMetadataMiddleware_opResetDeployments(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "ResetDeployments",
	}
}
