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

// Describe an existing Authorizers () resource. AWS CLI
// (https://docs.aws.amazon.com/cli/latest/reference/apigateway/get-authorizers.html)
func (c *Client) GetAuthorizers(ctx context.Context, params *GetAuthorizersInput, optFns ...func(*Options)) (*GetAuthorizersOutput, error) {
	stack := middleware.NewStack("GetAuthorizers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetAuthorizersMiddlewares(stack)
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
	addOpGetAuthorizersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAuthorizers(options.Region), middleware.Before)
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
			OperationName: "GetAuthorizers",
			Err:           err,
		}
	}
	out := result.(*GetAuthorizersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to describe an existing Authorizers () resource.
type GetAuthorizersInput struct {
	TemplateSkipList []*string
	Title            *string
	Template         *bool
	// The maximum number of returned results per page. The default value is 25 and the
	// maximum value is 500.
	Limit *int32
	Name  *string
	// The current pagination position in the paged result set.
	Position *string
	// [Required] The string identifier of the associated RestApi ().
	RestApiId *string
}

// Represents a collection of Authorizer () resources. Use Lambda Function as
// Authorizer
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-use-lambda-authorizer.html)Use
// Cognito User Pool as Authorizer
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-integrate-with-cognito.html)
type GetAuthorizersOutput struct {
	// The current page of elements from this collection.
	Items []*types.Authorizer
	// The current pagination position in the paged result set.
	Position *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetAuthorizersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetAuthorizers{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAuthorizers{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetAuthorizers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetAuthorizers",
	}
}
