// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds or updates organization config rule for your entire organization evaluating
// whether your AWS resources comply with your desired configurations. Only a
// master account and a delegated administrator can create or update an
// organization config rule. When calling this API with a delegated administrator,
// you must ensure AWS Organizations ListDelegatedAdministrator permissions are
// added. This API enables organization service access through the
// EnableAWSServiceAccess action and creates a service linked role
// AWSServiceRoleForConfigMultiAccountSetup in the master or delegated
// administrator account of your organization. The service linked role is created
// only when the role does not exist in the caller account. AWS Config verifies the
// existence of role with GetRole action. To use this API with delegated
// administrator, register a delegated administrator by calling AWS Organization
// register-delegated-administrator for config-multiaccountsetup.amazonaws.com. You
// can use this action to create both custom AWS Config rules and AWS managed
// Config rules. If you are adding a new custom AWS Config rule, you must first
// create AWS Lambda function in the master account or a delegated administrator
// that the rule invokes to evaluate your resources. When you use the
// PutOrganizationConfigRule action to add the rule to AWS Config, you must specify
// the Amazon Resource Name (ARN) that AWS Lambda assigns to the function. If you
// are adding an AWS managed Config rule, specify the rule's identifier for the
// RuleIdentifier key. The maximum number of organization config rules that AWS
// Config supports is 150 and 3 delegated administrator per organization.
// Prerequisite: Ensure you call EnableAllFeatures API to enable all features in an
// organization. Specify either OrganizationCustomRuleMetadata or
// OrganizationManagedRuleMetadata.
func (c *Client) PutOrganizationConfigRule(ctx context.Context, params *PutOrganizationConfigRuleInput, optFns ...func(*Options)) (*PutOrganizationConfigRuleOutput, error) {
	stack := middleware.NewStack("PutOrganizationConfigRule", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutOrganizationConfigRuleMiddlewares(stack)
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
	addOpPutOrganizationConfigRuleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutOrganizationConfigRule(options.Region), middleware.Before)

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
			OperationName: "PutOrganizationConfigRule",
			Err:           err,
		}
	}
	out := result.(*PutOrganizationConfigRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutOrganizationConfigRuleInput struct {
	// An OrganizationCustomRuleMetadata object.
	OrganizationCustomRuleMetadata *types.OrganizationCustomRuleMetadata
	// A comma-separated list of accounts that you want to exclude from an organization
	// config rule.
	ExcludedAccounts []*string
	// An OrganizationManagedRuleMetadata object.
	OrganizationManagedRuleMetadata *types.OrganizationManagedRuleMetadata
	// The name that you assign to an organization config rule.
	OrganizationConfigRuleName *string
}

type PutOrganizationConfigRuleOutput struct {
	// The Amazon Resource Name (ARN) of an organization config rule.
	OrganizationConfigRuleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutOrganizationConfigRuleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutOrganizationConfigRule{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutOrganizationConfigRule{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutOrganizationConfigRule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "PutOrganizationConfigRule",
	}
}