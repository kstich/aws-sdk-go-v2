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

// Deletes a job for a branch of an Amplify app.
func (c *Client) DeleteJob(ctx context.Context, params *DeleteJobInput, optFns ...func(*Options)) (*DeleteJobOutput, error) {
	stack := middleware.NewStack("DeleteJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteJobMiddlewares(stack)
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
	addOpDeleteJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteJob(options.Region), middleware.Before)

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
			OperationName: "DeleteJob",
			Err:           err,
		}
	}
	out := result.(*DeleteJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request structure for the delete job request.
type DeleteJobInput struct {
	// The unique ID for the job.
	JobId *string
	// The name for the branch, for the job.
	BranchName *string
	// The unique ID for an Amplify app.
	AppId *string
}

// The result structure for the delete job request.
type DeleteJobOutput struct {
	// Describes the summary for an execution job for an Amplify app.
	JobSummary *types.JobSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteJob{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplify",
		OperationName: "DeleteJob",
	}
}
