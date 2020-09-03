// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the entire tag set from the specified Amazon S3 Batch Operations job. To
// use this operation, you must have permission to perform the s3:DeleteJobTagging
// action. For more information, see Using Job Tags
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/batch-ops-managing-jobs.html#batch-ops-job-tags)
// in the Amazon Simple Storage Service Developer Guide. Related actions include:
//
//
// * CreateJob ()
//
//     * GetJobTagging ()
//
//     * PutJobTagging ()
func (c *Client) DeleteJobTagging(ctx context.Context, params *DeleteJobTaggingInput, optFns ...func(*Options)) (*DeleteJobTaggingOutput, error) {
	stack := middleware.NewStack("DeleteJobTagging", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpDeleteJobTaggingMiddlewares(stack)
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
	addOpDeleteJobTaggingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteJobTagging(options.Region), middleware.Before)

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
			OperationName: "DeleteJobTagging",
			Err:           err,
		}
	}
	out := result.(*DeleteJobTaggingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteJobTaggingInput struct {
	// The ID for the Amazon S3 Batch Operations job whose tags you want to delete.
	JobId *string
	// The AWS account ID associated with the Amazon S3 Batch Operations job.
	AccountId *string
}

type DeleteJobTaggingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpDeleteJobTaggingMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpDeleteJobTagging{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteJobTagging{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteJobTagging(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "DeleteJobTagging",
	}
}
