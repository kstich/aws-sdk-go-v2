// Code generated by smithy-go-codegen DO NOT EDIT.

package rdsdata

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts a SQL transaction.  <important> <p>A transaction can run for a maximum of
// 24 hours. A transaction is terminated and rolled back automatically after 24
// hours.</p> <p>A transaction times out if no calls use its transaction ID in
// three minutes. If a transaction times out before it's committed, it's rolled
// back automatically.</p> <p>DDL statements inside a transaction cause an implicit
// commit. We recommend that you run each DDL statement in a separate
// <code>ExecuteStatement</code> call with <code>continueAfterTimeout</code>
// enabled.</p> </important>
func (c *Client) BeginTransaction(ctx context.Context, params *BeginTransactionInput, optFns ...func(*Options)) (*BeginTransactionOutput, error) {
	stack := middleware.NewStack("BeginTransaction", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpBeginTransactionMiddlewares(stack)
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
	addOpBeginTransactionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBeginTransaction(options.Region), middleware.Before)

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
			OperationName: "BeginTransaction",
			Err:           err,
		}
	}
	out := result.(*BeginTransactionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request parameters represent the input of a request to start a SQL
// transaction.
type BeginTransactionInput struct {
	// The name of the database.
	Database *string
	// The Amazon Resource Name (ARN) of the Aurora Serverless DB cluster.
	ResourceArn *string
	// The name of the database schema.
	Schema *string
	// The name or ARN of the secret that enables access to the DB cluster.
	SecretArn *string
}

// The response elements represent the output of a request to start a SQL
// transaction.
type BeginTransactionOutput struct {
	// The transaction ID of the transaction started by the call.
	TransactionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpBeginTransactionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpBeginTransaction{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpBeginTransaction{}, middleware.After)
}

func newServiceMetadataMiddleware_opBeginTransaction(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "",
		OperationName: "BeginTransaction",
	}
}
