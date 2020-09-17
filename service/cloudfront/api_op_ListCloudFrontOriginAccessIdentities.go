// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists origin access identities.
func (c *Client) ListCloudFrontOriginAccessIdentities(ctx context.Context, params *ListCloudFrontOriginAccessIdentitiesInput, optFns ...func(*Options)) (*ListCloudFrontOriginAccessIdentitiesOutput, error) {
	stack := middleware.NewStack("ListCloudFrontOriginAccessIdentities", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpListCloudFrontOriginAccessIdentitiesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListCloudFrontOriginAccessIdentities(options.Region), middleware.Before)
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
			OperationName: "ListCloudFrontOriginAccessIdentities",
			Err:           err,
		}
	}
	out := result.(*ListCloudFrontOriginAccessIdentitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to list origin access identities.
type ListCloudFrontOriginAccessIdentitiesInput struct {
	// The maximum number of origin access identities you want in the response body.
	MaxItems *string
	// Use this when paginating results to indicate where to begin in your list of
	// origin access identities. The results include identities in the list that occur
	// after the marker. To get the next page of results, set the Marker to the value
	// of the NextMarker from the current page's response (which is also the ID of the
	// last identity on that page).
	Marker *string
}

// The returned result of the corresponding request.
type ListCloudFrontOriginAccessIdentitiesOutput struct {
	// The CloudFrontOriginAccessIdentityList type.
	CloudFrontOriginAccessIdentityList *types.CloudFrontOriginAccessIdentityList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpListCloudFrontOriginAccessIdentitiesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpListCloudFrontOriginAccessIdentities{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpListCloudFrontOriginAccessIdentities{}, middleware.After)
}

func newServiceMetadataMiddleware_opListCloudFrontOriginAccessIdentities(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "ListCloudFrontOriginAccessIdentities",
	}
}
