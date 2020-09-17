// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Requests the details of a job for a third party action. Used for partner actions
// only. When this API is called, AWS CodePipeline returns temporary credentials
// for the S3 bucket used to store artifacts for the pipeline, if the action
// requires access to that S3 bucket for input or output artifacts. This API also
// returns any secret values defined for the action.
func (c *Client) GetThirdPartyJobDetails(ctx context.Context, params *GetThirdPartyJobDetailsInput, optFns ...func(*Options)) (*GetThirdPartyJobDetailsOutput, error) {
	stack := middleware.NewStack("GetThirdPartyJobDetails", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetThirdPartyJobDetailsMiddlewares(stack)
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
	addOpGetThirdPartyJobDetailsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetThirdPartyJobDetails(options.Region), middleware.Before)
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
			OperationName: "GetThirdPartyJobDetails",
			Err:           err,
		}
	}
	out := result.(*GetThirdPartyJobDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a GetThirdPartyJobDetails action.
type GetThirdPartyJobDetailsInput struct {
	// The unique system-generated ID used for identifying the job.
	JobId *string
	// The clientToken portion of the clientId and clientToken pair used to verify that
	// the calling entity is allowed access to the job and its details.
	ClientToken *string
}

// Represents the output of a GetThirdPartyJobDetails action.
type GetThirdPartyJobDetailsOutput struct {
	// The details of the job, including any protected values defined for the job.
	JobDetails *types.ThirdPartyJobDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetThirdPartyJobDetailsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetThirdPartyJobDetails{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetThirdPartyJobDetails{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetThirdPartyJobDetails(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "GetThirdPartyJobDetails",
	}
}
