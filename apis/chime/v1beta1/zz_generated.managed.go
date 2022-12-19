/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this VoiceConnector.
func (mg *VoiceConnector) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this VoiceConnector.
func (mg *VoiceConnector) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this VoiceConnector.
func (mg *VoiceConnector) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this VoiceConnector.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *VoiceConnector) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this VoiceConnector.
func (mg *VoiceConnector) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this VoiceConnector.
func (mg *VoiceConnector) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this VoiceConnector.
func (mg *VoiceConnector) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this VoiceConnector.
func (mg *VoiceConnector) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this VoiceConnector.
func (mg *VoiceConnector) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this VoiceConnector.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *VoiceConnector) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this VoiceConnector.
func (mg *VoiceConnector) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this VoiceConnector.
func (mg *VoiceConnector) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this VoiceConnectorGroup.
func (mg *VoiceConnectorGroup) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this VoiceConnectorGroup.
func (mg *VoiceConnectorGroup) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this VoiceConnectorGroup.
func (mg *VoiceConnectorGroup) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this VoiceConnectorGroup.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *VoiceConnectorGroup) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this VoiceConnectorGroup.
func (mg *VoiceConnectorGroup) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this VoiceConnectorGroup.
func (mg *VoiceConnectorGroup) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this VoiceConnectorGroup.
func (mg *VoiceConnectorGroup) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this VoiceConnectorGroup.
func (mg *VoiceConnectorGroup) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this VoiceConnectorGroup.
func (mg *VoiceConnectorGroup) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this VoiceConnectorGroup.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *VoiceConnectorGroup) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this VoiceConnectorGroup.
func (mg *VoiceConnectorGroup) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this VoiceConnectorGroup.
func (mg *VoiceConnectorGroup) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this VoiceConnectorLogging.
func (mg *VoiceConnectorLogging) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this VoiceConnectorLogging.
func (mg *VoiceConnectorLogging) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this VoiceConnectorLogging.
func (mg *VoiceConnectorLogging) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this VoiceConnectorLogging.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *VoiceConnectorLogging) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this VoiceConnectorLogging.
func (mg *VoiceConnectorLogging) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this VoiceConnectorLogging.
func (mg *VoiceConnectorLogging) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this VoiceConnectorLogging.
func (mg *VoiceConnectorLogging) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this VoiceConnectorLogging.
func (mg *VoiceConnectorLogging) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this VoiceConnectorLogging.
func (mg *VoiceConnectorLogging) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this VoiceConnectorLogging.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *VoiceConnectorLogging) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this VoiceConnectorLogging.
func (mg *VoiceConnectorLogging) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this VoiceConnectorLogging.
func (mg *VoiceConnectorLogging) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this VoiceConnectorOrigination.
func (mg *VoiceConnectorOrigination) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this VoiceConnectorOrigination.
func (mg *VoiceConnectorOrigination) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this VoiceConnectorOrigination.
func (mg *VoiceConnectorOrigination) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this VoiceConnectorOrigination.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *VoiceConnectorOrigination) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this VoiceConnectorOrigination.
func (mg *VoiceConnectorOrigination) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this VoiceConnectorOrigination.
func (mg *VoiceConnectorOrigination) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this VoiceConnectorOrigination.
func (mg *VoiceConnectorOrigination) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this VoiceConnectorOrigination.
func (mg *VoiceConnectorOrigination) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this VoiceConnectorOrigination.
func (mg *VoiceConnectorOrigination) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this VoiceConnectorOrigination.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *VoiceConnectorOrigination) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this VoiceConnectorOrigination.
func (mg *VoiceConnectorOrigination) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this VoiceConnectorOrigination.
func (mg *VoiceConnectorOrigination) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this VoiceConnectorStreaming.
func (mg *VoiceConnectorStreaming) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this VoiceConnectorStreaming.
func (mg *VoiceConnectorStreaming) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this VoiceConnectorStreaming.
func (mg *VoiceConnectorStreaming) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this VoiceConnectorStreaming.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *VoiceConnectorStreaming) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this VoiceConnectorStreaming.
func (mg *VoiceConnectorStreaming) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this VoiceConnectorStreaming.
func (mg *VoiceConnectorStreaming) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this VoiceConnectorStreaming.
func (mg *VoiceConnectorStreaming) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this VoiceConnectorStreaming.
func (mg *VoiceConnectorStreaming) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this VoiceConnectorStreaming.
func (mg *VoiceConnectorStreaming) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this VoiceConnectorStreaming.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *VoiceConnectorStreaming) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this VoiceConnectorStreaming.
func (mg *VoiceConnectorStreaming) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this VoiceConnectorStreaming.
func (mg *VoiceConnectorStreaming) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this VoiceConnectorTermination.
func (mg *VoiceConnectorTermination) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this VoiceConnectorTermination.
func (mg *VoiceConnectorTermination) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this VoiceConnectorTermination.
func (mg *VoiceConnectorTermination) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this VoiceConnectorTermination.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *VoiceConnectorTermination) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this VoiceConnectorTermination.
func (mg *VoiceConnectorTermination) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this VoiceConnectorTermination.
func (mg *VoiceConnectorTermination) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this VoiceConnectorTermination.
func (mg *VoiceConnectorTermination) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this VoiceConnectorTermination.
func (mg *VoiceConnectorTermination) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this VoiceConnectorTermination.
func (mg *VoiceConnectorTermination) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this VoiceConnectorTermination.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *VoiceConnectorTermination) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this VoiceConnectorTermination.
func (mg *VoiceConnectorTermination) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this VoiceConnectorTermination.
func (mg *VoiceConnectorTermination) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this VoiceConnectorTerminationCredentials.
func (mg *VoiceConnectorTerminationCredentials) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this VoiceConnectorTerminationCredentials.
func (mg *VoiceConnectorTerminationCredentials) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this VoiceConnectorTerminationCredentials.
func (mg *VoiceConnectorTerminationCredentials) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this VoiceConnectorTerminationCredentials.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *VoiceConnectorTerminationCredentials) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this VoiceConnectorTerminationCredentials.
func (mg *VoiceConnectorTerminationCredentials) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this VoiceConnectorTerminationCredentials.
func (mg *VoiceConnectorTerminationCredentials) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this VoiceConnectorTerminationCredentials.
func (mg *VoiceConnectorTerminationCredentials) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this VoiceConnectorTerminationCredentials.
func (mg *VoiceConnectorTerminationCredentials) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this VoiceConnectorTerminationCredentials.
func (mg *VoiceConnectorTerminationCredentials) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this VoiceConnectorTerminationCredentials.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *VoiceConnectorTerminationCredentials) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this VoiceConnectorTerminationCredentials.
func (mg *VoiceConnectorTerminationCredentials) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this VoiceConnectorTerminationCredentials.
func (mg *VoiceConnectorTerminationCredentials) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}