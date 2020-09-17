// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a bucket in Amazon S3 to store application versions, logs, and other
// files used by Elastic Beanstalk environments. The Elastic Beanstalk console and
// EB CLI call this API the first time you create an environment in a region. If
// the storage location already exists, CreateStorageLocation still returns the
// bucket name but does not create a new bucket.
func (c *Client) CreateStorageLocation(ctx context.Context, params *CreateStorageLocationInput, optFns ...func(*Options)) (*CreateStorageLocationOutput, error) {
	stack := middleware.NewStack("CreateStorageLocation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateStorageLocationMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateStorageLocation(options.Region), middleware.Before)
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
			OperationName: "CreateStorageLocation",
			Err:           err,
		}
	}
	out := result.(*CreateStorageLocationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateStorageLocationInput struct {
}

// Results of a CreateStorageLocationResult () call.
type CreateStorageLocationOutput struct {
	// The name of the Amazon S3 bucket created.
	S3Bucket *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateStorageLocationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateStorageLocation{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateStorageLocation{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateStorageLocation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "CreateStorageLocation",
	}
}
