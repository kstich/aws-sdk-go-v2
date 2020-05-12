// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53resolver

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AssociateResolverEndpointIpAddressInput struct {
	_ struct{} `type:"structure"`

	// Either the IPv4 address that you want to add to a resolver endpoint or a
	// subnet ID. If you specify a subnet ID, Resolver chooses an IP address for
	// you from the available IPs in the specified subnet.
	//
	// IpAddress is a required field
	IpAddress *IpAddressUpdate `type:"structure" required:"true"`

	// The ID of the resolver endpoint that you want to associate IP addresses with.
	//
	// ResolverEndpointId is a required field
	ResolverEndpointId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateResolverEndpointIpAddressInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateResolverEndpointIpAddressInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateResolverEndpointIpAddressInput"}

	if s.IpAddress == nil {
		invalidParams.Add(aws.NewErrParamRequired("IpAddress"))
	}

	if s.ResolverEndpointId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResolverEndpointId"))
	}
	if s.ResolverEndpointId != nil && len(*s.ResolverEndpointId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResolverEndpointId", 1))
	}
	if s.IpAddress != nil {
		if err := s.IpAddress.Validate(); err != nil {
			invalidParams.AddNested("IpAddress", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AssociateResolverEndpointIpAddressOutput struct {
	_ struct{} `type:"structure"`

	// The response to an AssociateResolverEndpointIpAddress request.
	ResolverEndpoint *ResolverEndpoint `type:"structure"`
}

// String returns the string representation
func (s AssociateResolverEndpointIpAddressOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssociateResolverEndpointIpAddress = "AssociateResolverEndpointIpAddress"

// AssociateResolverEndpointIpAddressRequest returns a request value for making API operation for
// Amazon Route 53 Resolver.
//
// Adds IP addresses to an inbound or an outbound resolver endpoint. If you
// want to adding more than one IP address, submit one AssociateResolverEndpointIpAddress
// request for each IP address.
//
// To remove an IP address from an endpoint, see DisassociateResolverEndpointIpAddress.
//
//    // Example sending a request using AssociateResolverEndpointIpAddressRequest.
//    req := client.AssociateResolverEndpointIpAddressRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/AssociateResolverEndpointIpAddress
func (c *Client) AssociateResolverEndpointIpAddressRequest(input *AssociateResolverEndpointIpAddressInput) AssociateResolverEndpointIpAddressRequest {
	op := &aws.Operation{
		Name:       opAssociateResolverEndpointIpAddress,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateResolverEndpointIpAddressInput{}
	}

	req := c.newRequest(op, input, &AssociateResolverEndpointIpAddressOutput{})

	return AssociateResolverEndpointIpAddressRequest{Request: req, Input: input, Copy: c.AssociateResolverEndpointIpAddressRequest}
}

// AssociateResolverEndpointIpAddressRequest is the request type for the
// AssociateResolverEndpointIpAddress API operation.
type AssociateResolverEndpointIpAddressRequest struct {
	*aws.Request
	Input *AssociateResolverEndpointIpAddressInput
	Copy  func(*AssociateResolverEndpointIpAddressInput) AssociateResolverEndpointIpAddressRequest
}

// Send marshals and sends the AssociateResolverEndpointIpAddress API request.
func (r AssociateResolverEndpointIpAddressRequest) Send(ctx context.Context) (*AssociateResolverEndpointIpAddressResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateResolverEndpointIpAddressResponse{
		AssociateResolverEndpointIpAddressOutput: r.Request.Data.(*AssociateResolverEndpointIpAddressOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateResolverEndpointIpAddressResponse is the response type for the
// AssociateResolverEndpointIpAddress API operation.
type AssociateResolverEndpointIpAddressResponse struct {
	*AssociateResolverEndpointIpAddressOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateResolverEndpointIpAddress request.
func (r *AssociateResolverEndpointIpAddressResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}