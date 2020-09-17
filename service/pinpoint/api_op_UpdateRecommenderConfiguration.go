// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates an Amazon Pinpoint configuration for a recommender model.
func (c *Client) UpdateRecommenderConfiguration(ctx context.Context, params *UpdateRecommenderConfigurationInput, optFns ...func(*Options)) (*UpdateRecommenderConfigurationOutput, error) {
	stack := middleware.NewStack("UpdateRecommenderConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateRecommenderConfigurationMiddlewares(stack)
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
	addOpUpdateRecommenderConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRecommenderConfiguration(options.Region), middleware.Before)
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
			OperationName: "UpdateRecommenderConfiguration",
			Err:           err,
		}
	}
	out := result.(*UpdateRecommenderConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRecommenderConfigurationInput struct {
	// Specifies Amazon Pinpoint configuration settings for retrieving and processing
	// recommendation data from a recommender model.
	UpdateRecommenderConfiguration *types.UpdateRecommenderConfigurationShape
	// The unique identifier for the recommender model configuration. This identifier
	// is displayed as the Recommender ID on the Amazon Pinpoint console.
	RecommenderId *string
}

type UpdateRecommenderConfigurationOutput struct {
	// Provides information about Amazon Pinpoint configuration settings for retrieving
	// and processing data from a recommender model.
	RecommenderConfigurationResponse *types.RecommenderConfigurationResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateRecommenderConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateRecommenderConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateRecommenderConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateRecommenderConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "UpdateRecommenderConfiguration",
	}
}
