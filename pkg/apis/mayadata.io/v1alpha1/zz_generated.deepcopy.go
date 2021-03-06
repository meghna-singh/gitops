// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitTrack) DeepCopyInto(out *GitTrack) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitTrack.
func (in *GitTrack) DeepCopy() *GitTrack {
	if in == nil {
		return nil
	}
	out := new(GitTrack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GitTrack) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitTrackDeployKey) DeepCopyInto(out *GitTrackDeployKey) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitTrackDeployKey.
func (in *GitTrackDeployKey) DeepCopy() *GitTrackDeployKey {
	if in == nil {
		return nil
	}
	out := new(GitTrackDeployKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitTrackList) DeepCopyInto(out *GitTrackList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GitTrack, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitTrackList.
func (in *GitTrackList) DeepCopy() *GitTrackList {
	if in == nil {
		return nil
	}
	out := new(GitTrackList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GitTrackList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitTrackSpec) DeepCopyInto(out *GitTrackSpec) {
	*out = *in
	out.DeployKey = in.DeployKey
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitTrackSpec.
func (in *GitTrackSpec) DeepCopy() *GitTrackSpec {
	if in == nil {
		return nil
	}
	out := new(GitTrackSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitTrackStatus) DeepCopyInto(out *GitTrackStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitTrackStatus.
func (in *GitTrackStatus) DeepCopy() *GitTrackStatus {
	if in == nil {
		return nil
	}
	out := new(GitTrackStatus)
	in.DeepCopyInto(out)
	return out
}
