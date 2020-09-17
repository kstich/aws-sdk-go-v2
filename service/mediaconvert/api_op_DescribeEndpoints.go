// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconvert

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Send an request with an empty body to the regional API endpoint to get your
// account API endpoint.
func (c *Client) DescribeEndpoints(ctx context.Context, params *DescribeEndpointsInput, optFns ...func(*Options)) (*DescribeEndpointsOutput, error) {
	stack := middleware.NewStack("DescribeEndpoints", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeEndpointsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEndpoints(options.Region), middleware.Before)
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
			OperationName: "DescribeEndpoints",
			Err:           err,
		}
	}
	out := result.(*DescribeEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeEndpointsRequest
type DescribeEndpointsInput struct {
	// Use this string, provided with the response to a previous request, to request
	// the next batch of endpoints.
	NextToken *string
	// Optional. Max number of endpoints, up to twenty, that will be returned at one
	// time.
	MaxResults *int32
	// Optional field, defaults to DEFAULT. Specify DEFAULT for this operation to
	// return your endpoints if any exist, or to create an endpoint for you and return
	// it if one doesn't already exist. Specify GET_ONLY to return your endpoints if
	// any exist, or an empty list if none exist.
	Mode types.DescribeEndpointsMode
}

type DescribeEndpointsOutput struct {
	// Use this string to request the next batch of endpoints.
	NextToken *string
	// List of endpoints
	Endpoints []*types.Endpoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeEndpointsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeEndpoints{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeEndpoints{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeEndpoints(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconvert",
		OperationName: "DescribeEndpoints",
	}
}
