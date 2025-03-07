// +build !ignore_autogenerated

// Copyright Contributors to the Open Cluster Management project

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/openshift/hive/apis/hive/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrintClusterClaim) DeepCopyInto(out *PrintClusterClaim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrintClusterClaim.
func (in *PrintClusterClaim) DeepCopy() *PrintClusterClaim {
	if in == nil {
		return nil
	}
	out := new(PrintClusterClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrintClusterClaim) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrintClusterClaimCredential) DeepCopyInto(out *PrintClusterClaimCredential) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrintClusterClaimCredential.
func (in *PrintClusterClaimCredential) DeepCopy() *PrintClusterClaimCredential {
	if in == nil {
		return nil
	}
	out := new(PrintClusterClaimCredential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrintClusterClaimCredential) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrintClusterClaimCredentialList) DeepCopyInto(out *PrintClusterClaimCredentialList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PrintClusterClaimCredential, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrintClusterClaimCredentialList.
func (in *PrintClusterClaimCredentialList) DeepCopy() *PrintClusterClaimCredentialList {
	if in == nil {
		return nil
	}
	out := new(PrintClusterClaimCredentialList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrintClusterClaimCredentialList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrintClusterClaimCredentialSpec) DeepCopyInto(out *PrintClusterClaimCredentialSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrintClusterClaimCredentialSpec.
func (in *PrintClusterClaimCredentialSpec) DeepCopy() *PrintClusterClaimCredentialSpec {
	if in == nil {
		return nil
	}
	out := new(PrintClusterClaimCredentialSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrintClusterClaimList) DeepCopyInto(out *PrintClusterClaimList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PrintClusterClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrintClusterClaimList.
func (in *PrintClusterClaimList) DeepCopy() *PrintClusterClaimList {
	if in == nil {
		return nil
	}
	out := new(PrintClusterClaimList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrintClusterClaimList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrintClusterClaimSpec) DeepCopyInto(out *PrintClusterClaimSpec) {
	*out = *in
	if in.ClusterClaim != nil {
		in, out := &in.ClusterClaim, &out.ClusterClaim
		*out = new(v1.ClusterClaim)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrintClusterClaimSpec.
func (in *PrintClusterClaimSpec) DeepCopy() *PrintClusterClaimSpec {
	if in == nil {
		return nil
	}
	out := new(PrintClusterClaimSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrintClusterPool) DeepCopyInto(out *PrintClusterPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrintClusterPool.
func (in *PrintClusterPool) DeepCopy() *PrintClusterPool {
	if in == nil {
		return nil
	}
	out := new(PrintClusterPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrintClusterPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrintClusterPoolHost) DeepCopyInto(out *PrintClusterPoolHost) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrintClusterPoolHost.
func (in *PrintClusterPoolHost) DeepCopy() *PrintClusterPoolHost {
	if in == nil {
		return nil
	}
	out := new(PrintClusterPoolHost)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrintClusterPoolHost) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrintClusterPoolHostList) DeepCopyInto(out *PrintClusterPoolHostList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PrintClusterPoolHost, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrintClusterPoolHostList.
func (in *PrintClusterPoolHostList) DeepCopy() *PrintClusterPoolHostList {
	if in == nil {
		return nil
	}
	out := new(PrintClusterPoolHostList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrintClusterPoolHostList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrintClusterPoolHostSpec) DeepCopyInto(out *PrintClusterPoolHostSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrintClusterPoolHostSpec.
func (in *PrintClusterPoolHostSpec) DeepCopy() *PrintClusterPoolHostSpec {
	if in == nil {
		return nil
	}
	out := new(PrintClusterPoolHostSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrintClusterPoolList) DeepCopyInto(out *PrintClusterPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PrintClusterPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrintClusterPoolList.
func (in *PrintClusterPoolList) DeepCopy() *PrintClusterPoolList {
	if in == nil {
		return nil
	}
	out := new(PrintClusterPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrintClusterPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrintClusterPoolSpec) DeepCopyInto(out *PrintClusterPoolSpec) {
	*out = *in
	if in.ClusterPool != nil {
		in, out := &in.ClusterPool, &out.ClusterPool
		*out = new(v1.ClusterPool)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrintClusterPoolSpec.
func (in *PrintClusterPoolSpec) DeepCopy() *PrintClusterPoolSpec {
	if in == nil {
		return nil
	}
	out := new(PrintClusterPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrintPolicies) DeepCopyInto(out *PrintPolicies) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrintPolicies.
func (in *PrintPolicies) DeepCopy() *PrintPolicies {
	if in == nil {
		return nil
	}
	out := new(PrintPolicies)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrintPolicies) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrintPoliciesList) DeepCopyInto(out *PrintPoliciesList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PrintPolicies, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrintPoliciesList.
func (in *PrintPoliciesList) DeepCopy() *PrintPoliciesList {
	if in == nil {
		return nil
	}
	out := new(PrintPoliciesList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrintPoliciesList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrintPoliciesSpec) DeepCopyInto(out *PrintPoliciesSpec) {
	*out = *in
	in.Policy.DeepCopyInto(&out.Policy)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrintPoliciesSpec.
func (in *PrintPoliciesSpec) DeepCopy() *PrintPoliciesSpec {
	if in == nil {
		return nil
	}
	out := new(PrintPoliciesSpec)
	in.DeepCopyInto(out)
	return out
}
