/*
Copyright 2018 caicloud authors. All rights reserved.
*/

package v1alpha1

import (
	scheme "github.com/caicloud/clientset/kubernetes/scheme"
	v1alpha1 "github.com/caicloud/clientset/pkg/apis/resource/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// StorageServicesGetter has a method to return a StorageServiceInterface.
// A group's client should implement this interface.
type StorageServicesGetter interface {
	StorageServices() StorageServiceInterface
}

// StorageServiceInterface has methods to work with StorageService resources.
type StorageServiceInterface interface {
	Create(*v1alpha1.StorageService) (*v1alpha1.StorageService, error)
	Update(*v1alpha1.StorageService) (*v1alpha1.StorageService, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.StorageService, error)
	List(opts v1.ListOptions) (*v1alpha1.StorageServiceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StorageService, err error)
	StorageServiceExpansion
}

// storageServices implements StorageServiceInterface
type storageServices struct {
	client rest.Interface
}

// newStorageServices returns a StorageServices
func newStorageServices(c *ResourceV1alpha1Client) *storageServices {
	return &storageServices{
		client: c.RESTClient(),
	}
}

// Create takes the representation of a storageService and creates it.  Returns the server's representation of the storageService, and an error, if there is any.
func (c *storageServices) Create(storageService *v1alpha1.StorageService) (result *v1alpha1.StorageService, err error) {
	result = &v1alpha1.StorageService{}
	err = c.client.Post().
		Resource("storageservices").
		Body(storageService).
		Do().
		Into(result)
	return
}

// Update takes the representation of a storageService and updates it. Returns the server's representation of the storageService, and an error, if there is any.
func (c *storageServices) Update(storageService *v1alpha1.StorageService) (result *v1alpha1.StorageService, err error) {
	result = &v1alpha1.StorageService{}
	err = c.client.Put().
		Resource("storageservices").
		Name(storageService.Name).
		Body(storageService).
		Do().
		Into(result)
	return
}

// Delete takes name of the storageService and deletes it. Returns an error if one occurs.
func (c *storageServices) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("storageservices").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *storageServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("storageservices").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the storageService, and returns the corresponding storageService object, and an error if there is any.
func (c *storageServices) Get(name string, options v1.GetOptions) (result *v1alpha1.StorageService, err error) {
	result = &v1alpha1.StorageService{}
	err = c.client.Get().
		Resource("storageservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of StorageServices that match those selectors.
func (c *storageServices) List(opts v1.ListOptions) (result *v1alpha1.StorageServiceList, err error) {
	result = &v1alpha1.StorageServiceList{}
	err = c.client.Get().
		Resource("storageservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested storageServices.
func (c *storageServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("storageservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched storageService.
func (c *storageServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StorageService, err error) {
	result = &v1alpha1.StorageService{}
	err = c.client.Patch(pt).
		Resource("storageservices").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}