// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about a previously requested handshake. The handshake ID
// comes from the response to the original InviteAccountToOrganization () operation
// that generated the handshake. You can access handshakes that are ACCEPTED,
// DECLINED, or CANCELED for only 30 days after they change to that state. They're
// then deleted and no longer accessible. This operation can be called from any
// account in the organization.
func (c *Client) DescribeHandshake(ctx context.Context, params *DescribeHandshakeInput, optFns ...func(*Options)) (*DescribeHandshakeOutput, error) {
	stack := middleware.NewStack("DescribeHandshake", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeHandshakeMiddlewares(stack)
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
	addOpDescribeHandshakeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeHandshake(options.Region), middleware.Before)
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
			OperationName: "DescribeHandshake",
			Err:           err,
		}
	}
	out := result.(*DescribeHandshakeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeHandshakeInput struct {
	// The unique identifier (ID) of the handshake that you want information about. You
	// can get the ID from the original call to InviteAccountToOrganization (), or from
	// a call to ListHandshakesForAccount () or ListHandshakesForOrganization (). The
	// regex pattern (http://wikipedia.org/wiki/regex) for handshake ID string requires
	// "h-" followed by from 8 to 32 lowercase letters or digits.
	HandshakeId *string
}

type DescribeHandshakeOutput struct {
	// A structure that contains information about the specified handshake.
	Handshake *types.Handshake

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeHandshakeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeHandshake{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeHandshake{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeHandshake(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "DescribeHandshake",
	}
}
