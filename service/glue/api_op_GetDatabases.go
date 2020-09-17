// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves all databases defined in a given Data Catalog.
func (c *Client) GetDatabases(ctx context.Context, params *GetDatabasesInput, optFns ...func(*Options)) (*GetDatabasesOutput, error) {
	stack := middleware.NewStack("GetDatabases", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetDatabasesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDatabases(options.Region), middleware.Before)
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
			OperationName: "GetDatabases",
			Err:           err,
		}
	}
	out := result.(*GetDatabasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDatabasesInput struct {
	// A continuation token, if this is a continuation call.
	NextToken *string
	// Allows you to specify that you want to list the databases shared with your
	// account. The allowable values are FOREIGN or ALL.  <ul> <li> <p>If set to
	// <code>FOREIGN</code>, will list the databases shared with your account. </p>
	// </li> <li> <p>If set to <code>ALL</code>, will list the databases shared with
	// your account, as well as the databases in yor local account. </p> </li> </ul>
	ResourceShareType types.ResourceShareType
	// The maximum number of databases to return in one response.
	MaxResults *int32
	// The ID of the Data Catalog from which to retrieve Databases. If none is
	// provided, the AWS account ID is used by default.
	CatalogId *string
}

type GetDatabasesOutput struct {
	// A list of Database objects from the specified catalog.
	DatabaseList []*types.Database
	// A continuation token for paginating the returned list of tokens, returned if the
	// current segment of the list is not the last.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetDatabasesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetDatabases{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDatabases{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDatabases(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetDatabases",
	}
}
