// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retreives the CA associated with a group. Returns the public key of the CA.
func (c *Client) GetGroupCertificateAuthority(ctx context.Context, params *GetGroupCertificateAuthorityInput, optFns ...func(*Options)) (*GetGroupCertificateAuthorityOutput, error) {
	stack := middleware.NewStack("GetGroupCertificateAuthority", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetGroupCertificateAuthorityMiddlewares(stack)
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
	addOpGetGroupCertificateAuthorityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetGroupCertificateAuthority(options.Region), middleware.Before)

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
			OperationName: "GetGroupCertificateAuthority",
			Err:           err,
		}
	}
	out := result.(*GetGroupCertificateAuthorityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetGroupCertificateAuthorityInput struct {
	// The ID of the certificate authority.
	CertificateAuthorityId *string
	// The ID of the Greengrass group.
	GroupId *string
}

type GetGroupCertificateAuthorityOutput struct {
	// The ARN of the certificate authority for the group.
	GroupCertificateAuthorityArn *string
	// The PEM encoded certificate for the group.
	PemEncodedCertificate *string
	// The ID of the certificate authority for the group.
	GroupCertificateAuthorityId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetGroupCertificateAuthorityMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetGroupCertificateAuthority{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetGroupCertificateAuthority{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetGroupCertificateAuthority(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "GetGroupCertificateAuthority",
	}
}
