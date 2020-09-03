// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Contains all of the attributes of a specific DAX cluster.
type Cluster struct {
	// A range of time when maintenance of DAX cluster software will be performed. For
	// example: sun:01:00-sun:09:00. Cluster maintenance normally takes less than 30
	// minutes, and is performed automatically within the maintenance window.
	PreferredMaintenanceWindow *string
	// The parameter group being used by nodes in the cluster.
	ParameterGroup *ParameterGroupStatus
	// The description of the server-side encryption status on the specified DAX
	// cluster.
	SSEDescription *SSEDescription
	// The configuration endpoint for this DAX cluster, consisting of a DNS name and a
	// port number. Client applications can specify this endpoint, rather than an
	// individual node endpoint, and allow the DAX client software to intelligently
	// route requests and responses to nodes in the DAX cluster.
	ClusterDiscoveryEndpoint *Endpoint
	// The description of the cluster.
	Description *string
	// The number of nodes in the cluster that are active (i.e., capable of serving
	// requests).
	ActiveNodes *int32
	// A list of nodes that are currently in the cluster.
	Nodes []*Node
	// The subnet group where the DAX cluster is running.
	SubnetGroup *string
	// Describes a notification topic and its status. Notification topics are used for
	// publishing DAX events to subscribers using Amazon Simple Notification Service
	// (SNS).
	NotificationConfiguration *NotificationConfiguration
	// A valid Amazon Resource Name (ARN) that identifies an IAM role. At runtime, DAX
	// will assume this role and use the role's permissions to access DynamoDB on your
	// behalf.
	IamRoleArn *string
	// A list of nodes to be removed from the cluster.
	NodeIdsToRemove []*string
	// The current status of the cluster.
	Status *string
	// A list of security groups, and the status of each, for the nodes in the cluster.
	SecurityGroups []*SecurityGroupMembership
	// The Amazon Resource Name (ARN) that uniquely identifies the cluster.
	ClusterArn *string
	// The node type for the nodes in the cluster. (All nodes in a DAX cluster are of
	// the same type.)
	NodeType *string
	// The total number of nodes in the cluster.
	TotalNodes *int32
	// The name of the DAX cluster.
	ClusterName *string
}

// Represents the information required for client programs to connect to the
// configuration endpoint for a DAX cluster, or to an individual node within the
// cluster.
type Endpoint struct {
	// The port number that applications should use to connect to the endpoint.
	Port *int32
	// The DNS hostname of the endpoint.
	Address *string
}

// Represents a single occurrence of something interesting within the system. Some
// examples of events are creating a DAX cluster, adding or removing a node, or
// rebooting a node.
type Event struct {
	// A user-defined message associated with the event.
	Message *string
	// Specifies the origin of this event - a cluster, a parameter group, a node ID,
	// etc.
	SourceType SourceType
	// The date and time when the event occurred.
	Date *time.Time
	// The source of the event. For example, if the event occurred at the node level,
	// the source would be the node ID.
	SourceName *string
}

// Represents an individual node within a DAX cluster.
type Node struct {
	// A system-generated identifier for the node.
	NodeId *string
	// The Availability Zone (AZ) in which the node has been deployed.
	AvailabilityZone *string
	// The date and time (in UNIX epoch format) when the node was launched.
	NodeCreateTime *time.Time
	// The current status of the node. For example: available.
	NodeStatus *string
	// The endpoint for the node, consisting of a DNS name and a port number. Client
	// applications can connect directly to a node endpoint, if desired (as an
	// alternative to allowing DAX client software to intelligently route requests and
	// responses to nodes in the DAX cluster.
	Endpoint *Endpoint
	// The status of the parameter group associated with this node. For example,
	// in-sync.
	ParameterGroupStatus *string
}

// Represents a parameter value that is applicable to a particular node type.
type NodeTypeSpecificValue struct {
	// A node type to which the parameter value applies.
	NodeType *string
	// The parameter value for this node type.
	Value *string
}

// Describes a notification topic and its status. Notification topics are used for
// publishing DAX events to subscribers using Amazon Simple Notification Service
// (SNS).
type NotificationConfiguration struct {
	// The Amazon Resource Name (ARN) that identifies the topic.
	TopicArn *string
	// The current state of the topic.
	TopicStatus *string
}

