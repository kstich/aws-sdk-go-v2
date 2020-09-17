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

// Simulate the execution of a Method () in your RestApi () with headers,
// parameters, and an incoming request body.
func (c *Client) TestInvokeMethod(ctx context.Context, params *TestInvokeMethodInput, optFns ...func(*Options)) (*TestInvokeMethodOutput, error) {
	stack := middleware.NewStack("TestInvokeMethod", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpTestInvokeMethodMiddlewares(stack)
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
	addOpTestInvokeMethodValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTestInvokeMethod(options.Region), middleware.Before)
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
			OperationName: "TestInvokeMethod",
			Err:           err,
		}
	}
	out := result.(*TestInvokeMethodOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Make a request to simulate the execution of a Method ().
type TestInvokeMethodInput struct {
	// [Required] The string identifier of the associated RestApi ().
	RestApiId *string
	// [Required] Specifies a test invoke method request's resource ID.
	ResourceId *string
	// The headers as a map from string to list of values to simulate an incoming
	// invocation request.
	MultiValueHeaders map[string][]*string
	// A key-value map of stage variables to simulate an invocation on a deployed Stage
	// ().
	StageVariables map[string]*string
	// The simulated request body of an incoming invocation request.
	Body *string
	// A ClientCertificate () identifier to use in the test invocation. API Gateway
	// will use the certificate when making the HTTPS request to the defined back-end
	// endpoint.
	ClientCertificateId *string
	// [Required] Specifies a test invoke method request's HTTP method.
	HttpMethod *string
	// The URI path, including query string, of the simulated invocation request. Use
	// this to specify path parameters and query string parameters.
	PathWithQueryString *string
	// A key-value map of headers to simulate an incoming invocation request.
	Headers map[string]*string
}

// Represents the response of the test invoke request in the HTTP method. Test API
// using the API Gateway console
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-test-method.html#how-to-test-method-console)
type TestInvokeMethodOutput struct {
	// The headers of the HTTP response.
	Headers map[string]*string
	// The API Gateway execution log for the test invoke request.
	Log *string
	// The HTTP status code.
	Status *int32
	// The body of the HTTP response.
	Body *string
	// The headers of the HTTP response as a map from string to list of values.
	MultiValueHeaders map[string][]*string
	// The execution latency of the test invoke request.
	Latency *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpTestInvokeMethodMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpTestInvokeMethod{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpTestInvokeMethod{}, middleware.After)
}

func newServiceMetadataMiddleware_opTestInvokeMethod(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "TestInvokeMethod",
	}
}
