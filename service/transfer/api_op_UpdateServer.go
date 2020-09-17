// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the file transfer protocol-enabled server's properties after that server
// has been created.  <p>The <code>UpdateServer</code> call returns the
// <code>ServerId</code> of the server you updated.</p>
func (c *Client) UpdateServer(ctx context.Context, params *UpdateServerInput, optFns ...func(*Options)) (*UpdateServerOutput, error) {
	stack := middleware.NewStack("UpdateServer", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateServerMiddlewares(stack)
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
	addOpUpdateServerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateServer(options.Region), middleware.Before)
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
			OperationName: "UpdateServer",
			Err:           err,
		}
	}
	out := result.(*UpdateServerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateServerInput struct {
	// The virtual private cloud (VPC) endpoint settings that are configured for your
	// file transfer protocol-enabled server. With a VPC endpoint, you can restrict
	// access to your server to resources only within your VPC. To control incoming
	// internet traffic, you will need to associate one or more Elastic IP addresses
	// with your server's endpoint.
	EndpointDetails *types.EndpointDetails
	// The type of endpoint that you want your file transfer protocol-enabled server to
	// connect to. You can choose to connect to the public internet or a VPC endpoint.
	// With a VPC endpoint, you can restrict access to your server and resources only
	// within your VPC.  <note> <p>It is recommended that you use <code>VPC</code> as
	// the <code>EndpointType</code>. With this endpoint type, you have the option to
	// directly associate up to three Elastic IPv4 addresses (BYO IP included) with
	// your server's endpoint and use VPC security groups to restrict traffic by the
	// client's public IP address. This is not possible with <code>EndpointType</code>
	// set to <code>VPC_ENDPOINT</code>.</p> </note>
	EndpointType types.EndpointType
	// An array containing all of the information required to call a customer's
	// authentication API method.
	IdentityProviderDetails *types.IdentityProviderDetails
	// A system-assigned unique identifier for a file transfer protocol-enabled server
	// instance that the user account is assigned to.
	ServerId *string
	// Changes the AWS Identity and Access Management (IAM) role that allows Amazon S3
	// events to be logged in Amazon CloudWatch, turning logging on or off.
	LoggingRole *string
	// The RSA private key as generated by ssh-keygen -N "" -m PEM -f
	// my-new-server-key.  <important> <p>If you aren't planning to migrate existing
	// users from an existing file transfer protocol-enabled server to a new server,
	// don't update the host key. Accidentally changing a server's host key can be
	// disruptive.</p> </important> <p>For more information, see <a
	// href="https://docs.aws.amazon.com/transfer/latest/userguide/edit-server-config.html#configuring-servers-change-host-key">Change
	// the host key for your SFTP-enabled server</a> in the <i>AWS Transfer Family User
	// Guide</i>.</p>
	HostKey *string
	// Specifies the file transfer protocol or protocols over which your file transfer
	// protocol client can connect to your server's endpoint. The available protocols
	// are:  <ul> <li> <p>Secure Shell (SSH) File Transfer Protocol (SFTP): File
	// transfer over SSH</p> </li> <li> <p>File Transfer Protocol Secure (FTPS): File
	// transfer with TLS encryption</p> </li> <li> <p>File Transfer Protocol (FTP):
	// Unencrypted file transfer</p> </li> </ul> <note> <p>If you select
	// <code>FTPS</code>, you must choose a certificate stored in AWS Certificate
	// Manager (ACM) which will be used to identify your server when clients connect to
	// it over FTPS.</p> <p>If <code>Protocol</code> includes either <code>FTP</code>
	// or <code>FTPS</code>, then the <code>EndpointType</code> must be
	// <code>VPC</code> and the <code>IdentityProviderType</code> must be
	// <code>API_GATEWAY</code>.</p> <p>If <code>Protocol</code> includes
	// <code>FTP</code>, then <code>AddressAllocationIds</code> cannot be
	// associated.</p> <p>If <code>Protocol</code> is set only to <code>SFTP</code>,
	// the <code>EndpointType</code> can be set to <code>PUBLIC</code> and the
	// <code>IdentityProviderType</code> can be set to
	// <code>SERVICE_MANAGED</code>.</p> </note>
	Protocols []types.Protocol
	// The Amazon Resource Name (ARN) of the AWS Certificate Manager (ACM) certificate.
	// Required when Protocols is set to FTPS.  <p>To request a new public certificate,
	// see <a
	// href="https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-request-public.html">Request
	// a public certificate</a> in the <i> AWS Certificate Manager User Guide</i>.</p>
	// <p>To import an existing certificate into ACM, see <a
	// href="https://docs.aws.amazon.com/acm/latest/userguide/import-certificate.html">Importing
	// certificates into ACM</a> in the <i> AWS Certificate Manager User Guide</i>.</p>
	// <p>To request a private certificate to use FTPS through private IP addresses,
	// see <a
	// href="https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-request-private.html">Request
	// a private certificate</a> in the <i> AWS Certificate Manager User Guide</i>.</p>
	// <p>Certificates with the following cryptographic algorithms and key sizes are
	// supported:</p> <ul> <li> <p>2048-bit RSA (RSA_2048)</p> </li> <li> <p>4096-bit
	// RSA (RSA_4096)</p> </li> <li> <p>Elliptic Prime Curve 256 bit
	// (EC_prime256v1)</p> </li> <li> <p>Elliptic Prime Curve 384 bit
	// (EC_secp384r1)</p> </li> <li> <p>Elliptic Prime Curve 521 bit (EC_secp521r1)</p>
	// </li> </ul> <note> <p>The certificate must be a valid SSL/TLS X.509 version 3
	// certificate with FQDN or IP address specified and information about the
	// issuer.</p> </note>
	Certificate *string
}

type UpdateServerOutput struct {
	// A system-assigned unique identifier for a file transfer protocol-enabled server
	// that the user account is assigned to.
	ServerId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateServerMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateServer{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateServer{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateServer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "UpdateServer",
	}
}
