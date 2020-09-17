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

// Adds a new Model () resource to an existing RestApi () resource.
func (c *Client) CreateModel(ctx context.Context, params *CreateModelInput, optFns ...func(*Options)) (*CreateModelOutput, error) {
	stack := middleware.NewStack("CreateModel", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateModelMiddlewares(stack)
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
	addOpCreateModelValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateModel(options.Region), middleware.Before)
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
			OperationName: "CreateModel",
			Err:           err,
		}
	}
	out := result.(*CreateModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to add a new Model () to an existing RestApi () resource.
type CreateModelInput struct {
	// [Required] The RestApi () identifier under which the Model () will be created.
	RestApiId *string
	// [Required] The name of the model. Must be alphanumeric.
	Name             *string
	TemplateSkipList []*string
	Title            *string
	Template         *bool
	// [Required] The content-type for the model.
	ContentType *string
	// The schema for the model. For application/json models, this should be JSON
	// schema draft 4 (https://tools.ietf.org/html/draft-zyp-json-schema-04) model.
	Schema *string
	// The description of the model.
	Description *string
}

// Represents the data structure of a method's request or response payload. A
// request model defines the data structure of the client-supplied request payload.
// A response model defines the data structure of the response payload returned by
// the back end. Although not required, models are useful for mapping payloads
// between the front end and back end. A model is used for generating an API's SDK,
// validating the input request body, and creating a skeletal mapping template.
// Method (), MethodResponse (), Models and Mappings
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/models-mappings.html)
type CreateModelOutput struct {
	// The identifier for the model resource.
	Id *string
	// The name of the model. Must be an alphanumeric string.
	Name *string
	// The content-type for the model.
	ContentType *string
	// The description of the model.
	Description *string
	// The schema for the model. For application/json models, this should be JSON
	// schema draft 4 (https://tools.ietf.org/html/draft-zyp-json-schema-04) model. Do
	// not include "\*/" characters in the description of any properties because such
	// "\*/" characters may be interpreted as the closing marker for comments in some
	// languages, such as Java or JavaScript, causing the installation of your API's
	// SDK generated by API Gateway to fail.
	Schema *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateModelMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateModel{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateModel{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateModel(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "CreateModel",
	}
}
