// Code generated by smithy-go-codegen DO NOT EDIT.

package computeoptimizer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns Auto Scaling group recommendations. AWS Compute Optimizer currently
// generates recommendations for Auto Scaling groups that are configured to run
// instances of the M, C, R, T, and X instance families. The service does not
// generate recommendations for Auto Scaling groups that have a scaling policy
// attached to them, or that do not have the same values for desired, minimum, and
// maximum capacity. In order for Compute Optimizer to analyze your Auto Scaling
// groups, they must be of a fixed size. For more information, see the AWS Compute
// Optimizer User Guide
// (https://docs.aws.amazon.com/compute-optimizer/latest/ug/what-is.html).
func (c *Client) GetAutoScalingGroupRecommendations(ctx context.Context, params *GetAutoScalingGroupRecommendationsInput, optFns ...func(*Options)) (*GetAutoScalingGroupRecommendationsOutput, error) {
	stack := middleware.NewStack("GetAutoScalingGroupRecommendations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpGetAutoScalingGroupRecommendationsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAutoScalingGroupRecommendations(options.Region), middleware.Before)
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
			OperationName: "GetAutoScalingGroupRecommendations",
			Err:           err,
		}
	}
	out := result.(*GetAutoScalingGroupRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAutoScalingGroupRecommendationsInput struct {
	// The maximum number of Auto Scaling group recommendations to return with a single
	// request. To retrieve the remaining results, make another request with the
	// returned NextToken value.
	MaxResults *int32
	// An array of objects that describe a filter that returns a more specific list of
	// Auto Scaling group recommendations.
	Filters []*types.Filter
	// The token to advance to the next page of Auto Scaling group recommendations.
	NextToken *string
	// The Amazon Resource Name (ARN) of the Auto Scaling groups for which to return
	// recommendations.
	AutoScalingGroupArns []*string
	// The IDs of the AWS accounts for which to return Auto Scaling group
	// recommendations. If your account is the master account of an organization, use
	// this parameter to specify the member accounts for which you want to return Auto
	// Scaling group recommendations. Only one account ID can be specified per request.
	AccountIds []*string
}

type GetAutoScalingGroupRecommendationsOutput struct {
	// An array of objects that describe errors of the request. For example, an error
	// is returned if you request recommendations for an unsupported Auto Scaling
	// group.
	Errors []*types.GetRecommendationError
	// The token to use to advance to the next page of Auto Scaling group
	// recommendations. This value is null when there are no more pages of Auto Scaling
	// group recommendations to return.
	NextToken *string
	// An array of objects that describe Auto Scaling group recommendations.
	AutoScalingGroupRecommendations []*types.AutoScalingGroupRecommendation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpGetAutoScalingGroupRecommendationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpGetAutoScalingGroupRecommendations{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetAutoScalingGroupRecommendations{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetAutoScalingGroupRecommendations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "compute-optimizer",
		OperationName: "GetAutoScalingGroupRecommendations",
	}
}
