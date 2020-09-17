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

// Tests the functionality of repository triggers by sending information to the
// trigger target. If real data is available in the repository, the test sends data
// from the last commit. If no data is available, sample data is generated.
func (c *Client) TestRepositoryTriggers(ctx context.Context, params *TestRepositoryTriggersInput, optFns ...func(*Options)) (*TestRepositoryTriggersOutput, error) {
	stack := middleware.NewStack("TestRepositoryTriggers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpTestRepositoryTriggersMiddlewares(stack)
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
	addOpTestRepositoryTriggersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTestRepositoryTriggers(options.Region), middleware.Before)
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
			OperationName: "TestRepositoryTriggers",
			Err:           err,
		}
	}
	out := result.(*TestRepositoryTriggersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a test repository triggers operation.
type TestRepositoryTriggersInput struct {
	// The name of the repository in which to test the triggers.
	RepositoryName *string
	// The list of triggers to test.
	Triggers []*types.RepositoryTrigger
}

// Represents the output of a test repository triggers operation.
type TestRepositoryTriggersOutput struct {
	// The list of triggers that were successfully tested. This list provides the names
	// of the triggers that were successfully tested, separated by commas.
	SuccessfulExecutions []*string
	// The list of triggers that were not tested. This list provides the names of the
	// triggers that could not be tested, separated by commas.
	FailedExecutions []*types.RepositoryTriggerExecutionFailure

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpTestRepositoryTriggersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpTestRepositoryTriggers{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpTestRepositoryTriggers{}, middleware.After)
}

func newServiceMetadataMiddleware_opTestRepositoryTriggers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "TestRepositoryTriggers",
	}
}
