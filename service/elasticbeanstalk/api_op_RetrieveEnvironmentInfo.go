// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the compiled information from a RequestEnvironmentInfo () request.
// Related Topics
//
//     * RequestEnvironmentInfo ()
func (c *Client) RetrieveEnvironmentInfo(ctx context.Context, params *RetrieveEnvironmentInfoInput, optFns ...func(*Options)) (*RetrieveEnvironmentInfoOutput, error) {
	stack := middleware.NewStack("RetrieveEnvironmentInfo", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpRetrieveEnvironmentInfoMiddlewares(stack)
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
	addOpRetrieveEnvironmentInfoValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRetrieveEnvironmentInfo(options.Region), middleware.Before)

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
			OperationName: "RetrieveEnvironmentInfo",
			Err:           err,
		}
	}
	out := result.(*RetrieveEnvironmentInfoOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to download logs retrieved with RequestEnvironmentInfo ().
type RetrieveEnvironmentInfoInput struct {
	// The ID of the data's environment. If no such environment is found, returns an
	// InvalidParameterValue error. Condition: You must specify either this or an
	// EnvironmentName, or both. If you do not specify either, AWS Elastic Beanstalk
	// returns MissingRequiredParameter error.
	EnvironmentId *string
	// The name of the data's environment. If no such environment is found, returns an
	// InvalidParameterValue error. Condition: You must specify either this or an
	// EnvironmentId, or both. If you do not specify either, AWS Elastic Beanstalk
	// returns MissingRequiredParameter error.
	EnvironmentName *string
	// The type of information to retrieve.
	InfoType types.EnvironmentInfoType
}

// Result message containing a description of the requested environment info.
type RetrieveEnvironmentInfoOutput struct {
	// The EnvironmentInfoDescription () of the environment.
	EnvironmentInfo []*types.EnvironmentInfoDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpRetrieveEnvironmentInfoMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpRetrieveEnvironmentInfo{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpRetrieveEnvironmentInfo{}, middleware.After)
}

func newServiceMetadataMiddleware_opRetrieveEnvironmentInfo(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "RetrieveEnvironmentInfo",
	}
}
