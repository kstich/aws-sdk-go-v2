// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists Amazon GuardDuty findings for the specified detector ID.
func (c *Client) ListFindings(ctx context.Context, params *ListFindingsInput, optFns ...func(*Options)) (*ListFindingsOutput, error) {
	stack := middleware.NewStack("ListFindings", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListFindingsMiddlewares(stack)
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
	addOpListFindingsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListFindings(options.Region), middleware.Before)
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
			OperationName: "ListFindings",
			Err:           err,
		}
	}
	out := result.(*ListFindingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFindingsInput struct {
	// The ID of the detector that specifies the GuardDuty service whose findings you
	// want to list.
	DetectorId *string
	// You can use this parameter to indicate the maximum number of items you want in
	// the response. The default value is 50. The maximum value is 50.
	MaxResults *int32
	// You can use this parameter when paginating results. Set the value of this
	// parameter to null on your first call to the list action. For subsequent calls to
	// the action, fill nextToken in the request with the value of NextToken from the
	// previous response to continue listing data.
	NextToken *string
	// Represents the criteria used for querying findings. Valid values include:
	//
	//     *
	// JSON field name
	//
	//     * accountId
	//
	//     * region
	//
	//     * confidence
	//
	//     * id
	//
	//
	// * resource.accessKeyDetails.accessKeyId
	//
	//     *
	// resource.accessKeyDetails.principalId
	//
	//     *
	// resource.accessKeyDetails.userName
	//
	//     * resource.accessKeyDetails.userType
	//
	//
	// * resource.instanceDetails.iamInstanceProfile.id
	//
	//     *
	// resource.instanceDetails.imageId
	//
	//     * resource.instanceDetails.instanceId
	//
	//
	// * resource.instanceDetails.networkInterfaces.ipv6Addresses
	//
	//     *
	// resource.instanceDetails.networkInterfaces.privateIpAddresses.privateIpAddress
	//
	//
	// * resource.instanceDetails.networkInterfaces.publicDnsName
	//
	//     *
	// resource.instanceDetails.networkInterfaces.publicIp
	//
	//     *
	// resource.instanceDetails.networkInterfaces.securityGroups.groupId
	//
	//     *
	// resource.instanceDetails.networkInterfaces.securityGroups.groupName
	//
	//     *
	// resource.instanceDetails.networkInterfaces.subnetId
	//
	//     *
	// resource.instanceDetails.networkInterfaces.vpcId
	//
	//     *
	// resource.instanceDetails.tags.key
	//
	//     * resource.instanceDetails.tags.value
	//
	//
	// * resource.resourceType
	//
	//     * service.action.actionType
	//
	//     *
	// service.action.awsApiCallAction.api
	//
	//     *
	// service.action.awsApiCallAction.callerType
	//
	//     *
	// service.action.awsApiCallAction.remoteIpDetails.city.cityName
	//
	//     *
	// service.action.awsApiCallAction.remoteIpDetails.country.countryName
	//
	//     *
	// service.action.awsApiCallAction.remoteIpDetails.ipAddressV4
	//
	//     *
	// service.action.awsApiCallAction.remoteIpDetails.organization.asn
	//
	//     *
	// service.action.awsApiCallAction.remoteIpDetails.organization.asnOrg
	//
	//     *
	// service.action.awsApiCallAction.serviceName
	//
	//     *
	// service.action.dnsRequestAction.domain
	//
	//     *
	// service.action.networkConnectionAction.blocked
	//
	//     *
	// service.action.networkConnectionAction.connectionDirection
	//
	//     *
	// service.action.networkConnectionAction.localPortDetails.port
	//
	//     *
	// service.action.networkConnectionAction.protocol
	//
	//     *
	// service.action.networkConnectionAction.remoteIpDetails.city.cityName
	//
	//     *
	// service.action.networkConnectionAction.remoteIpDetails.country.countryName
	//
	//
	// * service.action.networkConnectionAction.remoteIpDetails.ipAddressV4
	//
	//     *
	// service.action.networkConnectionAction.remoteIpDetails.organization.asn
	//
	//     *
	// service.action.networkConnectionAction.remoteIpDetails.organization.asnOrg
	//
	//
	// * service.action.networkConnectionAction.remotePortDetails.port
	//
	//     *
	// service.additionalInfo.threatListName
	//
	//     * service.archived When this
	// attribute is set to 'true', only archived findings are listed. When it's set to
	// 'false', only unarchived findings are listed. When this attribute is not set,
	// all existing findings are listed.
	//
	//     * service.resourceRole
	//
	//     * severity
	//
	//
	// * type
	//
	//     * updatedAt Type: Timestamp in Unix Epoch millisecond format:
	// 1486685375000
	FindingCriteria *types.FindingCriteria
	// Represents the criteria used for sorting findings.
	SortCriteria *types.SortCriteria
}

type ListFindingsOutput struct {
	// The IDs of the findings that you're listing.
	FindingIds []*string
	// The pagination parameter to be used on the next list operation to retrieve more
	// items.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListFindingsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListFindings{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListFindings{}, middleware.After)
}

func newServiceMetadataMiddleware_opListFindings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "ListFindings",
	}
}
