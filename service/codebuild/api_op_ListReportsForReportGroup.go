// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of ARNs for the reports that belong to a ReportGroup.
func (c *Client) ListReportsForReportGroup(ctx context.Context, params *ListReportsForReportGroupInput, optFns ...func(*Options)) (*ListReportsForReportGroupOutput, error) {
	stack := middleware.NewStack("ListReportsForReportGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListReportsForReportGroupMiddlewares(stack)
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
	addOpListReportsForReportGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListReportsForReportGroup(options.Region), middleware.Before)

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
			OperationName: "ListReportsForReportGroup",
			Err:           err,
		}
	}
	out := result.(*ListReportsForReportGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListReportsForReportGroupInput struct {
	// During a previous call, the maximum number of items that can be returned is the
	// value specified in maxResults. If there more items in the list, then a unique
	// string called a nextToken is returned. To get the next batch of items in the
	// list, call this operation again, adding the next token to the call. To get all
	// of the items in the list, keep calling this operation with each subsequent next
	// token that is returned, until no more next tokens are returned.
	NextToken *string
	// Use to specify whether the results are returned in ascending or descending
	// order.
	SortOrder types.SortOrderType
	// The maximum number of paginated reports in this report group returned per
	// response. Use nextToken to iterate pages in the list of returned Report objects.
	// The default value is 100.
	MaxResults *int32
	// The ARN of the report group for which you want to return report ARNs.
	ReportGroupArn *string
	// A ReportFilter object used to filter the returned reports.
	Filter *types.ReportFilter
}

type ListReportsForReportGroupOutput struct {
	// The list of report ARNs.
	Reports []*string
	// During a previous call, the maximum number of items that can be returned is the
	// value specified in maxResults. If there more items in the list, then a unique
	// string called a nextToken is returned. To get the next batch of items in the
	// list, call this operation again, adding the next token to the call. To get all
	// of the items in the list, keep calling this operation with each subsequent next
	// token that is returned, until no more next tokens are returned.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListReportsForReportGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListReportsForReportGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListReportsForReportGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opListReportsForReportGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "ListReportsForReportGroup",
	}
}