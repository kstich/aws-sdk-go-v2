// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of AssetSummary
// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_AssetSummary.html)
// objects for assets in a package version.
func (c *Client) ListPackageVersionAssets(ctx context.Context, params *ListPackageVersionAssetsInput, optFns ...func(*Options)) (*ListPackageVersionAssetsOutput, error) {
	stack := middleware.NewStack("ListPackageVersionAssets", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListPackageVersionAssetsMiddlewares(stack)
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
	addOpListPackageVersionAssetsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPackageVersionAssets(options.Region), middleware.Before)
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
			OperationName: "ListPackageVersionAssets",
			Err:           err,
		}
	}
	out := result.(*ListPackageVersionAssetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPackageVersionAssetsInput struct {
	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string
	// A string that contains the package version (for example, 3.5.2).
	PackageVersion *string
	// The name of the domain that contains the repository associated with the package
	// version assets.
	Domain *string
	// The 12-digit account number of the AWS account that owns the domain. It does not
	// include dashes or spaces.
	DomainOwner *string
	// The format of the package that contains the returned package version assets. The
	// valid package types are:
	//
	//     * npm: A Node Package Manager (npm) package.
	//
	//
	// * pypi: A Python Package Index (PyPI) package.
	//
	//     * maven: A Maven package
	// that contains compiled code in a distributable format, such as a JAR file.
	Format types.PackageFormat
	// The maximum number of results to return per page.
	MaxResults *int32
	// The name of the package that contains the returned package version assets.
	Package *string
	// The name of the repository that contains the package that contains the returned
	// package version assets.
	Repository *string
	// The namespace of the package. The package component that specifies its namespace
	// depends on its type. For example:
	//
	//     * The namespace of a Maven package is its
	// groupId.
	//
	//     * The namespace of an npm package is its scope.
	//
	//     * A Python
	// package does not contain a corresponding component, so Python packages do not
	// have a namespace.
	Namespace *string
}

type ListPackageVersionAssetsOutput struct {
	// The version of the package associated with the returned assets.
	Version *string
	// If there are additional results, this is the token for the next set of results.
	NextToken *string
	// The returned list of AssetSummary
	// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_AssetSummary.html)
	// objects.
	Assets []*types.AssetSummary
	// The namespace of the package. The package component that specifies its namespace
	// depends on its type. For example:
	//
	//     * The namespace of a Maven package is its
	// groupId.
	//
	//     * The namespace of an npm package is its scope.
	//
	//     * A Python
	// package does not contain a corresponding component, so Python packages do not
	// have a namespace.
	Namespace *string
	// The format of the package that contains the returned package version assets.
	Format types.PackageFormat
	// The current revision associated with the package version.
	VersionRevision *string
	// The name of the package that contains the returned package version assets.
	Package *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListPackageVersionAssetsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListPackageVersionAssets{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListPackageVersionAssets{}, middleware.After)
}

func newServiceMetadataMiddleware_opListPackageVersionAssets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "ListPackageVersionAssets",
	}
}
