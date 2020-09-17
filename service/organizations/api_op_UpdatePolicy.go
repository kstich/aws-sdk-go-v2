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

// Updates an existing policy with a new name, description, or content. If you
// don't supply any parameter, that value remains unchanged. You can't change a
// policy's type. This operation can be called only from the organization's master
// account.
func (c *Client) UpdatePolicy(ctx context.Context, params *UpdatePolicyInput, optFns ...func(*Options)) (*UpdatePolicyOutput, error) {
	stack := middleware.NewStack("UpdatePolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdatePolicyMiddlewares(stack)
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
	addOpUpdatePolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePolicy(options.Region), middleware.Before)
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
			OperationName: "UpdatePolicy",
			Err:           err,
		}
	}
	out := result.(*UpdatePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePolicyInput struct {
	// If provided, the new content for the policy. The text must be correctly
	// formatted JSON that complies with the syntax for the policy's type. For more
	// information, see Service Control Policy Syntax
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_reference_scp-syntax.html)
	// in the AWS Organizations User Guide.
	Content *string
	// The unique identifier (ID) of the policy that you want to update. The regex
	// pattern (http://wikipedia.org/wiki/regex) for a policy ID string requires "p-"
	// followed by from 8 to 128 lowercase or uppercase letters, digits, or the
	// underscore character (_).
	PolicyId *string
	// If provided, the new name for the policy. The regex pattern
	// (http://wikipedia.org/wiki/regex) that is used to validate this parameter is a
	// string of any of the characters in the ASCII character range.
	Name *string
	// If provided, the new description for the policy.
	Description *string
}

type UpdatePolicyOutput struct {
	// A structure that contains details about the updated policy, showing the
	// requested changes.
	Policy *types.Policy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdatePolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdatePolicy{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdatePolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdatePolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "UpdatePolicy",
	}
}
