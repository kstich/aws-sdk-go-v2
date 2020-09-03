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

// Gets information about triggers configured for a repository.
func (c *Client) GetRepositoryTriggers(ctx context.Context, params *GetRepositoryTriggersInput, optFns ...func(*Options)) (*GetRepositoryTriggersOutput, error) {
	stack := middleware.NewStack("GetRepositoryTriggers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetRepositoryTriggersMiddlewares(stack)
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
	addOpGetRepositoryTriggersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetRepositoryTriggers(options.Region), middleware.Before)

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
			OperationName: "GetRepositoryTriggers",
			Err:           err,
		}
	}
	out := result.(*GetRepositoryTriggersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a get repository triggers operation.
type GetRepositoryTriggersInput struct {
	// The name of the repository for which the trigger is configured.
	RepositoryName *string
}

// Represents the output of a get repository triggers operation.
type GetRepositoryTriggersOutput struct {
	// The system-generated unique ID for the trigger.
	ConfigurationId *string
	// The JSON block of configuration information for each trigger.
	Triggers []*types.RepositoryTrigger

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetRepositoryTriggersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetRepositoryTriggers{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRepositoryTriggers{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetRepositoryTriggers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "GetRepositoryTriggers",
	}
}
