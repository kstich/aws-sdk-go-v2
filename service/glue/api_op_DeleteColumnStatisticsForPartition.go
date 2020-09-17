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

// Delete the partition column statistics of a column.
func (c *Client) DeleteColumnStatisticsForPartition(ctx context.Context, params *DeleteColumnStatisticsForPartitionInput, optFns ...func(*Options)) (*DeleteColumnStatisticsForPartitionOutput, error) {
	stack := middleware.NewStack("DeleteColumnStatisticsForPartition", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteColumnStatisticsForPartitionMiddlewares(stack)
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
	addOpDeleteColumnStatisticsForPartitionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteColumnStatisticsForPartition(options.Region), middleware.Before)
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
			OperationName: "DeleteColumnStatisticsForPartition",
			Err:           err,
		}
	}
	out := result.(*DeleteColumnStatisticsForPartitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteColumnStatisticsForPartitionInput struct {
	// The name of the partitions' table.
	TableName *string
	// Name of the column.
	ColumnName *string
	// The name of the catalog database where the partitions reside.
	DatabaseName *string
	// The ID of the Data Catalog where the partitions in question reside. If none is
	// supplied, the AWS account ID is used by default.
	CatalogId *string
	// A list of partition values identifying the partition.
	PartitionValues []*string
}

type DeleteColumnStatisticsForPartitionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteColumnStatisticsForPartitionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteColumnStatisticsForPartition{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteColumnStatisticsForPartition{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteColumnStatisticsForPartition(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "DeleteColumnStatisticsForPartition",
	}
}
