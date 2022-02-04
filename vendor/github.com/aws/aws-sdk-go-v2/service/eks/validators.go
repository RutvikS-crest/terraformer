// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpAssociateEncryptionConfig struct {
}

func (*validateOpAssociateEncryptionConfig) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAssociateEncryptionConfig) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AssociateEncryptionConfigInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAssociateEncryptionConfigInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpAssociateIdentityProviderConfig struct {
}

func (*validateOpAssociateIdentityProviderConfig) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAssociateIdentityProviderConfig) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AssociateIdentityProviderConfigInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAssociateIdentityProviderConfigInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateAddon struct {
}

func (*validateOpCreateAddon) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateAddon) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateAddonInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateAddonInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateCluster struct {
}

func (*validateOpCreateCluster) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateCluster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateClusterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateClusterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateFargateProfile struct {
}

func (*validateOpCreateFargateProfile) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateFargateProfile) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateFargateProfileInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateFargateProfileInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateNodegroup struct {
}

func (*validateOpCreateNodegroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateNodegroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateNodegroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateNodegroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteAddon struct {
}

func (*validateOpDeleteAddon) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteAddon) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteAddonInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteAddonInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteCluster struct {
}

func (*validateOpDeleteCluster) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteCluster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteClusterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteClusterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteFargateProfile struct {
}

func (*validateOpDeleteFargateProfile) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteFargateProfile) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteFargateProfileInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteFargateProfileInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteNodegroup struct {
}

func (*validateOpDeleteNodegroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteNodegroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteNodegroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteNodegroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeAddon struct {
}

func (*validateOpDescribeAddon) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeAddon) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeAddonInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeAddonInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeCluster struct {
}

func (*validateOpDescribeCluster) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeCluster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeClusterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeClusterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeFargateProfile struct {
}

func (*validateOpDescribeFargateProfile) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeFargateProfile) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeFargateProfileInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeFargateProfileInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeIdentityProviderConfig struct {
}

func (*validateOpDescribeIdentityProviderConfig) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeIdentityProviderConfig) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeIdentityProviderConfigInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeIdentityProviderConfigInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeNodegroup struct {
}

func (*validateOpDescribeNodegroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeNodegroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeNodegroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeNodegroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeUpdate struct {
}

func (*validateOpDescribeUpdate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeUpdate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeUpdateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeUpdateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisassociateIdentityProviderConfig struct {
}

func (*validateOpDisassociateIdentityProviderConfig) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisassociateIdentityProviderConfig) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisassociateIdentityProviderConfigInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisassociateIdentityProviderConfigInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListAddons struct {
}

func (*validateOpListAddons) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListAddons) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListAddonsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListAddonsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListFargateProfiles struct {
}

func (*validateOpListFargateProfiles) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListFargateProfiles) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListFargateProfilesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListFargateProfilesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListIdentityProviderConfigs struct {
}

func (*validateOpListIdentityProviderConfigs) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListIdentityProviderConfigs) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListIdentityProviderConfigsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListIdentityProviderConfigsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListNodegroups struct {
}

func (*validateOpListNodegroups) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListNodegroups) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListNodegroupsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListNodegroupsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListUpdates struct {
}

func (*validateOpListUpdates) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListUpdates) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListUpdatesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListUpdatesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateAddon struct {
}

func (*validateOpUpdateAddon) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateAddon) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateAddonInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateAddonInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateClusterConfig struct {
}

func (*validateOpUpdateClusterConfig) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateClusterConfig) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateClusterConfigInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateClusterConfigInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateClusterVersion struct {
}

func (*validateOpUpdateClusterVersion) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateClusterVersion) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateClusterVersionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateClusterVersionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateNodegroupConfig struct {
}

func (*validateOpUpdateNodegroupConfig) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateNodegroupConfig) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateNodegroupConfigInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateNodegroupConfigInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateNodegroupVersion struct {
}

func (*validateOpUpdateNodegroupVersion) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateNodegroupVersion) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateNodegroupVersionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateNodegroupVersionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAssociateEncryptionConfigValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAssociateEncryptionConfig{}, middleware.After)
}

func addOpAssociateIdentityProviderConfigValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAssociateIdentityProviderConfig{}, middleware.After)
}

func addOpCreateAddonValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateAddon{}, middleware.After)
}

func addOpCreateClusterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateCluster{}, middleware.After)
}

func addOpCreateFargateProfileValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateFargateProfile{}, middleware.After)
}

func addOpCreateNodegroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateNodegroup{}, middleware.After)
}

func addOpDeleteAddonValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteAddon{}, middleware.After)
}

func addOpDeleteClusterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteCluster{}, middleware.After)
}

func addOpDeleteFargateProfileValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteFargateProfile{}, middleware.After)
}

func addOpDeleteNodegroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteNodegroup{}, middleware.After)
}

func addOpDescribeAddonValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeAddon{}, middleware.After)
}

func addOpDescribeClusterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeCluster{}, middleware.After)
}

func addOpDescribeFargateProfileValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeFargateProfile{}, middleware.After)
}

func addOpDescribeIdentityProviderConfigValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeIdentityProviderConfig{}, middleware.After)
}

func addOpDescribeNodegroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeNodegroup{}, middleware.After)
}

func addOpDescribeUpdateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeUpdate{}, middleware.After)
}

func addOpDisassociateIdentityProviderConfigValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDisassociateIdentityProviderConfig{}, middleware.After)
}

func addOpListAddonsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListAddons{}, middleware.After)
}

func addOpListFargateProfilesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListFargateProfiles{}, middleware.After)
}

func addOpListIdentityProviderConfigsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListIdentityProviderConfigs{}, middleware.After)
}

func addOpListNodegroupsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListNodegroups{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpListUpdatesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListUpdates{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateAddonValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateAddon{}, middleware.After)
}

func addOpUpdateClusterConfigValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateClusterConfig{}, middleware.After)
}

func addOpUpdateClusterVersionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateClusterVersion{}, middleware.After)
}

func addOpUpdateNodegroupConfigValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateNodegroupConfig{}, middleware.After)
}

func addOpUpdateNodegroupVersionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateNodegroupVersion{}, middleware.After)
}

func validateIdentityProviderConfig(v *types.IdentityProviderConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "IdentityProviderConfig"}
	if v.Type == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOidcIdentityProviderConfigRequest(v *types.OidcIdentityProviderConfigRequest) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "OidcIdentityProviderConfigRequest"}
	if v.IdentityProviderConfigName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityProviderConfigName"))
	}
	if v.IssuerUrl == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IssuerUrl"))
	}
	if v.ClientId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClientId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAssociateEncryptionConfigInput(v *AssociateEncryptionConfigInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssociateEncryptionConfigInput"}
	if v.ClusterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterName"))
	}
	if v.EncryptionConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EncryptionConfig"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAssociateIdentityProviderConfigInput(v *AssociateIdentityProviderConfigInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssociateIdentityProviderConfigInput"}
	if v.ClusterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterName"))
	}
	if v.Oidc == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Oidc"))
	} else if v.Oidc != nil {
		if err := validateOidcIdentityProviderConfigRequest(v.Oidc); err != nil {
			invalidParams.AddNested("Oidc", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateAddonInput(v *CreateAddonInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateAddonInput"}
	if v.ClusterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterName"))
	}
	if v.AddonName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AddonName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateClusterInput(v *CreateClusterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateClusterInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.RoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleArn"))
	}
	if v.ResourcesVpcConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourcesVpcConfig"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateFargateProfileInput(v *CreateFargateProfileInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateFargateProfileInput"}
	if v.FargateProfileName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FargateProfileName"))
	}
	if v.ClusterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterName"))
	}
	if v.PodExecutionRoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PodExecutionRoleArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateNodegroupInput(v *CreateNodegroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateNodegroupInput"}
	if v.ClusterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterName"))
	}
	if v.NodegroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NodegroupName"))
	}
	if v.Subnets == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Subnets"))
	}
	if v.NodeRole == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NodeRole"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteAddonInput(v *DeleteAddonInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteAddonInput"}
	if v.ClusterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterName"))
	}
	if v.AddonName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AddonName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteClusterInput(v *DeleteClusterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteClusterInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteFargateProfileInput(v *DeleteFargateProfileInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteFargateProfileInput"}
	if v.ClusterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterName"))
	}
	if v.FargateProfileName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FargateProfileName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteNodegroupInput(v *DeleteNodegroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteNodegroupInput"}
	if v.ClusterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterName"))
	}
	if v.NodegroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NodegroupName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeAddonInput(v *DescribeAddonInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeAddonInput"}
	if v.ClusterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterName"))
	}
	if v.AddonName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AddonName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeClusterInput(v *DescribeClusterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeClusterInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeFargateProfileInput(v *DescribeFargateProfileInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeFargateProfileInput"}
	if v.ClusterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterName"))
	}
	if v.FargateProfileName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FargateProfileName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeIdentityProviderConfigInput(v *DescribeIdentityProviderConfigInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeIdentityProviderConfigInput"}
	if v.ClusterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterName"))
	}
	if v.IdentityProviderConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityProviderConfig"))
	} else if v.IdentityProviderConfig != nil {
		if err := validateIdentityProviderConfig(v.IdentityProviderConfig); err != nil {
			invalidParams.AddNested("IdentityProviderConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeNodegroupInput(v *DescribeNodegroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeNodegroupInput"}
	if v.ClusterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterName"))
	}
	if v.NodegroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NodegroupName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeUpdateInput(v *DescribeUpdateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeUpdateInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.UpdateId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UpdateId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisassociateIdentityProviderConfigInput(v *DisassociateIdentityProviderConfigInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisassociateIdentityProviderConfigInput"}
	if v.ClusterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterName"))
	}
	if v.IdentityProviderConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityProviderConfig"))
	} else if v.IdentityProviderConfig != nil {
		if err := validateIdentityProviderConfig(v.IdentityProviderConfig); err != nil {
			invalidParams.AddNested("IdentityProviderConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListAddonsInput(v *ListAddonsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListAddonsInput"}
	if v.ClusterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListFargateProfilesInput(v *ListFargateProfilesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListFargateProfilesInput"}
	if v.ClusterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListIdentityProviderConfigsInput(v *ListIdentityProviderConfigsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListIdentityProviderConfigsInput"}
	if v.ClusterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListNodegroupsInput(v *ListNodegroupsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListNodegroupsInput"}
	if v.ClusterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListUpdatesInput(v *ListUpdatesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListUpdatesInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateAddonInput(v *UpdateAddonInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateAddonInput"}
	if v.ClusterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterName"))
	}
	if v.AddonName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AddonName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateClusterConfigInput(v *UpdateClusterConfigInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateClusterConfigInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateClusterVersionInput(v *UpdateClusterVersionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateClusterVersionInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Version == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Version"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateNodegroupConfigInput(v *UpdateNodegroupConfigInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateNodegroupConfigInput"}
	if v.ClusterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterName"))
	}
	if v.NodegroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NodegroupName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateNodegroupVersionInput(v *UpdateNodegroupVersionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateNodegroupVersionInput"}
	if v.ClusterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterName"))
	}
	if v.NodegroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NodegroupName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
