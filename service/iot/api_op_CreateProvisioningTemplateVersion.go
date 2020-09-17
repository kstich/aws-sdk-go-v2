// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new version of a fleet provisioning template.
func (c *Client) CreateProvisioningTemplateVersion(ctx context.Context, params *CreateProvisioningTemplateVersionInput, optFns ...func(*Options)) (*CreateProvisioningTemplateVersionOutput, error) {
	stack := middleware.NewStack("CreateProvisioningTemplateVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateProvisioningTemplateVersionMiddlewares(stack)
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
	addOpCreateProvisioningTemplateVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProvisioningTemplateVersion(options.Region), middleware.Before)
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
			OperationName: "CreateProvisioningTemplateVersion",
			Err:           err,
		}
	}
	out := result.(*CreateProvisioningTemplateVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProvisioningTemplateVersionInput struct {
	// Sets a fleet provision template version as the default version.
	SetAsDefault *bool
	// The name of the fleet provisioning template.
	TemplateName *string
	// The JSON formatted contents of the fleet provisioning template.
	TemplateBody *string
}

type CreateProvisioningTemplateVersionOutput struct {
	// True if the fleet provisioning template version is the default version,
	// otherwise false.
	IsDefaultVersion *bool
	// The name of the fleet provisioning template.
	TemplateName *string
	// The version of the fleet provisioning template.
	VersionId *int32
	// The ARN that identifies the provisioning template.
	TemplateArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateProvisioningTemplateVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateProvisioningTemplateVersion{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateProvisioningTemplateVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateProvisioningTemplateVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CreateProvisioningTemplateVersion",
	}
}
