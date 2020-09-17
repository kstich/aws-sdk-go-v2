// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a signing certificate associated with the specified IAM user. If you do
// not specify a user name, IAM determines the user name implicitly based on the
// AWS access key ID signing the request. This operation works for access keys
// under the AWS account. Consequently, you can use this operation to manage AWS
// account root user credentials even if the AWS account has no associated IAM
// users.
func (c *Client) DeleteSigningCertificate(ctx context.Context, params *DeleteSigningCertificateInput, optFns ...func(*Options)) (*DeleteSigningCertificateOutput, error) {
	stack := middleware.NewStack("DeleteSigningCertificate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeleteSigningCertificateMiddlewares(stack)
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
	addOpDeleteSigningCertificateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSigningCertificate(options.Region), middleware.Before)
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
			OperationName: "DeleteSigningCertificate",
			Err:           err,
		}
	}
	out := result.(*DeleteSigningCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteSigningCertificateInput struct {
	// The name of the user the signing certificate belongs to. This parameter allows
	// (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of
	// characters consisting of upper and lowercase alphanumeric characters with no
	// spaces. You can also include any of the following characters: _+=,.@-
	UserName *string
	// The ID of the signing certificate to delete. The format of this parameter, as
	// described by its regex (http://wikipedia.org/wiki/regex) pattern, is a string of
	// characters that can be upper- or lower-cased letters or digits.
	CertificateId *string
}

type DeleteSigningCertificateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeleteSigningCertificateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeleteSigningCertificate{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteSigningCertificate{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteSigningCertificate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "DeleteSigningCertificate",
	}
}
