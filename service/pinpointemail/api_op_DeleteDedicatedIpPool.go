// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointemail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Delete a dedicated IP pool.
func (c *Client) DeleteDedicatedIpPool(ctx context.Context, params *DeleteDedicatedIpPoolInput, optFns ...func(*Options)) (*DeleteDedicatedIpPoolOutput, error) {
	stack := middleware.NewStack("DeleteDedicatedIpPool", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteDedicatedIpPoolMiddlewares(stack)
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
	addOpDeleteDedicatedIpPoolValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDedicatedIpPool(options.Region), middleware.Before)
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
			OperationName: "DeleteDedicatedIpPool",
			Err:           err,
		}
	}
	out := result.(*DeleteDedicatedIpPoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to delete a dedicated IP pool.
type DeleteDedicatedIpPoolInput struct {
	// The name of the dedicated IP pool that you want to delete.
	PoolName *string
}

// An HTTP 200 response if the request succeeds, or an error message if the request
// fails.
type DeleteDedicatedIpPoolOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteDedicatedIpPoolMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteDedicatedIpPool{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteDedicatedIpPool{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteDedicatedIpPool(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "DeleteDedicatedIpPool",
	}
}
