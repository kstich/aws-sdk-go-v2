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

// Gets the text detection results of a Amazon Rekognition Video analysis started
// by StartTextDetection (). Text detection with Amazon Rekognition Video is an
// asynchronous operation. You start text detection by calling StartTextDetection
// () which returns a job identifier (JobId) When the text detection operation
// finishes, Amazon Rekognition publishes a completion status to the Amazon Simple
// Notification Service topic registered in the initial call to StartTextDetection.
// To get the results of the text detection operation, first check that the status
// value published to the Amazon SNS topic is SUCCEEDED. if so, call
// GetTextDetection and pass the job identifier (JobId) from the initial call of
// StartLabelDetection. GetTextDetection returns an array of detected text
// (TextDetections) sorted by the time the text was detected, up to 50 words per
// frame of video. Each element of the array includes the detected text, the
// precentage confidence in the acuracy of the detected text, the time the text was
// detected, bounding box information for where the text was located, and unique
// identifiers for words and their lines. Use MaxResults parameter to limit the
// number of text detections returned. If there are more results than specified in
// MaxResults, the value of NextToken in the operation response contains a
// pagination token for getting the next set of results. To get the next page of
// results, call GetTextDetection and populate the NextToken request parameter with
// the token value returned from the previous call to GetTextDetection.
func (c *Client) GetTextDetection(ctx context.Context, params *GetTextDetectionInput, optFns ...func(*Options)) (*GetTextDetectionOutput, error) {
	stack := middleware.NewStack("GetTextDetection", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetTextDetectionMiddlewares(stack)
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
	addOpGetTextDetectionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetTextDetection(options.Region), middleware.Before)
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
			OperationName: "GetTextDetection",
			Err:           err,
		}
	}
	out := result.(*GetTextDetectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTextDetectionInput struct {
	// Job identifier for the text detection operation for which you want results
	// returned. You get the job identifer from an initial call to StartTextDetection.
	JobId *string
	// Maximum number of results to return per paginated call. The largest value you
	// can specify is 1000.
	MaxResults *int32
	// If the previous response was incomplete (because there are more labels to
	// retrieve), Amazon Rekognition Video returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of text.
	NextToken *string
}

type GetTextDetectionOutput struct {
	// If the response is truncated, Amazon Rekognition Video returns this token that
	// you can use in the subsequent request to retrieve the next set of text.
	NextToken *string
	// Version number of the text detection model that was used to detect text.
	TextModelVersion *string
	// If the job fails, StatusMessage provides a descriptive error message.
	StatusMessage *string
	// Information about a video that Amazon Rekognition analyzed. Videometadata is
	// returned in every page of paginated responses from a Amazon Rekognition video
	// operation.
	VideoMetadata *types.VideoMetadata
	// Current status of the text detection job.
	JobStatus types.VideoJobStatus
	// An array of text detected in the video. Each element contains the detected text,
	// the time in milliseconds from the start of the video that the text was detected,
	// and where it was detected on the screen.
	TextDetections []*types.TextDetectionResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetTextDetectionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetTextDetection{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetTextDetection{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetTextDetection(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "GetTextDetection",
	}
}
