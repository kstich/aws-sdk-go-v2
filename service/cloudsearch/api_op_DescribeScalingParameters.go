// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudsearch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearch/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the scaling parameters configured for a domain. A domain's scaling
// parameters specify the desired search instance type and replication count. For
// more information, see Configuring Scaling Options
// (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-scaling-options.html)
// in the Amazon CloudSearch Developer Guide.
func (c *Client) DescribeScalingParameters(ctx context.Context, params *DescribeScalingParametersInput, optFns ...func(*Options)) (*DescribeScalingParametersOutput, error) {
	stack := middleware.NewStack("DescribeScalingParameters", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeScalingParametersMiddlewares(stack)
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
	addOpDescribeScalingParametersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeScalingParameters(options.Region), middleware.Before)

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
			OperationName: "DescribeScalingParameters",
			Err:           err,
		}
	}
	out := result.(*DescribeScalingParametersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DescribeScalingParameters () operation.
// Specifies the name of the domain you want to describe.
type DescribeScalingParametersInput struct {
	// A string that represents the name of a domain. Domain names are unique across
	// the domains owned by an account within an AWS region. Domain names start with a
	// letter or number and can contain the following characters: a-z (lowercase), 0-9,
	// and - (hyphen).
	DomainName *string
}

// The result of a DescribeScalingParameters request. Contains the scaling
// parameters configured for the domain specified in the request.
type DescribeScalingParametersOutput struct {
	// The status and configuration of a search domain's scaling parameters.
	ScalingParameters *types.ScalingParametersStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeScalingParametersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeScalingParameters{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeScalingParameters{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeScalingParameters(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudsearch",
		OperationName: "DescribeScalingParameters",
	}
}
