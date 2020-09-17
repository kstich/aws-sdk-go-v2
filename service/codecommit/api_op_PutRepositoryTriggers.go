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

// Replaces all triggers for a repository. Used to create or delete triggers.
func (c *Client) PutRepositoryTriggers(ctx context.Context, params *PutRepositoryTriggersInput, optFns ...func(*Options)) (*PutRepositoryTriggersOutput, error) {
	stack := middleware.NewStack("PutRepositoryTriggers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutRepositoryTriggersMiddlewares(stack)
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
	addOpPutRepositoryTriggersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutRepositoryTriggers(options.Region), middleware.Before)
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
			OperationName: "PutRepositoryTriggers",
			Err:           err,
		}
	}
	out := result.(*PutRepositoryTriggersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a put repository triggers operation.
type PutRepositoryTriggersInput struct {
	// The JSON block of configuration information for each trigger.
	Triggers []*types.RepositoryTrigger
	// The name of the repository where you want to create or update the trigger.
	RepositoryName *string
}

// Represents the output of a put repository triggers operation.
type PutRepositoryTriggersOutput struct {
	// The system-generated unique ID for the create or update operation.
	ConfigurationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutRepositoryTriggersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutRepositoryTriggers{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutRepositoryTriggers{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutRepositoryTriggers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "PutRepositoryTriggers",
	}
}
