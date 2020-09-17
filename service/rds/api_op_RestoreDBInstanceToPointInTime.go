// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Restores a DB instance to an arbitrary point in time. You can restore to any
// point in time before the time identified by the LatestRestorableTime property.
// You can restore to a point up to the number of days specified by the
// BackupRetentionPeriod property. The target database is created with most of the
// original configuration, but in a system-selected Availability Zone, with the
// default security group, the default subnet group, and the default DB parameter
// group. By default, the new DB instance is created as a single-AZ deployment
// except when the instance is a SQL Server instance that has an option group that
// is associated with mirroring; in this case, the instance becomes a mirrored
// deployment and not a single-AZ deployment. This command doesn't apply to Aurora
// MySQL and Aurora PostgreSQL. For Aurora, use RestoreDBClusterToPointInTime.
func (c *Client) RestoreDBInstanceToPointInTime(ctx context.Context, params *RestoreDBInstanceToPointInTimeInput, optFns ...func(*Options)) (*RestoreDBInstanceToPointInTimeOutput, error) {
	stack := middleware.NewStack("RestoreDBInstanceToPointInTime", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpRestoreDBInstanceToPointInTimeMiddlewares(stack)
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
	addOpRestoreDBInstanceToPointInTimeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreDBInstanceToPointInTime(options.Region), middleware.Before)
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
			OperationName: "RestoreDBInstanceToPointInTime",
			Err:           err,
		}
	}
	out := result.(*RestoreDBInstanceToPointInTimeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type RestoreDBInstanceToPointInTimeInput struct {
	// The identifier of the source DB instance from which to restore. Constraints:
	//
	//
	// * Must match the identifier of an existing DB instance.
	SourceDBInstanceIdentifier *string
	// License model information for the restored DB instance. Default: Same as source.
	// Valid values: license-included | bring-your-own-license | general-public-license
	LicenseModel *string
	// The database engine to use for the new instance. Default: The same as source
	// Constraint: Must be compatible with the engine of the source  <p>Valid
	// Values:</p> <ul> <li> <p> <code>mariadb</code> </p> </li> <li> <p>
	// <code>mysql</code> </p> </li> <li> <p> <code>oracle-ee</code> </p> </li> <li>
	// <p> <code>oracle-se2</code> </p> </li> <li> <p> <code>oracle-se1</code> </p>
	// </li> <li> <p> <code>oracle-se</code> </p> </li> <li> <p> <code>postgres</code>
	// </p> </li> <li> <p> <code>sqlserver-ee</code> </p> </li> <li> <p>
	// <code>sqlserver-se</code> </p> </li> <li> <p> <code>sqlserver-ex</code> </p>
	// </li> <li> <p> <code>sqlserver-web</code> </p> </li> </ul>
	Engine *string
	// The amount of Provisioned IOPS (input/output operations per second) to be
	// initially allocated for the DB instance. Constraints: Must be an integer greater
	// than 1000. SQL Server Setting the IOPS value for the SQL Server database engine
	// isn't supported.
	Iops *int32
	// The list of logs that the restored DB instance is to export to CloudWatch Logs.
	// The values in the list depend on the DB engine being used. For more information,
	// see Publishing Database Logs to Amazon CloudWatch Logs
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)
	// in the Amazon RDS User Guide.
	EnableCloudwatchLogsExports []*string
	// Specifies the storage type to be associated with the DB instance. Valid values:
	// standard | gp2 | io1 If you specify io1, you must also include a value for the
	// Iops parameter. Default: io1 if the Iops parameter is specified, otherwise gp2
	StorageType *string
	// The password for the given ARN from the key store in order to access the device.
	TdeCredentialPassword *string
	// The DB subnet group name to use for the new instance. Constraints: If supplied,
	// must match the name of an existing DBSubnetGroup. Example: mySubnetgroup
	DBSubnetGroupName *string
	// A value that indicates whether the DB instance is restored from the latest
	// backup time. By default, the DB instance isn't restored from the latest backup
	// time. Constraints: Can't be specified if the RestoreTime parameter is provided.
	UseLatestRestorableTime *bool
	// A value that indicates whether the DB instance has deletion protection enabled.
	// The database can't be deleted when deletion protection is enabled. By default,
	// deletion protection is disabled. For more information, see  Deleting a DB
	// Instance
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_DeleteInstance.html).
	DeletionProtection *bool
	// The name of the option group to be used for the restored DB instance.
	// <p>Permanent options, such as the TDE option for Oracle Advanced Security TDE,
	// can't be removed from an option group, and that option group can't be removed
	// from a DB instance once it is associated with a DB instance</p>
	OptionGroupName *string
	// The database name for the restored DB instance. This parameter isn't used for
	// the MySQL or MariaDB engines.
	DBName *string
	// A list of EC2 VPC security groups to associate with this DB instance. Default:
	// The default EC2 VPC security group for the DB subnet group's VPC.
	VpcSecurityGroupIds []*string
	// The Availability Zone (AZ) where the DB instance will be created. Default: A
	// random, system-chosen Availability Zone. Constraint: You can't specify the
	// AvailabilityZone parameter if the DB instance is a Multi-AZ deployment. Example:
	// us-east-1a
	AvailabilityZone *string
	// The resource ID of the source DB instance from which to restore.
	SourceDbiResourceId *string
	// Specify the Active Directory directory ID to restore the DB instance in. The
	// domain must be created prior to this operation. Currently, only Microsoft SQL
	// Server and Oracle DB instances can be created in an Active Directory Domain. For
	// Microsoft SQL Server DB instances, Amazon RDS can use Windows Authentication to
	// authenticate users that connect to the DB instance. For more information, see
	// Using Windows Authentication with an Amazon RDS DB Instance Running Microsoft
	// SQL Server
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_SQLServerWinAuth.html)
	// in the Amazon RDS User Guide. For Oracle DB instances, Amazon RDS can use
	// Kerberos authentication to authenticate users that connect to the DB instance.
	// For more information, see  Using Kerberos Authentication with Amazon RDS for
	// Oracle
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-kerberos.html) in
	// the Amazon RDS User Guide.
	Domain *string
	// A value that indicates whether to enable mapping of AWS Identity and Access
	// Management (IAM) accounts to database accounts. By default, mapping is disabled.
	// For information about the supported DB engines, see CreateDBInstance ().  <p>For
	// more information about IAM database authentication, see <a
	// href="https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html">
	// IAM Database Authentication for MySQL and PostgreSQL</a> in the <i>Amazon RDS
	// User Guide.</i> </p>
	EnableIAMDatabaseAuthentication *bool
	// A value that indicates whether the DB instance class of the DB instance uses its
	// default processor features.
	UseDefaultProcessorFeatures *bool
	// A value that indicates whether the DB instance is publicly accessible. When the
	// DB instance is publicly accessible, its DNS endpoint resolves to the private IP
	// address from within the DB instance's VPC, and to the public IP address from
	// outside of the DB instance's VPC. Access to the DB instance is ultimately
	// controlled by the security group it uses, and that public access is not
	// permitted if the security group assigned to the DB instance doesn't permit it.
	// When the DB instance isn't publicly accessible, it is an internal DB instance
	// with a DNS name that resolves to a private IP address. For more information, see
	// CreateDBInstance ().
	PubliclyAccessible *bool
	// The port number on which the database accepts connections. Constraints: Value
	// must be 1150-65535 Default: The same port as the original DB instance.
	Port *int32
	// A value that indicates whether the DB instance is a Multi-AZ deployment.
	// Constraint: You can't specify the AvailabilityZone parameter if the DB instance
	// is a Multi-AZ deployment.
	MultiAZ *bool
	// Specify the name of the IAM role to be used when making API calls to the
	// Directory Service.
	DomainIAMRoleName *string
	// The date and time to restore from. Valid Values: Value must be a time in
	// Universal Coordinated Time (UTC) format Constraints:
	//
	//     * Must be before the
	// latest restorable time for the DB instance
	//
	//     * Can't be specified if the
	// UseLatestRestorableTime parameter is enabled
	//
	// Example: 2009-09-07T23:45:00Z
	RestoreTime *time.Time
	// The ARN from the key store with which to associate the instance for TDE
	// encryption.
	TdeCredentialArn *string
	// The name of the new DB instance to be created. Constraints:
	//
	//     * Must contain
	// from 1 to 63 letters, numbers, or hyphens
	//
	//     * First character must be a
	// letter
	//
	//     * Can't end with a hyphen or contain two consecutive hyphens
	TargetDBInstanceIdentifier *string
	// The number of CPU cores and the number of threads per core for the DB instance
	// class of the DB instance.
	ProcessorFeatures []*types.ProcessorFeature
	// A list of tags. For more information, see Tagging Amazon RDS Resources
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in
	// the Amazon RDS User Guide.
	Tags []*types.Tag
	// The name of the DB parameter group to associate with this DB instance. If you do
	// not specify a value for DBParameterGroupName, then the default DBParameterGroup
	// for the specified DB engine is used. Constraints:
	//
	//     * If supplied, must match
	// the name of an existing DBParameterGroup.
	//
	//     * Must be 1 to 255 letters,
	// numbers, or hyphens.
	//
	//     * First character must be a letter.
	//
	//     * Can't end
	// with a hyphen or contain two consecutive hyphens.
	DBParameterGroupName *string
	// A value that indicates whether minor version upgrades are applied automatically
	// to the DB instance during the maintenance window.
	AutoMinorVersionUpgrade *bool
	// A value that indicates whether to copy all tags from the restored DB instance to
	// snapshots of the DB instance. By default, tags are not copied.
	CopyTagsToSnapshot *bool
	// The compute and memory capacity of the Amazon RDS DB instance, for example,
	// db.m4.large. Not all DB instance classes are available in all AWS Regions, or
	// for all database engines. For the full list of DB instance classes, and
	// availability for your engine, see DB Instance Class
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html)
	// in the Amazon RDS User Guide. Default: The same DBInstanceClass as the original
	// DB instance.
	DBInstanceClass *string
}

type RestoreDBInstanceToPointInTimeOutput struct {
	// Contains the details of an Amazon RDS DB instance. This data type is used as a
	// response element in the DescribeDBInstances action.
	DBInstance *types.DBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpRestoreDBInstanceToPointInTimeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpRestoreDBInstanceToPointInTime{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpRestoreDBInstanceToPointInTime{}, middleware.After)
}

func newServiceMetadataMiddleware_opRestoreDBInstanceToPointInTime(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "RestoreDBInstanceToPointInTime",
	}
}
