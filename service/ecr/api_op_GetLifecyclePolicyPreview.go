// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the results of the lifecycle policy preview request for the specified
// repository.
func (c *Client) GetLifecyclePolicyPreview(ctx context.Context, params *GetLifecyclePolicyPreviewInput, optFns ...func(*Options)) (*GetLifecyclePolicyPreviewOutput, error) {
	stack := middleware.NewStack("GetLifecyclePolicyPreview", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetLifecyclePolicyPreviewMiddlewares(stack)
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
	addOpGetLifecyclePolicyPreviewValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetLifecyclePolicyPreview(options.Region), middleware.Before)
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
			OperationName: "GetLifecyclePolicyPreview",
			Err:           err,
		}
	}
	out := result.(*GetLifecyclePolicyPreviewOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLifecyclePolicyPreviewInput struct {
	// The name of the repository.
	RepositoryName *string
	// The maximum number of repository results returned by
	// GetLifecyclePolicyPreviewRequest in  paginated output. When this parameter is
	// used, GetLifecyclePolicyPreviewRequest only returns  maxResults results in a
	// single page along with a nextToken  response element. The remaining results of
	// the initial request can be seen by sending  another
	// GetLifecyclePolicyPreviewRequest request with the returned nextToken  value.
	// This value can be between 1 and 1000. If this  parameter is not used, then
	// GetLifecyclePolicyPreviewRequest returns up to  100 results and a nextToken
	// value, if  applicable. This option cannot be used when you specify images with
	// imageIds.
	MaxResults *int32
	// An optional parameter that filters results based on image tag status and all
	// tags, if tagged.
	Filter *types.LifecyclePolicyPreviewFilter
	// The list of imageIDs to be included.
	ImageIds []*types.ImageIdentifier
	// The nextToken value returned from a previous paginated
	// GetLifecyclePolicyPreviewRequest request where maxResults was used and the
	// results exceeded the value of that parameter. Pagination continues from the end
	// of the  previous results that returned the nextToken value. This value is  null
	// when there are no more results to return. This option cannot be used when you
	// specify images with imageIds.
	NextToken *string
	// The AWS account ID associated with the registry that contains the repository. If
	// you do not specify a registry, the default registry is assumed.
	RegistryId *string
}

type GetLifecyclePolicyPreviewOutput struct {
	// The results of the lifecycle policy preview request.
	PreviewResults []*types.LifecyclePolicyPreviewResult
	// The repository name associated with the request.
	RepositoryName *string
	// The JSON lifecycle policy text.
	LifecyclePolicyText *string
	// The nextToken value to include in a future GetLifecyclePolicyPreview request.
	// When the results of a GetLifecyclePolicyPreview request exceed maxResults, this
	// value can be used to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string
	// The registry ID associated with the request.
	RegistryId *string
	// The status of the lifecycle policy preview request.
	Status types.LifecyclePolicyPreviewStatus
	// The list of images that is returned as a result of the action.
	Summary *types.LifecyclePolicyPreviewSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetLifecyclePolicyPreviewMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetLifecyclePolicyPreview{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetLifecyclePolicyPreview{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetLifecyclePolicyPreview(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr",
		OperationName: "GetLifecyclePolicyPreview",
	}
}
