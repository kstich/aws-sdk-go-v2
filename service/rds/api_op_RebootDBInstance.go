// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// You might need to reboot your DB instance, usually for maintenance reasons. For
// example, if you make certain modifications, or if you change the DB parameter
// group associated with the DB instance, you must reboot the instance for the
// changes to take effect.  <p>Rebooting a DB instance restarts the database engine
// service. Rebooting a DB instance results in a momentary outage, during which the
// DB instance status is set to rebooting. </p> <p>For more information about
// rebooting, see <a
// href="https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_RebootInstance.html">Rebooting
// a DB Instance</a> in the <i>Amazon RDS User Guide.</i> </p>
func (c *Client) RebootDBInstance(ctx context.Context, params *RebootDBInstanceInput, optFns ...func(*Options)) (*RebootDBInstanceOutput, error) {
	stack := middleware.NewStack("RebootDBInstance", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpRebootDBInstanceMiddlewares(stack)
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
	addOpRebootDBInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRebootDBInstance(options.Region), middleware.Before)
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
			OperationName: "RebootDBInstance",
			Err:           err,
		}
	}
	out := result.(*RebootDBInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type RebootDBInstanceInput struct {
	// The DB instance identifier. This parameter is stored as a lowercase string.
	// Constraints:
	//
	//     * Must match the identifier of an existing DBInstance.
	DBInstanceIdentifier *string
	// A value that indicates whether the reboot is conducted through a Multi-AZ
	// failover. Constraint: You can't enable force failover if the instance isn't
	// configured for Multi-AZ.
	ForceFailover *bool
}

type RebootDBInstanceOutput struct {
	// Contains the details of an Amazon RDS DB instance. This data type is used as a
	// response element in the DescribeDBInstances action.
	DBInstance *types.DBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpRebootDBInstanceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpRebootDBInstance{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpRebootDBInstance{}, middleware.After)
}

func newServiceMetadataMiddleware_opRebootDBInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "RebootDBInstance",
	}
}
