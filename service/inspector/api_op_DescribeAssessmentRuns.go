// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the assessment runs that are specified by the ARNs of the assessment
// runs.
func (c *Client) DescribeAssessmentRuns(ctx context.Context, params *DescribeAssessmentRunsInput, optFns ...func(*Options)) (*DescribeAssessmentRunsOutput, error) {
	stack := middleware.NewStack("DescribeAssessmentRuns", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeAssessmentRunsMiddlewares(stack)
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
	addOpDescribeAssessmentRunsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAssessmentRuns(options.Region), middleware.Before)
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
			OperationName: "DescribeAssessmentRuns",
			Err:           err,
		}
	}
	out := result.(*DescribeAssessmentRunsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAssessmentRunsInput struct {
	// The ARN that specifies the assessment run that you want to describe.
	AssessmentRunArns []*string
}

type DescribeAssessmentRunsOutput struct {
	// Information about the assessment run.
	AssessmentRuns []*types.AssessmentRun
	// Assessment run details that cannot be described. An error code is provided for
	// each failed item.
	FailedItems map[string]*types.FailedItemDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeAssessmentRunsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAssessmentRuns{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAssessmentRuns{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAssessmentRuns(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "DescribeAssessmentRuns",
	}
}
