// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts asynchronous detection of segment detection in a stored video. Amazon
// Rekognition Video can detect segments in a video stored in an Amazon S3 bucket.
// Use Video () to specify the bucket name and the filename of the video.
// StartSegmentDetection returns a job identifier (JobId) which you use to get the
// results of the operation. When segment detection is finished, Amazon Rekognition
// Video publishes a completion status to the Amazon Simple Notification Service
// topic that you specify in NotificationChannel. You can use the Filters
// (StartSegmentDetectionFilters ()) input parameter to specify the minimum
// detection confidence returned in the response. Within Filters, use ShotFilter
// (StartShotDetectionFilter ()) to filter detected shots. Use TechnicalCueFilter
// (StartTechnicalCueDetectionFilter ()) to filter technical cues. To get the
// results of the segment detection operation, first check that the status value
// published to the Amazon SNS topic is SUCCEEDED. if so, call GetSegmentDetection
// () and pass the job identifier (JobId) from the initial call to
// StartSegmentDetection.  <p>For more information, see Detecting Video Segments in
// Stored Video in the Amazon Rekognition Developer Guide.</p>
func (c *Client) StartSegmentDetection(ctx context.Context, params *StartSegmentDetectionInput, optFns ...func(*Options)) (*StartSegmentDetectionOutput, error) {
	stack := middleware.NewStack("StartSegmentDetection", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartSegmentDetectionMiddlewares(stack)
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
	addOpStartSegmentDetectionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartSegmentDetection(options.Region), middleware.Before)
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
			OperationName: "StartSegmentDetection",
			Err:           err,
		}
	}
	out := result.(*StartSegmentDetectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartSegmentDetectionInput struct {
	// An array of segment types to detect in the video. Valid values are TECHNICAL_CUE
	// and SHOT.
	SegmentTypes []types.SegmentType
	// Idempotent token used to identify the start request. If you use the same token
	// with multiple StartSegmentDetection requests, the same JobId is returned. Use
	// ClientRequestToken to prevent the same job from being accidently started more
	// than once.
	ClientRequestToken *string
	// The ARN of the Amazon SNS topic to which you want Amazon Rekognition Video to
	// publish the completion status of the segment detection operation.
	NotificationChannel *types.NotificationChannel
	// Filters for technical cue or shot detection.
	Filters *types.StartSegmentDetectionFilters
	// Video file stored in an Amazon S3 bucket. Amazon Rekognition video start
	// operations such as StartLabelDetection () use Video to specify a video for
	// analysis. The supported file formats are .mp4, .mov and .avi.
	Video *types.Video
	// An identifier you specify that's returned in the completion notification that's
	// published to your Amazon Simple Notification Service topic. For example, you can
	// use JobTag to group related jobs and identify them in the completion
	// notification.
	JobTag *string
}

type StartSegmentDetectionOutput struct {
	// Unique identifier for the segment detection job. The JobId is returned from
	// StartSegmentDetection.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartSegmentDetectionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartSegmentDetection{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartSegmentDetection{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartSegmentDetection(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "StartSegmentDetection",
	}
}
