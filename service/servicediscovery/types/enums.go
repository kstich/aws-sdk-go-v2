// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type CustomHealthStatus string

// Enum values for CustomHealthStatus
const (
	CustomHealthStatusHealthy   CustomHealthStatus = "HEALTHY"
	CustomHealthStatusUnhealthy CustomHealthStatus = "UNHEALTHY"
)

type FilterCondition string

// Enum values for FilterCondition
const (
	FilterConditionEq      FilterCondition = "EQ"
	FilterConditionIn      FilterCondition = "IN"
	FilterConditionBetween FilterCondition = "BETWEEN"
)

type HealthCheckType string

// Enum values for HealthCheckType
const (
	HealthCheckTypeHttp  HealthCheckType = "HTTP"
	HealthCheckTypeHttps HealthCheckType = "HTTPS"
	HealthCheckTypeTcp   HealthCheckType = "TCP"
)

type HealthStatus string

// Enum values for HealthStatus
const (
	HealthStatusHealthy   HealthStatus = "HEALTHY"
	HealthStatusUnhealthy HealthStatus = "UNHEALTHY"
	HealthStatusUnknown   HealthStatus = "UNKNOWN"
)

type HealthStatusFilter string

// Enum values for HealthStatusFilter
const (
	HealthStatusFilterHealthy   HealthStatusFilter = "HEALTHY"
	HealthStatusFilterUnhealthy HealthStatusFilter = "UNHEALTHY"
	HealthStatusFilterAll       HealthStatusFilter = "ALL"
)

type NamespaceFilterName string

// Enum values for NamespaceFilterName
const (
	NamespaceFilterNameType NamespaceFilterName = "TYPE"
)

type NamespaceType string

// Enum values for NamespaceType
const (
	NamespaceTypeDns_public  NamespaceType = "DNS_PUBLIC"
	NamespaceTypeDns_private NamespaceType = "DNS_PRIVATE"
	NamespaceTypeHttp        NamespaceType = "HTTP"
)

type OperationFilterName string

// Enum values for OperationFilterName
const (
	OperationFilterNameNamespace_id OperationFilterName = "NAMESPACE_ID"
	OperationFilterNameService_id   OperationFilterName = "SERVICE_ID"
	OperationFilterNameStatus       OperationFilterName = "STATUS"
	OperationFilterNameType         OperationFilterName = "TYPE"
	OperationFilterNameUpdate_date  OperationFilterName = "UPDATE_DATE"
)

type OperationStatus string

// Enum values for OperationStatus
const (
	OperationStatusSubmitted OperationStatus = "SUBMITTED"
	OperationStatusPending   OperationStatus = "PENDING"
	OperationStatusSuccess   OperationStatus = "SUCCESS"
	OperationStatusFail      OperationStatus = "FAIL"
)

type OperationTargetType string

// Enum values for OperationTargetType
const (
	OperationTargetTypeNamespace OperationTargetType = "NAMESPACE"
	OperationTargetTypeService   OperationTargetType = "SERVICE"
	OperationTargetTypeInstance  OperationTargetType = "INSTANCE"
)

type OperationType string

// Enum values for OperationType
const (
	OperationTypeCreate_namespace    OperationType = "CREATE_NAMESPACE"
	OperationTypeDelete_namespace    OperationType = "DELETE_NAMESPACE"
	OperationTypeUpdate_service      OperationType = "UPDATE_SERVICE"
	OperationTypeRegister_instance   OperationType = "REGISTER_INSTANCE"
	OperationTypeDeregister_instance OperationType = "DEREGISTER_INSTANCE"
)

type RecordType string

// Enum values for RecordType
const (
	RecordTypeSrv   RecordType = "SRV"
	RecordTypeA     RecordType = "A"
	RecordTypeAaaa  RecordType = "AAAA"
	RecordTypeCname RecordType = "CNAME"
)

type RoutingPolicy string

// Enum values for RoutingPolicy
const (
	RoutingPolicyMultivalue RoutingPolicy = "MULTIVALUE"
	RoutingPolicyWeighted   RoutingPolicy = "WEIGHTED"
)

type ServiceFilterName string

// Enum values for ServiceFilterName
const (
	ServiceFilterNameNamespace_id ServiceFilterName = "NAMESPACE_ID"
)
