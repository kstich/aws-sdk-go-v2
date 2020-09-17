// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes an existing Amazon Kendra index
func (c *Client) DescribeIndex(ctx context.Context, params *DescribeIndexInput, optFns ...func(*Options)) (*DescribeIndexOutput, error) {
	stack := middleware.NewStack("DescribeIndex", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeIndexMiddlewares(stack)
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
	addOpDescribeIndexValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeIndex(options.Region), middleware.Before)
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
			OperationName: "DescribeIndex",
			Err:           err,
		}
	}
	out := result.(*DescribeIndexOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeIndexInput struct {
	// The name of the index to describe.
	Id *string
}

type DescribeIndexOutput struct {
	// When th eStatus field value is FAILED, the ErrorMessage field contains a message
	// that explains why.
	ErrorMessage *string
	// The Amazon Resource Name (ARN) of the IAM role that gives Amazon Kendra
	// permission to write to your Amazon Cloudwatch logs.
	RoleArn *string
	// Configuration settings for any metadata applied to the documents in the index.
	DocumentMetadataConfigurations []*types.DocumentMetadataConfiguration
	// the name of the index.
	Id *string
	// The description of the index.
	Description *string
	// For enterprise edtion indexes, you can choose to use additional capacity to meet
	// the needs of your application. This contains the capacity units used for the
	// index. A 0 for the query capacity or the storage capacity indicates that the
	// index is using the default capacity for the index.
	CapacityUnits *types.CapacityUnitsConfiguration
	// Provides information about the number of FAQ questions and answers and the
	// number of text documents indexed.
	IndexStatistics *types.IndexStatistics
	// The Unix datetime that the index was created.
	CreatedAt *time.Time
	// The Unix datetime that the index was last updated.
	UpdatedAt *time.Time
	// The name of the index.
	Name *string
	// The Amazon Kendra edition used for the index. You decide the edition when you
	// create the index.
	Edition types.IndexEdition
	// The identifier of the AWS KMS customer master key (CMK) used to encrypt your
	// data. Amazon Kendra doesn't support asymmetric CMKs.
	ServerSideEncryptionConfiguration *types.ServerSideEncryptionConfiguration
	// The current status of the index. When the value is ACTIVE, the index is ready
	// for use. If the Status field value is FAILED, the ErrorMessage field contains a
	// message that explains why.
	Status types.IndexStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeIndexMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeIndex{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeIndex{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeIndex(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "DescribeIndex",
	}
}
