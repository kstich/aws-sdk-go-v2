// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all your Direct Connect gateways or only the specified Direct Connect
// gateway. Deleted Direct Connect gateways are not returned.
func (c *Client) DescribeDirectConnectGateways(ctx context.Context, params *DescribeDirectConnectGatewaysInput, optFns ...func(*Options)) (*DescribeDirectConnectGatewaysOutput, error) {
	stack := middleware.NewStack("DescribeDirectConnectGateways", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeDirectConnectGatewaysMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDirectConnectGateways(options.Region), middleware.Before)

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
			OperationName: "DescribeDirectConnectGateways",
			Err:           err,
		}
	}
	out := result.(*DescribeDirectConnectGatewaysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDirectConnectGatewaysInput struct {
	// The ID of the Direct Connect gateway.
	DirectConnectGatewayId *string
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value. If
	// MaxResults is given a value larger than 100, only 100 results are returned.
	MaxResults *int32
	// The token provided in the previous call to retrieve the next page.
	NextToken *string
}

type DescribeDirectConnectGatewaysOutput struct {
	// The Direct Connect gateways.
	DirectConnectGateways []*types.DirectConnectGateway
	// The token to retrieve the next page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeDirectConnectGatewaysMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDirectConnectGateways{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDirectConnectGateways{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeDirectConnectGateways(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "DescribeDirectConnectGateways",
	}
}
