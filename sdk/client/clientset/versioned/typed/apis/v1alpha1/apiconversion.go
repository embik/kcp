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

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1alpha1 "github.com/kcp-dev/kcp/sdk/apis/apis/v1alpha1"
	scheme "github.com/kcp-dev/kcp/sdk/client/clientset/versioned/scheme"
)

// APIConversionsGetter has a method to return a APIConversionInterface.
// A group's client should implement this interface.
type APIConversionsGetter interface {
	APIConversions() APIConversionInterface
}

// APIConversionInterface has methods to work with APIConversion resources.
type APIConversionInterface interface {
	Create(ctx context.Context, aPIConversion *v1alpha1.APIConversion, opts v1.CreateOptions) (*v1alpha1.APIConversion, error)
	Update(ctx context.Context, aPIConversion *v1alpha1.APIConversion, opts v1.UpdateOptions) (*v1alpha1.APIConversion, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.APIConversion, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.APIConversionList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.APIConversion, err error)
	APIConversionExpansion
}

// aPIConversions implements APIConversionInterface
type aPIConversions struct {
	client rest.Interface
}

// newAPIConversions returns a APIConversions
func newAPIConversions(c *ApisV1alpha1Client) *aPIConversions {
	return &aPIConversions{
		client: c.RESTClient(),
	}
}

// Get takes name of the aPIConversion, and returns the corresponding aPIConversion object, and an error if there is any.
func (c *aPIConversions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.APIConversion, err error) {
	result = &v1alpha1.APIConversion{}
	err = c.client.Get().
		Resource("apiconversions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of APIConversions that match those selectors.
func (c *aPIConversions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.APIConversionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.APIConversionList{}
	err = c.client.Get().
		Resource("apiconversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested aPIConversions.
func (c *aPIConversions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("apiconversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a aPIConversion and creates it.  Returns the server's representation of the aPIConversion, and an error, if there is any.
func (c *aPIConversions) Create(ctx context.Context, aPIConversion *v1alpha1.APIConversion, opts v1.CreateOptions) (result *v1alpha1.APIConversion, err error) {
	result = &v1alpha1.APIConversion{}
	err = c.client.Post().
		Resource("apiconversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aPIConversion).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a aPIConversion and updates it. Returns the server's representation of the aPIConversion, and an error, if there is any.
func (c *aPIConversions) Update(ctx context.Context, aPIConversion *v1alpha1.APIConversion, opts v1.UpdateOptions) (result *v1alpha1.APIConversion, err error) {
	result = &v1alpha1.APIConversion{}
	err = c.client.Put().
		Resource("apiconversions").
		Name(aPIConversion.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aPIConversion).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the aPIConversion and deletes it. Returns an error if one occurs.
func (c *aPIConversions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("apiconversions").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *aPIConversions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("apiconversions").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched aPIConversion.
func (c *aPIConversions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.APIConversion, err error) {
	result = &v1alpha1.APIConversion{}
	err = c.client.Patch(pt).
		Resource("apiconversions").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
