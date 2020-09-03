// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes a member from a group.
func (c *Client) DisassociateMemberFromGroup(ctx context.Context, params *DisassociateMemberFromGroupInput, optFns ...func(*Options)) (*DisassociateMemberFromGroupOutput, error) {
	stack := middleware.NewStack("DisassociateMemberFromGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDisassociateMemberFromGroupMiddlewares(stack)
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
	addOpDisassociateMemberFromGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateMemberFromGroup(options.Region), middleware.Before)

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
			OperationName: "DisassociateMemberFromGroup",
			Err:           err,
		}
	}
	out := result.(*DisassociateMemberFromGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateMemberFromGroupInput struct {
	// The identifier for the member to be removed to the group.
	MemberId *string
	// The identifier for the organization under which the group exists.
	OrganizationId *string
	// The identifier for the group from which members are removed.
	GroupId *string
}

type DisassociateMemberFromGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDisassociateMemberFromGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateMemberFromGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateMemberFromGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisassociateMemberFromGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "DisassociateMemberFromGroup",
	}
}
