/*
Copyright (c) 2016-2017 Bitnami

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

package v1beta1

import (
	kubeless_v1beta1 "github.com/kubeless/kubeless/pkg/apis/kubeless/v1beta1"
	versioned "github.com/kubeless/kubeless/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kubeless/kubeless/pkg/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/kubeless/kubeless/pkg/client/listers/kubeless/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// KafkaTriggerInformer provides access to a shared informer and lister for
// KafkaTriggers.
type KafkaTriggerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.KafkaTriggerLister
}

type kafkaTriggerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewKafkaTriggerInformer constructs a new informer for KafkaTrigger type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKafkaTriggerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKafkaTriggerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredKafkaTriggerInformer constructs a new informer for KafkaTrigger type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKafkaTriggerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubelessV1beta1().KafkaTriggers(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubelessV1beta1().KafkaTriggers(namespace).Watch(options)
			},
		},
		&kubeless_v1beta1.KafkaTrigger{},
		resyncPeriod,
		indexers,
	)
}

func (f *kafkaTriggerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKafkaTriggerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *kafkaTriggerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubeless_v1beta1.KafkaTrigger{}, f.defaultInformer)
}

func (f *kafkaTriggerInformer) Lister() v1beta1.KafkaTriggerLister {
	return v1beta1.NewKafkaTriggerLister(f.Informer().GetIndexer())
}
