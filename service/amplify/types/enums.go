// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type DomainStatus string

// Enum values for DomainStatus
const (
	DomainStatusPending_verification   DomainStatus = "PENDING_VERIFICATION"
	DomainStatusIn_progress            DomainStatus = "IN_PROGRESS"
	DomainStatusAvailable              DomainStatus = "AVAILABLE"
	DomainStatusPending_deployment     DomainStatus = "PENDING_DEPLOYMENT"
	DomainStatusFailed                 DomainStatus = "FAILED"
	DomainStatusCreating               DomainStatus = "CREATING"
	DomainStatusRequesting_certificate DomainStatus = "REQUESTING_CERTIFICATE"
	DomainStatusUpdating               DomainStatus = "UPDATING"
)

type JobStatus string

// Enum values for JobStatus
const (
	JobStatusPending      JobStatus = "PENDING"
	JobStatusProvisioning JobStatus = "PROVISIONING"
	JobStatusRunning      JobStatus = "RUNNING"
	JobStatusFailed       JobStatus = "FAILED"
	JobStatusSucceed      JobStatus = "SUCCEED"
	JobStatusCancelling   JobStatus = "CANCELLING"
	JobStatusCancelled    JobStatus = "CANCELLED"
)

type JobType string

// Enum values for JobType
const (
	JobTypeRelease  JobType = "RELEASE"
	JobTypeRetry    JobType = "RETRY"
	JobTypeManual   JobType = "MANUAL"
	JobTypeWeb_hook JobType = "WEB_HOOK"
)

type Platform string

// Enum values for Platform
const (
	PlatformWeb Platform = "WEB"
)

type Stage string

// Enum values for Stage
const (
	StageProduction   Stage = "PRODUCTION"
	StageBeta         Stage = "BETA"
	StageDevelopment  Stage = "DEVELOPMENT"
	StageExperimental Stage = "EXPERIMENTAL"
	StagePull_request Stage = "PULL_REQUEST"
)
