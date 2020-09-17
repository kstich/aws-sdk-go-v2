// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the list of bundles that are available in Amazon Lightsail. A bundle
// describes the performance specifications for a database. You can use a bundle ID
// to create a new database with explicit performance specifications.
func (c *Client) GetRelationalDatabaseBundles(ctx context.Context, params *GetRelationalDatabaseBundlesInput, optFns ...func(*Options)) (*GetRelationalDatabaseBundlesOutput, error) {
	stack := middleware.NewStack("GetRelationalDatabaseBundles", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetRelationalDatabaseBundlesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetRelationalDatabaseBundles(options.Region), middleware.Before)
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
			OperationName: "GetRelationalDatabaseBundles",
			Err:           err,
		}
	}
	out := result.(*GetRelationalDatabaseBundlesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRelationalDatabaseBundlesInput struct {
	// The token to advance to the next page of results from your request. To get a
	// page token, perform an initial GetRelationalDatabaseBundles request. If your
	// results are paginated, the response will return a next page token that you can
	// specify as the page token in a subsequent request.
	PageToken *string
}

type GetRelationalDatabaseBundlesOutput struct {
	// The token to advance to the next page of resutls from your request. A next page
	// token is not returned if there are no more results to display. To get the next
	// page of results, perform another GetRelationalDatabaseBundles request and
	// specify the next page token using the pageToken parameter.
	NextPageToken *string
	// An object describing the result of your get relational database bundles request.
	Bundles []*types.RelationalDatabaseBundle

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetRelationalDatabaseBundlesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetRelationalDatabaseBundles{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRelationalDatabaseBundles{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetRelationalDatabaseBundles(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetRelationalDatabaseBundles",
	}
}
