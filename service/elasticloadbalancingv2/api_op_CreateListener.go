// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a listener for the specified Application Load Balancer or Network Load
// Balancer. To update a listener, use ModifyListener (). When you are finished
// with a listener, you can delete it using DeleteListener (). If you are finished
// with both the listener and the load balancer, you can delete them both using
// DeleteLoadBalancer (). This operation is idempotent, which means that it
// completes at most one time. If you attempt to create multiple listeners with the
// same settings, each call succeeds. For more information, see Listeners for Your
// Application Load Balancers
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html)
// in the Application Load Balancers Guide and Listeners for Your Network Load
// Balancers
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-listeners.html)
// in the Network Load Balancers Guide.
func (c *Client) CreateListener(ctx context.Context, params *CreateListenerInput, optFns ...func(*Options)) (*CreateListenerOutput, error) {
	stack := middleware.NewStack("CreateListener", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateListenerMiddlewares(stack)
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
	addOpCreateListenerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateListener(options.Region), middleware.Before)
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
			OperationName: "CreateListener",
			Err:           err,
		}
	}
	out := result.(*CreateListenerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateListenerInput struct {
	// The Amazon Resource Name (ARN) of the load balancer.
	LoadBalancerArn *string
	// [HTTPS and TLS listeners] The security policy that defines which protocols and
	// ciphers are supported. The following are the possible values:
	//
	//     *
	// ELBSecurityPolicy-2016-08
	//
	//     * ELBSecurityPolicy-TLS-1-0-2015-04
	//
	//     *
	// ELBSecurityPolicy-TLS-1-1-2017-01
	//
	//     * ELBSecurityPolicy-TLS-1-2-2017-01
	//
	//
	// * ELBSecurityPolicy-TLS-1-2-Ext-2018-06
	//
	//     * ELBSecurityPolicy-FS-2018-06
	//
	//
	// * ELBSecurityPolicy-FS-1-1-2019-08
	//
	//     * ELBSecurityPolicy-FS-1-2-2019-08
	//
	//
	// * ELBSecurityPolicy-FS-1-2-Res-2019-08
	//
	// For more information, see Security
	// Policies
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html#describe-ssl-policies)
	// in the Application Load Balancers Guide and Security Policies
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/network/create-tls-listener.html#describe-ssl-policies)
	// in the Network Load Balancers Guide.
	SslPolicy *string
	// The port on which the load balancer is listening.
	Port *int32
	// The actions for the default rule. The rule must include one forward action or
	// one or more fixed-response actions. If the action type is forward, you specify
	// one or more target groups. The protocol of the target group must be HTTP or
	// HTTPS for an Application Load Balancer. The protocol of the target group must be
	// TCP, TLS, UDP, or TCP_UDP for a Network Load Balancer. [HTTPS listeners] If the
	// action type is authenticate-oidc, you authenticate users through an identity
	// provider that is OpenID Connect (OIDC) compliant. [HTTPS listeners] If the
	// action type is authenticate-cognito, you authenticate users through the user
	// pools supported by Amazon Cognito. [Application Load Balancer] If the action
	// type is redirect, you redirect specified client requests from one URL to
	// another. [Application Load Balancer] If the action type is fixed-response, you
	// drop specified client requests and return a custom HTTP response.
	DefaultActions []*types.Action
	// [TLS listeners] The name of the Application-Layer Protocol Negotiation (ALPN)
	// policy. You can specify one policy name. The following are the possible
	// values:
	//
	//     * HTTP1Only
	//
	//     * HTTP2Only
	//
	//     * HTTP2Optional
	//
	//     *
	// HTTP2Preferred
	//
	//     * None
	//
	// For more information, see ALPN Policies
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/network/create-tls-listener.html#alpn-policies)
	// in the Network Load Balancers Guide.
	AlpnPolicy []*string
	// The protocol for connections from clients to the load balancer. For Application
	// Load Balancers, the supported protocols are HTTP and HTTPS. For Network Load
	// Balancers, the supported protocols are TCP, TLS, UDP, and TCP_UDP.
	Protocol types.ProtocolEnum
	// [HTTPS and TLS listeners] The default certificate for the listener. You must
	// provide exactly one certificate. Set CertificateArn to the certificate ARN but
	// do not set IsDefault. To create a certificate list for the listener, use
	// AddListenerCertificates ().
	Certificates []*types.Certificate
}

type CreateListenerOutput struct {
	// Information about the listener.
	Listeners []*types.Listener

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateListenerMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateListener{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateListener{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateListener(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "CreateListener",
	}
}
