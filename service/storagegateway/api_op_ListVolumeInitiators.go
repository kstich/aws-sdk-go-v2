// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists iSCSI initiators that are connected to a volume. You can use this
// operation to determine whether a volume is being used or not. This operation is
// only supported in the cached volume and stored volume gateway types.
func (c *Client) ListVolumeInitiators(ctx context.Context, params *ListVolumeInitiatorsInput, optFns ...func(*Options)) (*ListVolumeInitiatorsOutput, error) {
	stack := middleware.NewStack("ListVolumeInitiators", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListVolumeInitiatorsMiddlewares(stack)
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
	addOpListVolumeInitiatorsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListVolumeInitiators(options.Region), middleware.Before)
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
			OperationName: "ListVolumeInitiators",
			Err:           err,
		}
	}
	out := result.(*ListVolumeInitiatorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// ListVolumeInitiatorsInput
type ListVolumeInitiatorsInput struct {
	// The Amazon Resource Name (ARN) of the volume. Use the ListVolumes () operation
	// to return a list of gateway volumes for the gateway.
	VolumeARN *string
}

// ListVolumeInitiatorsOutput
type ListVolumeInitiatorsOutput struct {
	// The host names and port numbers of all iSCSI initiators that are connected to
	// the gateway.
	Initiators []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListVolumeInitiatorsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListVolumeInitiators{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListVolumeInitiators{}, middleware.After)
}

func newServiceMetadataMiddleware_opListVolumeInitiators(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "ListVolumeInitiators",
	}
}
