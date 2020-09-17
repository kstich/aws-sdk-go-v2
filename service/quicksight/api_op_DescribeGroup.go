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

// Returns an Amazon QuickSight group's description and Amazon Resource Name (ARN).
func (c *Client) DescribeGroup(ctx context.Context, params *DescribeGroupInput, optFns ...func(*Options)) (*DescribeGroupOutput, error) {
	stack := middleware.NewStack("DescribeGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeGroupMiddlewares(stack)
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
	addOpDescribeGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeGroup(options.Region), middleware.Before)
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
			OperationName: "DescribeGroup",
			Err:           err,
		}
	}
	out := result.(*DescribeGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeGroupInput struct {
	// The name of the group that you want to describe.
	GroupName *string
	// The namespace. Currently, you should set this to default.
	Namespace *string
	// The ID for the AWS account that the group is in. Currently, you use the ID for
	// the AWS account that contains your Amazon QuickSight account.
	AwsAccountId *string
}

type DescribeGroupOutput struct {
	// The name of the group.
	Group *types.Group
	// The AWS request ID for this operation.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeGroup{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DescribeGroup",
	}
}
