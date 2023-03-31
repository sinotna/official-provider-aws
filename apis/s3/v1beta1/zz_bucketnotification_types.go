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

type BucketNotificationObservation struct {

	// Unique identifier for each of the notification configurations.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BucketNotificationParameters struct {

	// Name of the bucket for notification configuration.
	// +crossplane:generate:reference:type=github.com/dkb-bank/official-provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Whether to enable Amazon EventBridge notifications.
	// +kubebuilder:validation:Optional
	Eventbridge *bool `json:"eventbridge,omitempty" tf:"eventbridge,omitempty"`

	// Used to configure notifications to a Lambda Function. See below.
	// +kubebuilder:validation:Optional
	LambdaFunction []LambdaFunctionParameters `json:"lambdaFunction,omitempty" tf:"lambda_function,omitempty"`

	// Notification configuration to SQS Queue. See below.
	// +kubebuilder:validation:Optional
	Queue []QueueParameters `json:"queue,omitempty" tf:"queue,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Notification configuration to SNS Topic. See below.
	// +kubebuilder:validation:Optional
	Topic []TopicParameters `json:"topic,omitempty" tf:"topic,omitempty"`
}

type LambdaFunctionObservation struct {
}

type LambdaFunctionParameters struct {

	// Event for which to send notifications.
	// +kubebuilder:validation:Required
	Events []*string `json:"events" tf:"events,omitempty"`

	// Object key name prefix.
	// +kubebuilder:validation:Optional
	FilterPrefix *string `json:"filterPrefix,omitempty" tf:"filter_prefix,omitempty"`

	// Object key name suffix.
	// +kubebuilder:validation:Optional
	FilterSuffix *string `json:"filterSuffix,omitempty" tf:"filter_suffix,omitempty"`

	// Unique identifier for each of the notification configurations.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Lambda function ARN.
	// +kubebuilder:validation:Optional
	LambdaFunctionArn *string `json:"lambdaFunctionArn,omitempty" tf:"lambda_function_arn,omitempty"`
}

type QueueObservation struct {
}

type QueueParameters struct {

	// Specifies event for which to send notifications.
	// +kubebuilder:validation:Required
	Events []*string `json:"events" tf:"events,omitempty"`

	// Object key name prefix.
	// +kubebuilder:validation:Optional
	FilterPrefix *string `json:"filterPrefix,omitempty" tf:"filter_prefix,omitempty"`

	// Object key name suffix.
	// +kubebuilder:validation:Optional
	FilterSuffix *string `json:"filterSuffix,omitempty" tf:"filter_suffix,omitempty"`

	// Unique identifier for each of the notification configurations.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// SQS queue ARN.
	// +kubebuilder:validation:Required
	QueueArn *string `json:"queueArn" tf:"queue_arn,omitempty"`
}

type TopicObservation struct {
}

type TopicParameters struct {

	// Event for which to send notifications.
	// +kubebuilder:validation:Required
	Events []*string `json:"events" tf:"events,omitempty"`

	// Object key name prefix.
	// +kubebuilder:validation:Optional
	FilterPrefix *string `json:"filterPrefix,omitempty" tf:"filter_prefix,omitempty"`

	// Object key name suffix.
	// +kubebuilder:validation:Optional
	FilterSuffix *string `json:"filterSuffix,omitempty" tf:"filter_suffix,omitempty"`

	// Unique identifier for each of the notification configurations.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// SNS topic ARN.
	// +kubebuilder:validation:Required
	TopicArn *string `json:"topicArn" tf:"topic_arn,omitempty"`
}

// BucketNotificationSpec defines the desired state of BucketNotification
type BucketNotificationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketNotificationParameters `json:"forProvider"`
}

// BucketNotificationStatus defines the observed state of BucketNotification.
type BucketNotificationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketNotificationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketNotification is the Schema for the BucketNotifications API. Manages a S3 Bucket Notification Configuration
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BucketNotification struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketNotificationSpec   `json:"spec"`
	Status            BucketNotificationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketNotificationList contains a list of BucketNotifications
type BucketNotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketNotification `json:"items"`
}

// Repository type metadata.
var (
	BucketNotification_Kind             = "BucketNotification"
	BucketNotification_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketNotification_Kind}.String()
	BucketNotification_KindAPIVersion   = BucketNotification_Kind + "." + CRDGroupVersion.String()
	BucketNotification_GroupVersionKind = CRDGroupVersion.WithKind(BucketNotification_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketNotification{}, &BucketNotificationList{})
}
