// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a conditional forwarder that has been set up for your AWS directory.
func (c *Client) UpdateConditionalForwarder(ctx context.Context, params *UpdateConditionalForwarderInput, optFns ...func(*Options)) (*UpdateConditionalForwarderOutput, error) {
	stack := middleware.NewStack("UpdateConditionalForwarder", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateConditionalForwarderMiddlewares(stack)
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
	addOpUpdateConditionalForwarderValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateConditionalForwarder(options.Region), middleware.Before)
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
			OperationName: "UpdateConditionalForwarder",
			Err:           err,
		}
	}
	out := result.(*UpdateConditionalForwarderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Updates a conditional forwarder.
type UpdateConditionalForwarderInput struct {
	// The fully qualified domain name (FQDN) of the remote domain with which you will
	// set up a trust relationship.
	RemoteDomainName *string
	// The directory ID of the AWS directory for which to update the conditional
	// forwarder.
	DirectoryId *string
	// The updated IP addresses of the remote DNS server associated with the
	// conditional forwarder.
	DnsIpAddrs []*string
}

// The result of an UpdateConditionalForwarder request.
type UpdateConditionalForwarderOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateConditionalForwarderMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateConditionalForwarder{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateConditionalForwarder{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateConditionalForwarder(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "UpdateConditionalForwarder",
	}
}
