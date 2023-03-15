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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1alpha1 "github.com/kcp-dev/kcp/sdk/apis/tenancy/v1alpha1"
)

// FakeWorkspaceTypes implements WorkspaceTypeInterface
type FakeWorkspaceTypes struct {
	Fake *FakeTenancyV1alpha1
}

var workspacetypesResource = schema.GroupVersionResource{Group: "tenancy.kcp.io", Version: "v1alpha1", Resource: "workspacetypes"}

var workspacetypesKind = schema.GroupVersionKind{Group: "tenancy.kcp.io", Version: "v1alpha1", Kind: "WorkspaceType"}

// Get takes name of the workspaceType, and returns the corresponding workspaceType object, and an error if there is any.
func (c *FakeWorkspaceTypes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.WorkspaceType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(workspacetypesResource, name), &v1alpha1.WorkspaceType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorkspaceType), err
}

// List takes label and field selectors, and returns the list of WorkspaceTypes that match those selectors.
func (c *FakeWorkspaceTypes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.WorkspaceTypeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(workspacetypesResource, workspacetypesKind, opts), &v1alpha1.WorkspaceTypeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.WorkspaceTypeList{ListMeta: obj.(*v1alpha1.WorkspaceTypeList).ListMeta}
	for _, item := range obj.(*v1alpha1.WorkspaceTypeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested workspaceTypes.
func (c *FakeWorkspaceTypes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(workspacetypesResource, opts))
}

// Create takes the representation of a workspaceType and creates it.  Returns the server's representation of the workspaceType, and an error, if there is any.
func (c *FakeWorkspaceTypes) Create(ctx context.Context, workspaceType *v1alpha1.WorkspaceType, opts v1.CreateOptions) (result *v1alpha1.WorkspaceType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(workspacetypesResource, workspaceType), &v1alpha1.WorkspaceType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorkspaceType), err
}

// Update takes the representation of a workspaceType and updates it. Returns the server's representation of the workspaceType, and an error, if there is any.
func (c *FakeWorkspaceTypes) Update(ctx context.Context, workspaceType *v1alpha1.WorkspaceType, opts v1.UpdateOptions) (result *v1alpha1.WorkspaceType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(workspacetypesResource, workspaceType), &v1alpha1.WorkspaceType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorkspaceType), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeWorkspaceTypes) UpdateStatus(ctx context.Context, workspaceType *v1alpha1.WorkspaceType, opts v1.UpdateOptions) (*v1alpha1.WorkspaceType, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(workspacetypesResource, "status", workspaceType), &v1alpha1.WorkspaceType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorkspaceType), err
}

// Delete takes name of the workspaceType and deletes it. Returns an error if one occurs.
func (c *FakeWorkspaceTypes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(workspacetypesResource, name, opts), &v1alpha1.WorkspaceType{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWorkspaceTypes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(workspacetypesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.WorkspaceTypeList{})
	return err
}

// Patch applies the patch and returns the patched workspaceType.
func (c *FakeWorkspaceTypes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.WorkspaceType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(workspacetypesResource, name, pt, data, subresources...), &v1alpha1.WorkspaceType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WorkspaceType), err
}
