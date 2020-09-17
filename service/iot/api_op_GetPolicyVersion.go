// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

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

// Gets information about the specified policy version.
func (c *Client) GetPolicyVersion(ctx context.Context, params *GetPolicyVersionInput, optFns ...func(*Options)) (*GetPolicyVersionOutput, error) {
	stack := middleware.NewStack("GetPolicyVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetPolicyVersionMiddlewares(stack)
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
	addOpGetPolicyVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetPolicyVersion(options.Region), middleware.Before)
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
			OperationName: "GetPolicyVersion",
			Err:           err,
		}
	}
	out := result.(*GetPolicyVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the GetPolicyVersion operation.
type GetPolicyVersionInput struct {
	// The policy version ID.
	PolicyVersionId *string
	// The name of the policy.
	PolicyName *string
}

// The output from the GetPolicyVersion operation.
type GetPolicyVersionOutput struct {
	// The policy version ID.
	PolicyVersionId *string
	// The policy ARN.
	PolicyArn *string
	// The policy name.
	PolicyName *string
	// Specifies whether the policy version is the default.
	IsDefaultVersion *bool
	// The date the policy was last modified.
	LastModifiedDate *time.Time
	// The generation ID of the policy version.
	GenerationId *string
	// The JSON document that describes the policy.
	PolicyDocument *string
	// The date the policy was created.
	CreationDate *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetPolicyVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetPolicyVersion{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetPolicyVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetPolicyVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "GetPolicyVersion",
	}
}
