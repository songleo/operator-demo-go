// +build !ignore_autogenerated

/*
Copyright 2021.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterObservability) DeepCopyInto(out *MultiClusterObservability) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterObservability.
func (in *MultiClusterObservability) DeepCopy() *MultiClusterObservability {
	if in == nil {
		return nil
	}
	out := new(MultiClusterObservability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiClusterObservability) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterObservabilityList) DeepCopyInto(out *MultiClusterObservabilityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MultiClusterObservability, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterObservabilityList.
func (in *MultiClusterObservabilityList) DeepCopy() *MultiClusterObservabilityList {
	if in == nil {
		return nil
	}
	out := new(MultiClusterObservabilityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiClusterObservabilityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterObservabilitySpec) DeepCopyInto(out *MultiClusterObservabilitySpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterObservabilitySpec.
func (in *MultiClusterObservabilitySpec) DeepCopy() *MultiClusterObservabilitySpec {
	if in == nil {
		return nil
	}
	out := new(MultiClusterObservabilitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiClusterObservabilityStatus) DeepCopyInto(out *MultiClusterObservabilityStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiClusterObservabilityStatus.
func (in *MultiClusterObservabilityStatus) DeepCopy() *MultiClusterObservabilityStatus {
	if in == nil {
		return nil
	}
	out := new(MultiClusterObservabilityStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservabilityAddon) DeepCopyInto(out *ObservabilityAddon) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservabilityAddon.
func (in *ObservabilityAddon) DeepCopy() *ObservabilityAddon {
	if in == nil {
		return nil
	}
	out := new(ObservabilityAddon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObservabilityAddon) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservabilityAddonList) DeepCopyInto(out *ObservabilityAddonList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ObservabilityAddon, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservabilityAddonList.
func (in *ObservabilityAddonList) DeepCopy() *ObservabilityAddonList {
	if in == nil {
		return nil
	}
	out := new(ObservabilityAddonList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObservabilityAddonList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservabilityAddonSpec) DeepCopyInto(out *ObservabilityAddonSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservabilityAddonSpec.
func (in *ObservabilityAddonSpec) DeepCopy() *ObservabilityAddonSpec {
	if in == nil {
		return nil
	}
	out := new(ObservabilityAddonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservabilityAddonStatus) DeepCopyInto(out *ObservabilityAddonStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservabilityAddonStatus.
func (in *ObservabilityAddonStatus) DeepCopy() *ObservabilityAddonStatus {
	if in == nil {
		return nil
	}
	out := new(ObservabilityAddonStatus)
	in.DeepCopyInto(out)
	return out
}
