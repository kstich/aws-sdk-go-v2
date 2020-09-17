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
	"time"
)

// Returns the block indexes and block tokens for blocks that are different between
// two Amazon Elastic Block Store snapshots of the same volume/snapshot lineage.
func (c *Client) ListChangedBlocks(ctx context.Context, params *ListChangedBlocksInput, optFns ...func(*Options)) (*ListChangedBlocksOutput, error) {
	stack := middleware.NewStack("ListChangedBlocks", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListChangedBlocksMiddlewares(stack)
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
	addOpListChangedBlocksValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListChangedBlocks(options.Region), middleware.Before)
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
			OperationName: "ListChangedBlocks",
			Err:           err,
		}
	}
	out := result.(*ListChangedBlocksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListChangedBlocksInput struct {
	// The ID of the first snapshot to use for the comparison. The FirstSnapshotID
	// parameter must be specified with a SecondSnapshotId parameter; otherwise, an
	// error occurs.
	FirstSnapshotId *string
	// The token to request the next page of results.
	NextToken *string
	// The number of results to return.
	MaxResults *int32
	// The ID of the second snapshot to use for the comparison. The SecondSnapshotId
	// parameter must be specified with a FirstSnapshotID parameter; otherwise, an
	// error occurs.
	SecondSnapshotId *string
	// The block index from which the comparison should start.  <p>The list in the
	// response will start from this block index or the next valid block index in the
	// snapshots.</p>
	StartingBlockIndex *int32
}

type ListChangedBlocksOutput struct {
	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string
	// The size of the volume in GB.
	VolumeSize *int64
	// The size of the block.
	BlockSize *int32
	// The time when the BlockToken expires.
	ExpiryTime *time.Time
	// An array of objects containing information about the changed blocks.
	ChangedBlocks []*types.ChangedBlock

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListChangedBlocksMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListChangedBlocks{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListChangedBlocks{}, middleware.After)
}

func newServiceMetadataMiddleware_opListChangedBlocks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ebs",
		OperationName: "ListChangedBlocks",
	}
}
