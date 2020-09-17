// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates one or more partitions in a batch operation.
func (c *Client) BatchCreatePartition(ctx context.Context, params *BatchCreatePartitionInput, optFns ...func(*Options)) (*BatchCreatePartitionOutput, error) {
	stack := middleware.NewStack("BatchCreatePartition", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpBatchCreatePartitionMiddlewares(stack)
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
	addOpBatchCreatePartitionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchCreatePartition(options.Region), middleware.Before)
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
			OperationName: "BatchCreatePartition",
			Err:           err,
		}
	}
	out := result.(*BatchCreatePartitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchCreatePartitionInput struct {
	// The ID of the catalog in which the partition is to be created. Currently, this
	// should be the AWS account ID.
	CatalogId *string
	// The name of the metadata database in which the partition is to be created.
	DatabaseName *string
	// The name of the metadata table in which the partition is to be created.
	TableName *string
	// A list of PartitionInput structures that define the partitions to be created.
	PartitionInputList []*types.PartitionInput
}

type BatchCreatePartitionOutput struct {
	// The errors encountered when trying to create the requested partitions.
	Errors []*types.PartitionError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpBatchCreatePartitionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpBatchCreatePartition{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchCreatePartition{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchCreatePartition(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "BatchCreatePartition",
	}
}
