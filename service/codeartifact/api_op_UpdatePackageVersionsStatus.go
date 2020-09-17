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

// Updates the status of one or more versions of a package.
func (c *Client) UpdatePackageVersionsStatus(ctx context.Context, params *UpdatePackageVersionsStatusInput, optFns ...func(*Options)) (*UpdatePackageVersionsStatusOutput, error) {
	stack := middleware.NewStack("UpdatePackageVersionsStatus", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdatePackageVersionsStatusMiddlewares(stack)
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
	addOpUpdatePackageVersionsStatusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePackageVersionsStatus(options.Region), middleware.Before)
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
			OperationName: "UpdatePackageVersionsStatus",
			Err:           err,
		}
	}
	out := result.(*UpdatePackageVersionsStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePackageVersionsStatusInput struct {
	// The domain that contains the repository that contains the package versions with
	// a status to be updated.
	Domain *string
	// An array of strings that specify the versions of the package with the statuses
	// to update.
	Versions []*string
	// The package version’s expected status before it is updated. If expectedStatus is
	// provided, the package version's status is updated only if its status at the time
	// UpdatePackageVersionsStatus is called matches expectedStatus.
	ExpectedStatus types.PackageVersionStatus
	// The name of the package with the version statuses to update.
	Package *string
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
	// A map of package versions and package version revisions. The map key is the
	// package version (for example, 3.5.2), and the map value is the package version
	// revision.
	VersionRevisions map[string]*string
	// A format that specifies the type of the package with the statuses to update. The
	// valid values are:
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
	// The repository that contains the package versions with the status you want to
	// update.
	Repository *string
	// The status you want to change the package version status to.
	TargetStatus types.PackageVersionStatus
}

type UpdatePackageVersionsStatusOutput struct {
	// A list of SuccessfulPackageVersionInfo objects, one for each package version
	// with a status that successfully updated.
	FailedVersions map[string]*types.PackageVersionError
	// A list of PackageVersionError objects, one for each package version with a
	// status that failed to update.
	SuccessfulVersions map[string]*types.SuccessfulPackageVersionInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdatePackageVersionsStatusMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdatePackageVersionsStatus{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdatePackageVersionsStatus{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdatePackageVersionsStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "UpdatePackageVersionsStatus",
	}
}
