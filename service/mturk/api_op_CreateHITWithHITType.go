// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The CreateHITWithHITType operation creates a new Human Intelligence Task (HIT)
// using an existing HITTypeID generated by the CreateHITType operation. This is an
// alternative way to create HITs from the CreateHIT operation. This is the
// recommended best practice for Requesters who are creating large numbers of HITs.
// CreateHITWithHITType also supports several ways to provide question data: by
// providing a value for the Question parameter that fully specifies the contents
// of the HIT, or by providing a HitLayoutId and associated HitLayoutParameters. If
// a HIT is created with 10 or more maximum assignments, there is an additional
// fee. For more information, see Amazon Mechanical Turk Pricing
// (https://requester.mturk.com/pricing).
func (c *Client) CreateHITWithHITType(ctx context.Context, params *CreateHITWithHITTypeInput, optFns ...func(*Options)) (*CreateHITWithHITTypeOutput, error) {
	stack := middleware.NewStack("CreateHITWithHITType", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateHITWithHITTypeMiddlewares(stack)
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
	addOpCreateHITWithHITTypeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateHITWithHITType(options.Region), middleware.Before)
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
			OperationName: "CreateHITWithHITType",
			Err:           err,
		}
	}
	out := result.(*CreateHITWithHITTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateHITWithHITTypeInput struct {
	// The Assignment-level Review Policy applies to the assignments under the HIT. You
	// can specify for Mechanical Turk to take various actions based on the policy.
	AssignmentReviewPolicy *types.ReviewPolicy
	// If the HITLayoutId is provided, any placeholder values must be filled in with
	// values using the HITLayoutParameter structure. For more information, see
	// HITLayout.
	HITLayoutParameters []*types.HITLayoutParameter
	// The number of times the HIT can be accepted and completed before the HIT becomes
	// unavailable.
	MaxAssignments *int32
	// A unique identifier for this request which allows you to retry the call on error
	// without creating duplicate HITs. This is useful in cases such as network
	// timeouts where it is unclear whether or not the call succeeded on the server. If
	// the HIT already exists in the system from a previous call using the same
	// UniqueRequestToken, subsequent calls will return a
	// AWS.MechanicalTurk.HitAlreadyExists error with a message containing the HITId.
	// Note: It is your responsibility to ensure uniqueness of the token. The unique
	// token expires after 24 hours. Subsequent calls using the same UniqueRequestToken
	// made after the 24 hour limit could create duplicate HITs.
	UniqueRequestToken *string
	// The HIT type ID you want to create this HIT with.
	HITTypeId *string
	// An arbitrary data field. The RequesterAnnotation parameter lets your application
	// attach arbitrary data to the HIT for tracking purposes. For example, this
	// parameter could be an identifier internal to the Requester's application that
	// corresponds with the HIT. The RequesterAnnotation parameter for a HIT is only
	// visible to the Requester who created the HIT. It is not shown to the Worker, or
	// any other Requester. The RequesterAnnotation parameter may be different for each
	// HIT you submit. It does not affect how your HITs are grouped.
	RequesterAnnotation *string
	// The data the person completing the HIT uses to produce the results. Constraints:
	// Must be a QuestionForm data structure, an ExternalQuestion data structure, or an
	// HTMLQuestion data structure. The XML question data must not be larger than 64
	// kilobytes (65,535 bytes) in size, including whitespace. Either a Question
	// parameter or a HITLayoutId parameter must be provided.
	Question *string
	// The HITLayoutId allows you to use a pre-existing HIT design with placeholder
	// values and create an additional HIT by providing those values as
	// HITLayoutParameters. Constraints: Either a Question parameter or a HITLayoutId
	// parameter must be provided.
	HITLayoutId *string
	// The HIT-level Review Policy applies to the HIT. You can specify for Mechanical
	// Turk to take various actions based on the policy.
	HITReviewPolicy *types.ReviewPolicy
	// An amount of time, in seconds, after which the HIT is no longer available for
	// users to accept. After the lifetime of the HIT elapses, the HIT no longer
	// appears in HIT searches, even if not all of the assignments for the HIT have
	// been accepted.
	LifetimeInSeconds *int64
}

type CreateHITWithHITTypeOutput struct {
	// Contains the newly created HIT data. For a description of the HIT data structure
	// as it appears in responses, see the HIT Data Structure documentation.
	HIT *types.HIT

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateHITWithHITTypeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateHITWithHITType{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateHITWithHITType{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateHITWithHITType(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "CreateHITWithHITType",
	}
}
