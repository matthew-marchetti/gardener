// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	garden "github.com/gardener/gardener/pkg/apis/garden"
	scheme "github.com/gardener/gardener/pkg/client/garden/clientset/internalversion/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BackupInfrastructuresGetter has a method to return a BackupInfrastructureInterface.
// A group's client should implement this interface.
type BackupInfrastructuresGetter interface {
	BackupInfrastructures(namespace string) BackupInfrastructureInterface
}

// BackupInfrastructureInterface has methods to work with BackupInfrastructure resources.
type BackupInfrastructureInterface interface {
	Create(*garden.BackupInfrastructure) (*garden.BackupInfrastructure, error)
	Update(*garden.BackupInfrastructure) (*garden.BackupInfrastructure, error)
	UpdateStatus(*garden.BackupInfrastructure) (*garden.BackupInfrastructure, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*garden.BackupInfrastructure, error)
	List(opts v1.ListOptions) (*garden.BackupInfrastructureList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *garden.BackupInfrastructure, err error)
	BackupInfrastructureExpansion
}

// backupInfrastructures implements BackupInfrastructureInterface
type backupInfrastructures struct {
	client rest.Interface
	ns     string
}

// newBackupInfrastructures returns a BackupInfrastructures
func newBackupInfrastructures(c *GardenClient, namespace string) *backupInfrastructures {
	return &backupInfrastructures{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the backupInfrastructure, and returns the corresponding backupInfrastructure object, and an error if there is any.
func (c *backupInfrastructures) Get(name string, options v1.GetOptions) (result *garden.BackupInfrastructure, err error) {
	result = &garden.BackupInfrastructure{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backupinfrastructures").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BackupInfrastructures that match those selectors.
func (c *backupInfrastructures) List(opts v1.ListOptions) (result *garden.BackupInfrastructureList, err error) {
	result = &garden.BackupInfrastructureList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backupinfrastructures").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested backupInfrastructures.
func (c *backupInfrastructures) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("backupinfrastructures").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a backupInfrastructure and creates it.  Returns the server's representation of the backupInfrastructure, and an error, if there is any.
func (c *backupInfrastructures) Create(backupInfrastructure *garden.BackupInfrastructure) (result *garden.BackupInfrastructure, err error) {
	result = &garden.BackupInfrastructure{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("backupinfrastructures").
		Body(backupInfrastructure).
		Do().
		Into(result)
	return
}

// Update takes the representation of a backupInfrastructure and updates it. Returns the server's representation of the backupInfrastructure, and an error, if there is any.
func (c *backupInfrastructures) Update(backupInfrastructure *garden.BackupInfrastructure) (result *garden.BackupInfrastructure, err error) {
	result = &garden.BackupInfrastructure{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backupinfrastructures").
		Name(backupInfrastructure.Name).
		Body(backupInfrastructure).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *backupInfrastructures) UpdateStatus(backupInfrastructure *garden.BackupInfrastructure) (result *garden.BackupInfrastructure, err error) {
	result = &garden.BackupInfrastructure{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backupinfrastructures").
		Name(backupInfrastructure.Name).
		SubResource("status").
		Body(backupInfrastructure).
		Do().
		Into(result)
	return
}

// Delete takes name of the backupInfrastructure and deletes it. Returns an error if one occurs.
func (c *backupInfrastructures) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backupinfrastructures").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *backupInfrastructures) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backupinfrastructures").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched backupInfrastructure.
func (c *backupInfrastructures) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *garden.BackupInfrastructure, err error) {
	result = &garden.BackupInfrastructure{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("backupinfrastructures").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
