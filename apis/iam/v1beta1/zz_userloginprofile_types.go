/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type UserLoginProfileObservation struct {

	// The encrypted password, base64 encoded.
	EncryptedPassword *string `json:"encryptedPassword,omitempty" tf:"encrypted_password,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The fingerprint of the PGP key used to encrypt the password.
	KeyFingerprint *string `json:"keyFingerprint,omitempty" tf:"key_fingerprint,omitempty"`

	// The plain text password, only available when pgp_key is not provided.
	Password *string `json:"password,omitempty" tf:"password,omitempty"`
}

type UserLoginProfileParameters struct {

	// The length of the generated password on resource creation. Only applies on resource creation. Drift detection is not possible with this argument. Default value is 20.
	// +kubebuilder:validation:Optional
	PasswordLength *float64 `json:"passwordLength,omitempty" tf:"password_length,omitempty"`

	// Whether the user should be forced to reset the generated password on resource creation. Only applies on resource creation.
	// +kubebuilder:validation:Optional
	PasswordResetRequired *bool `json:"passwordResetRequired,omitempty" tf:"password_reset_required,omitempty"`

	// Either a base-64 encoded PGP public key, or a keybase username in the form keybase:username. Only applies on resource creation. Drift detection is not possible with this argument.
	// +kubebuilder:validation:Optional
	PgpKey *string `json:"pgpKey,omitempty" tf:"pgp_key,omitempty"`

	// The IAM user's name.
	// +crossplane:generate:reference:type=github.com/dkb-bank/official-provider-aws/apis/iam/v1beta1.User
	// +kubebuilder:validation:Optional
	User *string `json:"user,omitempty" tf:"user,omitempty"`

	// Reference to a User in iam to populate user.
	// +kubebuilder:validation:Optional
	UserRef *v1.Reference `json:"userRef,omitempty" tf:"-"`

	// Selector for a User in iam to populate user.
	// +kubebuilder:validation:Optional
	UserSelector *v1.Selector `json:"userSelector,omitempty" tf:"-"`
}

// UserLoginProfileSpec defines the desired state of UserLoginProfile
type UserLoginProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserLoginProfileParameters `json:"forProvider"`
}

// UserLoginProfileStatus defines the observed state of UserLoginProfile.
type UserLoginProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserLoginProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UserLoginProfile is the Schema for the UserLoginProfiles API. Manages an IAM User Login Profile
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type UserLoginProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserLoginProfileSpec   `json:"spec"`
	Status            UserLoginProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserLoginProfileList contains a list of UserLoginProfiles
type UserLoginProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserLoginProfile `json:"items"`
}

// Repository type metadata.
var (
	UserLoginProfile_Kind             = "UserLoginProfile"
	UserLoginProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserLoginProfile_Kind}.String()
	UserLoginProfile_KindAPIVersion   = UserLoginProfile_Kind + "." + CRDGroupVersion.String()
	UserLoginProfile_GroupVersionKind = CRDGroupVersion.WithKind(UserLoginProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&UserLoginProfile{}, &UserLoginProfileList{})
}
