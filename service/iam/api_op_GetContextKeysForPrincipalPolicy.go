// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list of all of the context keys referenced in all the IAM policies that
// are attached to the specified IAM entity. The entity can be an IAM user, group,
// or role. If you specify a user, then the request also includes all of the
// policies attached to groups that the user is a member of. You can optionally
// include a list of one or more additional policies, specified as strings. If you
// want to include only a list of policies by string, use
// GetContextKeysForCustomPolicy () instead. Note: This API discloses information
// about the permissions granted to other users. If you do not want users to see
// other user's permissions, then consider allowing them to use
// GetContextKeysForCustomPolicy () instead. Context keys are variables maintained
// by AWS and its services that provide details about the context of an API query
// request. Context keys can be evaluated by testing against a value in an IAM
// policy. Use GetContextKeysForPrincipalPolicy () to understand what key names and
// values you must supply when you call SimulatePrincipalPolicy ().
func (c *Client) GetContextKeysForPrincipalPolicy(ctx context.Context, params *GetContextKeysForPrincipalPolicyInput, optFns ...func(*Options)) (*GetContextKeysForPrincipalPolicyOutput, error) {
	stack := middleware.NewStack("GetContextKeysForPrincipalPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpGetContextKeysForPrincipalPolicyMiddlewares(stack)
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
	addOpGetContextKeysForPrincipalPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetContextKeysForPrincipalPolicy(options.Region), middleware.Before)
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
			OperationName: "GetContextKeysForPrincipalPolicy",
			Err:           err,
		}
	}
	out := result.(*GetContextKeysForPrincipalPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetContextKeysForPrincipalPolicyInput struct {
	// The ARN of a user, group, or role whose policies contain the context keys that
	// you want listed. If you specify a user, the list includes context keys that are
	// found in all policies that are attached to the user. The list also includes all
	// groups that the user is a member of. If you pick a group or a role, then it
	// includes only those context keys that are found in policies attached to that
	// entity. Note that all parameters are shown in unencoded form here for clarity,
	// but must be URL encoded to be included as a part of a real HTML request. For
	// more information about ARNs, see Amazon Resource Names (ARNs) and AWS Service
	// Namespaces
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference.
	PolicySourceArn *string
	// An optional list of additional policies for which you want the list of context
	// keys that are referenced. The regex pattern (http://wikipedia.org/wiki/regex)
	// used to validate this parameter is a string of characters consisting of the
	// following:
	//
	//     * Any printable ASCII character ranging from the space character
	// (\u0020) through the end of the ASCII character range
	//
	//     * The printable
	// characters in the Basic Latin and Latin-1 Supplement character set (through
	// \u00FF)
	//
	//     * The special characters tab (\u0009), line feed (\u000A), and
	// carriage return (\u000D)
	PolicyInputList []*string
}

// Contains the response to a successful GetContextKeysForPrincipalPolicy () or
// GetContextKeysForCustomPolicy () request.
type GetContextKeysForPrincipalPolicyOutput struct {
	// The list of context keys that are referenced in the input policies.
	ContextKeyNames []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpGetContextKeysForPrincipalPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpGetContextKeysForPrincipalPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpGetContextKeysForPrincipalPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetContextKeysForPrincipalPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "GetContextKeysForPrincipalPolicy",
	}
}
