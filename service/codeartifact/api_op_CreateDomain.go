// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a domain. CodeArtifact domains make it easier to manage multiple
// repositories across an organization. You can use a domain to apply permissions
// across many repositories owned by different AWS accounts. An asset is stored
// only once in a domain, even if it's in multiple repositories.  <p>Although you
// can have multiple domains, we recommend a single production domain that contains
// all published artifacts so that your development teams can find and share
// packages. You can use a second pre-production domain to test changes to the
// production domain configuration. </p>
func (c *Client) CreateDomain(ctx context.Context, params *CreateDomainInput, optFns ...func(*Options)) (*CreateDomainOutput, error) {
	stack := middleware.NewStack("CreateDomain", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateDomainMiddlewares(stack)
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
	addOpCreateDomainValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDomain(options.Region), middleware.Before)

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
			OperationName: "CreateDomain",
			Err:           err,
		}
	}
	out := result.(*CreateDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDomainInput struct {
	// The name of the domain to create. All domain names in an AWS Region that are in
	// the same AWS account must be unique. The domain name is used as the prefix in
	// DNS hostnames. Do not use sensitive information in a domain name because it is
	// publicly discoverable.
	Domain *string
	// The encryption key for the domain. This is used to encrypt content stored in a
	// domain. An encryption key can be a key ID, a key Amazon Resource Name (ARN), a
	// key alias, or a key alias ARN. To specify an encryptionKey, your IAM role must
	// have kms:DescribeKey and kms:CreateGrant permissions on the encryption key that
	// is used. For more information, see DescribeKey
	// (https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html#API_DescribeKey_RequestSyntax)
	// in the AWS Key Management Service API Reference and AWS KMS API Permissions
	// Reference
	// (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
	// in the AWS Key Management Service Developer Guide. CodeArtifact supports only
	// symmetric CMKs. Do not associate an asymmetric CMK with your domain. For more
	// information, see Using symmetric and asymmetric keys
	// (https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html)
	// in the AWS Key Management Service Developer Guide.
	EncryptionKey *string
}

type CreateDomainOutput struct {
	// Contains information about the created domain after processing the request.
	Domain *types.DomainDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateDomainMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateDomain{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDomain{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDomain(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "CreateDomain",
	}
}
