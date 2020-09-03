// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a build. This action permanently deletes the build resource and any
// uploaded build files. Deleting a build does not affect the status of any active
// fleets using the build, but you can no longer create new fleets with the deleted
// build. To delete a build, specify the build ID. Learn more  Upload a Custom
// Server Build
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-build-intro.html)
// Related operations
//
//     * CreateBuild ()
//
//     * ListBuilds ()
//
//     *
// DescribeBuild ()
//
//     * UpdateBuild ()
//
//     * DeleteBuild ()
func (c *Client) DeleteBuild(ctx context.Context, params *DeleteBuildInput, optFns ...func(*Options)) (*DeleteBuildOutput, error) {
	stack := middleware.NewStack("DeleteBuild", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteBuildMiddlewares(stack)
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
	addOpDeleteBuildValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBuild(options.Region), middleware.Before)

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
			OperationName: "DeleteBuild",
			Err:           err,
		}
	}
	out := result.(*DeleteBuildOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type DeleteBuildInput struct {
	// A unique identifier for a build to delete. You can use either the build ID or
	// ARN value.
	BuildId *string
}

type DeleteBuildOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteBuildMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteBuild{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteBuild{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteBuild(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DeleteBuild",
	}
}
