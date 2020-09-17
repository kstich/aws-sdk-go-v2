// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all packages associated with the Amazon ES domain.
func (c *Client) ListPackagesForDomain(ctx context.Context, params *ListPackagesForDomainInput, optFns ...func(*Options)) (*ListPackagesForDomainOutput, error) {
	stack := middleware.NewStack("ListPackagesForDomain", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListPackagesForDomainMiddlewares(stack)
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
	addOpListPackagesForDomainValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPackagesForDomain(options.Region), middleware.Before)
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
			OperationName: "ListPackagesForDomain",
			Err:           err,
		}
	}
	out := result.(*ListPackagesForDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for request parameters to ListPackagesForDomain () operation.
type ListPackagesForDomainInput struct {
	// The name of the domain for which you want to list associated packages.
	DomainName *string
	// Limits results to a maximum number of packages.
	MaxResults *int32
	// Used for pagination. Only necessary if a previous API call includes a non-null
	// NextToken value. If provided, returns results for the next page.
	NextToken *string
}

// Container for response parameters to ListPackagesForDomain () operation.
type ListPackagesForDomainOutput struct {
	// Pagination token that needs to be supplied to the next call to get the next page
	// of results.
	NextToken *string
	// List of DomainPackageDetails objects.
	DomainPackageDetailsList []*types.DomainPackageDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListPackagesForDomainMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListPackagesForDomain{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListPackagesForDomain{}, middleware.After)
}

func newServiceMetadataMiddleware_opListPackagesForDomain(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "ListPackagesForDomain",
	}
}
