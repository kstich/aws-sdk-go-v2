// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the names of all trigger resources in this AWS account, or the
// resources with the specified tag. This operation allows you to see which
// resources are available in your account, and their names.  <p>This operation
// takes the optional <code>Tags</code> field, which you can use as a filter on the
// response so that tagged resources can be retrieved as a group. If you choose to
// use tags filtering, only resources with the tag are retrieved.</p>
func (c *Client) ListTriggers(ctx context.Context, params *ListTriggersInput, optFns ...func(*Options)) (*ListTriggersOutput, error) {
	stack := middleware.NewStack("ListTriggers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListTriggersMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTriggers(options.Region), middleware.Before)
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
			OperationName: "ListTriggers",
			Err:           err,
		}
	}
	out := result.(*ListTriggersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTriggersInput struct {
	// The name of the job for which to retrieve triggers. The trigger that can start
	// this job is returned. If there is no such trigger, all triggers are returned.
	DependentJobName *string
	// The maximum size of a list to return.
	MaxResults *int32
	// Specifies to return only these tagged resources.
	Tags map[string]*string
	// A continuation token, if this is a continuation request.
	NextToken *string
}

type ListTriggersOutput struct {
	// The names of all triggers in the account, or the triggers with the specified
	// tags.
	TriggerNames []*string
	// A continuation token, if the returned list does not contain the last metric
	// available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListTriggersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListTriggers{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTriggers{}, middleware.After)
}

func newServiceMetadataMiddleware_opListTriggers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "ListTriggers",
	}
}
