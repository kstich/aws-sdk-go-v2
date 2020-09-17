// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieve information about one or more parameters in a specific hierarchy.
// Request results are returned on a best-effort basis. If you specify MaxResults
// in the request, the response includes information up to the limit specified. The
// number of items returned, however, can be between zero and the value of
// MaxResults. If the service reaches an internal limit while processing the
// results, it stops the operation and returns the matching values up to that point
// and a NextToken. You can specify the NextToken in a subsequent call to get the
// next set of results.
func (c *Client) GetParametersByPath(ctx context.Context, params *GetParametersByPathInput, optFns ...func(*Options)) (*GetParametersByPathOutput, error) {
	stack := middleware.NewStack("GetParametersByPath", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetParametersByPathMiddlewares(stack)
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
	addOpGetParametersByPathValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetParametersByPath(options.Region), middleware.Before)
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
			OperationName: "GetParametersByPath",
			Err:           err,
		}
	}
	out := result.(*GetParametersByPathOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetParametersByPathInput struct {
	// A token to start the list. Use this token to get the next set of results.
	NextToken *string
	// Retrieve all parameters within a hierarchy. If a user has access to a path, then
	// the user can access all levels of that path. For example, if a user has
	// permission to access path /a, then the user can also access /a/b. Even if a user
	// has explicitly been denied access in IAM for parameter /a/b, they can still call
	// the GetParametersByPath API action recursively for /a and view /a/b.
	Recursive *bool
	// The hierarchy for the parameter. Hierarchies start with a forward slash (/) and
	// end with the parameter name. A parameter name hierarchy can have a maximum of 15
	// levels. Here is an example of a hierarchy:
	// /Finance/Prod/IAD/WinServ2016/license33
	Path *string
	// Filters to limit the request results.
	ParameterFilters []*types.ParameterStringFilter
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32
	// Retrieve all parameters in a hierarchy with their value decrypted.
	WithDecryption *bool
}

type GetParametersByPathOutput struct {
	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string
	// A list of parameters found in the specified hierarchy.
	Parameters []*types.Parameter

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetParametersByPathMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetParametersByPath{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetParametersByPath{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetParametersByPath(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "GetParametersByPath",
	}
}
