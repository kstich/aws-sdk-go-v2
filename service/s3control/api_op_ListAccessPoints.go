// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of the access points currently associated with the specified
// bucket. You can retrieve up to 1000 access points per call. If the specified
// bucket has more than 1,000 access points (or the number specified in maxResults,
// whichever is less), the response will include a continuation token that you can
// use to list the additional access points.
func (c *Client) ListAccessPoints(ctx context.Context, params *ListAccessPointsInput, optFns ...func(*Options)) (*ListAccessPointsOutput, error) {
	stack := middleware.NewStack("ListAccessPoints", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpListAccessPointsMiddlewares(stack)
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
	addOpListAccessPointsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListAccessPoints(options.Region), middleware.Before)
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
			OperationName: "ListAccessPoints",
			Err:           err,
		}
	}
	out := result.(*ListAccessPointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAccessPointsInput struct {
	// The name of the bucket whose associated access points you want to list.
	Bucket *string
	// The maximum number of access points that you want to include in the list. If the
	// specified bucket has more than this number of access points, then the response
	// will include a continuation token in the NextToken field that you can use to
	// retrieve the next page of access points.
	MaxResults *int32
	// The AWS account ID for owner of the bucket whose access points you want to list.
	AccountId *string
	// A continuation token. If a previous call to ListAccessPoints returned a
	// continuation token in the NextToken field, then providing that value here causes
	// Amazon S3 to retrieve the next page of results.
	NextToken *string
}

type ListAccessPointsOutput struct {
	// If the specified bucket has more access points than can be returned in one call
	// to this API, then this field contains a continuation token that you can provide
	// in subsequent calls to this API to retrieve additional access points.
	NextToken *string
	// Contains identification and configuration information for one or more access
	// points associated with the specified bucket.
	AccessPointList []*types.AccessPoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpListAccessPointsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpListAccessPoints{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpListAccessPoints{}, middleware.After)
}

func newServiceMetadataMiddleware_opListAccessPoints(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "ListAccessPoints",
	}
}
