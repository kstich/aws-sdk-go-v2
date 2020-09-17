// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the details for one or more Amazon SageMaker models that have been imported
// into the service. This is a paginated API. If you provide a null maxResults,
// this actions retrieves a maximum of 10 records per page. If you provide a
// maxResults, the value must be between 5 and 10. To get the next page results,
// provide the pagination token from the GetExternalModelsResult as part of your
// request. A null pagination token fetches the records from the beginning.
func (c *Client) GetExternalModels(ctx context.Context, params *GetExternalModelsInput, optFns ...func(*Options)) (*GetExternalModelsOutput, error) {
	stack := middleware.NewStack("GetExternalModels", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetExternalModelsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetExternalModels(options.Region), middleware.Before)
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
			OperationName: "GetExternalModels",
			Err:           err,
		}
	}
	out := result.(*GetExternalModelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetExternalModelsInput struct {
	// The maximum number of objects to return for the request.
	MaxResults *int32
	// The Amazon SageMaker model endpoint.
	ModelEndpoint *string
	// The next page token for the request.
	NextToken *string
}

type GetExternalModelsOutput struct {
	// The next page token to be used in subsequent requests.
	NextToken *string
	// Gets the Amazon SageMaker models.
	ExternalModels []*types.ExternalModel

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetExternalModelsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetExternalModels{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetExternalModels{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetExternalModels(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "GetExternalModels",
	}
}
