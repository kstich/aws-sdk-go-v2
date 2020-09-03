// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns an array of target IDs that are associated a deployment.
func (c *Client) ListDeploymentTargets(ctx context.Context, params *ListDeploymentTargetsInput, optFns ...func(*Options)) (*ListDeploymentTargetsOutput, error) {
	stack := middleware.NewStack("ListDeploymentTargets", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListDeploymentTargetsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListDeploymentTargets(options.Region), middleware.Before)

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
			OperationName: "ListDeploymentTargets",
			Err:           err,
		}
	}
	out := result.(*ListDeploymentTargetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDeploymentTargetsInput struct {
	// A key used to filter the returned targets. The two valid values are:
	//
	//     *
	// TargetStatus - A TargetStatus filter string can be Failed, InProgress, Pending,
	// Ready, Skipped, Succeeded, or Unknown.
	//
	//     * ServerInstanceLabel - A
	// ServerInstanceLabel filter string can be Blue or Green.
	TargetFilters map[string][]*string
	// A token identifier returned from the previous ListDeploymentTargets call. It can
	// be used to return the next set of deployment targets in the list.
	NextToken *string
	// The unique ID of a deployment.
	DeploymentId *string
}

type ListDeploymentTargetsOutput struct {
	// If a large amount of information is returned, a token identifier is also
	// returned. It can be used in a subsequent ListDeploymentTargets call to return
	// the next set of deployment targets in the list.
	NextToken *string
	// The unique IDs of deployment targets.
	TargetIds []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListDeploymentTargetsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListDeploymentTargets{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDeploymentTargets{}, middleware.After)
}

func newServiceMetadataMiddleware_opListDeploymentTargets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "ListDeploymentTargets",
	}
}
