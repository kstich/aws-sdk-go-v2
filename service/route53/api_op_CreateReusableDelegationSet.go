// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a delegation set (a group of four name servers) that can be reused by
// multiple hosted zones that were created by the same AWS account. You can also
// create a reusable delegation set that uses the four name servers that are
// associated with an existing hosted zone. Specify the hosted zone ID in the
// CreateReusableDelegationSet request. You can't associate a reusable delegation
// set with a private hosted zone. For information about using a reusable
// delegation set to configure white label name servers, see Configuring White
// Label Name Servers
// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/white-label-name-servers.html).
// <p>The process for migrating existing hosted zones to use a reusable delegation
// set is comparable to the process for configuring white label name servers. You
// need to perform the following steps:</p> <ol> <li> <p>Create a reusable
// delegation set.</p> </li> <li> <p>Recreate hosted zones, and reduce the TTL to
// 60 seconds or less.</p> </li> <li> <p>Recreate resource record sets in the new
// hosted zones.</p> </li> <li> <p>Change the registrar's name servers to use the
// name servers for the new hosted zones.</p> </li> <li> <p>Monitor traffic for the
// website or application.</p> </li> <li> <p>Change TTLs back to their original
// values.</p> </li> </ol> <p>If you want to migrate existing hosted zones to use a
// reusable delegation set, the existing hosted zones can't use any of the name
// servers that are assigned to the reusable delegation set. If one or more hosted
// zones do use one or more name servers that are assigned to the reusable
// delegation set, you can do one of the following:</p> <ul> <li> <p>For small
// numbers of hosted zones—up to a few hundred—it's relatively easy to create
// reusable delegation sets until you get one that has four name servers that don't
// overlap with any of the name servers in your hosted zones.</p> </li> <li> <p>For
// larger numbers of hosted zones, the easiest solution is to use more than one
// reusable delegation set.</p> </li> <li> <p>For larger numbers of hosted zones,
// you can also migrate hosted zones that have overlapping name servers to hosted
// zones that don't have overlapping name servers, then migrate the hosted zones
// again to use the reusable delegation set.</p> </li> </ul>
func (c *Client) CreateReusableDelegationSet(ctx context.Context, params *CreateReusableDelegationSetInput, optFns ...func(*Options)) (*CreateReusableDelegationSetOutput, error) {
	stack := middleware.NewStack("CreateReusableDelegationSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpCreateReusableDelegationSetMiddlewares(stack)
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
	addOpCreateReusableDelegationSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateReusableDelegationSet(options.Region), middleware.Before)
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
			OperationName: "CreateReusableDelegationSet",
			Err:           err,
		}
	}
	out := result.(*CreateReusableDelegationSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateReusableDelegationSetInput struct {
	// If you want to mark the delegation set for an existing hosted zone as reusable,
	// the ID for that hosted zone.
	HostedZoneId *string
	// A unique string that identifies the request, and that allows you to retry failed
	// CreateReusableDelegationSet requests without the risk of executing the operation
	// twice. You must use a unique CallerReference string every time you submit a
	// CreateReusableDelegationSet request. CallerReference can be any unique string,
	// for example a date/time stamp.
	CallerReference *string
}

type CreateReusableDelegationSetOutput struct {
	// The unique URL representing the new reusable delegation set.
	Location *string
	// A complex type that contains name server information.
	DelegationSet *types.DelegationSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpCreateReusableDelegationSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpCreateReusableDelegationSet{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpCreateReusableDelegationSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateReusableDelegationSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "CreateReusableDelegationSet",
	}
}
