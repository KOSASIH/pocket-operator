//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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

package v1alpha1

import (
	"github.com/nukleros/operator-builder-tools/pkg/status"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PocketSet) DeepCopyInto(out *PocketSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PocketSet.
func (in *PocketSet) DeepCopy() *PocketSet {
	if in == nil {
		return nil
	}
	out := new(PocketSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PocketSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PocketSetList) DeepCopyInto(out *PocketSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PocketSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PocketSetList.
func (in *PocketSetList) DeepCopy() *PocketSetList {
	if in == nil {
		return nil
	}
	out := new(PocketSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PocketSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PocketSetSpec) DeepCopyInto(out *PocketSetSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PocketSetSpec.
func (in *PocketSetSpec) DeepCopy() *PocketSetSpec {
	if in == nil {
		return nil
	}
	out := new(PocketSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PocketSetStatus) DeepCopyInto(out *PocketSetStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*status.PhaseCondition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(status.PhaseCondition)
				**out = **in
			}
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]*status.ChildResource, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(status.ChildResource)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PocketSetStatus.
func (in *PocketSetStatus) DeepCopy() *PocketSetStatus {
	if in == nil {
		return nil
	}
	out := new(PocketSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PocketValidator) DeepCopyInto(out *PocketValidator) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PocketValidator.
func (in *PocketValidator) DeepCopy() *PocketValidator {
	if in == nil {
		return nil
	}
	out := new(PocketValidator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PocketValidator) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PocketValidatorCollectionSpec) DeepCopyInto(out *PocketValidatorCollectionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PocketValidatorCollectionSpec.
func (in *PocketValidatorCollectionSpec) DeepCopy() *PocketValidatorCollectionSpec {
	if in == nil {
		return nil
	}
	out := new(PocketValidatorCollectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PocketValidatorList) DeepCopyInto(out *PocketValidatorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PocketValidator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PocketValidatorList.
func (in *PocketValidatorList) DeepCopy() *PocketValidatorList {
	if in == nil {
		return nil
	}
	out := new(PocketValidatorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PocketValidatorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PocketValidatorSpec) DeepCopyInto(out *PocketValidatorSpec) {
	*out = *in
	out.Collection = in.Collection
	out.Ports = in.Ports
	out.PrivateKey = in.PrivateKey
	out.Postgres = in.Postgres
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PocketValidatorSpec.
func (in *PocketValidatorSpec) DeepCopy() *PocketValidatorSpec {
	if in == nil {
		return nil
	}
	out := new(PocketValidatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PocketValidatorSpecPorts) DeepCopyInto(out *PocketValidatorSpecPorts) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PocketValidatorSpecPorts.
func (in *PocketValidatorSpecPorts) DeepCopy() *PocketValidatorSpecPorts {
	if in == nil {
		return nil
	}
	out := new(PocketValidatorSpecPorts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PocketValidatorSpecPostgres) DeepCopyInto(out *PocketValidatorSpecPostgres) {
	*out = *in
	out.User = in.User
	out.Password = in.Password
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PocketValidatorSpecPostgres.
func (in *PocketValidatorSpecPostgres) DeepCopy() *PocketValidatorSpecPostgres {
	if in == nil {
		return nil
	}
	out := new(PocketValidatorSpecPostgres)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PocketValidatorSpecPostgresPassword) DeepCopyInto(out *PocketValidatorSpecPostgresPassword) {
	*out = *in
	out.SecretKeyRef = in.SecretKeyRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PocketValidatorSpecPostgresPassword.
func (in *PocketValidatorSpecPostgresPassword) DeepCopy() *PocketValidatorSpecPostgresPassword {
	if in == nil {
		return nil
	}
	out := new(PocketValidatorSpecPostgresPassword)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PocketValidatorSpecPostgresPasswordSecretKeyRef) DeepCopyInto(out *PocketValidatorSpecPostgresPasswordSecretKeyRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PocketValidatorSpecPostgresPasswordSecretKeyRef.
func (in *PocketValidatorSpecPostgresPasswordSecretKeyRef) DeepCopy() *PocketValidatorSpecPostgresPasswordSecretKeyRef {
	if in == nil {
		return nil
	}
	out := new(PocketValidatorSpecPostgresPasswordSecretKeyRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PocketValidatorSpecPostgresUser) DeepCopyInto(out *PocketValidatorSpecPostgresUser) {
	*out = *in
	out.SecretKeyRef = in.SecretKeyRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PocketValidatorSpecPostgresUser.
func (in *PocketValidatorSpecPostgresUser) DeepCopy() *PocketValidatorSpecPostgresUser {
	if in == nil {
		return nil
	}
	out := new(PocketValidatorSpecPostgresUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PocketValidatorSpecPostgresUserSecretKeyRef) DeepCopyInto(out *PocketValidatorSpecPostgresUserSecretKeyRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PocketValidatorSpecPostgresUserSecretKeyRef.
func (in *PocketValidatorSpecPostgresUserSecretKeyRef) DeepCopy() *PocketValidatorSpecPostgresUserSecretKeyRef {
	if in == nil {
		return nil
	}
	out := new(PocketValidatorSpecPostgresUserSecretKeyRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PocketValidatorSpecPrivateKey) DeepCopyInto(out *PocketValidatorSpecPrivateKey) {
	*out = *in
	out.SecretKeyRef = in.SecretKeyRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PocketValidatorSpecPrivateKey.
func (in *PocketValidatorSpecPrivateKey) DeepCopy() *PocketValidatorSpecPrivateKey {
	if in == nil {
		return nil
	}
	out := new(PocketValidatorSpecPrivateKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PocketValidatorSpecPrivateKeySecretKeyRef) DeepCopyInto(out *PocketValidatorSpecPrivateKeySecretKeyRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PocketValidatorSpecPrivateKeySecretKeyRef.
func (in *PocketValidatorSpecPrivateKeySecretKeyRef) DeepCopy() *PocketValidatorSpecPrivateKeySecretKeyRef {
	if in == nil {
		return nil
	}
	out := new(PocketValidatorSpecPrivateKeySecretKeyRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PocketValidatorStatus) DeepCopyInto(out *PocketValidatorStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*status.PhaseCondition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(status.PhaseCondition)
				**out = **in
			}
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]*status.ChildResource, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(status.ChildResource)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PocketValidatorStatus.
func (in *PocketValidatorStatus) DeepCopy() *PocketValidatorStatus {
	if in == nil {
		return nil
	}
	out := new(PocketValidatorStatus)
	in.DeepCopyInto(out)
	return out
}