// Describes an individual setting that controls some aspect of DAX behavior.
type Parameter struct {
	// The data type of the parameter. For example, integer:
	DataType *string
	// How the parameter is defined. For example, system denotes a system-defined
	// parameter.
	Source *string
	// The conditions under which changes to this parameter can be applied. For
	// example, requires-reboot indicates that a new value for this parameter will only
	// take effect if a node is rebooted.
	ChangeType ChangeType
	// A description of the parameter
	Description *string
	// A list of node types, and specific parameter values for each node.
	NodeTypeSpecificValues []*NodeTypeSpecificValue
	// Determines whether the parameter can be applied to any nodes, or only nodes of a
	// particular type.
	ParameterType ParameterType
	// Whether the customer is allowed to modify the parameter.
	IsModifiable IsModifiable
	// The value for the parameter.
	ParameterValue *string
	// The name of the parameter.
	ParameterName *string
	// A range of values within which the parameter can be set.
	AllowedValues *string
}

// A named set of parameters that are applied to all of the nodes in a DAX cluster.
type ParameterGroup struct {
	// The name of the parameter group.
	ParameterGroupName *string
	// A description of the parameter group.
	Description *string
}

// The status of a parameter group.
type ParameterGroupStatus struct {
	// The name of the parameter group.
	ParameterGroupName *string
	// The node IDs of one or more nodes to be rebooted.
	NodeIdsToReboot []*string
	// The status of parameter updates.
	ParameterApplyStatus *string
}

// An individual DAX parameter.
type ParameterNameValue struct {
	// The value of the parameter.
	ParameterValue *string
	// The name of the parameter.
	ParameterName *string
}

// An individual VPC security group and its status.
type SecurityGroupMembership struct {
	// The status of this security group.
	Status *string
	// The unique ID for this security group.
	SecurityGroupIdentifier *string
}

// The description of the server-side encryption status on the specified DAX
// cluster.
type SSEDescription struct {
	// The current state of server-side encryption:
	//
	//     * ENABLING - Server-side
	// encryption is being enabled.
	//
	//     * ENABLED - Server-side encryption is
	// enabled.
	//
	//     * DISABLING - Server-side encryption is being disabled.
	//
	//     *
	// DISABLED - Server-side encryption is disabled.
	Status SSEStatus
}

// Represents the settings used to enable server-side encryption.
type SSESpecification struct {
	// Indicates whether server-side encryption is enabled (true) or disabled (false)
	// on the cluster.
	Enabled *bool
}

// Represents the subnet associated with a DAX cluster. This parameter refers to
// subnets defined in Amazon Virtual Private Cloud (Amazon VPC) and used with DAX.
type Subnet struct {
	// The system-assigned identifier for the subnet.
	SubnetIdentifier *string
	// The Availability Zone (AZ) for the subnet.
	SubnetAvailabilityZone *string
}

// Represents the output of one of the following actions:
//
//     *
// CreateSubnetGroup
//
//     * ModifySubnetGroup
type SubnetGroup struct {
	// A list of subnets associated with the subnet group.
	Subnets []*Subnet
	// The Amazon Virtual Private Cloud identifier (VPC ID) of the subnet group.
	VpcId *string
	// The description of the subnet group.
	Description *string
	// The name of the subnet group.
	SubnetGroupName *string
}

// A description of a tag. Every tag is a key-value pair. You can add up to 50 tags
// to a single DAX cluster. AWS-assigned tag names and values are automatically
// assigned the aws: prefix, which the user cannot assign. AWS-assigned tag names
// do not count towards the tag limit of 50. User-assigned tag names have the
// prefix user:. You cannot backdate the application of a tag.
type Tag struct {
	// The key for the tag. Tag keys are case sensitive. Every DAX cluster can only
	// have one tag with the same key. If you try to add an existing tag (same key),
	// the existing tag value will be updated to the new value.
	Key *string
	// The value of the tag. Tag values are case-sensitive and can be null.
	Value *string
}
