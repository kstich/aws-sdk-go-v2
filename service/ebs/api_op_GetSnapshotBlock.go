// Code generated by smithy-go-codegen DO NOT EDIT.

package ebs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ebs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"io"
)

// Returns the data in a block in an Amazon Elastic Block Store snapshot.
func (c *Client) GetSnapshotBlock(ctx context.Context, params *GetSnapshotBlockInput, optFns ...func(*Options)) (*GetSnapshotBlockOutput, error) {
	stack := middleware.NewStack("GetSnapshotBlock", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetSnapshotBlockMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	addOpGetSnapshotBlockValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSnapshotBlock(options.Region), middleware.Before)

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
			OperationName: "GetSnapshotBlock",
			Err:           err,
		}
	}
	out := result.(*GetSnapshotBlockOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSnapshotBlockInput struct {
	// The block token of the block from which to get data.  <p>Obtain the
	// <code>BlockToken</code> by running the <code>ListChangedBlocks</code> or
	// <code>ListSnapshotBlocks</code> operations.</p>
	BlockToken *string
	// The block index of the block from which to get data.  <p>Obtain the
	// <code>BlockIndex</code> by running the <code>ListChangedBlocks</code> or
	// <code>ListSnapshotBlocks</code> operations.</p>
	BlockIndex *int32
	// The ID of the snapshot containing the block from which to get data.
	SnapshotId *string
}

type GetSnapshotBlockOutput struct {
	// The data content of the block.
	BlockData io.ReadCloser
	// The checksum generated for the block, which is Base64 encoded.
	Checksum *string
	// The size of the data in the block.
	DataLength *int32
	// The algorithm used to generate the checksum for the block, such as SHA256.
	ChecksumAlgorithm types.ChecksumAlgorithm

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetSnapshotBlockMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetSnapshotBlock{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSnapshotBlock{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetSnapshotBlock(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ebs",
		OperationName: "GetSnapshotBlock",
	}
}
