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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSAccessRequestConfiguration) DeepCopyInto(out *AWSAccessRequestConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSAccessRequestConfiguration.
func (in *AWSAccessRequestConfiguration) DeepCopy() *AWSAccessRequestConfiguration {
	if in == nil {
		return nil
	}
	out := new(AWSAccessRequestConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSConfiguration) DeepCopyInto(out *AWSConfiguration) {
	*out = *in
	if in.MaxRetries != nil {
		in, out := &in.MaxRetries, &out.MaxRetries
		*out = new(int64)
		**out = **in
	}
	if in.LeaseConfig != nil {
		in, out := &in.LeaseConfig, &out.LeaseConfig
		*out = new(LeaseConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSConfiguration.
func (in *AWSConfiguration) DeepCopy() *AWSConfiguration {
	if in == nil {
		return nil
	}
	out := new(AWSConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSRole) DeepCopyInto(out *AWSRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSRole.
func (in *AWSRole) DeepCopy() *AWSRole {
	if in == nil {
		return nil
	}
	out := new(AWSRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSRoleList) DeepCopyInto(out *AWSRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSRoleList.
func (in *AWSRoleList) DeepCopy() *AWSRoleList {
	if in == nil {
		return nil
	}
	out := new(AWSRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSRoleSpec) DeepCopyInto(out *AWSRoleSpec) {
	*out = *in
	out.SecretEngineRef = in.SecretEngineRef
	if in.RoleARNs != nil {
		in, out := &in.RoleARNs, &out.RoleARNs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PolicyARNs != nil {
		in, out := &in.PolicyARNs, &out.PolicyARNs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSRoleSpec.
func (in *AWSRoleSpec) DeepCopy() *AWSRoleSpec {
	if in == nil {
		return nil
	}
	out := new(AWSRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureConfiguration) DeepCopyInto(out *AzureConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureConfiguration.
func (in *AzureConfiguration) DeepCopy() *AzureConfiguration {
	if in == nil {
		return nil
	}
	out := new(AzureConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureRole) DeepCopyInto(out *AzureRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureRole.
func (in *AzureRole) DeepCopy() *AzureRole {
	if in == nil {
		return nil
	}
	out := new(AzureRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureRoleList) DeepCopyInto(out *AzureRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AzureRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureRoleList.
func (in *AzureRoleList) DeepCopy() *AzureRoleList {
	if in == nil {
		return nil
	}
	out := new(AzureRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureRoleSpec) DeepCopyInto(out *AzureRoleSpec) {
	*out = *in
	out.SecretEngineRef = in.SecretEngineRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureRoleSpec.
func (in *AzureRoleSpec) DeepCopy() *AzureRoleSpec {
	if in == nil {
		return nil
	}
	out := new(AzureRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchConfiguration) DeepCopyInto(out *ElasticsearchConfiguration) {
	*out = *in
	in.DatabaseRef.DeepCopyInto(&out.DatabaseRef)
	if in.AllowedRoles != nil {
		in, out := &in.AllowedRoles, &out.AllowedRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchConfiguration.
func (in *ElasticsearchConfiguration) DeepCopy() *ElasticsearchConfiguration {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchRole) DeepCopyInto(out *ElasticsearchRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchRole.
func (in *ElasticsearchRole) DeepCopy() *ElasticsearchRole {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElasticsearchRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchRoleList) DeepCopyInto(out *ElasticsearchRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ElasticsearchRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchRoleList.
func (in *ElasticsearchRoleList) DeepCopy() *ElasticsearchRoleList {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElasticsearchRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchRoleSpec) DeepCopyInto(out *ElasticsearchRoleSpec) {
	*out = *in
	out.SecretEngineRef = in.SecretEngineRef
	if in.CreationStatements != nil {
		in, out := &in.CreationStatements, &out.CreationStatements
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RevocationStatements != nil {
		in, out := &in.RevocationStatements, &out.RevocationStatements
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchRoleSpec.
func (in *ElasticsearchRoleSpec) DeepCopy() *ElasticsearchRoleSpec {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPAccessRequestConfiguration) DeepCopyInto(out *GCPAccessRequestConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPAccessRequestConfiguration.
func (in *GCPAccessRequestConfiguration) DeepCopy() *GCPAccessRequestConfiguration {
	if in == nil {
		return nil
	}
	out := new(GCPAccessRequestConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPConfiguration) DeepCopyInto(out *GCPConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPConfiguration.
func (in *GCPConfiguration) DeepCopy() *GCPConfiguration {
	if in == nil {
		return nil
	}
	out := new(GCPConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPRole) DeepCopyInto(out *GCPRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPRole.
func (in *GCPRole) DeepCopy() *GCPRole {
	if in == nil {
		return nil
	}
	out := new(GCPRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GCPRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPRoleList) DeepCopyInto(out *GCPRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GCPRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPRoleList.
func (in *GCPRoleList) DeepCopy() *GCPRoleList {
	if in == nil {
		return nil
	}
	out := new(GCPRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GCPRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPRoleSpec) DeepCopyInto(out *GCPRoleSpec) {
	*out = *in
	out.SecretEngineRef = in.SecretEngineRef
	if in.TokenScopes != nil {
		in, out := &in.TokenScopes, &out.TokenScopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPRoleSpec.
func (in *GCPRoleSpec) DeepCopy() *GCPRoleSpec {
	if in == nil {
		return nil
	}
	out := new(GCPRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVConfiguration) DeepCopyInto(out *KVConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVConfiguration.
func (in *KVConfiguration) DeepCopy() *KVConfiguration {
	if in == nil {
		return nil
	}
	out := new(KVConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Lease) DeepCopyInto(out *Lease) {
	*out = *in
	out.Duration = in.Duration
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Lease.
func (in *Lease) DeepCopy() *Lease {
	if in == nil {
		return nil
	}
	out := new(Lease)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LeaseConfig) DeepCopyInto(out *LeaseConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LeaseConfig.
func (in *LeaseConfig) DeepCopy() *LeaseConfig {
	if in == nil {
		return nil
	}
	out := new(LeaseConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBConfiguration) DeepCopyInto(out *MongoDBConfiguration) {
	*out = *in
	in.DatabaseRef.DeepCopyInto(&out.DatabaseRef)
	if in.AllowedRoles != nil {
		in, out := &in.AllowedRoles, &out.AllowedRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBConfiguration.
func (in *MongoDBConfiguration) DeepCopy() *MongoDBConfiguration {
	if in == nil {
		return nil
	}
	out := new(MongoDBConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBRole) DeepCopyInto(out *MongoDBRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBRole.
func (in *MongoDBRole) DeepCopy() *MongoDBRole {
	if in == nil {
		return nil
	}
	out := new(MongoDBRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MongoDBRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBRoleList) DeepCopyInto(out *MongoDBRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MongoDBRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBRoleList.
func (in *MongoDBRoleList) DeepCopy() *MongoDBRoleList {
	if in == nil {
		return nil
	}
	out := new(MongoDBRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MongoDBRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBRoleSpec) DeepCopyInto(out *MongoDBRoleSpec) {
	*out = *in
	out.SecretEngineRef = in.SecretEngineRef
	if in.CreationStatements != nil {
		in, out := &in.CreationStatements, &out.CreationStatements
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RevocationStatements != nil {
		in, out := &in.RevocationStatements, &out.RevocationStatements
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBRoleSpec.
func (in *MongoDBRoleSpec) DeepCopy() *MongoDBRoleSpec {
	if in == nil {
		return nil
	}
	out := new(MongoDBRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLConfiguration) DeepCopyInto(out *MySQLConfiguration) {
	*out = *in
	in.DatabaseRef.DeepCopyInto(&out.DatabaseRef)
	if in.AllowedRoles != nil {
		in, out := &in.AllowedRoles, &out.AllowedRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLConfiguration.
func (in *MySQLConfiguration) DeepCopy() *MySQLConfiguration {
	if in == nil {
		return nil
	}
	out := new(MySQLConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLRole) DeepCopyInto(out *MySQLRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLRole.
func (in *MySQLRole) DeepCopy() *MySQLRole {
	if in == nil {
		return nil
	}
	out := new(MySQLRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MySQLRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLRoleList) DeepCopyInto(out *MySQLRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MySQLRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLRoleList.
func (in *MySQLRoleList) DeepCopy() *MySQLRoleList {
	if in == nil {
		return nil
	}
	out := new(MySQLRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MySQLRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLRoleSpec) DeepCopyInto(out *MySQLRoleSpec) {
	*out = *in
	out.SecretEngineRef = in.SecretEngineRef
	if in.CreationStatements != nil {
		in, out := &in.CreationStatements, &out.CreationStatements
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RevocationStatements != nil {
		in, out := &in.RevocationStatements, &out.RevocationStatements
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLRoleSpec.
func (in *MySQLRoleSpec) DeepCopy() *MySQLRoleSpec {
	if in == nil {
		return nil
	}
	out := new(MySQLRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresConfiguration) DeepCopyInto(out *PostgresConfiguration) {
	*out = *in
	in.DatabaseRef.DeepCopyInto(&out.DatabaseRef)
	if in.AllowedRoles != nil {
		in, out := &in.AllowedRoles, &out.AllowedRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresConfiguration.
func (in *PostgresConfiguration) DeepCopy() *PostgresConfiguration {
	if in == nil {
		return nil
	}
	out := new(PostgresConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresRole) DeepCopyInto(out *PostgresRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresRole.
func (in *PostgresRole) DeepCopy() *PostgresRole {
	if in == nil {
		return nil
	}
	out := new(PostgresRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresRoleList) DeepCopyInto(out *PostgresRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PostgresRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresRoleList.
func (in *PostgresRoleList) DeepCopy() *PostgresRoleList {
	if in == nil {
		return nil
	}
	out := new(PostgresRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresRoleSpec) DeepCopyInto(out *PostgresRoleSpec) {
	*out = *in
	out.SecretEngineRef = in.SecretEngineRef
	if in.CreationStatements != nil {
		in, out := &in.CreationStatements, &out.CreationStatements
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RevocationStatements != nil {
		in, out := &in.RevocationStatements, &out.RevocationStatements
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RollbackStatements != nil {
		in, out := &in.RollbackStatements, &out.RollbackStatements
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RenewStatements != nil {
		in, out := &in.RenewStatements, &out.RenewStatements
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresRoleSpec.
func (in *PostgresRoleSpec) DeepCopy() *PostgresRoleSpec {
	if in == nil {
		return nil
	}
	out := new(PostgresRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleStatus) DeepCopyInto(out *RoleStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PolicyRef != nil {
		in, out := &in.PolicyRef, &out.PolicyRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleStatus.
func (in *RoleStatus) DeepCopy() *RoleStatus {
	if in == nil {
		return nil
	}
	out := new(RoleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretAccessRequest) DeepCopyInto(out *SecretAccessRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretAccessRequest.
func (in *SecretAccessRequest) DeepCopy() *SecretAccessRequest {
	if in == nil {
		return nil
	}
	out := new(SecretAccessRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretAccessRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretAccessRequestConfiguration) DeepCopyInto(out *SecretAccessRequestConfiguration) {
	*out = *in
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		*out = new(AWSAccessRequestConfiguration)
		**out = **in
	}
	if in.GCP != nil {
		in, out := &in.GCP, &out.GCP
		*out = new(GCPAccessRequestConfiguration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretAccessRequestConfiguration.
func (in *SecretAccessRequestConfiguration) DeepCopy() *SecretAccessRequestConfiguration {
	if in == nil {
		return nil
	}
	out := new(SecretAccessRequestConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretAccessRequestList) DeepCopyInto(out *SecretAccessRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecretAccessRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretAccessRequestList.
func (in *SecretAccessRequestList) DeepCopy() *SecretAccessRequestList {
	if in == nil {
		return nil
	}
	out := new(SecretAccessRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretAccessRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretAccessRequestSpec) DeepCopyInto(out *SecretAccessRequestSpec) {
	*out = *in
	in.RoleRef.DeepCopyInto(&out.RoleRef)
	if in.Subjects != nil {
		in, out := &in.Subjects, &out.Subjects
		*out = make([]rbacv1.Subject, len(*in))
		copy(*out, *in)
	}
	in.SecretAccessRequestConfiguration.DeepCopyInto(&out.SecretAccessRequestConfiguration)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretAccessRequestSpec.
func (in *SecretAccessRequestSpec) DeepCopy() *SecretAccessRequestSpec {
	if in == nil {
		return nil
	}
	out := new(SecretAccessRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretAccessRequestStatus) DeepCopyInto(out *SecretAccessRequestStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Lease != nil {
		in, out := &in.Lease, &out.Lease
		*out = new(Lease)
		**out = **in
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(v1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretAccessRequestStatus.
func (in *SecretAccessRequestStatus) DeepCopy() *SecretAccessRequestStatus {
	if in == nil {
		return nil
	}
	out := new(SecretAccessRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretEngine) DeepCopyInto(out *SecretEngine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretEngine.
func (in *SecretEngine) DeepCopy() *SecretEngine {
	if in == nil {
		return nil
	}
	out := new(SecretEngine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretEngine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretEngineConfiguration) DeepCopyInto(out *SecretEngineConfiguration) {
	*out = *in
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		*out = new(AWSConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Azure != nil {
		in, out := &in.Azure, &out.Azure
		*out = new(AzureConfiguration)
		**out = **in
	}
	if in.GCP != nil {
		in, out := &in.GCP, &out.GCP
		*out = new(GCPConfiguration)
		**out = **in
	}
	if in.Postgres != nil {
		in, out := &in.Postgres, &out.Postgres
		*out = new(PostgresConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.MongoDB != nil {
		in, out := &in.MongoDB, &out.MongoDB
		*out = new(MongoDBConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.MySQL != nil {
		in, out := &in.MySQL, &out.MySQL
		*out = new(MySQLConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.KV != nil {
		in, out := &in.KV, &out.KV
		*out = new(KVConfiguration)
		**out = **in
	}
	if in.Elasticsearch != nil {
		in, out := &in.Elasticsearch, &out.Elasticsearch
		*out = new(ElasticsearchConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretEngineConfiguration.
func (in *SecretEngineConfiguration) DeepCopy() *SecretEngineConfiguration {
	if in == nil {
		return nil
	}
	out := new(SecretEngineConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretEngineList) DeepCopyInto(out *SecretEngineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecretEngine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretEngineList.
func (in *SecretEngineList) DeepCopy() *SecretEngineList {
	if in == nil {
		return nil
	}
	out := new(SecretEngineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretEngineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretEngineSpec) DeepCopyInto(out *SecretEngineSpec) {
	*out = *in
	out.VaultRef = in.VaultRef
	in.SecretEngineConfiguration.DeepCopyInto(&out.SecretEngineConfiguration)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretEngineSpec.
func (in *SecretEngineSpec) DeepCopy() *SecretEngineSpec {
	if in == nil {
		return nil
	}
	out := new(SecretEngineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretEngineStatus) DeepCopyInto(out *SecretEngineStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretEngineStatus.
func (in *SecretEngineStatus) DeepCopy() *SecretEngineStatus {
	if in == nil {
		return nil
	}
	out := new(SecretEngineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRoleBinding) DeepCopyInto(out *SecretRoleBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRoleBinding.
func (in *SecretRoleBinding) DeepCopy() *SecretRoleBinding {
	if in == nil {
		return nil
	}
	out := new(SecretRoleBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretRoleBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRoleBindingList) DeepCopyInto(out *SecretRoleBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecretRoleBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRoleBindingList.
func (in *SecretRoleBindingList) DeepCopy() *SecretRoleBindingList {
	if in == nil {
		return nil
	}
	out := new(SecretRoleBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretRoleBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRoleBindingSpec) DeepCopyInto(out *SecretRoleBindingSpec) {
	*out = *in
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]corev1.TypedLocalObjectReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Subjects != nil {
		in, out := &in.Subjects, &out.Subjects
		*out = make([]rbacv1.Subject, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRoleBindingSpec.
func (in *SecretRoleBindingSpec) DeepCopy() *SecretRoleBindingSpec {
	if in == nil {
		return nil
	}
	out := new(SecretRoleBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRoleBindingStatus) DeepCopyInto(out *SecretRoleBindingStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Lease != nil {
		in, out := &in.Lease, &out.Lease
		*out = new(Lease)
		**out = **in
	}
	if in.PolicyRef != nil {
		in, out := &in.PolicyRef, &out.PolicyRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.PolicyBindingRef != nil {
		in, out := &in.PolicyBindingRef, &out.PolicyBindingRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRoleBindingStatus.
func (in *SecretRoleBindingStatus) DeepCopy() *SecretRoleBindingStatus {
	if in == nil {
		return nil
	}
	out := new(SecretRoleBindingStatus)
	in.DeepCopyInto(out)
	return out
}
