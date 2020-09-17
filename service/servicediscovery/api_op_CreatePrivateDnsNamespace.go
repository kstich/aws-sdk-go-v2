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

// Creates a private namespace based on DNS, which will be visible only inside a
// specified Amazon VPC. The namespace defines your service naming scheme. For
// example, if you name your namespace example.com and name your service backend,
// the resulting DNS name for the service will be backend.example.com. For the
// current limit on the number of namespaces that you can create using the same AWS
// account, see AWS Cloud Map Limits
// (https://docs.aws.amazon.com/cloud-map/latest/dg/cloud-map-limits.html) in the
// AWS Cloud Map Developer Guide.
func (c *Client) CreatePrivateDnsNamespace(ctx context.Context, params *CreatePrivateDnsNamespaceInput, optFns ...func(*Options)) (*CreatePrivateDnsNamespaceOutput, error) {
	stack := middleware.NewStack("CreatePrivateDnsNamespace", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreatePrivateDnsNamespaceMiddlewares(stack)
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
	addIdempotencyToken_opCreatePrivateDnsNamespaceMiddleware(stack, options)
	addOpCreatePrivateDnsNamespaceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePrivateDnsNamespace(options.Region), middleware.Before)
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
			OperationName: "CreatePrivateDnsNamespace",
			Err:           err,
		}
	}
	out := result.(*CreatePrivateDnsNamespaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePrivateDnsNamespaceInput struct {
	// The name that you want to assign to this namespace. When you create a private
	// DNS namespace, AWS Cloud Map automatically creates an Amazon Route 53 private
	// hosted zone that has the same name as the namespace.
	Name *string
	// The ID of the Amazon VPC that you want to associate the namespace with.
	Vpc *string
	// A description for the namespace.
	Description *string
	// The tags to add to the namespace. Each tag consists of a key and an optional
	// value, both of which you define. Tag keys can have a maximum character length of
	// 128 characters, and tag values can have a maximum length of 256 characters.
	Tags []*types.Tag
	// A unique string that identifies the request and that allows failed
	// CreatePrivateDnsNamespace requests to be retried without the risk of executing
	// the operation twice. CreatorRequestId can be any unique string, for example, a
	// date/time stamp.
	CreatorRequestId *string
}

type CreatePrivateDnsNamespaceOutput struct {
	// A value that you can use to determine whether the request completed
	// successfully. To get the status of the operation, see GetOperation
	// (https://docs.aws.amazon.com/cloud-map/latest/api/API_GetOperation.html).
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreatePrivateDnsNamespaceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePrivateDnsNamespace{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePrivateDnsNamespace{}, middleware.After)
}

type idempotencyToken_initializeOpCreatePrivateDnsNamespace struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreatePrivateDnsNamespace) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreatePrivateDnsNamespace) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreatePrivateDnsNamespaceInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreatePrivateDnsNamespaceInput ")
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
func addIdempotencyToken_opCreatePrivateDnsNamespaceMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreatePrivateDnsNamespace{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreatePrivateDnsNamespace(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicediscovery",
		OperationName: "CreatePrivateDnsNamespace",
	}
}
