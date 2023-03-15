//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"

	corev1alpha1 "github.com/kcp-dev/kcp/sdk/apis/core/v1alpha1"
	conditionsv1alpha1 "github.com/kcp-dev/kcp/sdk/apis/third_party/conditions/apis/conditions/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIExportReference) DeepCopyInto(out *APIExportReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIExportReference.
func (in *APIExportReference) DeepCopy() *APIExportReference {
	if in == nil {
		return nil
	}
	out := new(APIExportReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualWorkspace) DeepCopyInto(out *VirtualWorkspace) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualWorkspace.
func (in *VirtualWorkspace) DeepCopy() *VirtualWorkspace {
	if in == nil {
		return nil
	}
	out := new(VirtualWorkspace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workspace) DeepCopyInto(out *Workspace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workspace.
func (in *Workspace) DeepCopy() *Workspace {
	if in == nil {
		return nil
	}
	out := new(Workspace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Workspace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceList) DeepCopyInto(out *WorkspaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Workspace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceList.
func (in *WorkspaceList) DeepCopy() *WorkspaceList {
	if in == nil {
		return nil
	}
	out := new(WorkspaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkspaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceLocation) DeepCopyInto(out *WorkspaceLocation) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceLocation.
func (in *WorkspaceLocation) DeepCopy() *WorkspaceLocation {
	if in == nil {
		return nil
	}
	out := new(WorkspaceLocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceSpec) DeepCopyInto(out *WorkspaceSpec) {
	*out = *in
	out.Type = in.Type
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(WorkspaceLocation)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceSpec.
func (in *WorkspaceSpec) DeepCopy() *WorkspaceSpec {
	if in == nil {
		return nil
	}
	out := new(WorkspaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceStatus) DeepCopyInto(out *WorkspaceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(conditionsv1alpha1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Initializers != nil {
		in, out := &in.Initializers, &out.Initializers
		*out = make([]corev1alpha1.LogicalClusterInitializer, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceStatus.
func (in *WorkspaceStatus) DeepCopy() *WorkspaceStatus {
	if in == nil {
		return nil
	}
	out := new(WorkspaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceType) DeepCopyInto(out *WorkspaceType) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceType.
func (in *WorkspaceType) DeepCopy() *WorkspaceType {
	if in == nil {
		return nil
	}
	out := new(WorkspaceType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkspaceType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceTypeExtension) DeepCopyInto(out *WorkspaceTypeExtension) {
	*out = *in
	if in.With != nil {
		in, out := &in.With, &out.With
		*out = make([]WorkspaceTypeReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceTypeExtension.
func (in *WorkspaceTypeExtension) DeepCopy() *WorkspaceTypeExtension {
	if in == nil {
		return nil
	}
	out := new(WorkspaceTypeExtension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceTypeList) DeepCopyInto(out *WorkspaceTypeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkspaceType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceTypeList.
func (in *WorkspaceTypeList) DeepCopy() *WorkspaceTypeList {
	if in == nil {
		return nil
	}
	out := new(WorkspaceTypeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkspaceTypeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceTypeReference) DeepCopyInto(out *WorkspaceTypeReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceTypeReference.
func (in *WorkspaceTypeReference) DeepCopy() *WorkspaceTypeReference {
	if in == nil {
		return nil
	}
	out := new(WorkspaceTypeReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceTypeSelector) DeepCopyInto(out *WorkspaceTypeSelector) {
	*out = *in
	if in.Types != nil {
		in, out := &in.Types, &out.Types
		*out = make([]WorkspaceTypeReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceTypeSelector.
func (in *WorkspaceTypeSelector) DeepCopy() *WorkspaceTypeSelector {
	if in == nil {
		return nil
	}
	out := new(WorkspaceTypeSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceTypeSpec) DeepCopyInto(out *WorkspaceTypeSpec) {
	*out = *in
	in.Extend.DeepCopyInto(&out.Extend)
	if in.AdditionalWorkspaceLabels != nil {
		in, out := &in.AdditionalWorkspaceLabels, &out.AdditionalWorkspaceLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DefaultChildWorkspaceType != nil {
		in, out := &in.DefaultChildWorkspaceType, &out.DefaultChildWorkspaceType
		*out = new(WorkspaceTypeReference)
		**out = **in
	}
	if in.LimitAllowedChildren != nil {
		in, out := &in.LimitAllowedChildren, &out.LimitAllowedChildren
		*out = new(WorkspaceTypeSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.LimitAllowedParents != nil {
		in, out := &in.LimitAllowedParents, &out.LimitAllowedParents
		*out = new(WorkspaceTypeSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.DefaultAPIBindings != nil {
		in, out := &in.DefaultAPIBindings, &out.DefaultAPIBindings
		*out = make([]APIExportReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceTypeSpec.
func (in *WorkspaceTypeSpec) DeepCopy() *WorkspaceTypeSpec {
	if in == nil {
		return nil
	}
	out := new(WorkspaceTypeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceTypeStatus) DeepCopyInto(out *WorkspaceTypeStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(conditionsv1alpha1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VirtualWorkspaces != nil {
		in, out := &in.VirtualWorkspaces, &out.VirtualWorkspaces
		*out = make([]VirtualWorkspace, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceTypeStatus.
func (in *WorkspaceTypeStatus) DeepCopy() *WorkspaceTypeStatus {
	if in == nil {
		return nil
	}
	out := new(WorkspaceTypeStatus)
	in.DeepCopyInto(out)
	return out
}
