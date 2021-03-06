/*
Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

package internalversion

import (
	"time"

	core "github.com/gardener/gardener/pkg/apis/core"
	scheme "github.com/gardener/gardener/pkg/client/core/clientset/internalversion/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BackupBucketsGetter has a method to return a BackupBucketInterface.
// A group's client should implement this interface.
type BackupBucketsGetter interface {
	BackupBuckets() BackupBucketInterface
}

// BackupBucketInterface has methods to work with BackupBucket resources.
type BackupBucketInterface interface {
	Create(*core.BackupBucket) (*core.BackupBucket, error)
	Update(*core.BackupBucket) (*core.BackupBucket, error)
	UpdateStatus(*core.BackupBucket) (*core.BackupBucket, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*core.BackupBucket, error)
	List(opts v1.ListOptions) (*core.BackupBucketList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.BackupBucket, err error)
	BackupBucketExpansion
}

// backupBuckets implements BackupBucketInterface
type backupBuckets struct {
	client rest.Interface
}

// newBackupBuckets returns a BackupBuckets
func newBackupBuckets(c *CoreClient) *backupBuckets {
	return &backupBuckets{
		client: c.RESTClient(),
	}
}

// Get takes name of the backupBucket, and returns the corresponding backupBucket object, and an error if there is any.
func (c *backupBuckets) Get(name string, options v1.GetOptions) (result *core.BackupBucket, err error) {
	result = &core.BackupBucket{}
	err = c.client.Get().
		Resource("backupbuckets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BackupBuckets that match those selectors.
func (c *backupBuckets) List(opts v1.ListOptions) (result *core.BackupBucketList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &core.BackupBucketList{}
	err = c.client.Get().
		Resource("backupbuckets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested backupBuckets.
func (c *backupBuckets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("backupbuckets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a backupBucket and creates it.  Returns the server's representation of the backupBucket, and an error, if there is any.
func (c *backupBuckets) Create(backupBucket *core.BackupBucket) (result *core.BackupBucket, err error) {
	result = &core.BackupBucket{}
	err = c.client.Post().
		Resource("backupbuckets").
		Body(backupBucket).
		Do().
		Into(result)
	return
}

// Update takes the representation of a backupBucket and updates it. Returns the server's representation of the backupBucket, and an error, if there is any.
func (c *backupBuckets) Update(backupBucket *core.BackupBucket) (result *core.BackupBucket, err error) {
	result = &core.BackupBucket{}
	err = c.client.Put().
		Resource("backupbuckets").
		Name(backupBucket.Name).
		Body(backupBucket).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *backupBuckets) UpdateStatus(backupBucket *core.BackupBucket) (result *core.BackupBucket, err error) {
	result = &core.BackupBucket{}
	err = c.client.Put().
		Resource("backupbuckets").
		Name(backupBucket.Name).
		SubResource("status").
		Body(backupBucket).
		Do().
		Into(result)
	return
}

// Delete takes name of the backupBucket and deletes it. Returns an error if one occurs.
func (c *backupBuckets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("backupbuckets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *backupBuckets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("backupbuckets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched backupBucket.
func (c *backupBuckets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core.BackupBucket, err error) {
	result = &core.BackupBucket{}
	err = c.client.Patch(pt).
		Resource("backupbuckets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
