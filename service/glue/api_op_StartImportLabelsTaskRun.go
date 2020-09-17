// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables you to provide additional labels (examples of truth) to be used to teach
// the machine learning transform and improve its quality. This API operation is
// generally used as part of the active learning workflow that starts with the
// StartMLLabelingSetGenerationTaskRun call and that ultimately results in
// improving the quality of your machine learning transform.  <p>After the
// <code>StartMLLabelingSetGenerationTaskRun</code> finishes, AWS Glue machine
// learning will have generated a series of questions for humans to answer.
// (Answering these questions is often called 'labeling' in the machine learning
// workflows). In the case of the <code>FindMatches</code> transform, these
// questions are of the form, “What is the correct way to group these rows together
// into groups composed entirely of matching records?” After the labeling process
// is finished, users upload their answers/labels with a call to
// <code>StartImportLabelsTaskRun</code>. After
// <code>StartImportLabelsTaskRun</code> finishes, all future runs of the machine
// learning transform use the new and improved labels and perform a higher-quality
// transformation.</p> <p>By default,
// <code>StartMLLabelingSetGenerationTaskRun</code> continually learns from and
// combines all labels that you upload unless you set <code>Replace</code> to true.
// If you set <code>Replace</code> to true, <code>StartImportLabelsTaskRun</code>
// deletes and forgets all previously uploaded labels and learns only from the
// exact set that you upload. Replacing labels can be helpful if you realize that
// you previously uploaded incorrect labels, and you believe that they are having a
// negative effect on your transform quality.</p> <p>You can check on the status of
// your task run by calling the <code>GetMLTaskRun</code> operation. </p>
func (c *Client) StartImportLabelsTaskRun(ctx context.Context, params *StartImportLabelsTaskRunInput, optFns ...func(*Options)) (*StartImportLabelsTaskRunOutput, error) {
	stack := middleware.NewStack("StartImportLabelsTaskRun", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartImportLabelsTaskRunMiddlewares(stack)
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
	addOpStartImportLabelsTaskRunValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartImportLabelsTaskRun(options.Region), middleware.Before)
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
			OperationName: "StartImportLabelsTaskRun",
			Err:           err,
		}
	}
	out := result.(*StartImportLabelsTaskRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartImportLabelsTaskRunInput struct {
	// Indicates whether to overwrite your existing labels.
	ReplaceAllLabels *bool
	// The Amazon Simple Storage Service (Amazon S3) path from where you import the
	// labels.
	InputS3Path *string
	// The unique identifier of the machine learning transform.
	TransformId *string
}

type StartImportLabelsTaskRunOutput struct {
	// The unique identifier for the task run.
	TaskRunId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartImportLabelsTaskRunMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartImportLabelsTaskRun{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartImportLabelsTaskRun{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartImportLabelsTaskRun(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "StartImportLabelsTaskRun",
	}
}
