// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets aside (overrides) all approval rule requirements for a specified pull
// request.
func (c *Client) OverridePullRequestApprovalRules(ctx context.Context, params *OverridePullRequestApprovalRulesInput, optFns ...func(*Options)) (*OverridePullRequestApprovalRulesOutput, error) {
	stack := middleware.NewStack("OverridePullRequestApprovalRules", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpOverridePullRequestApprovalRulesMiddlewares(stack)
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
	addOpOverridePullRequestApprovalRulesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opOverridePullRequestApprovalRules(options.Region), middleware.Before)
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
			OperationName: "OverridePullRequestApprovalRules",
			Err:           err,
		}
	}
	out := result.(*OverridePullRequestApprovalRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type OverridePullRequestApprovalRulesInput struct {
	// The system-generated ID of the most recent revision of the pull request. You
	// cannot override approval rules for anything but the most recent revision of a
	// pull request. To get the revision ID, use GetPullRequest.
	RevisionId *string
	// The system-generated ID of the pull request for which you want to override all
	// approval rule requirements. To get this information, use GetPullRequest ().
	PullRequestId *string
	// Whether you want to set aside approval rule requirements for the pull request
	// (OVERRIDE) or revoke a previous override and apply approval rule requirements
	// (REVOKE). REVOKE status is not stored.
	OverrideStatus types.OverrideStatus
}

type OverridePullRequestApprovalRulesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpOverridePullRequestApprovalRulesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpOverridePullRequestApprovalRules{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpOverridePullRequestApprovalRules{}, middleware.After)
}

func newServiceMetadataMiddleware_opOverridePullRequestApprovalRules(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "OverridePullRequestApprovalRules",
	}
}
