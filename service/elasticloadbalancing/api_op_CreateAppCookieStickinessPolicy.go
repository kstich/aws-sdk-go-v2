// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Generates a stickiness policy with sticky session lifetimes that follow that of
// an application-generated cookie. This policy can be associated only with
// HTTP/HTTPS listeners. This policy is similar to the policy created by
// CreateLBCookieStickinessPolicy (), except that the lifetime of the special
// Elastic Load Balancing cookie, AWSELB, follows the lifetime of the
// application-generated cookie specified in the policy configuration. The load
// balancer only inserts a new stickiness cookie when the application response
// includes a new application cookie. If the application cookie is explicitly
// removed or expires, the session stops being sticky until a new application
// cookie is issued. For more information, see Application-Controlled Session
// Stickiness
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-sticky-sessions.html#enable-sticky-sessions-application)
// in the Classic Load Balancers Guide.
func (c *Client) CreateAppCookieStickinessPolicy(ctx context.Context, params *CreateAppCookieStickinessPolicyInput, optFns ...func(*Options)) (*CreateAppCookieStickinessPolicyOutput, error) {
	stack := middleware.NewStack("CreateAppCookieStickinessPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateAppCookieStickinessPolicyMiddlewares(stack)
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
	addOpCreateAppCookieStickinessPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAppCookieStickinessPolicy(options.Region), middleware.Before)
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
			OperationName: "CreateAppCookieStickinessPolicy",
			Err:           err,
		}
	}
	out := result.(*CreateAppCookieStickinessPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CreateAppCookieStickinessPolicy.
type CreateAppCookieStickinessPolicyInput struct {
	// The name of the application cookie used for stickiness.
	CookieName *string
	// The name of the load balancer.
	LoadBalancerName *string
	// The name of the policy being created. Policy names must consist of alphanumeric
	// characters and dashes (-). This name must be unique within the set of policies
	// for this load balancer.
	PolicyName *string
}

// Contains the output for CreateAppCookieStickinessPolicy.
type CreateAppCookieStickinessPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateAppCookieStickinessPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateAppCookieStickinessPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateAppCookieStickinessPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateAppCookieStickinessPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "CreateAppCookieStickinessPolicy",
	}
}
