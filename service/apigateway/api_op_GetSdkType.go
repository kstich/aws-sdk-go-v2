// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

func (c *Client) GetSdkType(ctx context.Context, params *GetSdkTypeInput, optFns ...func(*Options)) (*GetSdkTypeOutput, error) {
	stack := middleware.NewStack("GetSdkType", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetSdkTypeMiddlewares(stack)
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
	addOpGetSdkTypeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSdkType(options.Region), middleware.Before)
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
			OperationName: "GetSdkType",
			Err:           err,
		}
	}
	out := result.(*GetSdkTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Get an SdkType () instance.
type GetSdkTypeInput struct {
	Template         *bool
	TemplateSkipList []*string
	// [Required] The identifier of the queried SdkType () instance.
	Id    *string
	Title *string
	Name  *string
}

// A type of SDK that API Gateway can generate.
type GetSdkTypeOutput struct {
	// The user-friendly name of an SdkType () instance.
	FriendlyName *string
	// The description of an SdkType ().
	Description *string
	// A list of configuration properties of an SdkType ().
	ConfigurationProperties []*types.SdkConfigurationProperty
	// The identifier of an SdkType () instance.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetSdkTypeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetSdkType{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSdkType{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetSdkType(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetSdkType",
	}
}
