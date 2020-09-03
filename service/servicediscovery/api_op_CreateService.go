// Code generated by smithy-go-codegen DO NOT EDIT.

package servicediscovery

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a service, which defines the configuration for the following entities:
//
//
// * For public and private DNS namespaces, one of the following combinations of
// DNS records in Amazon Route 53:
//
//         * A
//
//         * AAAA
//
//         * A and
// AAAA
//
//         * SRV
//
//         * CNAME
//
//     * Optionally, a health check
//
// After
// you create the service, you can submit a RegisterInstance
// (https://docs.aws.amazon.com/cloud-map/latest/api/API_RegisterInstance.html)
// request, and AWS Cloud Map uses the values in the configuration to create the
// specified entities. For the current limit on the number of instances that you
// can register using the same namespace and using the same service, see AWS Cloud
// Map Limits
// (https://docs.aws.amazon.com/cloud-map/latest/dg/cloud-map-limits.html) in the
// AWS Cloud Map Developer Guide.
func (c *Client) CreateService(ctx context.Context, params *CreateServiceInput, optFns ...func(*Options)) (*CreateServiceOutput, error) {
	stack := middleware.NewStack("CreateService", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateServiceMiddlewares(stack)
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
	addIdempotencyToken_opCreateServiceMiddleware(stack, options)
	addOpCreateServiceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateService(options.Region), middleware.Before)

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
			OperationName: "CreateService",
			Err:           err,
		}
	}
	out := result.(*CreateServiceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateServiceInput struct {
	// The tags to add to the service. Each tag consists of a key and an optional
	// value, both of which you define. Tag keys can have a maximum character length of
	// 128 characters, and tag values can have a maximum length of 256 characters.
	Tags []*types.Tag
	// A description for the service.
	Description *string
	// A complex type that contains information about the Amazon Route 53 records that
	// you want AWS Cloud Map to create when you register an instance.
	DnsConfig *types.DnsConfig
	// A complex type that contains information about an optional custom health check.
	// If you specify a health check configuration, you can specify either
	// HealthCheckCustomConfig or HealthCheckConfig but not both. You can't add,
	// update, or delete a HealthCheckCustomConfig configuration from an existing
	// service.
	HealthCheckCustomConfig *types.HealthCheckCustomConfig
	// The name that you want to assign to the service. If you want AWS Cloud Map to
	// create an SRV record when you register an instance, and if you're using a system
	// that requires a specific SRV format, such as HAProxy (http://www.haproxy.org/),
	// specify the following for Name:
	//
	//     * Start the name with an underscore (_),
	// such as _exampleservice
	//
	//     * End the name with ._protocol, such as ._tcp
	//
	// When
	// you register an instance, AWS Cloud Map creates an SRV record and assigns a name
	// to the record by concatenating the service name and the namespace name, for
	// example: _exampleservice._tcp.example.com
	Name *string
	// The ID of the namespace that you want to use to create the service.
	NamespaceId *string
	// Public DNS and HTTP namespaces only. A complex type that contains settings for
	// an optional Route 53 health check. If you specify settings for a health check,
	// AWS Cloud Map associates the health check with all the Route 53 DNS records that
	// you specify in DnsConfig. If you specify a health check configuration, you can
	// specify either HealthCheckCustomConfig or HealthCheckConfig but not both. For
	// information about the charges for health checks, see AWS Cloud Map Pricing
	// (http://aws.amazon.com/cloud-map/pricing/).
	HealthCheckConfig *types.HealthCheckConfig
	// A unique string that identifies the request and that allows failed CreateService
	// requests to be retried without the risk of executing the operation twice.
	// CreatorRequestId can be any unique string, for example, a date/time stamp.
	CreatorRequestId *string
}

type CreateServiceOutput struct {
	// A complex type that contains information about the new service.
	Service *types.Service

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateServiceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateService{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateService{}, middleware.After)
}

type idempotencyToken_initializeOpCreateService struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateService) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateService) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateServiceInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateServiceInput ")
	}

	if input.CreatorRequestId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.CreatorRequestId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateServiceMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateService{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateService(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicediscovery",
		OperationName: "CreateService",
	}
}
