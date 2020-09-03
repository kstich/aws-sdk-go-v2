// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a Lambda function. To create a function, you need a deployment package
// (https://docs.aws.amazon.com/lambda/latest/dg/deployment-package-v2.html) and an
// execution role
// (https://docs.aws.amazon.com/lambda/latest/dg/intro-permission-model.html#lambda-intro-execution-role).
// The deployment package contains your function code. The execution role grants
// the function permission to use AWS services, such as Amazon CloudWatch Logs for
// log streaming and AWS X-Ray for request tracing.  <p>When you create a function,
// Lambda provisions an instance of the function and its supporting resources. If
// your function connects to a VPC, this process can take a minute or so. During
// this time, you can't invoke or modify the function. The <code>State</code>,
// <code>StateReason</code>, and <code>StateReasonCode</code> fields in the
// response from <a>GetFunctionConfiguration</a> indicate when the function is
// ready to invoke. For more information, see <a
// href="https://docs.aws.amazon.com/lambda/latest/dg/functions-states.html">Function
// States</a>.</p> <p>A function has an unpublished version, and can have published
// versions and aliases. The unpublished version changes when you update your
// function's code and configuration. A published version is a snapshot of your
// function code and configuration that can't be changed. An alias is a named
// resource that maps to a version, and can be changed to map to a different
// version. Use the <code>Publish</code> parameter to create version <code>1</code>
// of your function from its initial configuration.</p> <p>The other parameters let
// you configure version-specific and function-level settings. You can modify
// version-specific settings later with <a>UpdateFunctionConfiguration</a>.
// Function-level settings apply to both the unpublished and published versions of
// the function, and include tags (<a>TagResource</a>) and per-function concurrency
// limits (<a>PutFunctionConcurrency</a>).</p> <p>If another account or an AWS
// service invokes your function, use <a>AddPermission</a> to grant permission by
// creating a resource-based IAM policy. You can grant permissions at the function
// level, on a version, or on an alias.</p> <p>To invoke your function directly,
// use <a>Invoke</a>. To invoke your function in response to events in other AWS
// services, create an event source mapping (<a>CreateEventSourceMapping</a>), or
// configure a function trigger in the other service. For more information, see <a
// href="https://docs.aws.amazon.com/lambda/latest/dg/lambda-invocation.html">Invoking
// Functions</a>.</p>
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
	// For network connectivity to AWS resources in a VPC, specify a list of security
	// groups and subnets in the VPC. When you connect a function to a VPC, it can only
	// access resources and the internet through that VPC. For more information, see
	// VPC Settings
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-vpc.html).
	VpcConfig *types.VpcConfig
	// A description of the function.
	Description *string
	// The amount of memory that your function has access to. Increasing the function's
	// memory also increases its CPU allocation. The default value is 128 MB. The value
	// must be a multiple of 64 MB.
	MemorySize *int32
	// The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt
	// your function's environment variables. If it's not provided, AWS Lambda uses a
	// default service key.
	KMSKeyArn *string
	// A list of function layers
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html) to add
	// to the function's execution environment. Specify each layer by its ARN,
	// including the version.
	Layers []*string
	// Set to true to publish the first version of the function during creation.
	Publish *bool
	// Set Mode to Active to sample and trace a subset of incoming requests with AWS
	// X-Ray.
	TracingConfig *types.TracingConfig
	// The name of the method within your code that Lambda calls to execute your
	// function. The format includes the file name. It can also include namespaces and
	// other qualifiers, depending on the runtime. For more information, see
	// Programming Model
	// (https://docs.aws.amazon.com/lambda/latest/dg/programming-model-v2.html).
	Handler *string
	// The amount of time that Lambda allows a function to run before stopping it. The
	// default is 3 seconds. The maximum allowed value is 900 seconds.
	Timeout *int32
	// The code for the function.
	Code *types.FunctionCode
	// The identifier of the function's runtime
	// (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html).
	Runtime types.Runtime
	// Connection settings for an Amazon EFS file system.
	FileSystemConfigs []*types.FileSystemConfig
	// The Amazon Resource Name (ARN) of the function's execution role.
	Role *string
	// Environment variables that are accessible from function code during execution.
	Environment *types.Environment
	// A list of tags (https://docs.aws.amazon.com/lambda/latest/dg/tagging.html) to
	// apply to the function.
	Tags map[string]*string
	// The name of the Lambda function. Name formats
	//
	//     * Function name -
	// my-function.
	//
	//     * Function ARN -
	// arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	//     * Partial ARN -
	// 123456789012:function:my-function.
	//
	// The length constraint applies only to the
	// full ARN. If you specify only the function name, it is limited to 64 characters
	// in length.
	FunctionName *string
	// A dead letter queue configuration that specifies the queue or topic where Lambda
	// sends asynchronous events when they fail processing. For more information, see
	// Dead Letter Queues
	// (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#dlq).
	DeadLetterConfig *types.DeadLetterConfig
}

// Details about a function's configuration.
type CreateFunctionOutput struct {
	// The name of the function.
	FunctionName *string
	// The latest updated revision of the function or alias.
	RevisionId *string
	// For Lambda@Edge functions, the ARN of the master function.
	MasterArn *string
	// The status of the last update that was performed on the function. This is first
	// set to Successful after function creation completes.
	LastUpdateStatus types.LastUpdateStatus
	// The function's  layers
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
	Layers []*types.Layer
	// Connection settings for an Amazon EFS file system.
	FileSystemConfigs []*types.FileSystemConfig
	// The function's networking configuration.
	VpcConfig *types.VpcConfigResponse
	// The runtime environment for the Lambda function.
	Runtime types.Runtime
	// The SHA256 hash of the function's deployment package.
	CodeSha256 *string
	// The date and time that the function was last updated, in ISO-8601 format
	// (https://www.w3.org/TR/NOTE-datetime) (YYYY-MM-DDThh:mm:ss.sTZD).
	LastModified *string
	// The memory that's allocated to the function.
	MemorySize *int32
	// The version of the Lambda function.
	Version *string
	// The amount of time in seconds that Lambda allows a function to run before
	// stopping it.
	Timeout *int32
	// The size of the function's deployment package, in bytes.
	CodeSize *int64
	// The function's Amazon Resource Name (ARN).
	FunctionArn *string
	// The function's environment variables.
	Environment *types.EnvironmentResponse
	// The reason for the last update that was performed on the function.
	LastUpdateStatusReason *string
	// The function that Lambda calls to begin executing your function.
	Handler *string
	// The function's AWS X-Ray tracing configuration.
	TracingConfig *types.TracingConfigResponse
	// The reason code for the function's current state. When the code is Creating, you
	// can't invoke or modify the function.
	StateReasonCode types.StateReasonCode
	// The reason code for the last update that was performed on the function.
	LastUpdateStatusReasonCode types.LastUpdateStatusReasonCode
	// The function's description.
	Description *string
	// The function's execution role.
	Role *string
	// The KMS key that's used to encrypt the function's environment variables. This
	// key is only returned if you've configured a customer managed CMK.
	KMSKeyArn *string
	// The function's dead letter queue.
	DeadLetterConfig *types.DeadLetterConfig
	// The reason for the function's current state.
	StateReason *string
	// The current state of the function. When the state is Inactive, you can
	// reactivate the function by invoking it.
	State types.State

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
		SigningName:   "lambda",
		OperationName: "CreateFunction",
	}
}
