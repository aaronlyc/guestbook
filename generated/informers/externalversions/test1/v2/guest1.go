/*
Copyright 2022 My name.

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
// Code generated by informer-gen. DO NOT EDIT.

package v2

import (
	"context"
	test1v2 "guestbook/apis/test1/v2"
	versioned "guestbook/generated/clientset/versioned"
	internalinterfaces "guestbook/generated/informers/externalversions/internalinterfaces"
	v2 "guestbook/generated/listers/test1/v2"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// Guest1Informer provides access to a shared informer and lister for
// Guest1s.
type Guest1Informer interface {
	Informer() cache.SharedIndexInformer
	Lister() v2.Guest1Lister
}

type guest1Informer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewGuest1Informer constructs a new informer for Guest1 type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewGuest1Informer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredGuest1Informer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredGuest1Informer constructs a new informer for Guest1 type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredGuest1Informer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Test1V2().Guest1s(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Test1V2().Guest1s(namespace).Watch(context.TODO(), options)
			},
		},
		&test1v2.Guest1{},
		resyncPeriod,
		indexers,
	)
}

func (f *guest1Informer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredGuest1Informer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *guest1Informer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&test1v2.Guest1{}, f.defaultInformer)
}

func (f *guest1Informer) Lister() v2.Guest1Lister {
	return v2.NewGuest1Lister(f.Informer().GetIndexer())
}