// Copyright 2024 Authors of infrastructure-io
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/infrastructure-io/topohub/pkg/k8s/apis/topohub.infrastructure.io/v1beta1"
	topohubinfrastructureiov1beta1 "github.com/infrastructure-io/topohub/pkg/k8s/client/clientset/versioned/typed/topohub.infrastructure.io/v1beta1"
	gentype "k8s.io/client-go/gentype"
)

// fakeHostStatuses implements HostStatusInterface
type fakeHostStatuses struct {
	*gentype.FakeClientWithList[*v1beta1.HostStatus, *v1beta1.HostStatusList]
	Fake *FakeTopohubV1beta1
}

func newFakeHostStatuses(fake *FakeTopohubV1beta1) topohubinfrastructureiov1beta1.HostStatusInterface {
	return &fakeHostStatuses{
		gentype.NewFakeClientWithList[*v1beta1.HostStatus, *v1beta1.HostStatusList](
			fake.Fake,
			"",
			v1beta1.SchemeGroupVersion.WithResource("hoststatuses"),
			v1beta1.SchemeGroupVersion.WithKind("HostStatus"),
			func() *v1beta1.HostStatus { return &v1beta1.HostStatus{} },
			func() *v1beta1.HostStatusList { return &v1beta1.HostStatusList{} },
			func(dst, src *v1beta1.HostStatusList) { dst.ListMeta = src.ListMeta },
			func(list *v1beta1.HostStatusList) []*v1beta1.HostStatus { return gentype.ToPointerSlice(list.Items) },
			func(list *v1beta1.HostStatusList, items []*v1beta1.HostStatus) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
