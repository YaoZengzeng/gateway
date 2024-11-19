// Copyright Envoy Gateway Authors
// SPDX-License-Identifier: Apache-2.0
// The full text of the Apache license is available in the LICENSE file at
// the root of the repo.

package v1alpha1

import (
	gwapiv1 "sigs.k8s.io/gateway-api/apis/v1"
)

// +kubebuilder:validation:Enum=Streamed;Buffered;BufferedPartial
type ExtProcBodyProcessingMode string

const (
	// StreamedExtProcBodyProcessingMode will stream the body to the server in pieces as they arrive at the proxy.
	// StreamedExtProcBodyProcessingMode会将stream流式发送给server，按照他们到达proxy的分片格式
	StreamedExtProcBodyProcessingMode ExtProcBodyProcessingMode = "Streamed"
	// BufferedExtProcBodyProcessingMode will buffer the message body in memory and send the entire body at once. If the body exceeds the configured buffer limit, then the downstream system will receive an error.
	// BufferedExtProcBodyProcessingMode会缓存message body在内存中并且一整个发送body，如果body超过了配置的buffer limit，那么downstream system会接收到一个错误
	BufferedExtProcBodyProcessingMode ExtProcBodyProcessingMode = "Buffered"
	// BufferedPartialExtBodyHeaderProcessingMode will buffer the message body in memory and send the entire body in one chunk. If the body exceeds the configured buffer limit, then the body contents up to the buffer limit will be sent.
	// BufferedPartialExtBodyHeaderProcessingMode会缓存message body在内存中并且在一个chunk中发送整个body，如果body超过了配置的buffer limit，那么整个body，直到buffer limit会被发送
	BufferedPartialExtBodyHeaderProcessingMode ExtProcBodyProcessingMode = "BufferedPartial"
)

// ProcessingModeOptions defines if headers or body should be processed by the external service
// ProcessingModeOptions定义了是否headers或者body应该被外部的service处理
type ProcessingModeOptions struct {
	// Defines body processing mode
	// 定义body processing模式
	//
	// +optional
	Body *ExtProcBodyProcessingMode `json:"body,omitempty"`
}

// ExtProcProcessingMode defines if and how headers and bodies are sent to the service.
// ExtProcProcessingMode定义了是否以及如何发送headers和bodies到service
// https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/http/ext_proc/v3/processing_mode.proto#envoy-v3-api-msg-extensions-filters-http-ext-proc-v3-processingmode
type ExtProcProcessingMode struct {
	// Defines processing mode for requests. If present, request headers are sent. Request body is processed according
	// to the specified mode.
	// 定义请求的processing mode，如果存在的话，request headers被发送，Request body被处理，根据指定的mode
	//
	// +optional
	Request *ProcessingModeOptions `json:"request,omitempty"`

	// Defines processing mode for responses. If present, response headers are sent. Response body is processed according
	// to the specified mode.
	// 定义response的处理模式，如果存在的话，response headers被发送，Response body根据指定的模式被处理
	//
	// +optional
	Response *ProcessingModeOptions `json:"response,omitempty"`
}

// ExtProc defines the configuration for External Processing filter.
// ExtProc定义了External Processing filter的配置
// +kubebuilder:validation:XValidation:message="BackendRefs must be used, backendRef is not supported.",rule="!has(self.backendRef)"
// +kubebuilder:validation:XValidation:message="BackendRefs only supports Service and Backend kind.",rule="has(self.backendRefs) ? self.backendRefs.all(f, f.kind == 'Service' || f.kind == 'Backend') : true"
// +kubebuilder:validation:XValidation:message="BackendRefs only supports Core and gateway.envoyproxy.io group.",rule="has(self.backendRefs) ? (self.backendRefs.all(f, f.group == \"\" || f.group == 'gateway.envoyproxy.io')) : true"
type ExtProc struct {
	BackendCluster `json:",inline"`

	// MessageTimeout is the timeout for a response to be returned from the external processor
	// MessageTimeout是默认的从external processor返回一个response的超时时间
	// Default: 200ms
	//
	// +optional
	MessageTimeout *gwapiv1.Duration `json:"messageTimeout,omitempty"`

	// FailOpen defines if requests or responses that cannot be processed due to connectivity to the
	// external processor are terminated or passed-through.
	// FailOpen定义了requests或者responses不能被处理，因为发往external processor的连接被终止或者passed-through
	// Default: false
	//
	// +optional
	FailOpen *bool `json:"failOpen,omitempty"`

	// ProcessingMode defines how request and response body is processed
	// Default: header and body are not sent to the external processor
	// ProcessingMode定义了request和response body怎样被处理，默认header和body不会发往external processor
	//
	// +optional
	ProcessingMode *ExtProcProcessingMode `json:"processingMode,omitempty"`
}
