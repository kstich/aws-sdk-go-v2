// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Archives GuardDuty findings that are specified by the list of finding IDs. Only
// the master account can archive findings. Member accounts don't have permission
// to archive findings from their accounts.
func (c *Client) ArchiveFindings(ctx context.Context, params *ArchiveFindingsInput, optFns ...func(*Options)) (*ArchiveFindingsOutput, error) {
	stack := middleware.NewStack("ArchiveFindings", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpArchiveFindingsMiddlewares(stack)
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
	addOpArchiveFindingsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opArchiveFindings(options.Region), middleware.Before)

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
			OperationName: "ArchiveFindings",
			Err:           err,
		}
	}
	out := result.(*ArchiveFindingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ArchiveFindingsInput struct {
	// The ID of the detector that specifies the GuardDuty service whose findings you
	// want to archive.
	DetectorId *string
	// The IDs of the findings that you want to archive.
	FindingIds []*string
}

type ArchiveFindingsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpArchiveFindingsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpArchiveFindings{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpArchiveFindings{}, middleware.After)
}

func newServiceMetadataMiddleware_opArchiveFindings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "ArchiveFindings",
	}
}
