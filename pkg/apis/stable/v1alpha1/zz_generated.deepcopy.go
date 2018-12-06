// +build !ignore_autogenerated

// Copyright 2018 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was autogenerated. Do not edit directly.

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1beta1 "k8s.io/api/admissionregistration/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BufferPolicy) DeepCopyInto(out *BufferPolicy) {
	*out = *in
	out.BufferSize = in.BufferSize
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BufferPolicy.
func (in *BufferPolicy) DeepCopy() *BufferPolicy {
	if in == nil {
		return nil
	}
	out := new(BufferPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Fleet) DeepCopyInto(out *Fleet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Fleet.
func (in *Fleet) DeepCopy() *Fleet {
	if in == nil {
		return nil
	}
	out := new(Fleet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Fleet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetAllocation) DeepCopyInto(out *FleetAllocation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetAllocation.
func (in *FleetAllocation) DeepCopy() *FleetAllocation {
	if in == nil {
		return nil
	}
	out := new(FleetAllocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FleetAllocation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetAllocationList) DeepCopyInto(out *FleetAllocationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FleetAllocation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetAllocationList.
func (in *FleetAllocationList) DeepCopy() *FleetAllocationList {
	if in == nil {
		return nil
	}
	out := new(FleetAllocationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FleetAllocationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetAllocationSpec) DeepCopyInto(out *FleetAllocationSpec) {
	*out = *in
	in.MetaPatch.DeepCopyInto(&out.MetaPatch)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetAllocationSpec.
func (in *FleetAllocationSpec) DeepCopy() *FleetAllocationSpec {
	if in == nil {
		return nil
	}
	out := new(FleetAllocationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetAllocationStatus) DeepCopyInto(out *FleetAllocationStatus) {
	*out = *in
	if in.GameServer != nil {
		in, out := &in.GameServer, &out.GameServer
		if *in == nil {
			*out = nil
		} else {
			*out = new(GameServer)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetAllocationStatus.
func (in *FleetAllocationStatus) DeepCopy() *FleetAllocationStatus {
	if in == nil {
		return nil
	}
	out := new(FleetAllocationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetAutoscaleRequest) DeepCopyInto(out *FleetAutoscaleRequest) {
	*out = *in
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetAutoscaleRequest.
func (in *FleetAutoscaleRequest) DeepCopy() *FleetAutoscaleRequest {
	if in == nil {
		return nil
	}
	out := new(FleetAutoscaleRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetAutoscaleResponse) DeepCopyInto(out *FleetAutoscaleResponse) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetAutoscaleResponse.
func (in *FleetAutoscaleResponse) DeepCopy() *FleetAutoscaleResponse {
	if in == nil {
		return nil
	}
	out := new(FleetAutoscaleResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetAutoscaleReview) DeepCopyInto(out *FleetAutoscaleReview) {
	*out = *in
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		if *in == nil {
			*out = nil
		} else {
			*out = new(FleetAutoscaleRequest)
			**out = **in
		}
	}
	if in.Response != nil {
		in, out := &in.Response, &out.Response
		if *in == nil {
			*out = nil
		} else {
			*out = new(FleetAutoscaleResponse)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetAutoscaleReview.
func (in *FleetAutoscaleReview) DeepCopy() *FleetAutoscaleReview {
	if in == nil {
		return nil
	}
	out := new(FleetAutoscaleReview)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetAutoscaler) DeepCopyInto(out *FleetAutoscaler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetAutoscaler.
func (in *FleetAutoscaler) DeepCopy() *FleetAutoscaler {
	if in == nil {
		return nil
	}
	out := new(FleetAutoscaler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FleetAutoscaler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetAutoscalerList) DeepCopyInto(out *FleetAutoscalerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FleetAutoscaler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetAutoscalerList.
func (in *FleetAutoscalerList) DeepCopy() *FleetAutoscalerList {
	if in == nil {
		return nil
	}
	out := new(FleetAutoscalerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FleetAutoscalerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetAutoscalerPolicy) DeepCopyInto(out *FleetAutoscalerPolicy) {
	*out = *in
	if in.Buffer != nil {
		in, out := &in.Buffer, &out.Buffer
		if *in == nil {
			*out = nil
		} else {
			*out = new(BufferPolicy)
			**out = **in
		}
	}
	if in.Webhook != nil {
		in, out := &in.Webhook, &out.Webhook
		if *in == nil {
			*out = nil
		} else {
			*out = new(WebhookPolicy)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetAutoscalerPolicy.
func (in *FleetAutoscalerPolicy) DeepCopy() *FleetAutoscalerPolicy {
	if in == nil {
		return nil
	}
	out := new(FleetAutoscalerPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetAutoscalerSpec) DeepCopyInto(out *FleetAutoscalerSpec) {
	*out = *in
	in.Policy.DeepCopyInto(&out.Policy)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetAutoscalerSpec.
func (in *FleetAutoscalerSpec) DeepCopy() *FleetAutoscalerSpec {
	if in == nil {
		return nil
	}
	out := new(FleetAutoscalerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetAutoscalerStatus) DeepCopyInto(out *FleetAutoscalerStatus) {
	*out = *in
	if in.LastScaleTime != nil {
		in, out := &in.LastScaleTime, &out.LastScaleTime
		if *in == nil {
			*out = nil
		} else {
			*out = (*in).DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetAutoscalerStatus.
func (in *FleetAutoscalerStatus) DeepCopy() *FleetAutoscalerStatus {
	if in == nil {
		return nil
	}
	out := new(FleetAutoscalerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetList) DeepCopyInto(out *FleetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Fleet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetList.
func (in *FleetList) DeepCopy() *FleetList {
	if in == nil {
		return nil
	}
	out := new(FleetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FleetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetSpec) DeepCopyInto(out *FleetSpec) {
	*out = *in
	in.Strategy.DeepCopyInto(&out.Strategy)
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetSpec.
func (in *FleetSpec) DeepCopy() *FleetSpec {
	if in == nil {
		return nil
	}
	out := new(FleetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetStatus) DeepCopyInto(out *FleetStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetStatus.
func (in *FleetStatus) DeepCopy() *FleetStatus {
	if in == nil {
		return nil
	}
	out := new(FleetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameServer) DeepCopyInto(out *GameServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameServer.
func (in *GameServer) DeepCopy() *GameServer {
	if in == nil {
		return nil
	}
	out := new(GameServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GameServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameServerAllocation) DeepCopyInto(out *GameServerAllocation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameServerAllocation.
func (in *GameServerAllocation) DeepCopy() *GameServerAllocation {
	if in == nil {
		return nil
	}
	out := new(GameServerAllocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GameServerAllocation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameServerAllocationList) DeepCopyInto(out *GameServerAllocationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GameServerAllocation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameServerAllocationList.
func (in *GameServerAllocationList) DeepCopy() *GameServerAllocationList {
	if in == nil {
		return nil
	}
	out := new(GameServerAllocationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GameServerAllocationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameServerAllocationSpec) DeepCopyInto(out *GameServerAllocationSpec) {
	*out = *in
	in.Required.DeepCopyInto(&out.Required)
	if in.Preferred != nil {
		in, out := &in.Preferred, &out.Preferred
		*out = make([]v1.LabelSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.MetaPatch.DeepCopyInto(&out.MetaPatch)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameServerAllocationSpec.
func (in *GameServerAllocationSpec) DeepCopy() *GameServerAllocationSpec {
	if in == nil {
		return nil
	}
	out := new(GameServerAllocationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameServerAllocationStatus) DeepCopyInto(out *GameServerAllocationStatus) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]GameServerStatusPort, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameServerAllocationStatus.
func (in *GameServerAllocationStatus) DeepCopy() *GameServerAllocationStatus {
	if in == nil {
		return nil
	}
	out := new(GameServerAllocationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameServerList) DeepCopyInto(out *GameServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GameServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameServerList.
func (in *GameServerList) DeepCopy() *GameServerList {
	if in == nil {
		return nil
	}
	out := new(GameServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GameServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameServerPort) DeepCopyInto(out *GameServerPort) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameServerPort.
func (in *GameServerPort) DeepCopy() *GameServerPort {
	if in == nil {
		return nil
	}
	out := new(GameServerPort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameServerSet) DeepCopyInto(out *GameServerSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameServerSet.
func (in *GameServerSet) DeepCopy() *GameServerSet {
	if in == nil {
		return nil
	}
	out := new(GameServerSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GameServerSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameServerSetList) DeepCopyInto(out *GameServerSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GameServerSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameServerSetList.
func (in *GameServerSetList) DeepCopy() *GameServerSetList {
	if in == nil {
		return nil
	}
	out := new(GameServerSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GameServerSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameServerSetSpec) DeepCopyInto(out *GameServerSetSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameServerSetSpec.
func (in *GameServerSetSpec) DeepCopy() *GameServerSetSpec {
	if in == nil {
		return nil
	}
	out := new(GameServerSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameServerSetStatus) DeepCopyInto(out *GameServerSetStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameServerSetStatus.
func (in *GameServerSetStatus) DeepCopy() *GameServerSetStatus {
	if in == nil {
		return nil
	}
	out := new(GameServerSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameServerSpec) DeepCopyInto(out *GameServerSpec) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]GameServerPort, len(*in))
		copy(*out, *in)
	}
	out.Health = in.Health
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameServerSpec.
func (in *GameServerSpec) DeepCopy() *GameServerSpec {
	if in == nil {
		return nil
	}
	out := new(GameServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameServerStatus) DeepCopyInto(out *GameServerStatus) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]GameServerStatusPort, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameServerStatus.
func (in *GameServerStatus) DeepCopy() *GameServerStatus {
	if in == nil {
		return nil
	}
	out := new(GameServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameServerStatusPort) DeepCopyInto(out *GameServerStatusPort) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameServerStatusPort.
func (in *GameServerStatusPort) DeepCopy() *GameServerStatusPort {
	if in == nil {
		return nil
	}
	out := new(GameServerStatusPort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GameServerTemplateSpec) DeepCopyInto(out *GameServerTemplateSpec) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GameServerTemplateSpec.
func (in *GameServerTemplateSpec) DeepCopy() *GameServerTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(GameServerTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Health) DeepCopyInto(out *Health) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Health.
func (in *Health) DeepCopy() *Health {
	if in == nil {
		return nil
	}
	out := new(Health)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetaPatch) DeepCopyInto(out *MetaPatch) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetaPatch.
func (in *MetaPatch) DeepCopy() *MetaPatch {
	if in == nil {
		return nil
	}
	out := new(MetaPatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookPolicy) DeepCopyInto(out *WebhookPolicy) {
	*out = *in
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1beta1.ServiceReference)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.CABundle != nil {
		in, out := &in.CABundle, &out.CABundle
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookPolicy.
func (in *WebhookPolicy) DeepCopy() *WebhookPolicy {
	if in == nil {
		return nil
	}
	out := new(WebhookPolicy)
	in.DeepCopyInto(out)
	return out
}
