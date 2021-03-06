// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package inspector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListEventSubscriptionsInput struct {
	_ struct{} `type:"structure"`

	// You can use this parameter to indicate the maximum number of items you want
	// in the response. The default value is 10. The maximum value is 500.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// You can use this parameter when paginating results. Set the value of this
	// parameter to null on your first call to the ListEventSubscriptions action.
	// Subsequent calls to the action fill nextToken in the request with the value
	// of NextToken from the previous response to continue listing data.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`

	// The ARN of the assessment template for which you want to list the existing
	// event subscriptions.
	ResourceArn *string `locationName:"resourceArn" min:"1" type:"string"`
}

// String returns the string representation
func (s ListEventSubscriptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListEventSubscriptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListEventSubscriptionsInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}
	if s.ResourceArn != nil && len(*s.ResourceArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListEventSubscriptionsOutput struct {
	_ struct{} `type:"structure"`

	// When a response is generated, if there is more data to be listed, this parameter
	// is present in the response and contains the value to use for the nextToken
	// parameter in a subsequent pagination request. If there is no more data to
	// be listed, this parameter is set to null.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`

	// Details of the returned event subscriptions.
	//
	// Subscriptions is a required field
	Subscriptions []Subscription `locationName:"subscriptions" type:"list" required:"true"`
}

// String returns the string representation
func (s ListEventSubscriptionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListEventSubscriptions = "ListEventSubscriptions"

// ListEventSubscriptionsRequest returns a request value for making API operation for
// Amazon Inspector.
//
// Lists all the event subscriptions for the assessment template that is specified
// by the ARN of the assessment template. For more information, see SubscribeToEvent
// and UnsubscribeFromEvent.
//
//    // Example sending a request using ListEventSubscriptionsRequest.
//    req := client.ListEventSubscriptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/ListEventSubscriptions
func (c *Client) ListEventSubscriptionsRequest(input *ListEventSubscriptionsInput) ListEventSubscriptionsRequest {
	op := &aws.Operation{
		Name:       opListEventSubscriptions,
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
		input = &ListEventSubscriptionsInput{}
	}

	req := c.newRequest(op, input, &ListEventSubscriptionsOutput{})

	return ListEventSubscriptionsRequest{Request: req, Input: input, Copy: c.ListEventSubscriptionsRequest}
}

// ListEventSubscriptionsRequest is the request type for the
// ListEventSubscriptions API operation.
type ListEventSubscriptionsRequest struct {
	*aws.Request
	Input *ListEventSubscriptionsInput
	Copy  func(*ListEventSubscriptionsInput) ListEventSubscriptionsRequest
}

// Send marshals and sends the ListEventSubscriptions API request.
func (r ListEventSubscriptionsRequest) Send(ctx context.Context) (*ListEventSubscriptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListEventSubscriptionsResponse{
		ListEventSubscriptionsOutput: r.Request.Data.(*ListEventSubscriptionsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListEventSubscriptionsRequestPaginator returns a paginator for ListEventSubscriptions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListEventSubscriptionsRequest(input)
//   p := inspector.NewListEventSubscriptionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListEventSubscriptionsPaginator(req ListEventSubscriptionsRequest) ListEventSubscriptionsPaginator {
	return ListEventSubscriptionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListEventSubscriptionsInput
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

// ListEventSubscriptionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListEventSubscriptionsPaginator struct {
	aws.Pager
}

func (p *ListEventSubscriptionsPaginator) CurrentPage() *ListEventSubscriptionsOutput {
	return p.Pager.CurrentPage().(*ListEventSubscriptionsOutput)
}

// ListEventSubscriptionsResponse is the response type for the
// ListEventSubscriptions API operation.
type ListEventSubscriptionsResponse struct {
	*ListEventSubscriptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListEventSubscriptions request.
func (r *ListEventSubscriptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
