// Copyright Envoy Gateway Authors
// SPDX-License-Identifier: Apache-2.0
// The full text of the Apache license is available in the LICENSE file at
// the root of the repo.

package v1alpha1

import (
	gwapiv1a2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
)

type PolicyTargetReferences struct {
	// TargetRef is the name of the resource this policy is being attached to.
	// This policy and the TargetRef MUST be in the same namespace for this
	// Policy to have effect
	//
	// Deprecated: use targetRefs/targetSelectors instead
	TargetRef *gwapiv1a2.LocalPolicyTargetReferenceWithSectionName `json:"targetRef,omitempty"`

	// TargetRefs are the names of the Gateway resources this policy
	// is being attached to.
	// TargetRefs是这个policy关联的Gateway resources的名字
	TargetRefs []gwapiv1a2.LocalPolicyTargetReferenceWithSectionName `json:"targetRefs,omitempty"`

	// TargetSelectors allow targeting resources for this policy based on labels
	// TargetSelectors允许这个policy基于labels选择resources
	TargetSelectors []TargetSelector `json:"targetSelectors,omitempty"`
}

// +kubebuilder:validation:XValidation:rule="has(self.group) ? self.group == 'gateway.networking.k8s.io' : true ", message="group must be gateway.networking.k8s.io"
type TargetSelector struct {
	// Group is the group that this selector targets. Defaults to gateway.networking.k8s.io
	// Group是这个selector选择的group
	//
	// +kubebuilder:default:="gateway.networking.k8s.io"
	Group *gwapiv1a2.Group `json:"group,omitempty"`

	// Kind is the resource kind that this selector targets.
	// Kind是这个selector选择的目标resource kind
	Kind gwapiv1a2.Kind `json:"kind"`

	// MatchLabels are the set of label selectors for identifying the targeted resource
	// MatchLabels是一系列的label selectors，用于标识targeted resource
	MatchLabels map[string]string `json:"matchLabels"`
}

func (p PolicyTargetReferences) GetTargetRefs() []gwapiv1a2.LocalPolicyTargetReferenceWithSectionName {
	if p.TargetRef != nil {
		return []gwapiv1a2.LocalPolicyTargetReferenceWithSectionName{*p.TargetRef}
	}
	return p.TargetRefs
}
