// +build !ignore_autogenerated

/*
Copyright 2019 (c) Microsoft and contributors. All rights reserved.

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

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonSpec) DeepCopyInto(out *AddonSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonSpec.
func (in *AddonSpec) DeepCopy() *AddonSpec {
	if in == nil {
		return nil
	}
	out := new(AddonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentitySpec) DeepCopyInto(out *IdentitySpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentitySpec.
func (in *IdentitySpec) DeepCopy() *IdentitySpec {
	if in == nil {
		return nil
	}
	out := new(IdentitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinuxSpec) DeepCopyInto(out *LinuxSpec) {
	*out = *in
	if in.AuthorizedKeys != nil {
		in, out := &in.AuthorizedKeys, &out.AuthorizedKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinuxSpec.
func (in *LinuxSpec) DeepCopy() *LinuxSpec {
	if in == nil {
		return nil
	}
	out := new(LinuxSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedCluster) DeepCopyInto(out *ManagedCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedCluster.
func (in *ManagedCluster) DeepCopy() *ManagedCluster {
	if in == nil {
		return nil
	}
	out := new(ManagedCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedClusterList) DeepCopyInto(out *ManagedClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagedCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedClusterList.
func (in *ManagedClusterList) DeepCopy() *ManagedClusterList {
	if in == nil {
		return nil
	}
	out := new(ManagedClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedClusterSpec) DeepCopyInto(out *ManagedClusterSpec) {
	*out = *in
	if in.Addons != nil {
		in, out := &in.Addons, &out.Addons
		*out = make(map[string]AddonSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	out.Network = in.Network
	if in.NodePools != nil {
		in, out := &in.NodePools, &out.NodePools
		*out = make([]NodePoolSpec, len(*in))
		copy(*out, *in)
	}
	in.Linux.DeepCopyInto(&out.Linux)
	out.Identity = in.Identity
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedClusterSpec.
func (in *ManagedClusterSpec) DeepCopy() *ManagedClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ManagedClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedClusterStatus) DeepCopyInto(out *ManagedClusterStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedClusterStatus.
func (in *ManagedClusterStatus) DeepCopy() *ManagedClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ManagedClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkSpec) DeepCopyInto(out *NetworkSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkSpec.
func (in *NetworkSpec) DeepCopy() *NetworkSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolSpec) DeepCopyInto(out *NodePoolSpec) {
	*out = *in
	out.NodeSpec = in.NodeSpec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolSpec.
func (in *NodePoolSpec) DeepCopy() *NodePoolSpec {
	if in == nil {
		return nil
	}
	out := new(NodePoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeSpec) DeepCopyInto(out *NodeSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeSpec.
func (in *NodeSpec) DeepCopy() *NodeSpec {
	if in == nil {
		return nil
	}
	out := new(NodeSpec)
	in.DeepCopyInto(out)
	return out
}