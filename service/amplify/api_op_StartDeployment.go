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

// Starts a deployment for a manually deployed app. Manually deployed apps are not
// connected to a repository.
func (c *Client) StartDeployment(ctx context.Context, params *StartDeploymentInput, optFns ...func(*Options)) (*StartDeploymentOutput, error) {
	stack := middleware.NewStack("StartDeployment", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpStartDeploymentMiddlewares(stack)
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
	addOpStartDeploymentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartDeployment(options.Region), middleware.Before)

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
			OperationName: "StartDeployment",
			Err:           err,
		}
	}
	out := result.(*StartDeploymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request structure for the start a deployment request.
type StartDeploymentInput struct {
	// The job ID for this deployment, generated by the create deployment request.
	JobId *string
	// The source URL for this deployment, used when calling start deployment without
	// create deployment. The source URL can be any HTTP GET URL that is publicly
	// accessible and downloads a single .zip file.
	SourceUrl *string
	// The unique ID for an Amplify app.
	AppId *string
	// The name for the branch, for the job.
	BranchName *string
}

// The result structure for the start a deployment request.
type StartDeploymentOutput struct {
	// The summary for the job.
	JobSummary *types.JobSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpStartDeploymentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpStartDeployment{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpStartDeployment{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartDeployment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplify",
		OperationName: "StartDeployment",
	}
}
