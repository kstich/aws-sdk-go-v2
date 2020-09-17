// Code generated by smithy-go-codegen DO NOT EDIT.

package dataexchange

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// This operation creates a revision for a data set.
func (c *Client) CreateRevision(ctx context.Context, params *CreateRevisionInput, optFns ...func(*Options)) (*CreateRevisionOutput, error) {
	stack := middleware.NewStack("CreateRevision", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateRevisionMiddlewares(stack)
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
	addOpCreateRevisionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRevision(options.Region), middleware.Before)
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
			OperationName: "CreateRevision",
			Err:           err,
		}
	}
	out := result.(*CreateRevisionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request body for CreateRevision.
type CreateRevisionInput struct {
	// An optional comment about the revision.
	Comment *string
	// The unique identifier for a data set.
	DataSetId *string
	// A revision tag is an optional label that you can assign to a revision when you
	// create it. Each tag consists of a key and an optional value, both of which you
	// define. When you use tagging, you can also use tag-based access control in IAM
	// policies to control access to these data sets and revisions.
	Tags map[string]*string
}

type CreateRevisionOutput struct {
	// An optional comment about the revision.
	Comment *string
	// The tags for the revision.
	Tags map[string]*string
	// The ARN for the revision
	Arn *string
	// To publish a revision to a data set in a product, the revision must first be
	// finalized. Finalizing a revision tells AWS Data Exchange that your changes to
	// the assets in the revision are complete. After it's in this read-only state, you
	// can publish the revision to your products. Finalized revisions can be published
	// through the AWS Data Exchange console or the AWS Marketplace Catalog API, using
	// the StartChangeSet AWS Marketplace Catalog API action. When using the API,
	// revisions are uniquely identified by their ARN.
	Finalized *bool
	// The unique identifier for the data set associated with this revision.
	DataSetId *string
	// The date and time that the revision was last updated, in ISO 8601 format.
	UpdatedAt *time.Time
	// The unique identifier for the revision.
	Id *string
	// The date and time that the revision was created, in ISO 8601 format.
	CreatedAt *time.Time
	// The revision ID of the owned revision corresponding to the entitled revision
	// being viewed. This parameter is returned when a revision owner is viewing the
	// entitled copy of its owned revision.
	SourceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateRevisionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateRevision{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRevision{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateRevision(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dataexchange",
		OperationName: "CreateRevision",
	}
}
