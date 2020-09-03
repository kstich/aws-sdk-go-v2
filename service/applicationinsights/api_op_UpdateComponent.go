// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationinsights

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the custom component name and/or the list of resources that make up the
// component.
func (c *Client) UpdateComponent(ctx context.Context, params *UpdateComponentInput, optFns ...func(*Options)) (*UpdateComponentOutput, error) {
	stack := middleware.NewStack("UpdateComponent", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateComponentMiddlewares(stack)
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
	addOpUpdateComponentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateComponent(options.Region), middleware.Before)

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
			OperationName: "UpdateComponent",
			Err:           err,
		}
	}
	out := result.(*UpdateComponentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateComponentInput struct {
	// The new name of the component.
	NewComponentName *string
	// The name of the resource group.
	ResourceGroupName *string
	// The name of the component.
	ComponentName *string
	// The list of resource ARNs that belong to the component.
	ResourceList []*string
}

type UpdateComponentOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateComponentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateComponent{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateComponent{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateComponent(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "applicationinsights",
		OperationName: "UpdateComponent",
	}
}
