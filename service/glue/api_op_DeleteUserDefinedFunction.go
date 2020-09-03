// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an existing function definition from the Data Catalog.
func (c *Client) DeleteUserDefinedFunction(ctx context.Context, params *DeleteUserDefinedFunctionInput, optFns ...func(*Options)) (*DeleteUserDefinedFunctionOutput, error) {
	stack := middleware.NewStack("DeleteUserDefinedFunction", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteUserDefinedFunctionMiddlewares(stack)
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
	addOpDeleteUserDefinedFunctionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteUserDefinedFunction(options.Region), middleware.Before)

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
			OperationName: "DeleteUserDefinedFunction",
			Err:           err,
		}
	}
	out := result.(*DeleteUserDefinedFunctionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteUserDefinedFunctionInput struct {
	// The ID of the Data Catalog where the function to be deleted is located. If none
	// is supplied, the AWS account ID is used by default.
	CatalogId *string
	// The name of the function definition to be deleted.
	FunctionName *string
	// The name of the catalog database where the function is located.
	DatabaseName *string
}

type DeleteUserDefinedFunctionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteUserDefinedFunctionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteUserDefinedFunction{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteUserDefinedFunction{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteUserDefinedFunction(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "DeleteUserDefinedFunction",
	}
}
