// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// RunJobFlow creates and starts running a new cluster (job flow). The cluster runs
// the steps specified. After the steps complete, the cluster stops and the HDFS
// partition is lost. To prevent loss of data, configure the last step of the job
// flow to store results in Amazon S3. If the JobFlowInstancesConfig
// ()KeepJobFlowAliveWhenNoSteps parameter is set to TRUE, the cluster transitions
// to the WAITING state rather than shutting down after the steps have completed.
// <p>For additional protection, you can set the <a>JobFlowInstancesConfig</a>
// <code>TerminationProtected</code> parameter to <code>TRUE</code> to lock the
// cluster and prevent it from being terminated by API call, user intervention, or
// in the event of a job flow error.</p> <p>A maximum of 256 steps are allowed in
// each job flow.</p> <p>If your cluster is long-running (such as a Hive data
// warehouse) or complex, you may require more than 256 steps to process your data.
// You can bypass the 256-step limitation in various ways, including using the SSH
// shell to connect to the master node and submitting queries directly to the
// software running on the master node, such as Hive and Hadoop. For more
// information on how to do this, see <a
// href="https://docs.aws.amazon.com/emr/latest/ManagementGuide/AddMoreThan256Steps.html">Add
// More than 256 Steps to a Cluster</a> in the <i>Amazon EMR Management
// Guide</i>.</p> <p>For long running clusters, we recommend that you periodically
// store your results.</p> <note> <p>The instance fleets configuration is available
// only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions. The
// RunJobFlow request can contain InstanceFleets parameters or InstanceGroups
// parameters, but not both.</p> </note>
func (c *Client) RunJobFlow(ctx context.Context, params *RunJobFlowInput, optFns ...func(*Options)) (*RunJobFlowOutput, error) {
	stack := middleware.NewStack("RunJobFlow", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRunJobFlowMiddlewares(stack)
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
	addOpRunJobFlowValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRunJobFlow(options.Region), middleware.Before)
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
			OperationName: "RunJobFlow",
			Err:           err,
		}
	}
	out := result.(*RunJobFlowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to the RunJobFlow () operation.
type RunJobFlowInput struct {
	// The AWS KMS customer master key (CMK) used for encrypting log files. If a value
	// is not provided, the logs will remain encrypted by AES-256. This attribute is
	// only available with EMR version 5.30.0 and later, excluding EMR 6.0.0.
	LogEncryptionKmsKeyId *string
	// Attributes for Kerberos configuration when Kerberos authentication is enabled
	// using a security configuration. For more information see Use Kerberos
	// Authentication
	// (https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-kerberos.html) in
	// the EMR Management Guide.
	KerberosAttributes *types.KerberosAttributes
	// A JSON string for selecting additional features.
	AdditionalInfo *string
	// Available only in Amazon EMR version 5.7.0 and later. The ID of a custom Amazon
	// EBS-backed Linux AMI. If specified, Amazon EMR uses this AMI when it launches
	// cluster EC2 instances. For more information about custom AMIs in Amazon EMR, see
	// Using a Custom AMI
	// (https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-custom-ami.html) in
	// the Amazon EMR Management Guide. If omitted, the cluster uses the base Linux AMI
	// for the ReleaseLabel specified. For Amazon EMR versions 2.x and 3.x, use
	// AmiVersion instead. For information about creating a custom AMI, see Creating an
	// Amazon EBS-Backed Linux AMI
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/creating-an-ami-ebs.html)
	// in the Amazon Elastic Compute Cloud User Guide for Linux Instances. For
	// information about finding an AMI ID, see Finding a Linux AMI
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/finding-an-ami.html).
	CustomAmiId *string
	// An IAM role for automatic scaling policies. The default role is
	// EMR_AutoScaling_DefaultRole. The IAM role provides permissions that the
	// automatic scaling feature requires to launch and terminate EC2 instances in an
	// instance group.
	AutoScalingRole *string
	// Applies to Amazon EMR releases 4.0 and later. A case-insensitive list of
	// applications for Amazon EMR to install and configure when launching the cluster.
	// For a list of applications available for each Amazon EMR release version, see
	// the Amazon EMR Release Guide
	// (https://docs.aws.amazon.com/emr/latest/ReleaseGuide/).
	Applications []*types.Application
	// Specifies the way that individual Amazon EC2 instances terminate when an
	// automatic scale-in activity occurs or an instance group is resized.
	// TERMINATE_AT_INSTANCE_HOUR indicates that Amazon EMR terminates nodes at the
	// instance-hour boundary, regardless of when the request to terminate the instance
	// was submitted. This option is only available with Amazon EMR 5.1.0 and later and
	// is the default for clusters created using that version.
	// TERMINATE_AT_TASK_COMPLETION indicates that Amazon EMR blacklists and drains
	// tasks from nodes before terminating the Amazon EC2 instances, regardless of the
	// instance-hour boundary. With either behavior, Amazon EMR removes the least
	// active nodes first and blocks instance termination if it could lead to HDFS
	// corruption. TERMINATE_AT_TASK_COMPLETION available only in Amazon EMR version
	// 4.1.0 and later, and is the default for versions of Amazon EMR earlier than
	// 5.1.0.
	ScaleDownBehavior types.ScaleDownBehavior
	// Applies only to Amazon EMR AMI versions 3.x and 2.x. For Amazon EMR releases 4.0
	// and later, ReleaseLabel is used. To specify a custom AMI, use CustomAmiID.
	AmiVersion *string
	// A list of tags to associate with a cluster and propagate to Amazon EC2
	// instances.
	Tags []*types.Tag
	// The IAM role that will be assumed by the Amazon EMR service to access AWS
	// resources on your behalf.
	ServiceRole *string
	// The location in Amazon S3 to write the log files of the job flow. If a value is
	// not provided, logs are not created.
	LogUri *string
	// The specified managed scaling policy for an Amazon EMR cluster.
	ManagedScalingPolicy *types.ManagedScalingPolicy
	// For Amazon EMR releases 4.0 and later. The list of configurations supplied for
	// the EMR cluster you are creating.
	Configurations []*types.Configuration
	// The Amazon EMR release label, which determines the version of open-source
	// application packages installed on the cluster. Release labels are in the form
	// emr-x.x.x, where x.x.x is an Amazon EMR release version such as emr-5.14.0. For
	// more information about Amazon EMR release versions and included application
	// versions and features, see https://docs.aws.amazon.com/emr/latest/ReleaseGuide/
	// (https://docs.aws.amazon.com/emr/latest/ReleaseGuide/). The release label
	// applies only to Amazon EMR releases version 4.0 and later. Earlier versions use
	// AmiVersion.
	ReleaseLabel *string
	// For Amazon EMR releases 3.x and 2.x. For Amazon EMR releases 4.x and later, use
	// Applications. A list of strings that indicates third-party software to use. For
	// more information, see the Amazon EMR Developer Guide
	// (https://docs.aws.amazon.com/emr/latest/DeveloperGuide/emr-dg.pdf). Currently
	// supported values are:
	//
	//     * "mapr-m3" - launch the job flow using MapR M3
	// Edition.
	//
	//     * "mapr-m5" - launch the job flow using MapR M5 Edition.
	SupportedProducts []*string
	// A list of steps to run.
	Steps []*types.StepConfig
	// A list of bootstrap actions to run before Hadoop starts on the cluster nodes.
	BootstrapActions []*types.BootstrapActionConfig
	// The name of the job flow.
	Name *string
	// Specifies the number of steps that can be executed concurrently. The default
	// value is 1. The maximum value is 256.
	StepConcurrencyLevel *int32
	// The size, in GiB, of the EBS root device volume of the Linux AMI that is used
	// for each EC2 instance. Available in Amazon EMR version 4.x and later.
	EbsRootVolumeSize *int32
	// A specification of the number and type of Amazon EC2 instances.
	Instances *types.JobFlowInstancesConfig
	// A value of true indicates that all IAM users in the AWS account can perform
	// cluster actions if they have the proper IAM policy permissions. This is the
	// default. A value of false indicates that only the IAM user who created the
	// cluster can perform actions.
	VisibleToAllUsers *bool
	// For Amazon EMR releases 3.x and 2.x. For Amazon EMR releases 4.x and later, use
	// Applications. A list of strings that indicates third-party software to use with
	// the job flow that accepts a user argument list. EMR accepts and forwards the
	// argument list to the corresponding installation script as bootstrap action
	// arguments. For more information, see "Launch a Job Flow on the MapR Distribution
	// for Hadoop" in the Amazon EMR Developer Guide
	// (https://docs.aws.amazon.com/emr/latest/DeveloperGuide/emr-dg.pdf). Supported
	// values are:
	//
	//     * "mapr-m3" - launch the cluster using MapR M3 Edition.
	//
	//     *
	// "mapr-m5" - launch the cluster using MapR M5 Edition.
	//
	//     * "mapr" with the
	// user arguments specifying "--edition,m3" or "--edition,m5" - launch the job flow
	// using MapR M3 or M5 Edition respectively.
	//
	//     * "mapr-m7" - launch the cluster
	// using MapR M7 Edition.
	//
	//     * "hunk" - launch the cluster with the Hunk Big Data
	// Analtics Platform.
	//
	//     * "hue"- launch the cluster with Hue installed.
	//
	//     *
	// "spark" - launch the cluster with Apache Spark installed.
	//
	//     * "ganglia" -
	// launch the cluster with the Ganglia Monitoring System installed.
	NewSupportedProducts []*types.SupportedProductConfig
	// The name of a security configuration to apply to the cluster.
	SecurityConfiguration *string
	// Also called instance profile and EC2 role. An IAM role for an EMR cluster. The
	// EC2 instances of the cluster assume this role. The default role is
	// EMR_EC2_DefaultRole. In order to use the default role, you must have already
	// created it using the CLI or console.
	JobFlowRole *string
	// Applies only when CustomAmiID is used. Specifies which updates from the Amazon
	// Linux AMI package repositories to apply automatically when the instance boots
	// using the AMI. If omitted, the default is SECURITY, which indicates that only
	// security updates are applied. If NONE is specified, no updates are applied, and
	// all updates must be applied manually.
	RepoUpgradeOnBoot types.RepoUpgradeOnBoot
}

// The result of the RunJobFlow () operation.
type RunJobFlowOutput struct {
	// The Amazon Resource Name of the cluster.
	ClusterArn *string
	// An unique identifier for the job flow.
	JobFlowId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRunJobFlowMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRunJobFlow{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRunJobFlow{}, middleware.After)
}

func newServiceMetadataMiddleware_opRunJobFlow(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "RunJobFlow",
	}
}
