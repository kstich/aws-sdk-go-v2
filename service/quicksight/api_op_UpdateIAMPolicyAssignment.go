// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates an existing IAM policy assignment. This operation updates only the
// optional parameter or parameters that are specified in the request.
func (c *Client) UpdateIAMPolicyAssignment(ctx context.Context, params *UpdateIAMPolicyAssignmentInput, optFns ...func(*Options)) (*UpdateIAMPolicyAssignmentOutput, error) {
	stack := middleware.NewStack("UpdateIAMPolicyAssignment", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateIAMPolicyAssignmentMiddlewares(stack)
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
	addOpUpdateIAMPolicyAssignmentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateIAMPolicyAssignment(options.Region), middleware.Before)
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
			OperationName: "UpdateIAMPolicyAssignment",
			Err:           err,
		}
	}
	out := result.(*UpdateIAMPolicyAssignmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateIAMPolicyAssignmentInput struct {
	// The ARN for the IAM policy to apply to the QuickSight users and groups specified
	// in this assignment.
	PolicyArn *string
	// The status of the assignment. Possible values are as follows:
	//
	//     * ENABLED -
	// Anything specified in this assignment is used when creating the data source.
	//
	//
	// * DISABLED - This assignment isn't used when creating the data source.
	//
	//     *
	// DRAFT - This assignment is an unfinished draft and isn't used when creating the
	// data source.
	AssignmentStatus types.AssignmentStatus
	// The name of the assignment. This name must be unique within an AWS account.
	AssignmentName *string
	// The QuickSight users, groups, or both that you want to assign the policy to.
	Identities map[string][]*string
	// The ID of the AWS account that contains the IAM policy assignment.
	AwsAccountId *string
	// The namespace of the assignment.
	Namespace *string
}

type UpdateIAMPolicyAssignmentOutput struct {
	// The ARN for the IAM policy applied to the QuickSight users and groups specified
	// in this assignment.
	PolicyArn *string
	// The status of the assignment. Possible values are as follows:
	//
	//     * ENABLED -
	// Anything specified in this assignment is used when creating the data source.
	//
	//
	// * DISABLED - This assignment isn't used when creating the data source.
	//
	//     *
	// DRAFT - This assignment is an unfinished draft and isn't used when creating the
	// data source.
	AssignmentStatus types.AssignmentStatus
	// The name of the assignment.
	AssignmentName *string
	// The QuickSight users, groups, or both that the IAM policy is assigned to.
	Identities map[string][]*string
	// The AWS request ID for this operation.
	RequestId *string
	// The ID of the assignment.
	AssignmentId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateIAMPolicyAssignmentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateIAMPolicyAssignment{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateIAMPolicyAssignment{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateIAMPolicyAssignment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "UpdateIAMPolicyAssignment",
	}
}
