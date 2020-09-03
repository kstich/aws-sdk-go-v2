// Code generated by smithy-go-codegen DO NOT EDIT.

package detective

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/detective/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
)

type validateOpAcceptInvitation struct {
}

func (*validateOpAcceptInvitation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAcceptInvitation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AcceptInvitationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAcceptInvitationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateMembers struct {
}

func (*validateOpCreateMembers) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateMembers) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateMembersInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateMembersInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteGraph struct {
}

func (*validateOpDeleteGraph) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteGraph) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteGraphInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteGraphInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteMembers struct {
}

func (*validateOpDeleteMembers) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteMembers) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteMembersInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteMembersInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisassociateMembership struct {
}

func (*validateOpDisassociateMembership) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisassociateMembership) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisassociateMembershipInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisassociateMembershipInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetMembers struct {
}

func (*validateOpGetMembers) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetMembers) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetMembersInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetMembersInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListMembers struct {
}

func (*validateOpListMembers) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListMembers) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListMembersInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListMembersInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRejectInvitation struct {
}

func (*validateOpRejectInvitation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRejectInvitation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RejectInvitationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRejectInvitationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartMonitoringMember struct {
}

func (*validateOpStartMonitoringMember) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartMonitoringMember) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartMonitoringMemberInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartMonitoringMemberInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAcceptInvitationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAcceptInvitation{}, middleware.After)
}

func addOpCreateMembersValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateMembers{}, middleware.After)
}

func addOpDeleteGraphValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteGraph{}, middleware.After)
}

func addOpDeleteMembersValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteMembers{}, middleware.After)
}

func addOpDisassociateMembershipValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDisassociateMembership{}, middleware.After)
}

func addOpGetMembersValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetMembers{}, middleware.After)
}

func addOpListMembersValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListMembers{}, middleware.After)
}

func addOpRejectInvitationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRejectInvitation{}, middleware.After)
}

func addOpStartMonitoringMemberValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartMonitoringMember{}, middleware.After)
}

func validateAccount(v *types.Account) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Account"}
	if v.EmailAddress == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EmailAddress"))
	}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateAccountList(v []*types.Account) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AccountList"}
	for i := range v {
		if err := validateAccount(v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAcceptInvitationInput(v *AcceptInvitationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AcceptInvitationInput"}
	if v.GraphArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateMembersInput(v *CreateMembersInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateMembersInput"}
	if v.GraphArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphArn"))
	}
	if v.Accounts == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Accounts"))
	} else if v.Accounts != nil {
		if err := validateAccountList(v.Accounts); err != nil {
			invalidParams.AddNested("Accounts", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteGraphInput(v *DeleteGraphInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteGraphInput"}
	if v.GraphArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteMembersInput(v *DeleteMembersInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteMembersInput"}
	if v.GraphArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphArn"))
	}
	if v.AccountIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisassociateMembershipInput(v *DisassociateMembershipInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisassociateMembershipInput"}
	if v.GraphArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetMembersInput(v *GetMembersInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetMembersInput"}
	if v.GraphArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphArn"))
	}
	if v.AccountIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListMembersInput(v *ListMembersInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListMembersInput"}
	if v.GraphArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRejectInvitationInput(v *RejectInvitationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RejectInvitationInput"}
	if v.GraphArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartMonitoringMemberInput(v *StartMonitoringMemberInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartMonitoringMemberInput"}
	if v.GraphArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GraphArn"))
	}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
