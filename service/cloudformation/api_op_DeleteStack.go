// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a specified stack. Once the call completes successfully, stack deletion
// starts. Deleted stacks do not show up in the DescribeStacks () API if the
// deletion has been completed successfully.
func (c *Client) DeleteStack(ctx context.Context, params *DeleteStackInput, optFns ...func(*Options)) (*DeleteStackOutput, error) {
	stack := middleware.NewStack("DeleteStack", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeleteStackMiddlewares(stack)
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
	addOpDeleteStackValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteStack(options.Region), middleware.Before)

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
			OperationName: "DeleteStack",
			Err:           err,
		}
	}
	out := result.(*DeleteStackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for DeleteStack () action.
type DeleteStackInput struct {
	// A unique identifier for this DeleteStack request. Specify this token if you plan
	// to retry requests so that AWS CloudFormation knows that you're not attempting to
	// delete a stack with the same name. You might retry DeleteStack requests to
	// ensure that AWS CloudFormation successfully received them. All events triggered
	// by a given stack operation are assigned the same client request token, which you
	// can use to track operations. For example, if you execute a CreateStack operation
	// with the token token1, then all the StackEvents generated by that operation will
	// have ClientRequestToken set as token1. In the console, stack operations display
	// the client request token on the Events tab. Stack operations that are initiated
	// from the console use the token format Console-StackOperation-ID, which helps you
	// easily identify the stack operation . For example, if you create a stack using
	// the console, each stack event would be assigned the same token in the following
	// format: Console-CreateStack-7f59c3cf-00d2-40c7-b2ff-e75db0987002.
	ClientRequestToken *string
	// The Amazon Resource Name (ARN) of an AWS Identity and Access Management (IAM)
	// role that AWS CloudFormation assumes to delete the stack. AWS CloudFormation
	// uses the role's credentials to make calls on your behalf. If you don't specify a
	// value, AWS CloudFormation uses the role that was previously associated with the
	// stack. If no role is available, AWS CloudFormation uses a temporary session that
	// is generated from your user credentials.
	RoleARN *string
	// For stacks in the DELETE_FAILED state, a list of resource logical IDs that are
	// associated with the resources you want to retain. During deletion, AWS
	// CloudFormation deletes the stack but does not delete the retained resources.
	// Retaining resources is useful when you cannot delete a resource, such as a
	// non-empty S3 bucket, but you want to delete the stack.
	RetainResources []*string
	// The name or the unique stack ID that is associated with the stack.
	StackName *string
}

type DeleteStackOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeleteStackMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeleteStack{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteStack{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteStack(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "DeleteStack",
	}
}
