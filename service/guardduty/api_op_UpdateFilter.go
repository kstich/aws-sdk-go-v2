// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the filter specified by the filter name.
func (c *Client) UpdateFilter(ctx context.Context, params *UpdateFilterInput, optFns ...func(*Options)) (*UpdateFilterOutput, error) {
	stack := middleware.NewStack("UpdateFilter", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateFilterMiddlewares(stack)
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
	addOpUpdateFilterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFilter(options.Region), middleware.Before)
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
			OperationName: "UpdateFilter",
			Err:           err,
		}
	}
	out := result.(*UpdateFilterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFilterInput struct {
	// Represents the criteria to be used in the filter for querying findings.
	FindingCriteria *types.FindingCriteria
	// Specifies the position of the filter in the list of current filters. Also
	// specifies the order in which this filter is applied to the findings.
	Rank *int32
	// The unique ID of the detector that specifies the GuardDuty service where you
	// want to update a filter.
	DetectorId *string
	// Specifies the action that is to be applied to the findings that match the
	// filter.
	Action types.FilterAction
	// The description of the filter.
	Description *string
	// The name of the filter.
	FilterName *string
}

type UpdateFilterOutput struct {
	// The name of the filter.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateFilterMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateFilter{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateFilter{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateFilter(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "UpdateFilter",
	}
}
