// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a usage plan key and remove the underlying API key from the associated
// usage plan.
func (c *Client) DeleteUsagePlanKey(ctx context.Context, params *DeleteUsagePlanKeyInput, optFns ...func(*Options)) (*DeleteUsagePlanKeyOutput, error) {
	stack := middleware.NewStack("DeleteUsagePlanKey", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteUsagePlanKeyMiddlewares(stack)
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
	addOpDeleteUsagePlanKeyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteUsagePlanKey(options.Region), middleware.Before)
	addResponseErrorWrapper(stack)
	addAcceptHeader(stack)

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
			OperationName: "DeleteUsagePlanKey",
			Err:           err,
		}
	}
	out := result.(*DeleteUsagePlanKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The DELETE request to delete a usage plan key and remove the underlying API key
// from the associated usage plan.
type DeleteUsagePlanKeyInput struct {
	// [Required] The Id of the UsagePlanKey () resource to be deleted.
	KeyId            *string
	Name             *string
	Title            *string
	TemplateSkipList []*string
	// [Required] The Id of the UsagePlan () resource representing the usage plan
	// containing the to-be-deleted UsagePlanKey () resource representing a plan
	// customer.
	UsagePlanId *string
	Template    *bool
}

type DeleteUsagePlanKeyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteUsagePlanKeyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteUsagePlanKey{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteUsagePlanKey{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteUsagePlanKey(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "DeleteUsagePlanKey",
	}
}
