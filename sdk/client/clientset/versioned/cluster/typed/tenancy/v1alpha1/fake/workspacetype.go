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

package v1alpha1

import (
	"context"

	"github.com/kcp-dev/logicalcluster/v3"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/testing"

	tenancyv1alpha1 "github.com/kcp-dev/kcp/sdk/apis/tenancy/v1alpha1"
	tenancyv1alpha1client "github.com/kcp-dev/kcp/sdk/client/clientset/versioned/typed/tenancy/v1alpha1"
)

var workspaceTypesResource = schema.GroupVersionResource{Group: "tenancy.kcp.io", Version: "v1alpha1", Resource: "workspacetypes"}
var workspaceTypesKind = schema.GroupVersionKind{Group: "tenancy.kcp.io", Version: "v1alpha1", Kind: "WorkspaceType"}

type workspaceTypesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *workspaceTypesClusterClient) Cluster(clusterPath logicalcluster.Path) tenancyv1alpha1client.WorkspaceTypeInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &workspaceTypesClient{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of WorkspaceTypes that match those selectors across all clusters.
func (c *workspaceTypesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*tenancyv1alpha1.WorkspaceTypeList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(workspaceTypesResource, workspaceTypesKind, logicalcluster.Wildcard, opts), &tenancyv1alpha1.WorkspaceTypeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &tenancyv1alpha1.WorkspaceTypeList{ListMeta: obj.(*tenancyv1alpha1.WorkspaceTypeList).ListMeta}
	for _, item := range obj.(*tenancyv1alpha1.WorkspaceTypeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested WorkspaceTypes across all clusters.
func (c *workspaceTypesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(workspaceTypesResource, logicalcluster.Wildcard, opts))
}

type workspaceTypesClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *workspaceTypesClient) Create(ctx context.Context, workspaceType *tenancyv1alpha1.WorkspaceType, opts metav1.CreateOptions) (*tenancyv1alpha1.WorkspaceType, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(workspaceTypesResource, c.ClusterPath, workspaceType), &tenancyv1alpha1.WorkspaceType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*tenancyv1alpha1.WorkspaceType), err
}

func (c *workspaceTypesClient) Update(ctx context.Context, workspaceType *tenancyv1alpha1.WorkspaceType, opts metav1.UpdateOptions) (*tenancyv1alpha1.WorkspaceType, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(workspaceTypesResource, c.ClusterPath, workspaceType), &tenancyv1alpha1.WorkspaceType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*tenancyv1alpha1.WorkspaceType), err
}

func (c *workspaceTypesClient) UpdateStatus(ctx context.Context, workspaceType *tenancyv1alpha1.WorkspaceType, opts metav1.UpdateOptions) (*tenancyv1alpha1.WorkspaceType, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(workspaceTypesResource, c.ClusterPath, "status", workspaceType), &tenancyv1alpha1.WorkspaceType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*tenancyv1alpha1.WorkspaceType), err
}

func (c *workspaceTypesClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(workspaceTypesResource, c.ClusterPath, name, opts), &tenancyv1alpha1.WorkspaceType{})
	return err
}

func (c *workspaceTypesClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(workspaceTypesResource, c.ClusterPath, listOpts)

	_, err := c.Fake.Invokes(action, &tenancyv1alpha1.WorkspaceTypeList{})
	return err
}

func (c *workspaceTypesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*tenancyv1alpha1.WorkspaceType, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(workspaceTypesResource, c.ClusterPath, name), &tenancyv1alpha1.WorkspaceType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*tenancyv1alpha1.WorkspaceType), err
}

// List takes label and field selectors, and returns the list of WorkspaceTypes that match those selectors.
func (c *workspaceTypesClient) List(ctx context.Context, opts metav1.ListOptions) (*tenancyv1alpha1.WorkspaceTypeList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(workspaceTypesResource, workspaceTypesKind, c.ClusterPath, opts), &tenancyv1alpha1.WorkspaceTypeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &tenancyv1alpha1.WorkspaceTypeList{ListMeta: obj.(*tenancyv1alpha1.WorkspaceTypeList).ListMeta}
	for _, item := range obj.(*tenancyv1alpha1.WorkspaceTypeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *workspaceTypesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(workspaceTypesResource, c.ClusterPath, opts))
}

func (c *workspaceTypesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*tenancyv1alpha1.WorkspaceType, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(workspaceTypesResource, c.ClusterPath, name, pt, data, subresources...), &tenancyv1alpha1.WorkspaceType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*tenancyv1alpha1.WorkspaceType), err
}
