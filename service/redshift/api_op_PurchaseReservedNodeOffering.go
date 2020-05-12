// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PurchaseReservedNodeOfferingInput struct {
	_ struct{} `type:"structure"`

	// The number of reserved nodes that you want to purchase.
	//
	// Default: 1
	NodeCount *int64 `type:"integer"`

	// The unique identifier of the reserved node offering you want to purchase.
	//
	// ReservedNodeOfferingId is a required field
	ReservedNodeOfferingId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s PurchaseReservedNodeOfferingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PurchaseReservedNodeOfferingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PurchaseReservedNodeOfferingInput"}

	if s.ReservedNodeOfferingId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReservedNodeOfferingId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PurchaseReservedNodeOfferingOutput struct {
	_ struct{} `type:"structure"`

	// Describes a reserved node. You can call the DescribeReservedNodeOfferings
	// API to obtain the available reserved node offerings.
	ReservedNode *ReservedNode `type:"structure"`
}

// String returns the string representation
func (s PurchaseReservedNodeOfferingOutput) String() string {
	return awsutil.Prettify(s)
}

const opPurchaseReservedNodeOffering = "PurchaseReservedNodeOffering"

// PurchaseReservedNodeOfferingRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Allows you to purchase reserved nodes. Amazon Redshift offers a predefined
// set of reserved node offerings. You can purchase one or more of the offerings.
// You can call the DescribeReservedNodeOfferings API to obtain the available
// reserved node offerings. You can call this API by providing a specific reserved
// node offering and the number of nodes you want to reserve.
//
// For more information about reserved node offerings, go to Purchasing Reserved
// Nodes (https://docs.aws.amazon.com/redshift/latest/mgmt/purchase-reserved-node-instance.html)
// in the Amazon Redshift Cluster Management Guide.
//
//    // Example sending a request using PurchaseReservedNodeOfferingRequest.
//    req := client.PurchaseReservedNodeOfferingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/PurchaseReservedNodeOffering
func (c *Client) PurchaseReservedNodeOfferingRequest(input *PurchaseReservedNodeOfferingInput) PurchaseReservedNodeOfferingRequest {
	op := &aws.Operation{
		Name:       opPurchaseReservedNodeOffering,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PurchaseReservedNodeOfferingInput{}
	}

	req := c.newRequest(op, input, &PurchaseReservedNodeOfferingOutput{})

	return PurchaseReservedNodeOfferingRequest{Request: req, Input: input, Copy: c.PurchaseReservedNodeOfferingRequest}
}

// PurchaseReservedNodeOfferingRequest is the request type for the
// PurchaseReservedNodeOffering API operation.
type PurchaseReservedNodeOfferingRequest struct {
	*aws.Request
	Input *PurchaseReservedNodeOfferingInput
	Copy  func(*PurchaseReservedNodeOfferingInput) PurchaseReservedNodeOfferingRequest
}

// Send marshals and sends the PurchaseReservedNodeOffering API request.
func (r PurchaseReservedNodeOfferingRequest) Send(ctx context.Context) (*PurchaseReservedNodeOfferingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PurchaseReservedNodeOfferingResponse{
		PurchaseReservedNodeOfferingOutput: r.Request.Data.(*PurchaseReservedNodeOfferingOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PurchaseReservedNodeOfferingResponse is the response type for the
// PurchaseReservedNodeOffering API operation.
type PurchaseReservedNodeOfferingResponse struct {
	*PurchaseReservedNodeOfferingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PurchaseReservedNodeOffering request.
func (r *PurchaseReservedNodeOfferingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}