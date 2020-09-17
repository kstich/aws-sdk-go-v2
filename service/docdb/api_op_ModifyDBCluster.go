// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies a setting for an Amazon DocumentDB cluster. You can change one or more
// database configuration parameters by specifying these parameters and the new
// values in the request.
func (c *Client) ModifyDBCluster(ctx context.Context, params *ModifyDBClusterInput, optFns ...func(*Options)) (*ModifyDBClusterOutput, error) {
	stack := middleware.NewStack("ModifyDBCluster", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpModifyDBClusterMiddlewares(stack)
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
	addOpModifyDBClusterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDBCluster(options.Region), middleware.Before)
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
			OperationName: "ModifyDBCluster",
			Err:           err,
		}
	}
	out := result.(*ModifyDBClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to ModifyDBCluster ().
type ModifyDBClusterInput struct {
	// The new cluster identifier for the cluster when renaming a cluster. This value
	// is stored as a lowercase string. Constraints:
	//
	//     * Must contain from 1 to 63
	// letters, numbers, or hyphens.
	//
	//     * The first character must be a letter.
	//
	//
	// * Cannot end with a hyphen or contain two consecutive hyphens.
	//
	// Example:
	// my-cluster2
	NewDBClusterIdentifier *string
	// Specifies whether this cluster can be deleted. If DeletionProtection is enabled,
	// the cluster cannot be deleted unless it is modified and DeletionProtection is
	// disabled. DeletionProtection protects clusters from being accidentally deleted.
	DeletionProtection *bool
	// The cluster identifier for the cluster that is being modified. This parameter is
	// not case sensitive. Constraints:
	//
	//     * Must match the identifier of an existing
	// DBCluster.
	DBClusterIdentifier *string
	// The weekly time range during which system maintenance can occur, in Universal
	// Coordinated Time (UTC). Format: ddd:hh24:mi-ddd:hh24:mi The default is a
	// 30-minute window selected at random from an 8-hour block of time for each AWS
	// Region, occurring on a random day of the week. Valid days: Mon, Tue, Wed, Thu,
	// Fri, Sat, Sun Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string
	// The port number on which the cluster accepts connections. Constraints: Must be a
	// value from 1150 to 65535. Default: The same port as the original cluster.
	Port *int32
	// The version number of the database engine to which you want to upgrade. Changing
	// this parameter results in an outage. The change is applied during the next
	// maintenance window unless the ApplyImmediately parameter is set to true.
	EngineVersion *string
	// The password for the master database user. This password can contain any
	// printable ASCII character except forward slash (/), double quote ("), or the
	// "at" symbol (@). Constraints: Must contain from 8 to 100 characters.
	MasterUserPassword *string
	// A list of virtual private cloud (VPC) security groups that the cluster will
	// belong to.
	VpcSecurityGroupIds []*string
	// The daily time range during which automated backups are created if automated
	// backups are enabled, using the BackupRetentionPeriod parameter. The default is a
	// 30-minute window selected at random from an 8-hour block of time for each AWS
	// Region. Constraints:
	//
	//     * Must be in the format hh24:mi-hh24:mi.
	//
	//     * Must
	// be in Universal Coordinated Time (UTC).
	//
	//     * Must not conflict with the
	// preferred maintenance window.
	//
	//     * Must be at least 30 minutes.
	PreferredBackupWindow *string
	// The name of the cluster parameter group to use for the cluster.
	DBClusterParameterGroupName *string
	// The number of days for which automated backups are retained. You must specify a
	// minimum value of 1. Default: 1 Constraints:
	//
	//     * Must be a value from 1 to 35.
	BackupRetentionPeriod *int32
	// A value that specifies whether the changes in this request and any pending
	// changes are asynchronously applied as soon as possible, regardless of the
	// PreferredMaintenanceWindow setting for the cluster. If this parameter is set to
	// false, changes to the cluster are applied during the next maintenance window.
	// The ApplyImmediately parameter affects only the NewDBClusterIdentifier and
	// MasterUserPassword values. If you set this parameter value to false, the changes
	// to the NewDBClusterIdentifier and MasterUserPassword values are applied during
	// the next maintenance window. All other changes are applied immediately,
	// regardless of the value of the ApplyImmediately parameter. Default: false
	ApplyImmediately *bool
	// The configuration setting for the log types to be enabled for export to Amazon
	// CloudWatch Logs for a specific instance or cluster. The EnableLogTypes and
	// DisableLogTypes arrays determine which logs are exported (or not exported) to
	// CloudWatch Logs.
	CloudwatchLogsExportConfiguration *types.CloudwatchLogsExportConfiguration
}

type ModifyDBClusterOutput struct {
	// Detailed information about a cluster.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpModifyDBClusterMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpModifyDBCluster{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyDBCluster{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyDBCluster(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ModifyDBCluster",
	}
}
