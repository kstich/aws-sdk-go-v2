// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List the listeners for an accelerator. To see an AWS CLI example of listing the
// listeners for an accelerator, scroll down to Example.
func (c *Client) ListListeners(ctx context.Context, params *ListListenersInput, optFns ...func(*Options)) (*ListListenersOutput, error) {
	stack := middleware.NewStack("ListListeners", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListListenersMiddlewares(stack)
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
	addOpListListenersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListListeners(options.Region), middleware.Before)
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
			OperationName: "ListListeners",
			Err:           err,
		}
	}
	out := result.(*ListListenersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListListenersInput struct {
	// The Amazon Resource Name (ARN) of the accelerator for which you want to list
	// listener objects.
	AcceleratorArn *string
	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string
	// The number of listener objects that you want to return with this call. The
	// default value is 10.
	MaxResults *int32
}

type ListListenersOutput struct {
	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string
	// The list of listeners for an accelerator.
	Listeners []*types.Listener

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListListenersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListListeners{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListListeners{}, middleware.After)
}

func newServiceMetadataMiddleware_opListListeners(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "globalaccelerator",
		OperationName: "ListListeners",
	}
}
