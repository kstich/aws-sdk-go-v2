// Code generated by smithy-go-codegen DO NOT EDIT.

package serverlessapplicationrepository

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an application version.
func (c *Client) CreateApplicationVersion(ctx context.Context, params *CreateApplicationVersionInput, optFns ...func(*Options)) (*CreateApplicationVersionOutput, error) {
	stack := middleware.NewStack("CreateApplicationVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateApplicationVersionMiddlewares(stack)
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
	addOpCreateApplicationVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateApplicationVersion(options.Region), middleware.Before)
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
			OperationName: "CreateApplicationVersion",
			Err:           err,
		}
	}
	out := result.(*CreateApplicationVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateApplicationVersionInput struct {
	// The semantic version of the new version.
	SemanticVersion *string
	// A link to the packaged AWS SAM template of your application.
	TemplateUrl *string
	// A link to the S3 object that contains the ZIP archive of the source code for
	// this version of your application.Maximum size 50 MB
	SourceCodeArchiveUrl *string
	// A link to a public repository for the source code of your application, for
	// example the URL of a specific GitHub commit.
	SourceCodeUrl *string
	// The Amazon Resource Name (ARN) of the application.
	ApplicationId *string
	// The raw packaged AWS SAM template of your application.
	TemplateBody *string
}

type CreateApplicationVersionOutput struct {
	// A link to the S3 object that contains the ZIP archive of the source code for
	// this version of your application.Maximum size 50 MB
	SourceCodeArchiveUrl *string
	// A link to the packaged AWS SAM template of your application.
	TemplateUrl *string
	// An array of parameter types supported by the application.
	ParameterDefinitions []*types.ParameterDefinition
	// A link to a public repository for the source code of your application, for
	// example the URL of a specific GitHub commit.
	SourceCodeUrl *string
	// The semantic version of the application: https://semver.org/
	// (https://semver.org/)
	SemanticVersion *string
	// The date and time this resource was created.
	CreationTime *string
	// Whether all of the AWS resources contained in this application are supported in
	// the region in which it is being retrieved.
	ResourcesSupported *bool
	// The application Amazon Resource Name (ARN).
	ApplicationId *string
	// A list of values that you must specify before you can deploy certain
	// applications. Some applications might include resources that can affect
	// permissions in your AWS account, for example, by creating new AWS Identity and
	// Access Management (IAM) users. For those applications, you must explicitly
	// acknowledge their capabilities by specifying this parameter.The only valid
	// values are CAPABILITY_IAM, CAPABILITY_NAMED_IAM, CAPABILITY_RESOURCE_POLICY, and
	// CAPABILITY_AUTO_EXPAND.The following resources require you to specify
	// CAPABILITY_IAM or CAPABILITY_NAMED_IAM: AWS::IAM::Group
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html),
	// AWS::IAM::InstanceProfile
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html),
	// AWS::IAM::Policy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html),
	// and AWS::IAM::Role
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html).
	// If the application contains IAM resources, you can specify either CAPABILITY_IAM
	// or CAPABILITY_NAMED_IAM. If the application contains IAM resources with custom
	// names, you must specify CAPABILITY_NAMED_IAM.The following resources require you
	// to specify CAPABILITY_RESOURCE_POLICY: AWS::Lambda::Permission
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html),
	// AWS::IAM:Policy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html),
	// AWS::ApplicationAutoScaling::ScalingPolicy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html),
	// AWS::S3::BucketPolicy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-policy.html),
	// AWS::SQS::QueuePolicy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-policy.html),
	// and AWS::SNS::TopicPolicy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html).Applications
	// that contain one or more nested applications require you to specify
	// CAPABILITY_AUTO_EXPAND.If your application template contains any of the above
	// resources, we recommend that you review all permissions associated with the
	// application before deploying. If you don't specify this parameter for an
	// application that requires capabilities, the call will fail.
	RequiredCapabilities []types.Capability

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateApplicationVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateApplicationVersion{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateApplicationVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateApplicationVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "serverlessrepo",
		OperationName: "CreateApplicationVersion",
	}
}
