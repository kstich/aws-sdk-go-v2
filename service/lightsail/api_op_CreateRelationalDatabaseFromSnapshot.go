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
	"time"
)

// Creates a new database from an existing database snapshot in Amazon Lightsail.
// You can create a new database from a snapshot in if something goes wrong with
// your original database, or to change it to a different plan, such as a high
// availability or standard plan. The create relational database from snapshot
// operation supports tag-based access control via request tags and resource tags
// applied to the resource identified by relationalDatabaseSnapshotName. For more
// information, see the Lightsail Dev Guide
// (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) CreateRelationalDatabaseFromSnapshot(ctx context.Context, params *CreateRelationalDatabaseFromSnapshotInput, optFns ...func(*Options)) (*CreateRelationalDatabaseFromSnapshotOutput, error) {
	stack := middleware.NewStack("CreateRelationalDatabaseFromSnapshot", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateRelationalDatabaseFromSnapshotMiddlewares(stack)
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
	addOpCreateRelationalDatabaseFromSnapshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRelationalDatabaseFromSnapshot(options.Region), middleware.Before)

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
			OperationName: "CreateRelationalDatabaseFromSnapshot",
			Err:           err,
		}
	}
	out := result.(*CreateRelationalDatabaseFromSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRelationalDatabaseFromSnapshotInput struct {
	// Specifies the accessibility options for your new database. A value of true
	// specifies a database that is available to resources outside of your Lightsail
	// account. A value of false specifies a database that is available only to your
	// Lightsail resources in the same region as your database.
	PubliclyAccessible *bool
	// The name to use for your new database. Constraints:
	//
	//     * Must contain from 2
	// to 255 alphanumeric characters, or hyphens.
	//
	//     * The first and last character
	// must be a letter or number.
	RelationalDatabaseName *string
	// The name of the database snapshot from which to create your new database.
	RelationalDatabaseSnapshotName *string
	// The bundle ID for your new database. A bundle describes the performance
	// specifications for your database. You can get a list of database bundle IDs by
	// using the get relational database bundles operation. When creating a new
	// database from a snapshot, you cannot choose a bundle that is smaller than the
	// bundle of the source database.
	RelationalDatabaseBundleId *string
	// The date and time to restore your database from. Constraints:
	//
	//     * Must be
	// before the latest restorable time for the database.
	//
	//     * Cannot be specified
	// if the use latest restorable time parameter is true.
	//
	//     * Specified in
	// Coordinated Universal Time (UTC).
	//
	//     * Specified in the Unix time format. For
	// example, if you wish to use a restore time of October 1, 2018, at 8 PM UTC, then
	// you input 1538424000 as the restore time.
	RestoreTime *time.Time
	// The tag keys and optional values to add to the resource during create. Use the
	// TagResource action to tag a resource after it's created.
	Tags []*types.Tag
	// The name of the source database.
	SourceRelationalDatabaseName *string
	// The Availability Zone in which to create your new database. Use the us-east-2a
	// case-sensitive format. You can get a list of Availability Zones by using the get
	// regions operation. Be sure to add the include relational database Availability
	// Zones parameter to your request.
	AvailabilityZone *string
	// Specifies whether your database is restored from the latest backup time. A value
	// of true restores from the latest backup time. Default: false Constraints: Cannot
	// be specified if the restore time parameter is provided.
	UseLatestRestorableTime *bool
}

type CreateRelationalDatabaseFromSnapshotOutput struct {
	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []*types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateRelationalDatabaseFromSnapshotMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateRelationalDatabaseFromSnapshot{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateRelationalDatabaseFromSnapshot{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateRelationalDatabaseFromSnapshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "CreateRelationalDatabaseFromSnapshot",
	}
}
