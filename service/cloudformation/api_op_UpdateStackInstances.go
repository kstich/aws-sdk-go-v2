// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the parameter values for stack instances for the specified accounts,
// within the specified Regions. A stack instance refers to a stack in a specific
// account and Region. You can only update stack instances in Regions and accounts
// where they already exist; to create additional stack instances, use
// CreateStackInstances
// (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_CreateStackInstances.html).
// During stack set updates, any parameters overridden for a stack instance are not
// updated, but retain their overridden value. You can only update the parameter
// values that are specified in the stack set; to add or delete a parameter itself,
// use UpdateStackSet
// (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_UpdateStackSet.html)
// to update the stack set template. If you add a parameter to a template, before
// you can override the parameter value specified in the stack set you must first
// use UpdateStackSet
// (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_UpdateStackSet.html)
// to update all stack instances with the updated template and parameter value
// specified in the stack set. Once a stack instance has been updated with the new
// parameter, you can then override the parameter value using UpdateStackInstances.
func (c *Client) UpdateStackInstances(ctx context.Context, params *UpdateStackInstancesInput, optFns ...func(*Options)) (*UpdateStackInstancesOutput, error) {
	stack := middleware.NewStack("UpdateStackInstances", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpUpdateStackInstancesMiddlewares(stack)
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
	addIdempotencyToken_opUpdateStackInstancesMiddleware(stack, options)
	addOpUpdateStackInstancesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateStackInstances(options.Region), middleware.Before)

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
			OperationName: "UpdateStackInstances",
			Err:           err,
		}
	}
	out := result.(*UpdateStackInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateStackInstancesInput struct {
	// A list of input parameters whose values you want to update for the specified
	// stack instances. Any overridden parameter values will be applied to all stack
	// instances in the specified accounts and Regions. When specifying parameters and
	// their values, be aware of how AWS CloudFormation sets parameter values during
	// stack instance update operations:
	//
	//     * To override the current value for a
	// parameter, include the parameter and specify its value.
	//
	//     * To leave a
	// parameter set to its present value, you can do one of the following:
	//
	//         *
	// Do not include the parameter in the list.
	//
	//         * Include the parameter and
	// specify UsePreviousValue as true. (You cannot specify both a value and set
	// UsePreviousValue to true.)
	//
	//     * To set all overridden parameter back to the
	// values specified in the stack set, specify a parameter list but do not include
	// any parameters.
	//
	//     * To leave all parameters set to their present values, do
	// not specify this property at all.
	//
	// During stack set updates, any parameter
	// values overridden for a stack instance are not updated, but retain their
	// overridden value. You can only override the parameter values that are specified
	// in the stack set; to add or delete a parameter itself, use UpdateStackSet to
	// update the stack set template. If you add a parameter to a template, before you
	// can override the parameter value specified in the stack set you must first use
	// UpdateStackSet
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_UpdateStackSet.html)
	// to update all stack instances with the updated template and parameter value
	// specified in the stack set. Once a stack instance has been updated with the new
	// parameter, you can then override the parameter value using UpdateStackInstances.
	ParameterOverrides []*types.Parameter
	// Preferences for how AWS CloudFormation performs this stack set operation.
	OperationPreferences *types.StackSetOperationPreferences
	// [Service-managed permissions] The AWS Organizations accounts for which you want
	// to update parameter values for stack instances. If your update targets OUs, the
	// overridden parameter values only apply to the accounts that are currently in the
	// target OUs and their child OUs. Accounts added to the target OUs and their child
	// OUs in the future won't use the overridden values. You can specify Accounts or
	// DeploymentTargets, but not both.
	DeploymentTargets *types.DeploymentTargets
	// The names of one or more Regions in which you want to update parameter values
	// for stack instances. The overridden parameter values will be applied to all
	// stack instances in the specified accounts and Regions.
	Regions []*string
	// The name or unique ID of the stack set associated with the stack instances.
	StackSetName *string
	// [Self-managed permissions] The names of one or more AWS accounts for which you
	// want to update parameter values for stack instances. The overridden parameter
	// values will be applied to all stack instances in the specified accounts and
	// Regions. You can specify Accounts or DeploymentTargets, but not both.
	Accounts []*string
	// The unique identifier for this stack set operation. The operation ID also
	// functions as an idempotency token, to ensure that AWS CloudFormation performs
	// the stack set operation only once, even if you retry the request multiple times.
	// You might retry stack set operation requests to ensure that AWS CloudFormation
	// successfully received them. If you don't specify an operation ID, the SDK
	// generates one automatically.
	OperationId *string
}

type UpdateStackInstancesOutput struct {
	// The unique identifier for this stack set operation.
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpUpdateStackInstancesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpUpdateStackInstances{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateStackInstances{}, middleware.After)
}

type idempotencyToken_initializeOpUpdateStackInstances struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateStackInstances) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateStackInstances) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateStackInstancesInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateStackInstancesInput ")
	}

	if input.OperationId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.OperationId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateStackInstancesMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpUpdateStackInstances{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateStackInstances(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "UpdateStackInstances",
	}
}
