// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deploys an application revision through the specified deployment group.
func (c *Client) CreateDeployment(ctx context.Context, params *CreateDeploymentInput, optFns ...func(*Options)) (*CreateDeploymentOutput, error) {
	stack := middleware.NewStack("CreateDeployment", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateDeploymentMiddlewares(stack)
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
	addOpCreateDeploymentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDeployment(options.Region), middleware.Before)

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
			OperationName: "CreateDeployment",
			Err:           err,
		}
	}
	out := result.(*CreateDeploymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a CreateDeployment operation.
type CreateDeploymentInput struct {
	// The name of a deployment configuration associated with the IAM user or AWS
	// account. If not specified, the value configured in the deployment group is used
	// as the default. If the deployment group does not have a deployment configuration
	// associated with it, CodeDeployDefault.OneAtATime is used by default.
	DeploymentConfigName *string
	// A comment about the deployment.
	Description *string
	// Information about how AWS CodeDeploy handles files that already exist in a
	// deployment target location but weren't part of the previous successful
	// deployment. The fileExistsBehavior parameter takes any of the following
	// values:
	//
	//     * DISALLOW: The deployment fails. This is also the default behavior
	// if no option is specified.
	//
	//     * OVERWRITE: The version of the file from the
	// application revision currently being deployed replaces the version already on
	// the instance.
	//
	//     * RETAIN: The version of the file already on the instance is
	// kept and used as part of the new deployment.
	FileExistsBehavior types.FileExistsBehavior
	// Information about the instances that belong to the replacement environment in a
	// blue/green deployment.
	TargetInstances *types.TargetInstances
	// If true, then if an ApplicationStop, BeforeBlockTraffic, or AfterBlockTraffic
	// deployment lifecycle event to an instance fails, then the deployment continues
	// to the next deployment lifecycle event. For example, if ApplicationStop fails,
	// the deployment continues with DownloadBundle. If BeforeBlockTraffic fails, the
	// deployment continues with BlockTraffic. If AfterBlockTraffic fails, the
	// deployment continues with ApplicationStop.  <p> If false or not specified, then
	// if a lifecycle event fails during a deployment to an instance, that deployment
	// fails. If deployment to that instance is part of an overall deployment and the
	// number of healthy hosts is not less than the minimum number of healthy hosts,
	// then a deployment to the next instance is attempted. </p> <p> During a
	// deployment, the AWS CodeDeploy agent runs the scripts specified for
	// <code>ApplicationStop</code>, <code>BeforeBlockTraffic</code>, and
	// <code>AfterBlockTraffic</code> in the AppSpec file from the previous successful
	// deployment. (All other scripts are run from the AppSpec file in the current
	// deployment.) If one of these scripts contains an error and does not run
	// successfully, the deployment can fail. </p> <p> If the cause of the failure is a
	// script from the last successful deployment that will never run successfully,
	// create a new deployment and use <code>ignoreApplicationStopFailures</code> to
	// specify that the <code>ApplicationStop</code>, <code>BeforeBlockTraffic</code>,
	// and <code>AfterBlockTraffic</code> failures should be ignored. </p>
	IgnoreApplicationStopFailures *bool
	// Indicates whether to deploy to all instances or only to instances that are not
	// running the latest application revision.
	UpdateOutdatedInstancesOnly *bool
	// The name of the deployment group.
	DeploymentGroupName *string
	// The name of an AWS CodeDeploy application associated with the IAM user or AWS
	// account.
	ApplicationName *string
	// The type and location of the revision to deploy.
	Revision *types.RevisionLocation
	// Configuration information for an automatic rollback that is added when a
	// deployment is created.
	AutoRollbackConfiguration *types.AutoRollbackConfiguration
}

// Represents the output of a CreateDeployment operation.
type CreateDeploymentOutput struct {
	// The unique ID of a deployment.
	DeploymentId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateDeploymentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDeployment{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDeployment{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDeployment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "CreateDeployment",
	}
}
