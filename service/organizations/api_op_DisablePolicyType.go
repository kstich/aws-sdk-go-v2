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

// Disables an organizational policy type in a root. A policy of a certain type can
// be attached to entities in a root only if that type is enabled in the root.
// After you perform this operation, you no longer can attach policies of the
// specified type to that root or to any organizational unit (OU) or account in
// that root. You can undo this by using the EnablePolicyType () operation. This is
// an asynchronous request that AWS performs in the background. If you disable a
// policy type for a root, it still appears enabled for the organization if all
// features
// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html)
// are enabled for the organization. AWS recommends that you first use ListRoots ()
// to see the status of policy types for a specified root, and then use this
// operation. This operation can be called only from the organization's master
// account. To view the status of available policy types in the organization, use
// DescribeOrganization ().
func (c *Client) DisablePolicyType(ctx context.Context, params *DisablePolicyTypeInput, optFns ...func(*Options)) (*DisablePolicyTypeOutput, error) {
	stack := middleware.NewStack("DisablePolicyType", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDisablePolicyTypeMiddlewares(stack)
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
	addOpDisablePolicyTypeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisablePolicyType(options.Region), middleware.Before)

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
			OperationName: "DisablePolicyType",
			Err:           err,
		}
	}
	out := result.(*DisablePolicyTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisablePolicyTypeInput struct {
	// The unique identifier (ID) of the root in which you want to disable a policy
	// type. You can get the ID from the ListRoots () operation. The regex pattern
	// (http://wikipedia.org/wiki/regex) for a root ID string requires "r-" followed by
	// from 4 to 32 lowercase letters or digits.
	RootId *string
	// The policy type that you want to disable in this root. You can specify one of
	// the following values:
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
}

type DisablePolicyTypeOutput struct {
	// A structure that shows the root with the updated list of enabled policy types.
	Root *types.Root

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDisablePolicyTypeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDisablePolicyType{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisablePolicyType{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisablePolicyType(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "DisablePolicyType",
	}
}
