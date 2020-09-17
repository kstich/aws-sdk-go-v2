// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the description of specific Amazon FSx file systems, if a FileSystemIds
// value is provided for that file system. Otherwise, it returns descriptions of
// all file systems owned by your AWS account in the AWS Region of the endpoint
// that you're calling.  <p>When retrieving all file system descriptions, you can
// optionally specify the <code>MaxResults</code> parameter to limit the number of
// descriptions in a response. If more file system descriptions remain, Amazon FSx
// returns a <code>NextToken</code> value in the response. In this case, send a
// later request with the <code>NextToken</code> request parameter set to the value
// of <code>NextToken</code> from the last response.</p> <p>This action is used in
// an iterative process to retrieve a list of your file system descriptions.
// <code>DescribeFileSystems</code> is called first without a
// <code>NextToken</code>value. Then the action continues to be called with the
// <code>NextToken</code> parameter set to the value of the last
// <code>NextToken</code> value until a response has no <code>NextToken</code>.</p>
// <p>When using this action, keep the following in mind:</p> <ul> <li> <p>The
// implementation might return fewer than <code>MaxResults</code> file system
// descriptions while still including a <code>NextToken</code> value.</p> </li>
// <li> <p>The order of file systems returned in the response of one
// <code>DescribeFileSystems</code> call and the order of file systems returned
// across the responses of a multicall iteration is unspecified.</p> </li> </ul>
func (c *Client) DescribeFileSystems(ctx context.Context, params *DescribeFileSystemsInput, optFns ...func(*Options)) (*DescribeFileSystemsOutput, error) {
	stack := middleware.NewStack("DescribeFileSystems", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeFileSystemsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFileSystems(options.Region), middleware.Before)
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
			OperationName: "DescribeFileSystems",
			Err:           err,
		}
	}
	out := result.(*DescribeFileSystemsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request object for DescribeFileSystems operation.
type DescribeFileSystemsInput struct {
	// IDs of the file systems whose descriptions you want to retrieve (String).
	FileSystemIds []*string
	// Maximum number of file systems to return in the response (integer). This
	// parameter value must be greater than 0. The number of items that Amazon FSx
	// returns is the minimum of the MaxResults parameter specified in the request and
	// the service's internal maximum number of items per page.
	MaxResults *int32
	// Opaque pagination token returned from a previous DescribeFileSystems operation
	// (String). If a token present, the action continues the list from where the
	// returning call left off.
	NextToken *string
}

// The response object for DescribeFileSystems operation.
type DescribeFileSystemsOutput struct {
	// An array of file system descriptions.
	FileSystems []*types.FileSystem
	// Present if there are more file systems than returned in the response (String).
	// You can use the NextToken value in the later request to fetch the descriptions.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeFileSystemsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeFileSystems{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeFileSystems{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeFileSystems(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fsx",
		OperationName: "DescribeFileSystems",
	}
}
