//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package informers

import (
	"reflect"
	"sync"
	"time"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/cache"

	scopedclientset "github.com/kcp-dev/kcp/sdk/client/clientset/versioned"
	clientset "github.com/kcp-dev/kcp/sdk/client/clientset/versioned/cluster"
	apiresourceinformers "github.com/kcp-dev/kcp/sdk/client/informers/externalversions/apiresource"
	apisinformers "github.com/kcp-dev/kcp/sdk/client/informers/externalversions/apis"
	coreinformers "github.com/kcp-dev/kcp/sdk/client/informers/externalversions/core"
	"github.com/kcp-dev/kcp/sdk/client/informers/externalversions/internalinterfaces"
	schedulinginformers "github.com/kcp-dev/kcp/sdk/client/informers/externalversions/scheduling"
	tenancyinformers "github.com/kcp-dev/kcp/sdk/client/informers/externalversions/tenancy"
	topologyinformers "github.com/kcp-dev/kcp/sdk/client/informers/externalversions/topology"
	workloadinformers "github.com/kcp-dev/kcp/sdk/client/informers/externalversions/workload"
)

// SharedInformerOption defines the functional option type for SharedInformerFactory.
type SharedInformerOption func(*SharedInformerOptions) *SharedInformerOptions

type SharedInformerOptions struct {
	customResync     map[reflect.Type]time.Duration
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

type sharedInformerFactory struct {
	client           clientset.ClusterInterface
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	lock             sync.Mutex
	defaultResync    time.Duration
	customResync     map[reflect.Type]time.Duration

	informers map[reflect.Type]kcpcache.ScopeableSharedIndexInformer
	// startedInformers is used for tracking which informers have been started.
	// This allows Start() to be called multiple times safely.
	startedInformers map[reflect.Type]bool
}

// WithCustomResyncConfig sets a custom resync period for the specified informer types.
func WithCustomResyncConfig(resyncConfig map[metav1.Object]time.Duration) SharedInformerOption {
	return func(opts *SharedInformerOptions) *SharedInformerOptions {
		for k, v := range resyncConfig {
			opts.customResync[reflect.TypeOf(k)] = v
		}
		return opts
	}
}

// WithTweakListOptions sets a custom filter on all listers of the configured SharedInformerFactory.
func WithTweakListOptions(tweakListOptions internalinterfaces.TweakListOptionsFunc) SharedInformerOption {
	return func(opts *SharedInformerOptions) *SharedInformerOptions {
		opts.tweakListOptions = tweakListOptions
		return opts
	}
}

// NewSharedInformerFactory constructs a new instance of SharedInformerFactory for all namespaces.
func NewSharedInformerFactory(client clientset.ClusterInterface, defaultResync time.Duration) SharedInformerFactory {
	return NewSharedInformerFactoryWithOptions(client, defaultResync)
}

// NewSharedInformerFactoryWithOptions constructs a new instance of a SharedInformerFactory with additional options.
func NewSharedInformerFactoryWithOptions(client clientset.ClusterInterface, defaultResync time.Duration, options ...SharedInformerOption) SharedInformerFactory {
	factory := &sharedInformerFactory{
		client:           client,
		defaultResync:    defaultResync,
		informers:        make(map[reflect.Type]kcpcache.ScopeableSharedIndexInformer),
		startedInformers: make(map[reflect.Type]bool),
		customResync:     make(map[reflect.Type]time.Duration),
	}

	opts := &SharedInformerOptions{
		customResync: make(map[reflect.Type]time.Duration),
	}

	// Apply all options
	for _, opt := range options {
		opts = opt(opts)
	}

	// Forward options to the factory
	factory.customResync = opts.customResync
	factory.tweakListOptions = opts.tweakListOptions

	return factory
}

// Start initializes all requested informers.
func (f *sharedInformerFactory) Start(stopCh <-chan struct{}) {
	f.lock.Lock()
	defer f.lock.Unlock()

	for informerType, informer := range f.informers {
		if !f.startedInformers[informerType] {
			go informer.Run(stopCh)
			f.startedInformers[informerType] = true
		}
	}
}

// WaitForCacheSync waits for all started informers' cache were synced.
func (f *sharedInformerFactory) WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool {
	informers := func() map[reflect.Type]kcpcache.ScopeableSharedIndexInformer {
		f.lock.Lock()
		defer f.lock.Unlock()

		informers := map[reflect.Type]kcpcache.ScopeableSharedIndexInformer{}
		for informerType, informer := range f.informers {
			if f.startedInformers[informerType] {
				informers[informerType] = informer
			}
		}
		return informers
	}()

	res := map[reflect.Type]bool{}
	for informType, informer := range informers {
		res[informType] = cache.WaitForCacheSync(stopCh, informer.HasSynced)
	}
	return res
}

// InformerFor returns the SharedIndexInformer for obj using an internal
// client.
func (f *sharedInformerFactory) InformerFor(obj runtime.Object, newFunc internalinterfaces.NewInformerFunc) kcpcache.ScopeableSharedIndexInformer {
	f.lock.Lock()
	defer f.lock.Unlock()

	informerType := reflect.TypeOf(obj)
	informer, exists := f.informers[informerType]
	if exists {
		return informer
	}

	resyncPeriod, exists := f.customResync[informerType]
	if !exists {
		resyncPeriod = f.defaultResync
	}

	informer = newFunc(f.client, resyncPeriod)
	f.informers[informerType] = informer

	return informer
}

type ScopedDynamicSharedInformerFactory interface {
	ForResource(resource schema.GroupVersionResource) (GenericInformer, error)
	Start(stopCh <-chan struct{})
}

// SharedInformerFactory provides shared informers for resources in all known
// API group versions.
type SharedInformerFactory interface {
	internalinterfaces.SharedInformerFactory
	Cluster(logicalcluster.Name) ScopedDynamicSharedInformerFactory
	ForResource(resource schema.GroupVersionResource) (GenericClusterInformer, error)
	WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool

	Apiresource() apiresourceinformers.ClusterInterface
	Apis() apisinformers.ClusterInterface
	Core() coreinformers.ClusterInterface
	Scheduling() schedulinginformers.ClusterInterface
	Tenancy() tenancyinformers.ClusterInterface
	Topology() topologyinformers.ClusterInterface
	Workload() workloadinformers.ClusterInterface
}

func (f *sharedInformerFactory) Apiresource() apiresourceinformers.ClusterInterface {
	return apiresourceinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Apis() apisinformers.ClusterInterface {
	return apisinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Core() coreinformers.ClusterInterface {
	return coreinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Scheduling() schedulinginformers.ClusterInterface {
	return schedulinginformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Tenancy() tenancyinformers.ClusterInterface {
	return tenancyinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Topology() topologyinformers.ClusterInterface {
	return topologyinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Workload() workloadinformers.ClusterInterface {
	return workloadinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Cluster(clusterName logicalcluster.Name) ScopedDynamicSharedInformerFactory {
	return &scopedDynamicSharedInformerFactory{
		sharedInformerFactory: f,
		clusterName:           clusterName,
	}
}

type scopedDynamicSharedInformerFactory struct {
	*sharedInformerFactory
	clusterName logicalcluster.Name
}

func (f *scopedDynamicSharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	clusterInformer, err := f.sharedInformerFactory.ForResource(resource)
	if err != nil {
		return nil, err
	}
	return clusterInformer.Cluster(f.clusterName), nil
}

func (f *scopedDynamicSharedInformerFactory) Start(stopCh <-chan struct{}) {
	f.sharedInformerFactory.Start(stopCh)
}

// WithNamespace limits the SharedInformerFactory to the specified namespace.
func WithNamespace(namespace string) SharedInformerOption {
	return func(opts *SharedInformerOptions) *SharedInformerOptions {
		opts.namespace = namespace
		return opts
	}
}

type sharedScopedInformerFactory struct {
	client           scopedclientset.Interface
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	lock             sync.Mutex
	defaultResync    time.Duration
	customResync     map[reflect.Type]time.Duration

	informers map[reflect.Type]cache.SharedIndexInformer
	// startedInformers is used for tracking which informers have been started.
	// This allows Start() to be called multiple times safely.
	startedInformers map[reflect.Type]bool
}

// NewSharedScopedInformerFactory constructs a new instance of SharedInformerFactory for some or all namespaces.
func NewSharedScopedInformerFactory(client scopedclientset.Interface, defaultResync time.Duration, namespace string) SharedScopedInformerFactory {
	return NewSharedScopedInformerFactoryWithOptions(client, defaultResync, WithNamespace(namespace))
}

// NewSharedScopedInformerFactoryWithOptions constructs a new instance of a SharedInformerFactory with additional options.
func NewSharedScopedInformerFactoryWithOptions(client scopedclientset.Interface, defaultResync time.Duration, options ...SharedInformerOption) SharedScopedInformerFactory {
	factory := &sharedScopedInformerFactory{
		client:           client,
		defaultResync:    defaultResync,
		informers:        make(map[reflect.Type]cache.SharedIndexInformer),
		startedInformers: make(map[reflect.Type]bool),
		customResync:     make(map[reflect.Type]time.Duration),
	}

	opts := &SharedInformerOptions{
		customResync: make(map[reflect.Type]time.Duration),
	}

	// Apply all options
	for _, opt := range options {
		opts = opt(opts)
	}

	// Forward options to the factory
	factory.customResync = opts.customResync
	factory.tweakListOptions = opts.tweakListOptions
	factory.namespace = opts.namespace

	return factory
}

// Start initializes all requested informers.
func (f *sharedScopedInformerFactory) Start(stopCh <-chan struct{}) {
	f.lock.Lock()
	defer f.lock.Unlock()

	for informerType, informer := range f.informers {
		if !f.startedInformers[informerType] {
			go informer.Run(stopCh)
			f.startedInformers[informerType] = true
		}
	}
}

// WaitForCacheSync waits for all started informers' cache were synced.
func (f *sharedScopedInformerFactory) WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool {
	informers := func() map[reflect.Type]cache.SharedIndexInformer {
		f.lock.Lock()
		defer f.lock.Unlock()

		informers := map[reflect.Type]cache.SharedIndexInformer{}
		for informerType, informer := range f.informers {
			if f.startedInformers[informerType] {
				informers[informerType] = informer
			}
		}
		return informers
	}()

	res := map[reflect.Type]bool{}
	for informType, informer := range informers {
		res[informType] = cache.WaitForCacheSync(stopCh, informer.HasSynced)
	}
	return res
}

// InformerFor returns the SharedIndexInformer for obj using an internal
// client.
func (f *sharedScopedInformerFactory) InformerFor(obj runtime.Object, newFunc internalinterfaces.NewScopedInformerFunc) cache.SharedIndexInformer {
	f.lock.Lock()
	defer f.lock.Unlock()

	informerType := reflect.TypeOf(obj)
	informer, exists := f.informers[informerType]
	if exists {
		return informer
	}

	resyncPeriod, exists := f.customResync[informerType]
	if !exists {
		resyncPeriod = f.defaultResync
	}

	informer = newFunc(f.client, resyncPeriod)
	f.informers[informerType] = informer

	return informer
}

// SharedScopedInformerFactory provides shared informers for resources in all known
// API group versions, scoped to one workspace.
type SharedScopedInformerFactory interface {
	internalinterfaces.SharedScopedInformerFactory
	ForResource(resource schema.GroupVersionResource) (GenericInformer, error)
	WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool

	Apiresource() apiresourceinformers.Interface
	Apis() apisinformers.Interface
	Core() coreinformers.Interface
	Scheduling() schedulinginformers.Interface
	Tenancy() tenancyinformers.Interface
	Topology() topologyinformers.Interface
	Workload() workloadinformers.Interface
}

func (f *sharedScopedInformerFactory) Apiresource() apiresourceinformers.Interface {
	return apiresourceinformers.NewScoped(f, f.namespace, f.tweakListOptions)
}

func (f *sharedScopedInformerFactory) Apis() apisinformers.Interface {
	return apisinformers.NewScoped(f, f.namespace, f.tweakListOptions)
}

func (f *sharedScopedInformerFactory) Core() coreinformers.Interface {
	return coreinformers.NewScoped(f, f.namespace, f.tweakListOptions)
}

func (f *sharedScopedInformerFactory) Scheduling() schedulinginformers.Interface {
	return schedulinginformers.NewScoped(f, f.namespace, f.tweakListOptions)
}

func (f *sharedScopedInformerFactory) Tenancy() tenancyinformers.Interface {
	return tenancyinformers.NewScoped(f, f.namespace, f.tweakListOptions)
}

func (f *sharedScopedInformerFactory) Topology() topologyinformers.Interface {
	return topologyinformers.NewScoped(f, f.namespace, f.tweakListOptions)
}

func (f *sharedScopedInformerFactory) Workload() workloadinformers.Interface {
	return workloadinformers.NewScoped(f, f.namespace, f.tweakListOptions)
}
