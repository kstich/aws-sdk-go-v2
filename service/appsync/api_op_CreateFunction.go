// Code generated by smithy-go-codegen DO NOT EDIT.

package appsync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a Function object. A function is a reusable entity. Multiple functions
// can be used to compose the resolver logic.
func (c *Client) CreateFunction(ctx context.Context, params *CreateFunctionInput, optFns ...func(*Options)) (*CreateFunctionOutput, error) {
	stack := middleware.NewStack("CreateFunction", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateFunctionMiddlewares(stack)
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
	addOpCreateFunctionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFunction(options.Region), middleware.Before)
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
			OperationName: "CreateFunction",
			Err:           err,
		}
	}
	out := result.(*CreateFunctionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFunctionInput struct {
	// The Function response mapping template.
	ResponseMappingTemplate *string
	// The version of the request mapping template. Currently the supported value is
	// 2018-05-29.
	FunctionVersion *string
	// The FunctionDataSource name.
	DataSourceName *string
	// The Function name. The function name does not have to be unique.
	Name *string
	// The Function description.
	Description *string
	// The GraphQL API ID.
	ApiId *string
	// The Function request mapping template. Functions support only the 2018-05-29
	// version of the request mapping template.
	RequestMappingTemplate *string
}

type CreateFunctionOutput struct {
	// The Function object.
	FunctionConfiguration *types.FunctionConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateFunctionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateFunction{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateFunction{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateFunction(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appsync",
		OperationName: "CreateFunction",
	}
}
