// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This method takes a user pool ID, and returns the signing certificate.
func (c *Client) GetSigningCertificate(ctx context.Context, params *GetSigningCertificateInput, optFns ...func(*Options)) (*GetSigningCertificateOutput, error) {
	stack := middleware.NewStack("GetSigningCertificate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetSigningCertificateMiddlewares(stack)
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
	addOpGetSigningCertificateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSigningCertificate(options.Region), middleware.Before)
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
			OperationName: "GetSigningCertificate",
			Err:           err,
		}
	}
	out := result.(*GetSigningCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to get a signing certificate from Cognito.
type GetSigningCertificateInput struct {
	// The user pool ID.
	UserPoolId *string
}

// Response from Cognito for a signing certificate request.
type GetSigningCertificateOutput struct {
	// The signing certificate.
	Certificate *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetSigningCertificateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetSigningCertificate{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSigningCertificate{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetSigningCertificate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "GetSigningCertificate",
	}
}
