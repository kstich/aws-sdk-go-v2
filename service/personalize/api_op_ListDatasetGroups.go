// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of dataset groups. The response provides the properties for each
// dataset group, including the Amazon Resource Name (ARN). For more information on
// dataset groups, see CreateDatasetGroup ().
func (c *Client) ListDatasetGroups(ctx context.Context, params *ListDatasetGroupsInput, optFns ...func(*Options)) (*ListDatasetGroupsOutput, error) {
	stack := middleware.NewStack("ListDatasetGroups", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListDatasetGroupsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListDatasetGroups(options.Region), middleware.Before)
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
			OperationName: "ListDatasetGroups",
			Err:           err,
		}
	}
	out := result.(*ListDatasetGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDatasetGroupsInput struct {
	// A token returned from the previous call to ListDatasetGroups for getting the
	// next set of dataset groups (if they exist).
	NextToken *string
	// The maximum number of dataset groups to return.
	MaxResults *int32
}

type ListDatasetGroupsOutput struct {
	// The list of your dataset groups.
	DatasetGroups []*types.DatasetGroupSummary
	// A token for getting the next set of dataset groups (if they exist).
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListDatasetGroupsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListDatasetGroups{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDatasetGroups{}, middleware.After)
}

func newServiceMetadataMiddleware_opListDatasetGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "ListDatasetGroups",
	}
}
