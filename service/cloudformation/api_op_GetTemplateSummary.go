// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about a new or existing template. The GetTemplateSummary
// action is useful for viewing parameter information, such as default parameter
// values and parameter types, before you create or update a stack or stack set.
// You can use the GetTemplateSummary action when you submit a template, or you can
// get template information for a stack set, or a running or deleted stack. For
// deleted stacks, GetTemplateSummary returns the template information for up to 90
// days after the stack has been deleted. If the template does not exist, a
// ValidationError is returned.
func (c *Client) GetTemplateSummary(ctx context.Context, params *GetTemplateSummaryInput, optFns ...func(*Options)) (*GetTemplateSummaryOutput, error) {
	stack := middleware.NewStack("GetTemplateSummary", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpGetTemplateSummaryMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetTemplateSummary(options.Region), middleware.Before)
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
			OperationName: "GetTemplateSummary",
			Err:           err,
		}
	}
	out := result.(*GetTemplateSummaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the GetTemplateSummary () action.
type GetTemplateSummaryInput struct {
	// Structure containing the template body with a minimum length of 1 byte and a
	// maximum length of 51,200 bytes. For more information about templates, see
	// Template Anatomy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html)
	// in the AWS CloudFormation User Guide. Conditional: You must specify only one of
	// the following parameters: StackName, StackSetName, TemplateBody, or TemplateURL.
	TemplateBody *string
	// Location of file containing the template body. The URL must point to a template
	// (max size: 460,800 bytes) that is located in an Amazon S3 bucket. For more
	// information about templates, see Template Anatomy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html)
	// in the AWS CloudFormation User Guide. Conditional: You must specify only one of
	// the following parameters: StackName, StackSetName, TemplateBody, or TemplateURL.
	TemplateURL *string
	// The name or unique ID of the stack set from which the stack was created.
	// Conditional: You must specify only one of the following parameters: StackName,
	// StackSetName, TemplateBody, or TemplateURL.
	StackSetName *string
	// The name or the stack ID that is associated with the stack, which are not always
	// interchangeable. For running stacks, you can specify either the stack's name or
	// its unique stack ID. For deleted stack, you must specify the unique stack ID.
	// Conditional: You must specify only one of the following parameters: StackName,
	// StackSetName, TemplateBody, or TemplateURL.
	StackName *string
}

// The output for the GetTemplateSummary () action.
type GetTemplateSummaryOutput struct {
	// A list of resource identifier summaries that describe the target resources of an
	// import operation and the properties you can provide during the import to
	// identify the target resources. For example, BucketName is a possible identifier
	// property for an AWS::S3::Bucket resource.
	ResourceIdentifierSummaries []*types.ResourceIdentifierSummary
	// The AWS template format version, which identifies the capabilities of the
	// template.
	Version *string
	// The value that is defined for the Metadata property of the template.
	Metadata *string
	// The value that is defined in the Description property of the template.
	Description *string
	// A list of the transforms that are declared in the template.
	DeclaredTransforms []*string
	// A list of all the template resource types that are defined in the template, such
	// as AWS::EC2::Instance, AWS::Dynamo::Table, and Custom::MyCustomInstance.
	ResourceTypes []*string
	// The capabilities found within the template. If your template contains IAM
	// resources, you must specify the CAPABILITY_IAM or CAPABILITY_NAMED_IAM value for
	// this parameter when you use the CreateStack () or UpdateStack () actions with
	// your template; otherwise, those actions return an InsufficientCapabilities
	// error. For more information, see Acknowledging IAM Resources in AWS
	// CloudFormation Templates
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#capabilities).
	Capabilities []types.Capability
	// The list of resources that generated the values in the Capabilities response
	// element.
	CapabilitiesReason *string
	// A list of parameter declarations that describe various properties for each
	// parameter.
	Parameters []*types.ParameterDeclaration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpGetTemplateSummaryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpGetTemplateSummary{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpGetTemplateSummary{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetTemplateSummary(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "GetTemplateSummary",
	}
}
