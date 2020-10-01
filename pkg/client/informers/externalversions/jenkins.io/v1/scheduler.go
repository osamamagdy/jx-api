/*
Copyright 2020 The Jenkins X Authors.

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

package v1

import (
	"context"
	time "time"

	jenkinsiov1 "github.com/jenkins-x/jx-api/v3/pkg/apis/jenkins.io/v1"
	versioned "github.com/jenkins-x/jx-api/v3/pkg/client/clientset/versioned"
	internalinterfaces "github.com/jenkins-x/jx-api/v3/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/jenkins-x/jx-api/v3/pkg/client/listers/jenkins.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SchedulerInformer provides access to a shared informer and lister for
// Schedulers.
type SchedulerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.SchedulerLister
}

type schedulerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSchedulerInformer constructs a new informer for Scheduler type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSchedulerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSchedulerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSchedulerInformer constructs a new informer for Scheduler type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSchedulerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.JenkinsV1().Schedulers(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.JenkinsV1().Schedulers(namespace).Watch(context.TODO(), options)
			},
		},
		&jenkinsiov1.Scheduler{},
		resyncPeriod,
		indexers,
	)
}

func (f *schedulerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSchedulerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *schedulerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&jenkinsiov1.Scheduler{}, f.defaultInformer)
}

func (f *schedulerInformer) Lister() v1.SchedulerLister {
	return v1.NewSchedulerLister(f.Informer().GetIndexer())
}
