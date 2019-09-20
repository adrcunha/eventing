// +build !ignore_autogenerated

/*
Copyright 2019 The Knative Authors

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
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	duckv1alpha1 "knative.dev/eventing/pkg/apis/duck/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Channel) DeepCopyInto(out *Channel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Channel.
func (in *Channel) DeepCopy() *Channel {
	if in == nil {
		return nil
	}
	out := new(Channel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Channel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChannelList) DeepCopyInto(out *ChannelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Channel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChannelList.
func (in *ChannelList) DeepCopy() *ChannelList {
	if in == nil {
		return nil
	}
	out := new(ChannelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChannelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChannelSpec) DeepCopyInto(out *ChannelSpec) {
	*out = *in
	if in.ChannelTemplate != nil {
		in, out := &in.ChannelTemplate, &out.ChannelTemplate
		*out = new(duckv1alpha1.ChannelTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Subscribable != nil {
		in, out := &in.Subscribable, &out.Subscribable
		*out = new(duckv1alpha1.Subscribable)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChannelSpec.
func (in *ChannelSpec) DeepCopy() *ChannelSpec {
	if in == nil {
		return nil
	}
	out := new(ChannelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChannelStatus) DeepCopyInto(out *ChannelStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	in.AddressStatus.DeepCopyInto(&out.AddressStatus)
	in.SubscribableTypeStatus.DeepCopyInto(&out.SubscribableTypeStatus)
	if in.Channel != nil {
		in, out := &in.Channel, &out.Channel
		*out = new(v1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChannelStatus.
func (in *ChannelStatus) DeepCopy() *ChannelStatus {
	if in == nil {
		return nil
	}
	out := new(ChannelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InMemoryChannel) DeepCopyInto(out *InMemoryChannel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InMemoryChannel.
func (in *InMemoryChannel) DeepCopy() *InMemoryChannel {
	if in == nil {
		return nil
	}
	out := new(InMemoryChannel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InMemoryChannel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InMemoryChannelList) DeepCopyInto(out *InMemoryChannelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InMemoryChannel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InMemoryChannelList.
func (in *InMemoryChannelList) DeepCopy() *InMemoryChannelList {
	if in == nil {
		return nil
	}
	out := new(InMemoryChannelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InMemoryChannelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InMemoryChannelSpec) DeepCopyInto(out *InMemoryChannelSpec) {
	*out = *in
	if in.Subscribable != nil {
		in, out := &in.Subscribable, &out.Subscribable
		*out = new(duckv1alpha1.Subscribable)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InMemoryChannelSpec.
func (in *InMemoryChannelSpec) DeepCopy() *InMemoryChannelSpec {
	if in == nil {
		return nil
	}
	out := new(InMemoryChannelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InMemoryChannelStatus) DeepCopyInto(out *InMemoryChannelStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	in.AddressStatus.DeepCopyInto(&out.AddressStatus)
	in.SubscribableTypeStatus.DeepCopyInto(&out.SubscribableTypeStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InMemoryChannelStatus.
func (in *InMemoryChannelStatus) DeepCopy() *InMemoryChannelStatus {
	if in == nil {
		return nil
	}
	out := new(InMemoryChannelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Parallel) DeepCopyInto(out *Parallel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Parallel.
func (in *Parallel) DeepCopy() *Parallel {
	if in == nil {
		return nil
	}
	out := new(Parallel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Parallel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParallelBranch) DeepCopyInto(out *ParallelBranch) {
	*out = *in
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(SubscriberSpec)
		(*in).DeepCopyInto(*out)
	}
	in.Subscriber.DeepCopyInto(&out.Subscriber)
	if in.Reply != nil {
		in, out := &in.Reply, &out.Reply
		*out = new(v1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParallelBranch.
func (in *ParallelBranch) DeepCopy() *ParallelBranch {
	if in == nil {
		return nil
	}
	out := new(ParallelBranch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParallelBranchStatus) DeepCopyInto(out *ParallelBranchStatus) {
	*out = *in
	in.FilterSubscriptionStatus.DeepCopyInto(&out.FilterSubscriptionStatus)
	in.FilterChannelStatus.DeepCopyInto(&out.FilterChannelStatus)
	in.SubscriptionStatus.DeepCopyInto(&out.SubscriptionStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParallelBranchStatus.
func (in *ParallelBranchStatus) DeepCopy() *ParallelBranchStatus {
	if in == nil {
		return nil
	}
	out := new(ParallelBranchStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParallelChannelStatus) DeepCopyInto(out *ParallelChannelStatus) {
	*out = *in
	out.Channel = in.Channel
	in.ReadyCondition.DeepCopyInto(&out.ReadyCondition)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParallelChannelStatus.
func (in *ParallelChannelStatus) DeepCopy() *ParallelChannelStatus {
	if in == nil {
		return nil
	}
	out := new(ParallelChannelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParallelList) DeepCopyInto(out *ParallelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Parallel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParallelList.
func (in *ParallelList) DeepCopy() *ParallelList {
	if in == nil {
		return nil
	}
	out := new(ParallelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ParallelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParallelSpec) DeepCopyInto(out *ParallelSpec) {
	*out = *in
	if in.Branches != nil {
		in, out := &in.Branches, &out.Branches
		*out = make([]ParallelBranch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ChannelTemplate != nil {
		in, out := &in.ChannelTemplate, &out.ChannelTemplate
		*out = new(duckv1alpha1.ChannelTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Reply != nil {
		in, out := &in.Reply, &out.Reply
		*out = new(v1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParallelSpec.
func (in *ParallelSpec) DeepCopy() *ParallelSpec {
	if in == nil {
		return nil
	}
	out := new(ParallelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParallelStatus) DeepCopyInto(out *ParallelStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	in.IngressChannelStatus.DeepCopyInto(&out.IngressChannelStatus)
	if in.BranchStatuses != nil {
		in, out := &in.BranchStatuses, &out.BranchStatuses
		*out = make([]ParallelBranchStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.AddressStatus.DeepCopyInto(&out.AddressStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParallelStatus.
func (in *ParallelStatus) DeepCopy() *ParallelStatus {
	if in == nil {
		return nil
	}
	out := new(ParallelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParallelSubscriptionStatus) DeepCopyInto(out *ParallelSubscriptionStatus) {
	*out = *in
	out.Subscription = in.Subscription
	in.ReadyCondition.DeepCopyInto(&out.ReadyCondition)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParallelSubscriptionStatus.
func (in *ParallelSubscriptionStatus) DeepCopy() *ParallelSubscriptionStatus {
	if in == nil {
		return nil
	}
	out := new(ParallelSubscriptionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplyStrategy) DeepCopyInto(out *ReplyStrategy) {
	*out = *in
	if in.Channel != nil {
		in, out := &in.Channel, &out.Channel
		*out = new(v1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplyStrategy.
func (in *ReplyStrategy) DeepCopy() *ReplyStrategy {
	if in == nil {
		return nil
	}
	out := new(ReplyStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sequence) DeepCopyInto(out *Sequence) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sequence.
func (in *Sequence) DeepCopy() *Sequence {
	if in == nil {
		return nil
	}
	out := new(Sequence)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Sequence) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SequenceChannelStatus) DeepCopyInto(out *SequenceChannelStatus) {
	*out = *in
	out.Channel = in.Channel
	in.ReadyCondition.DeepCopyInto(&out.ReadyCondition)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SequenceChannelStatus.
func (in *SequenceChannelStatus) DeepCopy() *SequenceChannelStatus {
	if in == nil {
		return nil
	}
	out := new(SequenceChannelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SequenceList) DeepCopyInto(out *SequenceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Sequence, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SequenceList.
func (in *SequenceList) DeepCopy() *SequenceList {
	if in == nil {
		return nil
	}
	out := new(SequenceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SequenceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SequenceSpec) DeepCopyInto(out *SequenceSpec) {
	*out = *in
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]SubscriberSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ChannelTemplate != nil {
		in, out := &in.ChannelTemplate, &out.ChannelTemplate
		*out = new(duckv1alpha1.ChannelTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Reply != nil {
		in, out := &in.Reply, &out.Reply
		*out = new(v1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SequenceSpec.
func (in *SequenceSpec) DeepCopy() *SequenceSpec {
	if in == nil {
		return nil
	}
	out := new(SequenceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SequenceStatus) DeepCopyInto(out *SequenceStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.SubscriptionStatuses != nil {
		in, out := &in.SubscriptionStatuses, &out.SubscriptionStatuses
		*out = make([]SequenceSubscriptionStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ChannelStatuses != nil {
		in, out := &in.ChannelStatuses, &out.ChannelStatuses
		*out = make([]SequenceChannelStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.AddressStatus.DeepCopyInto(&out.AddressStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SequenceStatus.
func (in *SequenceStatus) DeepCopy() *SequenceStatus {
	if in == nil {
		return nil
	}
	out := new(SequenceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SequenceSubscriptionStatus) DeepCopyInto(out *SequenceSubscriptionStatus) {
	*out = *in
	out.Subscription = in.Subscription
	in.ReadyCondition.DeepCopyInto(&out.ReadyCondition)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SequenceSubscriptionStatus.
func (in *SequenceSubscriptionStatus) DeepCopy() *SequenceSubscriptionStatus {
	if in == nil {
		return nil
	}
	out := new(SequenceSubscriptionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriberSpec) DeepCopyInto(out *SubscriberSpec) {
	*out = *in
	if in.Ref != nil {
		in, out := &in.Ref, &out.Ref
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.DeprecatedDNSName != nil {
		in, out := &in.DeprecatedDNSName, &out.DeprecatedDNSName
		*out = new(string)
		**out = **in
	}
	if in.URI != nil {
		in, out := &in.URI, &out.URI
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriberSpec.
func (in *SubscriberSpec) DeepCopy() *SubscriberSpec {
	if in == nil {
		return nil
	}
	out := new(SubscriberSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subscription) DeepCopyInto(out *Subscription) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subscription.
func (in *Subscription) DeepCopy() *Subscription {
	if in == nil {
		return nil
	}
	out := new(Subscription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Subscription) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionList) DeepCopyInto(out *SubscriptionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Subscription, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionList.
func (in *SubscriptionList) DeepCopy() *SubscriptionList {
	if in == nil {
		return nil
	}
	out := new(SubscriptionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubscriptionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionSpec) DeepCopyInto(out *SubscriptionSpec) {
	*out = *in
	out.Channel = in.Channel
	if in.Subscriber != nil {
		in, out := &in.Subscriber, &out.Subscriber
		*out = new(SubscriberSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Reply != nil {
		in, out := &in.Reply, &out.Reply
		*out = new(ReplyStrategy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionSpec.
func (in *SubscriptionSpec) DeepCopy() *SubscriptionSpec {
	if in == nil {
		return nil
	}
	out := new(SubscriptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionStatus) DeepCopyInto(out *SubscriptionStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	out.PhysicalSubscription = in.PhysicalSubscription
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionStatus.
func (in *SubscriptionStatus) DeepCopy() *SubscriptionStatus {
	if in == nil {
		return nil
	}
	out := new(SubscriptionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionStatusPhysicalSubscription) DeepCopyInto(out *SubscriptionStatusPhysicalSubscription) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionStatusPhysicalSubscription.
func (in *SubscriptionStatusPhysicalSubscription) DeepCopy() *SubscriptionStatusPhysicalSubscription {
	if in == nil {
		return nil
	}
	out := new(SubscriptionStatusPhysicalSubscription)
	in.DeepCopyInto(out)
	return out
}
