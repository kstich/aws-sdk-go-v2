// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns information about a saved resource, including the last time it was
// backed up, its Amazon Resource Name (ARN), and the AWS service type of the saved
// resource.
func (c *Client) DescribeProtectedResource(ctx context.Context, params *DescribeProtectedResourceInput, optFns ...func(*Options)) (*DescribeProtectedResourceOutput, error) {
	stack := middleware.NewStack("DescribeProtectedResource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeProtectedResourceMiddlewares(stack)
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
	addOpDescribeProtectedResourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeProtectedResource(options.Region), middleware.Before)
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
			OperationName: "DescribeProtectedResource",
			Err:           err,
		}
	}
	out := result.(*DescribeProtectedResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeProtectedResourceInput struct {
	// An Amazon Resource Name (ARN) that uniquely identifies a resource. The format of
	// the ARN depends on the resource type.
	ResourceArn *string
}

type DescribeProtectedResourceOutput struct {
	// An ARN that uniquely identifies a resource. The format of the ARN depends on the
	// resource type.
	ResourceArn *string
	// The date and time that a resource was last backed up, in Unix format and
	// Coordinated Universal Time (UTC). The value of LastBackupTime is accurate to
	// milliseconds. For example, the value 1516925490.087 represents Friday, January
	// 26, 2018 12:11:30.087 AM.
	LastBackupTime *time.Time
	// The type of AWS resource saved as a recovery point; for example, an EBS volume
	// or an Amazon RDS database.
	ResourceType *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeProtectedResourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeProtectedResource{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeProtectedResource{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeProtectedResource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "DescribeProtectedResource",
	}
}
