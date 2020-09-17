// Code generated by smithy-go-codegen DO NOT EDIT.

package firehose

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists your delivery streams in alphabetical order of their names. The number of
// delivery streams might be too large to return using a single call to
// ListDeliveryStreams. You can limit the number of delivery streams returned,
// using the Limit parameter. To determine whether there are more delivery streams
// to list, check the value of HasMoreDeliveryStreams in the output. If there are
// more delivery streams to list, you can request them by calling this operation
// again and setting the ExclusiveStartDeliveryStreamName parameter to the name of
// the last delivery stream returned in the last call.
func (c *Client) ListDeliveryStreams(ctx context.Context, params *ListDeliveryStreamsInput, optFns ...func(*Options)) (*ListDeliveryStreamsOutput, error) {
	stack := middleware.NewStack("ListDeliveryStreams", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListDeliveryStreamsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListDeliveryStreams(options.Region), middleware.Before)
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
			OperationName: "ListDeliveryStreams",
			Err:           err,
		}
	}
	out := result.(*ListDeliveryStreamsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDeliveryStreamsInput struct {
	// The list of delivery streams returned by this call to ListDeliveryStreams will
	// start with the delivery stream whose name comes alphabetically immediately after
	// the name you specify in ExclusiveStartDeliveryStreamName.
	ExclusiveStartDeliveryStreamName *string
	// The delivery stream type. This can be one of the following values:
	//
	//     *
	// DirectPut: Provider applications access the delivery stream directly.
	//
	//     *
	// KinesisStreamAsSource: The delivery stream uses a Kinesis data stream as a
	// source.
	//
	// This parameter is optional. If this parameter is omitted, delivery
	// streams of all types are returned.
	DeliveryStreamType types.DeliveryStreamType
	// The maximum number of delivery streams to list. The default value is 10.
	Limit *int32
}

type ListDeliveryStreamsOutput struct {
	// Indicates whether there are more delivery streams available to list.
	HasMoreDeliveryStreams *bool
	// The names of the delivery streams.
	DeliveryStreamNames []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListDeliveryStreamsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListDeliveryStreams{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDeliveryStreams{}, middleware.After)
}

func newServiceMetadataMiddleware_opListDeliveryStreams(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "firehose",
		OperationName: "ListDeliveryStreams",
	}
}
