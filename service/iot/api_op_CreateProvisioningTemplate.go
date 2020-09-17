// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a fleet provisioning template.
func (c *Client) CreateProvisioningTemplate(ctx context.Context, params *CreateProvisioningTemplateInput, optFns ...func(*Options)) (*CreateProvisioningTemplateOutput, error) {
	stack := middleware.NewStack("CreateProvisioningTemplate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateProvisioningTemplateMiddlewares(stack)
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
	addOpCreateProvisioningTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProvisioningTemplate(options.Region), middleware.Before)
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
			OperationName: "CreateProvisioningTemplate",
			Err:           err,
		}
	}
	out := result.(*CreateProvisioningTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProvisioningTemplateInput struct {
	// The description of the fleet provisioning template.
	Description *string
	// True to enable the fleet provisioning template, otherwise false.
	Enabled *bool
	// The role ARN for the role associated with the fleet provisioning template. This
	// IoT role grants permission to provision a device.
	ProvisioningRoleArn *string
	// Metadata which can be used to manage the fleet provisioning template. For URI
	// Request parameters use format: ...key1=value1&key2=value2... For the CLI
	// command-line parameter use format: &&tags "key1=value1&key2=value2..." For the
	// cli-input-json file use format: "tags": "key1=value1&key2=value2..."
	Tags []*types.Tag
	// Creates a pre-provisioning hook template.
	PreProvisioningHook *types.ProvisioningHook
	// The name of the fleet provisioning template.
	TemplateName *string
	// The JSON formatted contents of the fleet provisioning template.
	TemplateBody *string
}

type CreateProvisioningTemplateOutput struct {
	// The name of the fleet provisioning template.
	TemplateName *string
	// The default version of the fleet provisioning template.
	DefaultVersionId *int32
	// The ARN that identifies the provisioning template.
	TemplateArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateProvisioningTemplateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateProvisioningTemplate{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateProvisioningTemplate{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateProvisioningTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CreateProvisioningTemplate",
	}
}
