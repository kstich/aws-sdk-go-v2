// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disables ClassicLink for a VPC. You cannot disable ClassicLink for a VPC that
// has EC2-Classic instances linked to it.
func (c *Client) DisableVpcClassicLink(ctx context.Context, params *DisableVpcClassicLinkInput, optFns ...func(*Options)) (*DisableVpcClassicLinkOutput, error) {
	stack := middleware.NewStack("DisableVpcClassicLink", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDisableVpcClassicLinkMiddlewares(stack)
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
	addOpDisableVpcClassicLinkValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisableVpcClassicLink(options.Region), middleware.Before)
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
			OperationName: "DisableVpcClassicLink",
			Err:           err,
		}
	}
	out := result.(*DisableVpcClassicLinkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableVpcClassicLinkInput struct {
	// The ID of the VPC.
	VpcId *string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type DisableVpcClassicLinkOutput struct {
	// Returns true if the request succeeds; otherwise, it returns an error.
	Return *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDisableVpcClassicLinkMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDisableVpcClassicLink{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDisableVpcClassicLink{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisableVpcClassicLink(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DisableVpcClassicLink",
	}
}
