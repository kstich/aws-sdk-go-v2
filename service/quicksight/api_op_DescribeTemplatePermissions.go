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

// Describes read and write permissions on a template.
func (c *Client) DescribeTemplatePermissions(ctx context.Context, params *DescribeTemplatePermissionsInput, optFns ...func(*Options)) (*DescribeTemplatePermissionsOutput, error) {
	stack := middleware.NewStack("DescribeTemplatePermissions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeTemplatePermissionsMiddlewares(stack)
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
	addOpDescribeTemplatePermissionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTemplatePermissions(options.Region), middleware.Before)
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
			OperationName: "DescribeTemplatePermissions",
			Err:           err,
		}
	}
	out := result.(*DescribeTemplatePermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTemplatePermissionsInput struct {
	// The ID for the template.
	TemplateId *string
	// The ID of the AWS account that contains the template that you're describing.
	AwsAccountId *string
}

type DescribeTemplatePermissionsOutput struct {
	// A list of resource permissions to be set on the template.
	Permissions []*types.ResourcePermission
	// The ID for the template.
	TemplateId *string
	// The Amazon Resource Name (ARN) of the template.
	TemplateArn *string
	// The AWS request ID for this operation.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeTemplatePermissionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeTemplatePermissions{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeTemplatePermissions{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeTemplatePermissions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DescribeTemplatePermissions",
	}
}
