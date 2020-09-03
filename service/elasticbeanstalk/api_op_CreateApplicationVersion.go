// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an application version for the specified application. You can create an
// application version from a source bundle in Amazon S3, a commit in AWS
// CodeCommit, or the output of an AWS CodeBuild build as follows: Specify a commit
// in an AWS CodeCommit repository with SourceBuildInformation. Specify a build in
// an AWS CodeBuild with SourceBuildInformation and BuildConfiguration. Specify a
// source bundle in S3 with SourceBundle Omit both SourceBuildInformation and
// SourceBundle to use the default sample application. After you create an
// application version with a specified Amazon S3 bucket and key location, you
// can't change that Amazon S3 location. If you change the Amazon S3 location, you
// receive an exception when you attempt to launch an environment from the
// application version.
func (c *Client) CreateApplicationVersion(ctx context.Context, params *CreateApplicationVersionInput, optFns ...func(*Options)) (*CreateApplicationVersionOutput, error) {
	stack := middleware.NewStack("CreateApplicationVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateApplicationVersionMiddlewares(stack)
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

//
type CreateApplicationVersionInput struct {
	// A description of this application version.
	Description *string
	// Set to true to create an application with the specified name if it doesn't
	// already exist.
	AutoCreateApplication *bool
	// Settings for an AWS CodeBuild build.
	BuildConfiguration *types.BuildConfiguration
	// Pre-processes and validates the environment manifest (env.yaml) and
	// configuration files (*.config files in the .ebextensions folder) in the source
	// bundle. Validating configuration files can identify issues prior to deploying
	// the application version to an environment. You must turn processing on for
	// application versions that you create using AWS CodeBuild or AWS CodeCommit. For
	// application versions built from a source bundle in Amazon S3, processing is
	// optional. The Process option validates Elastic Beanstalk configuration files. It
	// doesn't validate your application's configuration files, like proxy server or
	// Docker configuration.
	Process *bool
	// Specifies the tags applied to the application version. Elastic Beanstalk applies
	// these tags only to the application version. Environments that use the
	// application version don't inherit the tags.
	Tags []*types.Tag
	// The name of the application. If no application is found with this name, and
	// AutoCreateApplication is false, returns an InvalidParameterValue error.
	ApplicationName *string
	// Specify a commit in an AWS CodeCommit Git repository to use as the source code
	// for the application version.
	SourceBuildInformation *types.SourceBuildInformation
	// The Amazon S3 bucket and key that identify the location of the source bundle for
	// this version. The Amazon S3 bucket must be in the same region as the
	// environment. Specify a source bundle in S3 or a commit in an AWS CodeCommit
	// repository (with SourceBuildInformation), but not both. If neither SourceBundle
	// nor SourceBuildInformation are provided, Elastic Beanstalk uses a sample
	// application.
	SourceBundle *types.S3Location
	// A label identifying this version. Constraint: Must be unique per application. If
	// an application version already exists with this label for the specified
	// application, AWS Elastic Beanstalk returns an InvalidParameterValue error.
	VersionLabel *string
}

// Result message wrapping a single description of an application version.
type CreateApplicationVersionOutput struct {
	// The ApplicationVersionDescription () of the application version.
	ApplicationVersion *types.ApplicationVersionDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateApplicationVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateApplicationVersion{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateApplicationVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateApplicationVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "CreateApplicationVersion",
	}
}
