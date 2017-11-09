/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package internalversion

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	admissionregistration "k8s.io/kubernetes/pkg/apis/admissionregistration"
	internalclientset "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
	internalinterfaces "k8s.io/kubernetes/pkg/client/informers/informers_generated/internalversion/internalinterfaces"
	internalversion "k8s.io/kubernetes/pkg/client/listers/admissionregistration/internalversion"
	time "time"
)

// ExternalAdmissionHookConfigurationInformer provides access to a shared informer and lister for
// ExternalAdmissionHookConfigurations.
type ExternalAdmissionHookConfigurationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.ExternalAdmissionHookConfigurationLister
}

type externalAdmissionHookConfigurationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewExternalAdmissionHookConfigurationInformer constructs a new informer for ExternalAdmissionHookConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewExternalAdmissionHookConfigurationInformer(client internalclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredExternalAdmissionHookConfigurationInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredExternalAdmissionHookConfigurationInformer constructs a new informer for ExternalAdmissionHookConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredExternalAdmissionHookConfigurationInformer(client internalclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Admissionregistration().ExternalAdmissionHookConfigurations().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Admissionregistration().ExternalAdmissionHookConfigurations().Watch(options)
			},
		},
		&admissionregistration.ExternalAdmissionHookConfiguration{},
		resyncPeriod,
		indexers,
	)
}

func (f *externalAdmissionHookConfigurationInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredExternalAdmissionHookConfigurationInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *externalAdmissionHookConfigurationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&admissionregistration.ExternalAdmissionHookConfiguration{}, f.defaultInformer)
}

func (f *externalAdmissionHookConfigurationInformer) Lister() internalversion.ExternalAdmissionHookConfigurationLister {
	return internalversion.NewExternalAdmissionHookConfigurationLister(f.Informer().GetIndexer())
}
