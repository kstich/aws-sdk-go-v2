// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Cancels a pending VPC peering authorization for the specified VPC. If you need
// to delete an existing VPC peering connection, call DeleteVpcPeeringConnection
// ().
//
//     * CreateVpcPeeringAuthorization ()
//
//     *
// DescribeVpcPeeringAuthorizations ()
//
//     * DeleteVpcPeeringAuthorization ()
//
//
// * CreateVpcPeeringConnection ()
//
//     * DescribeVpcPeeringConnections ()
//
//     *
// DeleteVpcPeeringConnection ()
func (c *Client) DeleteVpcPeeringAuthorization(ctx context.Context, params *DeleteVpcPeeringAuthorizationInput, optFns ...func(*Options)) (*DeleteVpcPeeringAuthorizationOutput, error) {
	stack := middleware.NewStack("DeleteVpcPeeringAuthorization", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteVpcPeeringAuthorizationMiddlewares(stack)
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
	addOpDeleteVpcPeeringAuthorizationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteVpcPeeringAuthorization(options.Region), middleware.Before)
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
			OperationName: "DeleteVpcPeeringAuthorization",
			Err:           err,
		}
	}
	out := result.(*DeleteVpcPeeringAuthorizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type DeleteVpcPeeringAuthorizationInput struct {
	// A unique identifier for the AWS account that you use to manage your Amazon
	// GameLift fleet. You can find your Account ID in the AWS Management Console under
	// account settings.
	GameLiftAwsAccountId *string
	// A unique identifier for a VPC with resources to be accessed by your Amazon
	// GameLift fleet. The VPC must be in the same Region where your fleet is deployed.
	// Look up a VPC ID using the VPC Dashboard (https://console.aws.amazon.com/vpc/)
	// in the AWS Management Console. Learn more about VPC peering in VPC Peering with
	// Amazon GameLift Fleets
	// (https://docs.aws.amazon.com/gamelift/latest/developerguide/vpc-peering.html).
	PeerVpcId *string
}

type DeleteVpcPeeringAuthorizationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteVpcPeeringAuthorizationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteVpcPeeringAuthorization{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteVpcPeeringAuthorization{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteVpcPeeringAuthorization(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DeleteVpcPeeringAuthorization",
	}
}
