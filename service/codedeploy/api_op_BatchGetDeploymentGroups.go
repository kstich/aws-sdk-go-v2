// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about one or more deployment groups.
func (c *Client) BatchGetDeploymentGroups(ctx context.Context, params *BatchGetDeploymentGroupsInput, optFns ...func(*Options)) (*BatchGetDeploymentGroupsOutput, error) {
	stack := middleware.NewStack("BatchGetDeploymentGroups", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpBatchGetDeploymentGroupsMiddlewares(stack)
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
	addOpBatchGetDeploymentGroupsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetDeploymentGroups(options.Region), middleware.Before)
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
			OperationName: "BatchGetDeploymentGroups",
			Err:           err,
		}
	}
	out := result.(*BatchGetDeploymentGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a BatchGetDeploymentGroups operation.
type BatchGetDeploymentGroupsInput struct {
	// The names of the deployment groups.
	DeploymentGroupNames []*string
	// The name of an AWS CodeDeploy application associated with the applicable IAM
	// user or AWS account.
	ApplicationName *string
}

// Represents the output of a BatchGetDeploymentGroups operation.
type BatchGetDeploymentGroupsOutput struct {
	// Information about errors that might have occurred during the API call.
	ErrorMessage *string
	// Information about the deployment groups.
	DeploymentGroupsInfo []*types.DeploymentGroupInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpBatchGetDeploymentGroupsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpBatchGetDeploymentGroups{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchGetDeploymentGroups{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchGetDeploymentGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "BatchGetDeploymentGroups",
	}
}
