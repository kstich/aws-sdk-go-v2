// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Provides more information regarding a given organization based on its
// identifier.
func (c *Client) DescribeOrganization(ctx context.Context, params *DescribeOrganizationInput, optFns ...func(*Options)) (*DescribeOrganizationOutput, error) {
	stack := middleware.NewStack("DescribeOrganization", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeOrganizationMiddlewares(stack)
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
	addOpDescribeOrganizationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOrganization(options.Region), middleware.Before)

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
			OperationName: "DescribeOrganization",
			Err:           err,
		}
	}
	out := result.(*DescribeOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeOrganizationInput struct {
	// The identifier for the organization to be described.
	OrganizationId *string
}

type DescribeOrganizationOutput struct {
	// The identifier for the directory associated with an Amazon WorkMail
	// organization.
	DirectoryId *string
	// (Optional) The error message indicating if unexpected behavior was encountered
	// with regards to the organization.
	ErrorMessage *string
	// The default mail domain associated with the organization.
	DefaultMailDomain *string
	// The identifier of an organization.
	OrganizationId *string
	// The date at which the organization became usable in the WorkMail context, in
	// UNIX epoch time format.
	CompletedDate *time.Time
	// The type of directory associated with the WorkMail organization.
	DirectoryType *string
	// The alias for an organization.
	Alias *string
	// The state of an organization.
	State *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeOrganizationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeOrganization{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeOrganization{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeOrganization(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "DescribeOrganization",
	}
}
