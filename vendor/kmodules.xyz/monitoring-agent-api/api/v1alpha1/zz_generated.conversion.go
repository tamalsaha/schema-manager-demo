//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "kmodules.xyz/monitoring-agent-api/api/v1"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v1.AgentSpec)(nil), (*AgentSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_AgentSpec_To_v1alpha1_AgentSpec(a.(*v1.AgentSpec), b.(*AgentSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1.PrometheusSpec)(nil), (*PrometheusSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PrometheusSpec_To_v1alpha1_PrometheusSpec(a.(*v1.PrometheusSpec), b.(*PrometheusSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*AgentSpec)(nil), (*v1.AgentSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_AgentSpec_To_v1_AgentSpec(a.(*AgentSpec), b.(*v1.AgentSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*PrometheusSpec)(nil), (*v1.PrometheusSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PrometheusSpec_To_v1_PrometheusSpec(a.(*PrometheusSpec), b.(*v1.PrometheusSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_AgentSpec_To_v1_AgentSpec(in *AgentSpec, out *v1.AgentSpec, s conversion.Scope) error {
	out.Agent = v1.AgentType(in.Agent)
	if in.Prometheus != nil {
		in, out := &in.Prometheus, &out.Prometheus
		*out = new(v1.PrometheusSpec)
		if err := Convert_v1alpha1_PrometheusSpec_To_v1_PrometheusSpec(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Prometheus = nil
	}
	// WARNING: in.Args requires manual conversion: does not exist in peer-type
	// WARNING: in.Env requires manual conversion: does not exist in peer-type
	// WARNING: in.Resources requires manual conversion: does not exist in peer-type
	// WARNING: in.SecurityContext requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1_AgentSpec_To_v1alpha1_AgentSpec(in *v1.AgentSpec, out *AgentSpec, s conversion.Scope) error {
	out.Agent = AgentType(in.Agent)
	if in.Prometheus != nil {
		in, out := &in.Prometheus, &out.Prometheus
		*out = new(PrometheusSpec)
		if err := Convert_v1_PrometheusSpec_To_v1alpha1_PrometheusSpec(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Prometheus = nil
	}
	return nil
}

// Convert_v1_AgentSpec_To_v1alpha1_AgentSpec is an autogenerated conversion function.
func Convert_v1_AgentSpec_To_v1alpha1_AgentSpec(in *v1.AgentSpec, out *AgentSpec, s conversion.Scope) error {
	return autoConvert_v1_AgentSpec_To_v1alpha1_AgentSpec(in, out, s)
}

func autoConvert_v1alpha1_PrometheusSpec_To_v1_PrometheusSpec(in *PrometheusSpec, out *v1.PrometheusSpec, s conversion.Scope) error {
	// WARNING: in.Port requires manual conversion: does not exist in peer-type
	// WARNING: in.Namespace requires manual conversion: does not exist in peer-type
	// WARNING: in.Labels requires manual conversion: does not exist in peer-type
	// WARNING: in.Interval requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1_PrometheusSpec_To_v1alpha1_PrometheusSpec(in *v1.PrometheusSpec, out *PrometheusSpec, s conversion.Scope) error {
	// WARNING: in.Exporter requires manual conversion: does not exist in peer-type
	// WARNING: in.ServiceMonitor requires manual conversion: does not exist in peer-type
	return nil
}
