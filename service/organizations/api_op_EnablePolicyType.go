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

// Enables a policy type in a root. After you enable a policy type in a root, you
// can attach policies of that type to the root, any organizational unit (OU), or
// account in that root. You can undo this by using the DisablePolicyType ()
// operation. This is an asynchronous request that AWS performs in the background.
// AWS recommends that you first use ListRoots () to see the status of policy types
// for a specified root, and then use this operation. This operation can be called
// only from the organization's master account. You can enable a policy type in a
// root only if that policy type is available in the organization. To view the
// status of available policy types in the organization, use DescribeOrganization
// ().
func (c *Client) EnablePolicyType(ctx context.Context, params *EnablePolicyTypeInput, optFns ...func(*Options)) (*EnablePolicyTypeOutput, error) {
	stack := middleware.NewStack("EnablePolicyType", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpEnablePolicyTypeMiddlewares(stack)
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
	addOpEnablePolicyTypeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opEnablePolicyType(options.Region), middleware.Before)
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
			OperationName: "EnablePolicyType",
			Err:           err,
		}
	}
	out := result.(*EnablePolicyTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnablePolicyTypeInput struct {
	// The policy type that you want to enable. You can specify one of the following
	// values:
	//
	//     * AISERVICES_OPT_OUT_POLICY
	// (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_ai-opt-out.html)
	//
	//
	// * BACKUP_POLICY
	// (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_backup.html)
	//
	//
	// * SERVICE_CONTROL_POLICY
	// (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scp.html)
	//
	//
	// * TAG_POLICY
	// (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_tag-policies.html)
	PolicyType types.PolicyType
	// The unique identifier (ID) of the root in which you want to enable a policy
	// type. You can get the ID from the ListRoots () operation. The regex pattern
	// (http://wikipedia.org/wiki/regex) for a root ID string requires "r-" followed by
	// from 4 to 32 lowercase letters or digits.
	RootId *string
}

type EnablePolicyTypeOutput struct {
	// A structure that shows the root with the updated list of enabled policy types.
	Root *types.Root

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpEnablePolicyTypeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpEnablePolicyType{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpEnablePolicyType{}, middleware.After)
}

func newServiceMetadataMiddleware_opEnablePolicyType(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "EnablePolicyType",
	}
}
