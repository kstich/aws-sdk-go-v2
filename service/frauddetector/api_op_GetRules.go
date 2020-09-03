// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Get all rules for a detector (paginated) if ruleId and ruleVersion are not
// specified. Gets all rules for the detector and the ruleId if present
// (paginated). Gets a specific rule if both the ruleId and the ruleVersion are
// specified. This is a paginated API. Providing null maxResults results in
// retrieving maximum of 100 records per page. If you provide maxResults the value
// must be between 50 and 100. To get the next page result, a provide a pagination
// token from GetRulesResult as part of your request. Null pagination token fetches
// the records from the beginning.
func (c *Client) GetRules(ctx context.Context, params *GetRulesInput, optFns ...func(*Options)) (*GetRulesOutput, error) {
	stack := middleware.NewStack("GetRules", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetRulesMiddlewares(stack)
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
	addOpGetRulesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetRules(options.Region), middleware.Before)

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
			OperationName: "GetRules",
			Err:           err,
		}
	}
	out := result.(*GetRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRulesInput struct {
	// The maximum number of rules to return for the request.
	MaxResults *int32
	// The detector ID.
	DetectorId *string
	// The rule version.
	RuleVersion *string
	// The next page token.
	NextToken *string
	// The rule ID.
	RuleId *string
}

type GetRulesOutput struct {
	// The details of the requested rule.
	RuleDetails []*types.RuleDetail
	// The next page token to be used in subsequent requests.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetRulesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetRules{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRules{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetRules(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "GetRules",
	}
}
