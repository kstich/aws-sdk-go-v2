// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Create a new invalidation.
func (c *Client) CreateInvalidation(ctx context.Context, params *CreateInvalidationInput, optFns ...func(*Options)) (*CreateInvalidationOutput, error) {
	stack := middleware.NewStack("CreateInvalidation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpCreateInvalidationMiddlewares(stack)
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
	addOpCreateInvalidationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateInvalidation(options.Region), middleware.Before)
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
			OperationName: "CreateInvalidation",
			Err:           err,
		}
	}
	out := result.(*CreateInvalidationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to create an invalidation.
type CreateInvalidationInput struct {
	// The batch information for the invalidation.
	InvalidationBatch *types.InvalidationBatch
	// The distribution's id.
	DistributionId *string
}

// The returned result of the corresponding request.
type CreateInvalidationOutput struct {
	// The invalidation's information.
	Invalidation *types.Invalidation
	// The fully qualified URI of the distribution and invalidation batch request,
	// including the Invalidation ID.
	Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpCreateInvalidationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpCreateInvalidation{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpCreateInvalidation{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateInvalidation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "CreateInvalidation",
	}
}
