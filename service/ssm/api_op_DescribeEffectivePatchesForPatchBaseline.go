// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the current effective patches (the patch and the approval state) for
// the specified patch baseline. Note that this API applies only to Windows patch
// baselines.
func (c *Client) DescribeEffectivePatchesForPatchBaseline(ctx context.Context, params *DescribeEffectivePatchesForPatchBaselineInput, optFns ...func(*Options)) (*DescribeEffectivePatchesForPatchBaselineOutput, error) {
	stack := middleware.NewStack("DescribeEffectivePatchesForPatchBaseline", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeEffectivePatchesForPatchBaselineMiddlewares(stack)
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
	addOpDescribeEffectivePatchesForPatchBaselineValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEffectivePatchesForPatchBaseline(options.Region), middleware.Before)
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
			OperationName: "DescribeEffectivePatchesForPatchBaseline",
			Err:           err,
		}
	}
	out := result.(*DescribeEffectivePatchesForPatchBaselineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEffectivePatchesForPatchBaselineInput struct {
	// The ID of the patch baseline to retrieve the effective patches for.
	BaselineId *string
	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string
	// The maximum number of patches to return (per page).
	MaxResults *int32
}

type DescribeEffectivePatchesForPatchBaselineOutput struct {
	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string
	// An array of patches and patch status.
	EffectivePatches []*types.EffectivePatch

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeEffectivePatchesForPatchBaselineMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEffectivePatchesForPatchBaseline{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEffectivePatchesForPatchBaseline{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeEffectivePatchesForPatchBaseline(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeEffectivePatchesForPatchBaseline",
	}
}
