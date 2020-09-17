// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a Redis (cluster mode disabled) or a Redis (cluster mode enabled)
// replication group. This API can be used to create a standalone regional
// replication group or a secondary replication group associated with a Global
// Datastore. A Redis (cluster mode disabled) replication group is a collection of
// clusters, where one of the clusters is a read/write primary and the others are
// read-only replicas. Writes to the primary are asynchronously propagated to the
// replicas. A Redis (cluster mode enabled) replication group is a collection of 1
// to 90 node groups (shards). Each node group (shard) has one read/write primary
// node and up to 5 read-only replica nodes. Writes to the primary are
// asynchronously propagated to the replicas. Redis (cluster mode enabled)
// replication groups partition the data across node groups (shards). When a Redis
// (cluster mode disabled) replication group has been successfully created, you can
// add one or more read replicas to it, up to a total of 5 read replicas. If you
// need to increase or decrease the number of node groups (console: shards), you
// can avail yourself of ElastiCache for Redis' scaling. For more information, see
// Scaling ElastiCache for Redis Clusters
// (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Scaling.html) in
// the ElastiCache User Guide. This operation is valid for Redis only.
func (c *Client) CreateReplicationGroup(ctx context.Context, params *CreateReplicationGroupInput, optFns ...func(*Options)) (*CreateReplicationGroupOutput, error) {
	stack := middleware.NewStack("CreateReplicationGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateReplicationGroupMiddlewares(stack)
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
	addOpCreateReplicationGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateReplicationGroup(options.Region), middleware.Before)
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
			OperationName: "CreateReplicationGroup",
			Err:           err,
		}
	}
	out := result.(*CreateReplicationGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a CreateReplicationGroup operation.
type CreateReplicationGroupInput struct {
	// A flag that enables in-transit encryption when set to true. You cannot modify
	// the value of TransitEncryptionEnabled after the cluster is created. To enable
	// in-transit encryption on a cluster you must set TransitEncryptionEnabled to true
	// when you create a cluster. This parameter is valid only if the Engine parameter
	// is redis, the EngineVersion parameter is 3.2.6, 4.x or later, and the cluster is
	// being created in an Amazon VPC. If you enable in-transit encryption, you must
	// also specify a value for CacheSubnetGroup. Required: Only available when
	// creating a replication group in an Amazon VPC using redis version 3.2.6, 4.x or
	// later. Default: false For HIPAA compliance, you must specify
	// TransitEncryptionEnabled as true, an AuthToken, and a CacheSubnetGroup.
	TransitEncryptionEnabled *bool
	// The name of the cache engine to be used for the clusters in this replication
	// group.
	Engine *string
	// The daily time range (in UTC) during which ElastiCache begins taking a daily
	// snapshot of your node group (shard). Example: 05:00-09:00 If you do not specify
	// this parameter, ElastiCache automatically chooses an appropriate time range.
	SnapshotWindow *string
	// Reserved parameter. The password used to access a password protected server.
	// AuthToken can be specified only on replication groups where
	// TransitEncryptionEnabled is true. For HIPAA compliance, you must specify
	// TransitEncryptionEnabled as true, an AuthToken, and a CacheSubnetGroup. Password
	// constraints:
	//
	//     * Must be only printable ASCII characters.
	//
	//     * Must be at
	// least 16 characters and no more than 128 characters in length.
	//
	//     * The only
	// permitted printable special characters are !, &, #, $, ^, <, >, and -. Other
	// printable special characters cannot be used in the AUTH token.
	//
	// For more
	// information, see AUTH password (http://redis.io/commands/AUTH) at
	// http://redis.io/commands/AUTH.
	AuthToken *string
	// The name of the cache subnet group to be used for the replication group. If
	// you're going to launch your cluster in an Amazon VPC, you need to create a
	// subnet group before you start creating a cluster. For more information, see
	// Subnets and Subnet Groups
	// (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/SubnetGroups.html).
	CacheSubnetGroupName *string
	// The port number on which each member of the replication group accepts
	// connections.
	Port *int32
	// A flag that enables encryption at rest when set to true. You cannot modify the
	// value of AtRestEncryptionEnabled after the replication group is created. To
	// enable encryption at rest on a replication group you must set
	// AtRestEncryptionEnabled to true when you create the replication group. Required:
	// Only available when creating a replication group in an Amazon VPC using redis
	// version 3.2.6, 4.x or later. Default: false
	AtRestEncryptionEnabled *bool
	// The identifier of the cluster that serves as the primary for this replication
	// group. This cluster must already exist and have a status of available. This
	// parameter is not required if NumCacheClusters, NumNodeGroups, or
	// ReplicasPerNodeGroup is specified.
	PrimaryClusterId *string
	// The name of the Global Datastore
	GlobalReplicationGroupId *string
	// Specifies whether a read-only replica is automatically promoted to read/write
	// primary if the existing primary fails.  <p>
	// <code>AutomaticFailoverEnabled</code> must be enabled for Redis (cluster mode
	// enabled) replication groups.</p> <p>Default: false</p>
	AutomaticFailoverEnabled *bool
	// A list of EC2 Availability Zones in which the replication group's clusters are
	// created. The order of the Availability Zones in the list is the order in which
	// clusters are allocated. The primary cluster is created in the first AZ in the
	// list. This parameter is not used if there is more than one node group (shard).
	// You should use NodeGroupConfiguration instead. If you are creating your
	// replication group in an Amazon VPC (recommended), you can only locate clusters
	// in Availability Zones associated with the subnets in the selected subnet group.
	// The number of Availability Zones listed must equal the value of
	// NumCacheClusters. Default: system chosen Availability Zones.
	PreferredCacheClusterAZs []*string
	// A user-created description for the replication group.
	ReplicationGroupDescription *string
	// One or more Amazon VPC security groups associated with this replication group.
	// Use this parameter only when you are creating a replication group in an Amazon
	// Virtual Private Cloud (Amazon VPC).
	SecurityGroupIds []*string
	// The version number of the cache engine to be used for the clusters in this
	// replication group. To view the supported cache engine versions, use the
	// DescribeCacheEngineVersions operation.  <p> <b>Important:</b> You can upgrade to
	// a newer engine version (see <a
	// href="https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/SelectEngine.html#VersionManagement">Selecting
	// a Cache Engine and Version</a>) in the <i>ElastiCache User Guide</i>, but you
	// cannot downgrade to an earlier engine version. If you want to use an earlier
	// engine version, you must delete the existing cluster or replication group and
	// create it anew with the earlier engine version. </p>
	EngineVersion *string
	// The compute and memory capacity of the nodes in the node group (shard). The
	// following node types are supported by ElastiCache. Generally speaking, the
	// current generation types provide more memory and computational power at lower
	// cost when compared to their equivalent previous generation counterparts.
	//
	//     *
	// General purpose:
	//
	//         * Current generation:  <p> <b>M5 node types:</b>
	// <code>cache.m5.large</code>, <code>cache.m5.xlarge</code>,
	// <code>cache.m5.2xlarge</code>, <code>cache.m5.4xlarge</code>,
	// <code>cache.m5.12xlarge</code>, <code>cache.m5.24xlarge</code> </p> <p> <b>M4
	// node types:</b> <code>cache.m4.large</code>, <code>cache.m4.xlarge</code>,
	// <code>cache.m4.2xlarge</code>, <code>cache.m4.4xlarge</code>,
	// <code>cache.m4.10xlarge</code> </p> <p> <b>T3 node types:</b>
	// <code>cache.t3.micro</code>, <code>cache.t3.small</code>,
	// <code>cache.t3.medium</code> </p> <p> <b>T2 node types:</b>
	// <code>cache.t2.micro</code>, <code>cache.t2.small</code>,
	// <code>cache.t2.medium</code> </p> </li> <li> <p>Previous generation: (not
	// recommended)</p> <p> <b>T1 node types:</b> <code>cache.t1.micro</code> </p> <p>
	// <b>M1 node types:</b> <code>cache.m1.small</code>, <code>cache.m1.medium</code>,
	// <code>cache.m1.large</code>, <code>cache.m1.xlarge</code> </p> <p> <b>M3 node
	// types:</b> <code>cache.m3.medium</code>, <code>cache.m3.large</code>,
	// <code>cache.m3.xlarge</code>, <code>cache.m3.2xlarge</code> </p> </li> </ul>
	// </li> <li> <p>Compute optimized:</p> <ul> <li> <p>Previous generation: (not
	// recommended)</p> <p> <b>C1 node types:</b> <code>cache.c1.xlarge</code> </p>
	// </li> </ul> </li> <li> <p>Memory optimized:</p> <ul> <li> <p>Current generation:
	// </p> <p> <b>R5 node types:</b> <code>cache.r5.large</code>,
	// <code>cache.r5.xlarge</code>, <code>cache.r5.2xlarge</code>,
	// <code>cache.r5.4xlarge</code>, <code>cache.r5.12xlarge</code>,
	// <code>cache.r5.24xlarge</code> </p> <p> <b>R4 node types:</b>
	// <code>cache.r4.large</code>, <code>cache.r4.xlarge</code>,
	// <code>cache.r4.2xlarge</code>, <code>cache.r4.4xlarge</code>,
	// <code>cache.r4.8xlarge</code>, <code>cache.r4.16xlarge</code> </p> </li> <li>
	// <p>Previous generation: (not recommended)</p> <p> <b>M2 node types:</b>
	// <code>cache.m2.xlarge</code>, <code>cache.m2.2xlarge</code>,
	// <code>cache.m2.4xlarge</code> </p> <p> <b>R3 node types:</b>
	// <code>cache.r3.large</code>, <code>cache.r3.xlarge</code>,
	// <code>cache.r3.2xlarge</code>, <code>cache.r3.4xlarge</code>,
	// <code>cache.r3.8xlarge</code> </p> </li> </ul> </li> </ul> <p> <b>Additional
	// node type info</b> </p> <ul> <li> <p>All current generation instance types are
	// created in Amazon VPC by default.</p> </li> <li> <p>Redis append-only files
	// (AOF) are not supported for T1 or T2 instances.</p> </li> <li> <p>Redis Multi-AZ
	// with automatic failover is not supported on T1 instances.</p> </li> <li>
	// <p>Redis configuration variables <code>appendonly</code> and
	// <code>appendfsync</code> are not supported on Redis version 2.8.22 and
	// later.</p> </li> </ul>
	CacheNodeType *string
	// The name of a snapshot from which to restore data into the new replication
	// group. The snapshot status changes to restoring while the new replication group
	// is being created.
	SnapshotName *string
	// The number of days for which ElastiCache retains automatic snapshots before
	// deleting them. For example, if you set SnapshotRetentionLimit to 5, a snapshot
	// that was taken today is retained for 5 days before being deleted. Default: 0
	// (i.e., automatic backups are disabled for this cluster).
	SnapshotRetentionLimit *int32
	// An optional parameter that specifies the number of replica nodes in each node
	// group (shard). Valid values are 0 to 5.
	ReplicasPerNodeGroup *int32
	// The number of clusters this replication group initially has. This parameter is
	// not used if there is more than one node group (shard). You should use
	// ReplicasPerNodeGroup instead. If AutomaticFailoverEnabled is true, the value of
	// this parameter must be at least 2. If AutomaticFailoverEnabled is false you can
	// omit this parameter (it will default to 1), or you can explicitly set it to a
	// value between 2 and 6. The maximum permitted value for NumCacheClusters is 6 (1
	// primary plus 5 replicas).
	NumCacheClusters *int32
	// A list of cost allocation tags to be added to this resource. Tags are
	// comma-separated key,value pairs (e.g. Key=myKey, Value=myKeyValue. You can
	// include multiple tags as shown following: Key=myKey, Value=myKeyValue
	// Key=mySecondKey, Value=mySecondKeyValue.
	Tags []*types.Tag
	// An optional parameter that specifies the number of node groups (shards) for this
	// Redis (cluster mode enabled) replication group. For Redis (cluster mode
	// disabled) either omit this parameter or set it to 1. Default: 1
	NumNodeGroups *int32
	// A list of cache security group names to associate with this replication group.
	CacheSecurityGroupNames []*string
	// A list of node group (shard) configuration options. Each node group (shard)
	// configuration has the following members: PrimaryAvailabilityZone,
	// ReplicaAvailabilityZones, ReplicaCount, and Slots. If you're creating a Redis
	// (cluster mode disabled) or a Redis (cluster mode enabled) replication group, you
	// can use this parameter to individually configure each node group (shard), or you
	// can omit this parameter. However, it is required when seeding a Redis (cluster
	// mode enabled) cluster from a S3 rdb file. You must configure each node group
	// (shard) using this parameter because you must specify the slots for each node
	// group.
	NodeGroupConfiguration []*types.NodeGroupConfiguration
	// The ID of the KMS key used to encrypt the disk in the cluster.
	KmsKeyId *string
	// A list of Amazon Resource Names (ARN) that uniquely identify the Redis RDB
	// snapshot files stored in Amazon S3. The snapshot files are used to populate the
	// new replication group. The Amazon S3 object name in the ARN cannot contain any
	// commas. The new replication group will have the number of node groups (console:
	// shards) specified by the parameter NumNodeGroups or the number of node groups
	// configured by NodeGroupConfiguration regardless of the number of ARNs specified
	// here. Example of an Amazon S3 ARN: arn:aws:s3:::my_bucket/snapshot1.rdb
	SnapshotArns []*string
	// The name of the parameter group to associate with this replication group. If
	// this argument is omitted, the default cache parameter group for the specified
	// engine is used. If you are restoring to an engine version that is different than
	// the original, you must specify the default version of that version. For example,
	// CacheParameterGroupName=default.redis4.0. If you are running Redis version 3.2.4
	// or later, only one node group (shard), and want to use a default parameter
	// group, we recommend that you specify the parameter group by name.
	//
	//     * To
	// create a Redis (cluster mode disabled) replication group, use
	// CacheParameterGroupName=default.redis3.2.
	//
	//     * To create a Redis (cluster mode
	// enabled) replication group, use
	// CacheParameterGroupName=default.redis3.2.cluster.on.
	CacheParameterGroupName *string
	// The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS)
	// topic to which notifications are sent. The Amazon SNS topic owner must be the
	// same as the cluster owner.
	NotificationTopicArn *string
	// A flag indicating if you have Multi-AZ enabled to enhance fault tolerance. For
	// more information, see Minimizing Downtime: Multi-AZ
	// (http://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/AutoFailover.html).
	MultiAZEnabled *bool
	// The replication group identifier. This parameter is stored as a lowercase
	// string.  <p>Constraints:</p> <ul> <li> <p>A name must contain from 1 to 40
	// alphanumeric characters or hyphens.</p> </li> <li> <p>The first character must
	// be a letter.</p> </li> <li> <p>A name cannot end with a hyphen or contain two
	// consecutive hyphens.</p> </li> </ul>
	ReplicationGroupId *string
	// Specifies the weekly time range during which maintenance on the cluster is
	// performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H
	// Clock UTC). The minimum maintenance window is a 60 minute period. Valid values
	// for ddd are:  <p>Specifies the weekly time range during which maintenance  on
	// the cluster is performed. It is specified as a range in the format
	// ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60
	// minute period. Valid values for ddd are:
	//
	//     * sun
	//
	//     * mon
	//
	//     * tue
	//
	//     *
	// wed
	//
	//     * thu
	//
	//     * fri
	//
	//     * sat
	//
	// Example: sun:23:00-mon:01:30
	PreferredMaintenanceWindow *string
	// This parameter is currently disabled.
	AutoMinorVersionUpgrade *bool
}

type CreateReplicationGroupOutput struct {
	// Contains all of the attributes of a specific Redis replication group.
	ReplicationGroup *types.ReplicationGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateReplicationGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateReplicationGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateReplicationGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateReplicationGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "CreateReplicationGroup",
	}
}
