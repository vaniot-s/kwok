//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023 The Kubernetes Authors.

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

// Code generated by defaulter-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&KwokConfiguration{}, func(obj interface{}) { SetObjectDefaults_KwokConfiguration(obj.(*KwokConfiguration)) })
	scheme.AddTypeDefaultingFunc(&KwokctlConfiguration{}, func(obj interface{}) { SetObjectDefaults_KwokctlConfiguration(obj.(*KwokctlConfiguration)) })
	scheme.AddTypeDefaultingFunc(&Stage{}, func(obj interface{}) { SetObjectDefaults_Stage(obj.(*Stage)) })
	return nil
}

func SetObjectDefaults_KwokConfiguration(in *KwokConfiguration) {
	if in.Options.CIDR == "" {
		in.Options.CIDR = "10.0.0.1/24"
	}
	if in.Options.ManageAllNodes == nil {
		var ptrVar1 bool = false
		in.Options.ManageAllNodes = &ptrVar1
	}
	if in.Options.EnableCNI == nil {
		var ptrVar1 bool = false
		in.Options.EnableCNI = &ptrVar1
	}
	if in.Options.EnableDebuggingHandlers == nil {
		var ptrVar1 bool = true
		in.Options.EnableDebuggingHandlers = &ptrVar1
	}
	if in.Options.EnableContentionProfiling == nil {
		var ptrVar1 bool = false
		in.Options.EnableContentionProfiling = &ptrVar1
	}
	if in.Options.EnableProfilingHandler == nil {
		var ptrVar1 bool = true
		in.Options.EnableProfilingHandler = &ptrVar1
	}
}

func SetObjectDefaults_KwokctlConfiguration(in *KwokctlConfiguration) {
	if in.Options.SecurePort == nil {
		var ptrVar1 bool = false
		in.Options.SecurePort = &ptrVar1
	}
	if in.Options.QuietPull == nil {
		var ptrVar1 bool = false
		in.Options.QuietPull = &ptrVar1
	}
	if in.Options.DisableKubeScheduler == nil {
		var ptrVar1 bool = false
		in.Options.DisableKubeScheduler = &ptrVar1
	}
	if in.Options.DisableKubeControllerManager == nil {
		var ptrVar1 bool = false
		in.Options.DisableKubeControllerManager = &ptrVar1
	}
	if in.Options.KubeAuthorization == nil {
		var ptrVar1 bool = false
		in.Options.KubeAuthorization = &ptrVar1
	}
	for i := range in.Components {
		a := &in.Components[i]
		for j := range a.Ports {
			b := &a.Ports[j]
			if b.Protocol == "" {
				b.Protocol = "TCP"
			}
		}
		for j := range a.Envs {
			b := &a.Envs[j]
			if b.Value == "" {
				b.Value = ""
			}
		}
	}
}

func SetObjectDefaults_Stage(in *Stage) {
	if in.Spec.ResourceRef.APIGroup == "" {
		in.Spec.ResourceRef.APIGroup = "v1"
	}
	if in.Spec.Weight == 0 {
		in.Spec.Weight = 0
	}
}
