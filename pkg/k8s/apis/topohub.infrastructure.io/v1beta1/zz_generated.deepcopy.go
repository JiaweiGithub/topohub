//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BasicInfo) DeepCopyInto(out *BasicInfo) {
	*out = *in
	if in.DhcpExpireTime != nil {
		in, out := &in.DhcpExpireTime, &out.DhcpExpireTime
		*out = new(string)
		**out = **in
	}
	if in.SubnetName != nil {
		in, out := &in.SubnetName, &out.SubnetName
		*out = new(string)
		**out = **in
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BasicInfo.
func (in *BasicInfo) DeepCopy() *BasicInfo {
	if in == nil {
		return nil
	}
	out := new(BasicInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingIp) DeepCopyInto(out *BindingIp) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingIp.
func (in *BindingIp) DeepCopy() *BindingIp {
	if in == nil {
		return nil
	}
	out := new(BindingIp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BindingIp) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingIpList) DeepCopyInto(out *BindingIpList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BindingIp, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingIpList.
func (in *BindingIpList) DeepCopy() *BindingIpList {
	if in == nil {
		return nil
	}
	out := new(BindingIpList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BindingIpList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingIpSpec) DeepCopyInto(out *BindingIpSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingIpSpec.
func (in *BindingIpSpec) DeepCopy() *BindingIpSpec {
	if in == nil {
		return nil
	}
	out := new(BindingIpSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingIpStatus) DeepCopyInto(out *BindingIpStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingIpStatus.
func (in *BindingIpStatus) DeepCopy() *BindingIpStatus {
	if in == nil {
		return nil
	}
	out := new(BindingIpStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DhcpStatusSpec) DeepCopyInto(out *DhcpStatusSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DhcpStatusSpec.
func (in *DhcpStatusSpec) DeepCopy() *DhcpStatusSpec {
	if in == nil {
		return nil
	}
	out := new(DhcpStatusSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnableSyncHoststatusSpec) DeepCopyInto(out *EnableSyncHoststatusSpec) {
	*out = *in
	if in.DefaultClusterName != nil {
		in, out := &in.DefaultClusterName, &out.DefaultClusterName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnableSyncHoststatusSpec.
func (in *EnableSyncHoststatusSpec) DeepCopy() *EnableSyncHoststatusSpec {
	if in == nil {
		return nil
	}
	out := new(EnableSyncHoststatusSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureSpec) DeepCopyInto(out *FeatureSpec) {
	*out = *in
	in.EnableSyncHoststatus.DeepCopyInto(&out.EnableSyncHoststatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureSpec.
func (in *FeatureSpec) DeepCopy() *FeatureSpec {
	if in == nil {
		return nil
	}
	out := new(FeatureSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostOperation) DeepCopyInto(out *HostOperation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostOperation.
func (in *HostOperation) DeepCopy() *HostOperation {
	if in == nil {
		return nil
	}
	out := new(HostOperation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HostOperation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostOperationList) DeepCopyInto(out *HostOperationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HostOperation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostOperationList.
func (in *HostOperationList) DeepCopy() *HostOperationList {
	if in == nil {
		return nil
	}
	out := new(HostOperationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HostOperationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostOperationSpec) DeepCopyInto(out *HostOperationSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostOperationSpec.
func (in *HostOperationSpec) DeepCopy() *HostOperationSpec {
	if in == nil {
		return nil
	}
	out := new(HostOperationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostOperationStatus) DeepCopyInto(out *HostOperationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostOperationStatus.
func (in *HostOperationStatus) DeepCopy() *HostOperationStatus {
	if in == nil {
		return nil
	}
	out := new(HostOperationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostStatus) DeepCopyInto(out *HostStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostStatus.
func (in *HostStatus) DeepCopy() *HostStatus {
	if in == nil {
		return nil
	}
	out := new(HostStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HostStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostStatusList) DeepCopyInto(out *HostStatusList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HostStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostStatusList.
func (in *HostStatusList) DeepCopy() *HostStatusList {
	if in == nil {
		return nil
	}
	out := new(HostStatusList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HostStatusList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostStatusStatus) DeepCopyInto(out *HostStatusStatus) {
	*out = *in
	in.Basic.DeepCopyInto(&out.Basic)
	if in.Info != nil {
		in, out := &in.Info, &out.Info
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Log.DeepCopyInto(&out.Log)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostStatusStatus.
func (in *HostStatusStatus) DeepCopy() *HostStatusStatus {
	if in == nil {
		return nil
	}
	out := new(HostStatusStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPv4SubnetSpec) DeepCopyInto(out *IPv4SubnetSpec) {
	*out = *in
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(string)
		**out = **in
	}
	if in.Dns != nil {
		in, out := &in.Dns, &out.Dns
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPv4SubnetSpec.
func (in *IPv4SubnetSpec) DeepCopy() *IPv4SubnetSpec {
	if in == nil {
		return nil
	}
	out := new(IPv4SubnetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfaceSpec) DeepCopyInto(out *InterfaceSpec) {
	*out = *in
	if in.VlanID != nil {
		in, out := &in.VlanID, &out.VlanID
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfaceSpec.
func (in *InterfaceSpec) DeepCopy() *InterfaceSpec {
	if in == nil {
		return nil
	}
	out := new(InterfaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogEntry) DeepCopyInto(out *LogEntry) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogEntry.
func (in *LogEntry) DeepCopy() *LogEntry {
	if in == nil {
		return nil
	}
	out := new(LogEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogStruct) DeepCopyInto(out *LogStruct) {
	*out = *in
	if in.LastestLog != nil {
		in, out := &in.LastestLog, &out.LastestLog
		*out = new(LogEntry)
		**out = **in
	}
	if in.LastestWarningLog != nil {
		in, out := &in.LastestWarningLog, &out.LastestWarningLog
		*out = new(LogEntry)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogStruct.
func (in *LogStruct) DeepCopy() *LogStruct {
	if in == nil {
		return nil
	}
	out := new(LogStruct)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subnet) DeepCopyInto(out *Subnet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subnet.
func (in *Subnet) DeepCopy() *Subnet {
	if in == nil {
		return nil
	}
	out := new(Subnet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Subnet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetList) DeepCopyInto(out *SubnetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Subnet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetList.
func (in *SubnetList) DeepCopy() *SubnetList {
	if in == nil {
		return nil
	}
	out := new(SubnetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubnetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetSpec) DeepCopyInto(out *SubnetSpec) {
	*out = *in
	in.IPv4Subnet.DeepCopyInto(&out.IPv4Subnet)
	in.Interface.DeepCopyInto(&out.Interface)
	if in.Feature != nil {
		in, out := &in.Feature, &out.Feature
		*out = new(FeatureSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetSpec.
func (in *SubnetSpec) DeepCopy() *SubnetSpec {
	if in == nil {
		return nil
	}
	out := new(SubnetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetStatus) DeepCopyInto(out *SubnetStatus) {
	*out = *in
	if in.DhcpStatus != nil {
		in, out := &in.DhcpStatus, &out.DhcpStatus
		*out = new(DhcpStatusSpec)
		**out = **in
	}
	if in.HostNode != nil {
		in, out := &in.HostNode, &out.HostNode
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetStatus.
func (in *SubnetStatus) DeepCopy() *SubnetStatus {
	if in == nil {
		return nil
	}
	out := new(SubnetStatus)
	in.DeepCopyInto(out)
	return out
}
