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

// Modifies the parameters of a DB parameter group. To modify more than one
// parameter, submit a list of the following: ParameterName, ParameterValue, and
// ApplyMethod. A maximum of 20 parameters can be modified in a single request.
// Changes to dynamic parameters are applied immediately. Changes to static
// parameters require a reboot without failover to the DB instance associated with
// the parameter group before the change can take effect. After you modify a DB
// parameter group, you should wait at least 5 minutes before creating your first
// DB instance that uses that DB parameter group as the default parameter group.
// This allows Amazon RDS to fully complete the modify action before the parameter
// group is used as the default for a new DB instance. This is especially important
// for parameters that are critical when creating the default database for a DB
// instance, such as the character set for the default database defined by the
// character_set_database parameter. You can use the Parameter Groups option of the
// Amazon RDS console (https://console.aws.amazon.com/rds/) or the
// DescribeDBParameters command to verify that your DB parameter group has been
// created or modified.
func (c *Client) ModifyDBParameterGroup(ctx context.Context, params *ModifyDBParameterGroupInput, optFns ...func(*Options)) (*ModifyDBParameterGroupOutput, error) {
	stack := middleware.NewStack("ModifyDBParameterGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpModifyDBParameterGroupMiddlewares(stack)
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
	addOpModifyDBParameterGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDBParameterGroup(options.Region), middleware.Before)
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
			OperationName: "ModifyDBParameterGroup",
			Err:           err,
		}
	}
	out := result.(*ModifyDBParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ModifyDBParameterGroupInput struct {
	// An array of parameter names, values, and the apply method for the parameter
	// update. At least one parameter name, value, and apply method must be supplied;
	// later arguments are optional. A maximum of 20 parameters can be modified in a
	// single request. Valid Values (for the application method): immediate |
	// pending-reboot You can use the immediate value with dynamic parameters only. You
	// can use the pending-reboot value for both dynamic and static parameters, and
	// changes are applied when you reboot the DB instance without failover.
	Parameters []*types.Parameter
	// The name of the DB parameter group. Constraints:
	//
	//     * If supplied, must match
	// the name of an existing DBParameterGroup.
	DBParameterGroupName *string
}

// Contains the result of a successful invocation of the ModifyDBParameterGroup or
// ResetDBParameterGroup action.
type ModifyDBParameterGroupOutput struct {
	// Provides the name of the DB parameter group.
	DBParameterGroupName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpModifyDBParameterGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpModifyDBParameterGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyDBParameterGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyDBParameterGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ModifyDBParameterGroup",
	}
}
