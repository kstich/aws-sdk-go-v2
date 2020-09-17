// Code generated by smithy-go-codegen DO NOT EDIT.

package snowball

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the UnlockCode code value for the specified job. A particular UnlockCode
// value can be accessed for up to 90 days after the associated job has been
// created.  <p>The <code>UnlockCode</code> value is a 29-character code with 25
// alphanumeric characters and 4 hyphens. This code is used to decrypt the manifest
// file when it is passed along with the manifest to the Snowball through the
// Snowball client when the client is started for the first time.</p> <p>As a best
// practice, we recommend that you don't save a copy of the <code>UnlockCode</code>
// in the same location as the manifest file for that job. Saving these separately
// helps prevent unauthorized parties from gaining access to the Snowball
// associated with that job.</p>
func (c *Client) GetJobUnlockCode(ctx context.Context, params *GetJobUnlockCodeInput, optFns ...func(*Options)) (*GetJobUnlockCodeOutput, error) {
	stack := middleware.NewStack("GetJobUnlockCode", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetJobUnlockCodeMiddlewares(stack)
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
	addOpGetJobUnlockCodeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetJobUnlockCode(options.Region), middleware.Before)
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
			OperationName: "GetJobUnlockCode",
			Err:           err,
		}
	}
	out := result.(*GetJobUnlockCodeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetJobUnlockCodeInput struct {
	// The ID for the job that you want to get the UnlockCode value for, for example
	// JID123e4567-e89b-12d3-a456-426655440000.
	JobId *string
}

type GetJobUnlockCodeOutput struct {
	// The UnlockCode value for the specified job. The UnlockCode value can be accessed
	// for up to 90 days after the job has been created.
	UnlockCode *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetJobUnlockCodeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetJobUnlockCode{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetJobUnlockCode{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetJobUnlockCode(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "snowball",
		OperationName: "GetJobUnlockCode",
	}
}
