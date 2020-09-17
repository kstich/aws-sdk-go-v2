// Code generated by smithy-go-codegen DO NOT EDIT.

package lakeformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Grants permissions to the principal to access metadata in the Data Catalog and
// data organized in underlying data storage such as Amazon S3. For information
// about permissions, see Security and Access Control to Metadata and Data
// (https://docs-aws.amazon.com/lake-formation/latest/dg/security-data-access.html).
func (c *Client) GrantPermissions(ctx context.Context, params *GrantPermissionsInput, optFns ...func(*Options)) (*GrantPermissionsOutput, error) {
	stack := middleware.NewStack("GrantPermissions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGrantPermissionsMiddlewares(stack)
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
	addOpGrantPermissionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGrantPermissions(options.Region), middleware.Before)
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
			OperationName: "GrantPermissions",
			Err:           err,
		}
	}
	out := result.(*GrantPermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GrantPermissionsInput struct {
	// The permissions granted to the principal on the resource. AWS Lake Formation
	// defines privileges to grant and revoke access to metadata in the Data Catalog
	// and data organized in underlying data storage such as Amazon S3. AWS Lake
	// Formation requires that each principal be authorized to perform a specific task
	// on AWS Lake Formation resources.
	Permissions []types.Permission
	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, and other control information to manage your AWS Lake
	// Formation environment.
	CatalogId *string
	// The principal to be granted the permissions on the resource. Supported
	// principals are IAM users or IAM roles, and they are defined by their principal
	// type and their ARN. Note that if you define a resource with a particular ARN,
	// then later delete, and recreate a resource with that same ARN, the resource
	// maintains the permissions already granted.
	Principal *types.DataLakePrincipal
	// Indicates a list of the granted permissions that the principal may pass to other
	// users. These permissions may only be a subset of the permissions granted in the
	// Privileges.
	PermissionsWithGrantOption []types.Permission
	// The resource to which permissions are to be granted. Resources in AWS Lake
	// Formation are the Data Catalog, databases, and tables.
	Resource *types.Resource
}

type GrantPermissionsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGrantPermissionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGrantPermissions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGrantPermissions{}, middleware.After)
}

func newServiceMetadataMiddleware_opGrantPermissions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lakeformation",
		OperationName: "GrantPermissions",
	}
}
