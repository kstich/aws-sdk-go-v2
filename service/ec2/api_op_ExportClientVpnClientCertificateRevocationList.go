// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Downloads the client certificate revocation list for the specified Client VPN
// endpoint.
func (c *Client) ExportClientVpnClientCertificateRevocationList(ctx context.Context, params *ExportClientVpnClientCertificateRevocationListInput, optFns ...func(*Options)) (*ExportClientVpnClientCertificateRevocationListOutput, error) {
	stack := middleware.NewStack("ExportClientVpnClientCertificateRevocationList", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpExportClientVpnClientCertificateRevocationListMiddlewares(stack)
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
	addOpExportClientVpnClientCertificateRevocationListValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opExportClientVpnClientCertificateRevocationList(options.Region), middleware.Before)
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
			OperationName: "ExportClientVpnClientCertificateRevocationList",
			Err:           err,
		}
	}
	out := result.(*ExportClientVpnClientCertificateRevocationListOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExportClientVpnClientCertificateRevocationListInput struct {
	// The ID of the Client VPN endpoint.
	ClientVpnEndpointId *string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type ExportClientVpnClientCertificateRevocationListOutput struct {
	// The current state of the client certificate revocation list.
	Status *types.ClientCertificateRevocationListStatus
	// Information about the client certificate revocation list.
	CertificateRevocationList *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpExportClientVpnClientCertificateRevocationListMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpExportClientVpnClientCertificateRevocationList{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpExportClientVpnClientCertificateRevocationList{}, middleware.After)
}

func newServiceMetadataMiddleware_opExportClientVpnClientCertificateRevocationList(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ExportClientVpnClientCertificateRevocationList",
	}
}
