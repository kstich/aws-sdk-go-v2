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
	"io"
)

// Returns an asset (or file) that is in a package. For example, for a Maven
// package version, use GetPackageVersionAsset to download a JAR file, a POM file,
// or any other assets in the package version.
func (c *Client) GetPackageVersionAsset(ctx context.Context, params *GetPackageVersionAssetInput, optFns ...func(*Options)) (*GetPackageVersionAssetOutput, error) {
	stack := middleware.NewStack("GetPackageVersionAsset", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetPackageVersionAssetMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	addOpGetPackageVersionAssetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetPackageVersionAsset(options.Region), middleware.Before)
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
			OperationName: "GetPackageVersionAsset",
			Err:           err,
		}
	}
	out := result.(*GetPackageVersionAssetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPackageVersionAssetInput struct {
	// A string that contains the package version (for example, 3.5.2).
	PackageVersion *string
	// The domain that contains the repository that contains the package version with
	// the requested asset.
	Domain *string
	// The name of the package that contains the requested asset.
	Package *string
	// The name of the package version revision that contains the requested asset.
	PackageVersionRevision *string
	// A format that specifies the type of the package version with the requested asset
	// file. The valid values are:
	//
	//     * npm
	//
	//     * pypi
	//
	//     * maven
	Format types.PackageFormat
	// The 12-digit account number of the AWS account that owns the domain. It does not
	// include dashes or spaces.
	DomainOwner *string
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
	// The name of the requested asset.
	Asset *string
	// The repository that contains the package version with the requested asset.
	Repository *string
}

type GetPackageVersionAssetOutput struct {
	// The name of the package version revision that contains the downloaded asset.
	PackageVersionRevision *string
	// A string that contains the package version (for example, 3.5.2).
	PackageVersion *string
	// The binary file, or asset, that is downloaded.
	Asset io.ReadCloser
	// The name of the asset that is downloaded.
	AssetName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetPackageVersionAssetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetPackageVersionAsset{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetPackageVersionAsset{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetPackageVersionAsset(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "GetPackageVersionAsset",
	}
}
