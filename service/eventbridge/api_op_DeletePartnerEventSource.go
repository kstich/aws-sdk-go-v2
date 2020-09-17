// Code generated by smithy-go-codegen DO NOT EDIT.

package eventbridge

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation is used by SaaS partners to delete a partner event source. This
// operation is not used by AWS customers. When you delete an event source, the
// status of the corresponding partner event bus in the AWS customer account
// becomes DELETED.
func (c *Client) DeletePartnerEventSource(ctx context.Context, params *DeletePartnerEventSourceInput, optFns ...func(*Options)) (*DeletePartnerEventSourceOutput, error) {
	stack := middleware.NewStack("DeletePartnerEventSource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeletePartnerEventSourceMiddlewares(stack)
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
	addOpDeletePartnerEventSourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePartnerEventSource(options.Region), middleware.Before)
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
			OperationName: "DeletePartnerEventSource",
			Err:           err,
		}
	}
	out := result.(*DeletePartnerEventSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePartnerEventSourceInput struct {
	// The AWS account ID of the AWS customer that the event source was created for.
	Account *string
	// The name of the event source to delete.
	Name *string
}

type DeletePartnerEventSourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeletePartnerEventSourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeletePartnerEventSource{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeletePartnerEventSource{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeletePartnerEventSource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "DeletePartnerEventSource",
	}
}
