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

// Detects text in the input image and converts it into machine-readable text. Pass
// the input image as base64-encoded image bytes or as a reference to an image in
// an Amazon S3 bucket. If you use the AWS CLI to call Amazon Rekognition
// operations, you must pass it as a reference to an image in an Amazon S3 bucket.
// For the AWS CLI, passing image bytes is not supported. The image must be either
// a .png or .jpeg formatted file. The DetectText operation returns text in an
// array of TextDetection () elements, TextDetections. Each TextDetection element
// provides information about a single word or line of text that was detected in
// the image. A word is one or more ISO basic latin script characters that are not
// separated by spaces. DetectText can detect up to 50 words in an image. A line is
// a string of equally spaced words. A line isn't necessarily a complete sentence.
// For example, a driver's license number is detected as a line. A line ends when
// there is no aligned text after it. Also, a line ends when there is a large gap
// between words, relative to the length of the words. This means, depending on the
// gap between words, Amazon Rekognition may detect multiple lines in text aligned
// in the same direction. Periods don't represent the end of a line. If a sentence
// spans multiple lines, the DetectText operation returns multiple lines. To
// determine whether a TextDetection element is a line of text or a word, use the
// TextDetection object Type field. To be detected, text must be within +/- 90
// degrees orientation of the horizontal axis.  <p>For more information, see
// DetectText in the Amazon Rekognition Developer Guide.</p>
func (c *Client) DetectText(ctx context.Context, params *DetectTextInput, optFns ...func(*Options)) (*DetectTextOutput, error) {
	stack := middleware.NewStack("DetectText", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDetectTextMiddlewares(stack)
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
	addOpDetectTextValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDetectText(options.Region), middleware.Before)

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
			OperationName: "DetectText",
			Err:           err,
		}
	}
	out := result.(*DetectTextOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectTextInput struct {
	// Optional parameters that let you set the criteria that the text must meet to be
	// included in your response.
	Filters *types.DetectTextFilters
	// The input image as base64-encoded bytes or an Amazon S3 object. If you use the
	// AWS CLI to call Amazon Rekognition operations, you can't pass image bytes. If
	// you are using an AWS SDK to call Amazon Rekognition, you might not need to
	// base64-encode image bytes passed using the Bytes field. For more information,
	// see Images in the Amazon Rekognition developer guide.
	Image *types.Image
}

type DetectTextOutput struct {
	// The model version used to detect text.
	TextModelVersion *string
	// An array of text that was detected in the input image.
	TextDetections []*types.TextDetection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDetectTextMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDetectText{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetectText{}, middleware.After)
}

func newServiceMetadataMiddleware_opDetectText(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "DetectText",
	}
}
