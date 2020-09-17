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

// Describes which data sources are enabled for the member account's detector.
func (c *Client) GetMemberDetectors(ctx context.Context, params *GetMemberDetectorsInput, optFns ...func(*Options)) (*GetMemberDetectorsOutput, error) {
	stack := middleware.NewStack("GetMemberDetectors", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetMemberDetectorsMiddlewares(stack)
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
	addOpGetMemberDetectorsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetMemberDetectors(options.Region), middleware.Before)
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
			OperationName: "GetMemberDetectors",
			Err:           err,
		}
	}
	out := result.(*GetMemberDetectorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMemberDetectorsInput struct {
	// The detector ID for the master account.
	DetectorId *string
	// The account ID of the member account.
	AccountIds []*string
}

type GetMemberDetectorsOutput struct {
	// A list of member account IDs that were unable to be processed along with an
	// explanation for why they were not processed.
	UnprocessedAccounts []*types.UnprocessedAccount
	// An object that describes which data sources are enabled for a member account.
	MemberDataSourceConfigurations []*types.MemberDataSourceConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetMemberDetectorsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetMemberDetectors{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMemberDetectors{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetMemberDetectors(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "GetMemberDetectors",
	}
}
