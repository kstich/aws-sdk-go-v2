// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation returns the current data retrieval policy for the account and
// region specified in the GET request. For more information about data retrieval
// policies, see Amazon Glacier Data Retrieval Policies
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/data-retrieval-policy.html).
func (c *Client) GetDataRetrievalPolicy(ctx context.Context, params *GetDataRetrievalPolicyInput, optFns ...func(*Options)) (*GetDataRetrievalPolicyOutput, error) {
	stack := middleware.NewStack("GetDataRetrievalPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetDataRetrievalPolicyMiddlewares(stack)
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
	addOpGetDataRetrievalPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDataRetrievalPolicy(options.Region), middleware.Before)
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
			OperationName: "GetDataRetrievalPolicy",
			Err:           err,
		}
	}
	out := result.(*GetDataRetrievalPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for GetDataRetrievalPolicy.
type GetDataRetrievalPolicyInput struct {
	// The AccountId value is the AWS account ID. This value must match the AWS account
	// ID associated with the credentials used to sign the request. You can either
	// specify an AWS account ID or optionally a single '-' (hyphen), in which case
	// Amazon Glacier uses the AWS account ID associated with the credentials used to
	// sign the request. If you specify your account ID, do not include any hyphens
	// ('-') in the ID.
	AccountId *string
}

// Contains the Amazon S3 Glacier response to the GetDataRetrievalPolicy request.
type GetDataRetrievalPolicyOutput struct {
	// Contains the returned data retrieval policy in JSON format.
	Policy *types.DataRetrievalPolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetDataRetrievalPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetDataRetrievalPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDataRetrievalPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDataRetrievalPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "GetDataRetrievalPolicy",
	}
}
