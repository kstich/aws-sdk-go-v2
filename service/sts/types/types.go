// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// The identifiers for the temporary security credentials that the operation
// returns.
type AssumedRoleUser struct {
	// The ARN of the temporary security credentials that are returned from the
	// AssumeRole () action. For more information about ARNs and how to use them in
	// policies, see IAM Identifiers
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html) in
	// the IAM User Guide.
	Arn *string
	// A unique identifier that contains the role ID and the role session name of the
	// role that is being assumed. The role ID is generated by AWS when the role is
	// created.
	AssumedRoleId *string
}

// AWS credentials for API authentication.
type Credentials struct {
	// The access key ID that identifies the temporary security credentials.
	AccessKeyId *string
	// The token that users must pass to the service API to use the temporary
	// credentials.
	SessionToken *string
	// The secret access key that can be used to sign requests.
	SecretAccessKey *string
	// The date on which the current credentials expire.
	Expiration *time.Time
}

// Identifiers for the federated user that is associated with the credentials.
type FederatedUser struct {
	// The ARN that specifies the federated user that is associated with the
	// credentials. For more information about ARNs and how to use them in policies,
	// see IAM Identifiers
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html) in
	// the IAM User Guide.
	Arn *string
	// The string that identifies the federated user associated with the credentials,
	// similar to the unique ID of an IAM user.
	FederatedUserId *string
}

// A reference to the IAM managed policy that is passed as a session policy for a
// role session or a federated user session.
type PolicyDescriptorType struct {
	// The Amazon Resource Name (ARN) of the IAM managed policy to use as a session
	// policy for the role. For more information about ARNs, see Amazon Resource Names
	// (ARNs) and AWS Service Namespaces
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference.
	Arn *string
}

// You can pass custom key-value pair attributes when you assume a role or federate
// a user. These are called session tags. You can then use the session tags to
// control access to resources. For more information, see Tagging AWS STS Sessions
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_session-tags.html) in the
// IAM User Guide.
type Tag struct {
	// The value for a session tag. You can pass up to 50 session tags. The plain text
	// session tag values can’t exceed 256 characters. For these and additional limits,
	// see IAM and STS Character Limits
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-limits.html#reference_iam-limits-entity-length)
	// in the IAM User Guide.
	Value *string
	// The key for a session tag. You can pass up to 50 session tags. The plain text
	// session tag keys can’t exceed 128 characters. For these and additional limits,
	// see IAM and STS Character Limits
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-limits.html#reference_iam-limits-entity-length)
	// in the IAM User Guide.
	Key *string
}
