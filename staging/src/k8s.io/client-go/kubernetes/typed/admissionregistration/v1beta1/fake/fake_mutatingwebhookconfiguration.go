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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "k8s.io/api/admissionregistration/v1beta1"
	admissionregistrationv1beta1 "k8s.io/client-go/applyconfigurations/admissionregistration/v1beta1"
	gentype "k8s.io/client-go/gentype"
	typedadmissionregistrationv1beta1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1"
)

// fakeMutatingWebhookConfigurations implements MutatingWebhookConfigurationInterface
type fakeMutatingWebhookConfigurations struct {
	*gentype.FakeClientWithListAndApply[*v1beta1.MutatingWebhookConfiguration, *v1beta1.MutatingWebhookConfigurationList, *admissionregistrationv1beta1.MutatingWebhookConfigurationApplyConfiguration]
	Fake *FakeAdmissionregistrationV1beta1
}

func newFakeMutatingWebhookConfigurations(fake *FakeAdmissionregistrationV1beta1) typedadmissionregistrationv1beta1.MutatingWebhookConfigurationInterface {
	return &fakeMutatingWebhookConfigurations{
		gentype.NewFakeClientWithListAndApply[*v1beta1.MutatingWebhookConfiguration, *v1beta1.MutatingWebhookConfigurationList, *admissionregistrationv1beta1.MutatingWebhookConfigurationApplyConfiguration](
			fake.Fake,
			"",
			v1beta1.SchemeGroupVersion.WithResource("mutatingwebhookconfigurations"),
			v1beta1.SchemeGroupVersion.WithKind("MutatingWebhookConfiguration"),
			func() *v1beta1.MutatingWebhookConfiguration { return &v1beta1.MutatingWebhookConfiguration{} },
			func() *v1beta1.MutatingWebhookConfigurationList { return &v1beta1.MutatingWebhookConfigurationList{} },
			func(dst, src *v1beta1.MutatingWebhookConfigurationList) { dst.ListMeta = src.ListMeta },
			func(list *v1beta1.MutatingWebhookConfigurationList) []*v1beta1.MutatingWebhookConfiguration {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1beta1.MutatingWebhookConfigurationList, items []*v1beta1.MutatingWebhookConfiguration) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
