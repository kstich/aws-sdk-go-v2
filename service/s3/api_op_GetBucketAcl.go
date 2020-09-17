// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This implementation of the GET operation uses the acl subresource to return the
// access control list (ACL) of a bucket. To use GET to return the ACL of the
// bucket, you must have READ_ACP access to the bucket. If READ_ACP permission is
// granted to the anonymous user, you can return the ACL of the bucket without
// using an authorization header.  <p class="title"> <b>Related Resources</b> </p>
// <ul> <li> <p> </p> </li> </ul>
func (c *Client) GetBucketAcl(ctx context.Context, params *GetBucketAclInput, optFns ...func(*Options)) (*GetBucketAclOutput, error) {
	stack := middleware.NewStack("GetBucketAcl", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetBucketAclMiddlewares(stack)
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
	addOpGetBucketAclValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketAcl(options.Region), middleware.Before)
	addResponseErrorWrapper(stack)
	addUpdateEndpointMiddleware(stack, options)

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
			OperationName: "GetBucketAcl",
			Err:           err,
		}
	}
	out := result.(*GetBucketAclOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketAclInput struct {
	// Specifies the S3 bucket whose ACL is being requested.
	Bucket *string
}

type GetBucketAclOutput struct {
	// A list of grants.
	Grants []*types.Grant
	// Container for the bucket owner's display name and ID.
	Owner *types.Owner

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetBucketAclMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetBucketAcl{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketAcl{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetBucketAcl(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketAcl",
	}
}
