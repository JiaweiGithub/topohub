// Copyright 2024 Authors of infrastructure-io
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/infrastructure-io/topohub/pkg/k8s/apis/topohub.infrastructure.io/v1beta1"
	topohubinfrastructureiov1beta1 "github.com/infrastructure-io/topohub/pkg/k8s/client/clientset/versioned/typed/topohub.infrastructure.io/v1beta1"
	gentype "k8s.io/client-go/gentype"
)

// fakeBindingIps implements BindingIpInterface
type fakeBindingIps struct {
	*gentype.FakeClientWithList[*v1beta1.BindingIp, *v1beta1.BindingIpList]
	Fake *FakeTopohubV1beta1
}

func newFakeBindingIps(fake *FakeTopohubV1beta1) topohubinfrastructureiov1beta1.BindingIpInterface {
	return &fakeBindingIps{
		gentype.NewFakeClientWithList[*v1beta1.BindingIp, *v1beta1.BindingIpList](
			fake.Fake,
			"",
			v1beta1.SchemeGroupVersion.WithResource("bindingips"),
			v1beta1.SchemeGroupVersion.WithKind("BindingIp"),
			func() *v1beta1.BindingIp { return &v1beta1.BindingIp{} },
			func() *v1beta1.BindingIpList { return &v1beta1.BindingIpList{} },
			func(dst, src *v1beta1.BindingIpList) { dst.ListMeta = src.ListMeta },
			func(list *v1beta1.BindingIpList) []*v1beta1.BindingIp { return gentype.ToPointerSlice(list.Items) },
			func(list *v1beta1.BindingIpList, items []*v1beta1.BindingIp) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
