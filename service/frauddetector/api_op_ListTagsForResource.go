// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package frauddetector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListTagsForResourceInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of objects to return for the request.
	MaxResults *int64 `locationName:"maxResults" min:"50" type:"integer"`

	// The next token from the previous results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The ARN that specifies the resource whose tags you want to list.
	//
	// ResourceARN is a required field
	ResourceARN *string `locationName:"resourceARN" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListTagsForResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTagsForResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTagsForResourceInput"}
	if s.MaxResults != nil && *s.MaxResults < 50 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 50))
	}

	if s.ResourceARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceARN"))
	}
	if s.ResourceARN != nil && len(*s.ResourceARN) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceARN", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListTagsForResourceOutput struct {
	_ struct{} `type:"structure"`

	// The next token for subsequent requests.
	NextToken *string `locationName:"nextToken" type:"string"`

	// A collection of key and value pairs.
	Tags []Tag `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s ListTagsForResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opListTagsForResource = "ListTagsForResource"

// ListTagsForResourceRequest returns a request value for making API operation for
// Amazon Fraud Detector.
//
// Lists all tags associated with the resource. This is a paginated API. To
// get the next page results, provide the pagination token from the response
// as part of your request. A null pagination token fetches the records from
// the beginning.
//
//    // Example sending a request using ListTagsForResourceRequest.
//    req := client.ListTagsForResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/frauddetector-2019-11-15/ListTagsForResource
func (c *Client) ListTagsForResourceRequest(input *ListTagsForResourceInput) ListTagsForResourceRequest {
	op := &aws.Operation{
		Name:       opListTagsForResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListTagsForResourceInput{}
	}

	req := c.newRequest(op, input, &ListTagsForResourceOutput{})

	return ListTagsForResourceRequest{Request: req, Input: input, Copy: c.ListTagsForResourceRequest}
}

// ListTagsForResourceRequest is the request type for the
// ListTagsForResource API operation.
type ListTagsForResourceRequest struct {
	*aws.Request
	Input *ListTagsForResourceInput
	Copy  func(*ListTagsForResourceInput) ListTagsForResourceRequest
}

// Send marshals and sends the ListTagsForResource API request.
func (r ListTagsForResourceRequest) Send(ctx context.Context) (*ListTagsForResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTagsForResourceResponse{
		ListTagsForResourceOutput: r.Request.Data.(*ListTagsForResourceOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListTagsForResourceRequestPaginator returns a paginator for ListTagsForResource.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListTagsForResourceRequest(input)
//   p := frauddetector.NewListTagsForResourceRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListTagsForResourcePaginator(req ListTagsForResourceRequest) ListTagsForResourcePaginator {
	return ListTagsForResourcePaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListTagsForResourceInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListTagsForResourcePaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListTagsForResourcePaginator struct {
	aws.Pager
}

func (p *ListTagsForResourcePaginator) CurrentPage() *ListTagsForResourceOutput {
	return p.Pager.CurrentPage().(*ListTagsForResourceOutput)
}

// ListTagsForResourceResponse is the response type for the
// ListTagsForResource API operation.
type ListTagsForResourceResponse struct {
	*ListTagsForResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTagsForResource request.
func (r *ListTagsForResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}