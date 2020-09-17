// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provides a summarized description of the specified Kinesis data stream without
// the shard list. The information returned includes the stream name, Amazon
// Resource Name (ARN), status, record retention period, approximate creation time,
// monitoring, encryption details, and open shard count.
func (c *Client) DescribeStreamSummary(ctx context.Context, params *DescribeStreamSummaryInput, optFns ...func(*Options)) (*DescribeStreamSummaryOutput, error) {
	stack := middleware.NewStack("DescribeStreamSummary", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeStreamSummaryMiddlewares(stack)
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
	addOpDescribeStreamSummaryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStreamSummary(options.Region), middleware.Before)
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
			OperationName: "DescribeStreamSummary",
			Err:           err,
		}
	}
	out := result.(*DescribeStreamSummaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStreamSummaryInput struct {
	// The name of the stream to describe.
	StreamName *string
}

type DescribeStreamSummaryOutput struct {
	// A StreamDescriptionSummary () containing information about the stream.
	StreamDescriptionSummary *types.StreamDescriptionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeStreamSummaryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeStreamSummary{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeStreamSummary{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeStreamSummary(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "DescribeStreamSummary",
	}
}
