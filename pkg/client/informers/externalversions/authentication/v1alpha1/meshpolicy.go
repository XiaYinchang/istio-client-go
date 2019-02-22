/*
Portions Copyright 2018 The Kubernetes Authors.
Portions Copyright 2018 Aspen Mesh Authors.

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

package v1alpha1

import (
	time "time"

	authentication_v1alpha1 "github.com/XiaYinchang/istio-client-go/pkg/apis/authentication/v1alpha1"
	versioned "github.com/XiaYinchang/istio-client-go/pkg/client/clientset/versioned"
	internalinterfaces "github.com/XiaYinchang/istio-client-go/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/XiaYinchang/istio-client-go/pkg/client/listers/authentication/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MeshPolicyInformer provides access to a shared informer and lister for
// MeshPolicies.
type MeshPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.MeshPolicyLister
}

type meshPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewMeshPolicyInformer constructs a new informer for MeshPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMeshPolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMeshPolicyInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredMeshPolicyInformer constructs a new informer for MeshPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMeshPolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AuthenticationV1alpha1().MeshPolicies().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AuthenticationV1alpha1().MeshPolicies().Watch(options)
			},
		},
		&authentication_v1alpha1.MeshPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *meshPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMeshPolicyInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *meshPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&authentication_v1alpha1.MeshPolicy{}, f.defaultInformer)
}

func (f *meshPolicyInformer) Lister() v1alpha1.MeshPolicyLister {
	return v1alpha1.NewMeshPolicyLister(f.Informer().GetIndexer())
}
