// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deploys conformance packs across member accounts in an AWS Organization. Only a
// master account and a delegated administrator can call this API. When calling
// this API with a delegated administrator, you must ensure AWS Organizations
// ListDelegatedAdministrator permissions are added. This API enables organization
// service access for config-multiaccountsetup.amazonaws.com through the
// EnableAWSServiceAccess action and creates a service linked role
// AWSServiceRoleForConfigMultiAccountSetup in the master or delegated
// administrator account of your organization. The service linked role is created
// only when the role does not exist in the caller account. To use this API with
// delegated administrator, register a delegated administrator by calling AWS
// Organization register-delegate-admin for config-multiaccountsetup.amazonaws.com.
// <note> <p>Prerequisite: Ensure you call <code>EnableAllFeatures</code> API to
// enable all features in an organization.</p> <p>You must specify either the
// <code>TemplateS3Uri</code> or the <code>TemplateBody</code> parameter, but not
// both. If you provide both AWS Config uses the <code>TemplateS3Uri</code>
// parameter and ignores the <code>TemplateBody</code> parameter.</p> <p>AWS Config
// sets the state of a conformance pack to CREATE_IN_PROGRESS and
// UPDATE_IN_PROGRESS until the conformance pack is created or updated. You cannot
// update a conformance pack while it is in this state.</p> <p>You can create 6
// conformance packs with 25 AWS Config rules in each pack and 3 delegated
// administrator per organization. </p> </note>
func (c *Client) PutOrganizationConformancePack(ctx context.Context, params *PutOrganizationConformancePackInput, optFns ...func(*Options)) (*PutOrganizationConformancePackOutput, error) {
	stack := middleware.NewStack("PutOrganizationConformancePack", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutOrganizationConformancePackMiddlewares(stack)
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
	addOpPutOrganizationConformancePackValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutOrganizationConformancePack(options.Region), middleware.Before)
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
			OperationName: "PutOrganizationConformancePack",
			Err:           err,
		}
	}
	out := result.(*PutOrganizationConformancePackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutOrganizationConformancePackInput struct {
	// Location of an Amazon S3 bucket where AWS Config can deliver evaluation results.
	// AWS Config stores intermediate files while processing conformance pack template.
	// The delivery bucket name should start with awsconfigconforms. For example:
	// "Resource": "arn:aws:s3:::your_bucket_name/*". For more information, see
	// Permissions for cross account bucket access
	// (https://docs.aws.amazon.com/config/latest/developerguide/conformance-pack-organization-apis.html).
	DeliveryS3Bucket *string
	// A string containing full conformance pack template body. Structure containing
	// the template body with a minimum length of 1 byte and a maximum length of 51,200
	// bytes.
	TemplateBody *string
	// Location of file containing the template body. The uri must point to the
	// conformance pack template (max size: 300 KB). You must have access to read
	// Amazon S3 bucket.
	TemplateS3Uri *string
	// Name of the organization conformance pack you want to create.
	OrganizationConformancePackName *string
	// A list of ConformancePackInputParameter objects.
	ConformancePackInputParameters []*types.ConformancePackInputParameter
	// The prefix for the Amazon S3 bucket.
	DeliveryS3KeyPrefix *string
	// A list of AWS accounts to be excluded from an organization conformance pack
	// while deploying a conformance pack.
	ExcludedAccounts []*string
}

type PutOrganizationConformancePackOutput struct {
	// ARN of the organization conformance pack.
	OrganizationConformancePackArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutOrganizationConformancePackMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutOrganizationConformancePack{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutOrganizationConformancePack{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutOrganizationConformancePack(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "PutOrganizationConformancePack",
	}
}
