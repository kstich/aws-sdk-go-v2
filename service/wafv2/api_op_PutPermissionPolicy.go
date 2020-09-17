// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Attaches an IAM policy to the specified resource. Use this to share a rule group
// across accounts. You must be the owner of the rule group to perform this
// operation. This action is subject to the following restrictions:
//
//     * You can
// attach only one policy with each PutPermissionPolicy request.
//
//     * The ARN in
// the request must be a valid WAF RuleGroup () ARN and the rule group must exist
// in the same region.
//
//     * The user making the request must be the owner of the
// rule group.
func (c *Client) PutPermissionPolicy(ctx context.Context, params *PutPermissionPolicyInput, optFns ...func(*Options)) (*PutPermissionPolicyOutput, error) {
	stack := middleware.NewStack("PutPermissionPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutPermissionPolicyMiddlewares(stack)
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
	addOpPutPermissionPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutPermissionPolicy(options.Region), middleware.Before)
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
			OperationName: "PutPermissionPolicy",
			Err:           err,
		}
	}
	out := result.(*PutPermissionPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutPermissionPolicyInput struct {
	// The Amazon Resource Name (ARN) of the RuleGroup () to which you want to attach
	// the policy.
	ResourceArn *string
	// The policy to attach to the specified rule group.  <p>The policy specifications
	// must conform to the following:</p> <ul> <li> <p>The policy must be composed
	// using IAM Policy version 2012-10-17 or version 2015-01-01.</p> </li> <li> <p>The
	// policy must include specifications for <code>Effect</code>, <code>Action</code>,
	// and <code>Principal</code>.</p> </li> <li> <p> <code>Effect</code> must specify
	// <code>Allow</code>.</p> </li> <li> <p> <code>Action</code> must specify
	// <code>wafv2:CreateWebACL</code>, <code>wafv2:UpdateWebACL</code>, and
	// <code>wafv2:PutFirewallManagerRuleGroups</code>. AWS WAF rejects any extra
	// actions or wildcard actions in the policy.</p> </li> <li> <p>The policy must not
	// include a <code>Resource</code> parameter.</p> </li> </ul> <p>For more
	// information, see <a
	// href="https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html">IAM
	// Policies</a>. </p>
	Policy *string
}

type PutPermissionPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutPermissionPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutPermissionPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutPermissionPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutPermissionPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wafv2",
		OperationName: "PutPermissionPolicy",
	}
}
