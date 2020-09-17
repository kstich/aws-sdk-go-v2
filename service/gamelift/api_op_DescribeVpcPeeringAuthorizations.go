// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves valid VPC peering authorizations that are pending for the AWS account.
// This operation returns all VPC peering authorizations and requests for peering.
// This includes those initiated and received by this account.
//
//     *
// CreateVpcPeeringAuthorization ()
//
//     * DescribeVpcPeeringAuthorizations ()
//
//
// * DeleteVpcPeeringAuthorization ()
//
//     * CreateVpcPeeringConnection ()
//
//     *
// DescribeVpcPeeringConnections ()
//
//     * DeleteVpcPeeringConnection ()
func (c *Client) DescribeVpcPeeringAuthorizations(ctx context.Context, params *DescribeVpcPeeringAuthorizationsInput, optFns ...func(*Options)) (*DescribeVpcPeeringAuthorizationsOutput, error) {
	stack := middleware.NewStack("DescribeVpcPeeringAuthorizations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeVpcPeeringAuthorizationsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVpcPeeringAuthorizations(options.Region), middleware.Before)
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
			OperationName: "DescribeVpcPeeringAuthorizations",
			Err:           err,
		}
	}
	out := result.(*DescribeVpcPeeringAuthorizationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVpcPeeringAuthorizationsInput struct {
}

type DescribeVpcPeeringAuthorizationsOutput struct {
	// A collection of objects that describe all valid VPC peering operations for the
	// current AWS account.
	VpcPeeringAuthorizations []*types.VpcPeeringAuthorization

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeVpcPeeringAuthorizationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeVpcPeeringAuthorizations{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeVpcPeeringAuthorizations{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeVpcPeeringAuthorizations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeVpcPeeringAuthorizations",
	}
}
