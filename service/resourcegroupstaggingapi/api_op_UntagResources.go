// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroupstaggingapi

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the specified tags from the specified resources. When you specify a tag
// key, the action removes both that key and its associated value. The operation
// succeeds even if you attempt to remove tags from a resource that were already
// removed. Note the following:
//
//     * To remove tags from a resource, you need the
// necessary permissions for the service that the resource belongs to as well as
// permissions for removing tags. For more information, see this list
// (http://docs.aws.amazon.com/resourcegroupstagging/latest/APIReference/Welcome.html).
//
//
// * You can only tag resources that are located in the specified Region for the
// AWS account.
func (c *Client) UntagResources(ctx context.Context, params *UntagResourcesInput, optFns ...func(*Options)) (*UntagResourcesOutput, error) {
	stack := middleware.NewStack("UntagResources", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUntagResourcesMiddlewares(stack)
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
	addOpUntagResourcesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUntagResources(options.Region), middleware.Before)
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
			OperationName: "UntagResources",
			Err:           err,
		}
	}
	out := result.(*UntagResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UntagResourcesInput struct {
	// A list of the tag keys that you want to remove from the specified resources.
	TagKeys []*string
	// A list of ARNs. An ARN (Amazon Resource Name) uniquely identifies a resource.
	// For more information, see Amazon Resource Names (ARNs) and AWS Service
	// Namespaces
	// (http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference.
	ResourceARNList []*string
}

type UntagResourcesOutput struct {
	// Details of resources that could not be untagged. An error code, status code, and
	// error message are returned for each failed item.
	FailedResourcesMap map[string]*types.FailureInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUntagResourcesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUntagResources{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUntagResources{}, middleware.After)
}

func newServiceMetadataMiddleware_opUntagResources(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "tagging",
		OperationName: "UntagResources",
	}
}
