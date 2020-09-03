// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates a resolver rule with a VPC. When you associate a rule with a VPC,
// Resolver forwards all DNS queries for the domain name that is specified in the
// rule and that originate in the VPC. The queries are forwarded to the IP
// addresses for the DNS resolvers that are specified in the rule. For more
// information about rules, see CreateResolverRule ().
func (c *Client) AssociateResolverRule(ctx context.Context, params *AssociateResolverRuleInput, optFns ...func(*Options)) (*AssociateResolverRuleOutput, error) {
	stack := middleware.NewStack("AssociateResolverRule", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAssociateResolverRuleMiddlewares(stack)
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
	addOpAssociateResolverRuleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateResolverRule(options.Region), middleware.Before)

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
			OperationName: "AssociateResolverRule",
			Err:           err,
		}
	}
	out := result.(*AssociateResolverRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateResolverRuleInput struct {
	// The ID of the resolver rule that you want to associate with the VPC. To list the
	// existing resolver rules, use ListResolverRules ().
	ResolverRuleId *string
	// A name for the association that you're creating between a resolver rule and a
	// VPC.
	Name *string
	// The ID of the VPC that you want to associate the resolver rule with.
	VPCId *string
}

type AssociateResolverRuleOutput struct {
	// Information about the AssociateResolverRule request, including the status of the
	// request.
	ResolverRuleAssociation *types.ResolverRuleAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAssociateResolverRuleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateResolverRule{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateResolverRule{}, middleware.After)
}

func newServiceMetadataMiddleware_opAssociateResolverRule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "AssociateResolverRule",
	}
}
