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

// Returns the Lake Formation permissions for a specified table or database
// resource located at a path in Amazon S3. GetEffectivePermissionsForPath will not
// return databases and tables if the catalog is encrypted.
func (c *Client) GetEffectivePermissionsForPath(ctx context.Context, params *GetEffectivePermissionsForPathInput, optFns ...func(*Options)) (*GetEffectivePermissionsForPathOutput, error) {
	stack := middleware.NewStack("GetEffectivePermissionsForPath", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetEffectivePermissionsForPathMiddlewares(stack)
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
	addOpGetEffectivePermissionsForPathValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetEffectivePermissionsForPath(options.Region), middleware.Before)
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
			OperationName: "GetEffectivePermissionsForPath",
			Err:           err,
		}
	}
	out := result.(*GetEffectivePermissionsForPathOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEffectivePermissionsForPathInput struct {
	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, and other control information to manage your AWS Lake
	// Formation environment.
	CatalogId *string
	// The Amazon Resource Name (ARN) of the resource for which you want to get
	// permissions.
	ResourceArn *string
	// The maximum number of results to return.
	MaxResults *int32
	// A continuation token, if this is not the first call to retrieve this list.
	NextToken *string
}

type GetEffectivePermissionsForPathOutput struct {
	// A list of the permissions for the specified table or database resource located
	// at the path in Amazon S3.
	Permissions []*types.PrincipalResourcePermissions
	// A continuation token, if this is not the first call to retrieve this list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetEffectivePermissionsForPathMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetEffectivePermissionsForPath{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetEffectivePermissionsForPath{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetEffectivePermissionsForPath(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lakeformation",
		OperationName: "GetEffectivePermissionsForPath",
	}
}
