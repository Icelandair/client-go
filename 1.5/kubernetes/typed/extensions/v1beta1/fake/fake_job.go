/*
Copyright 2016 The Kubernetes Authors.

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

package fake

import (
	api "github.com/Icelandair/client-go/1.5/pkg/api"
	unversioned "github.com/Icelandair/client-go/1.5/pkg/api/unversioned"
	v1beta1 "github.com/Icelandair/client-go/1.5/pkg/apis/extensions/v1beta1"
	labels "github.com/Icelandair/client-go/1.5/pkg/labels"
	watch "github.com/Icelandair/client-go/1.5/pkg/watch"
	testing "github.com/Icelandair/client-go/1.5/testing"
)

// FakeJobs implements JobInterface
type FakeJobs struct {
	Fake *FakeExtensions
	ns   string
}

var jobsResource = unversioned.GroupVersionResource{Group: "extensions", Version: "v1beta1", Resource: "jobs"}

func (c *FakeJobs) Create(job *v1beta1.Job) (result *v1beta1.Job, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(jobsResource, c.ns, job), &v1beta1.Job{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Job), err
}

func (c *FakeJobs) Update(job *v1beta1.Job) (result *v1beta1.Job, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(jobsResource, c.ns, job), &v1beta1.Job{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Job), err
}

func (c *FakeJobs) UpdateStatus(job *v1beta1.Job) (*v1beta1.Job, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(jobsResource, "status", c.ns, job), &v1beta1.Job{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Job), err
}

func (c *FakeJobs) Delete(name string, options *api.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(jobsResource, c.ns, name), &v1beta1.Job{})

	return err
}

func (c *FakeJobs) DeleteCollection(options *api.DeleteOptions, listOptions api.ListOptions) error {
	action := testing.NewDeleteCollectionAction(jobsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.JobList{})
	return err
}

func (c *FakeJobs) Get(name string) (result *v1beta1.Job, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(jobsResource, c.ns, name), &v1beta1.Job{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Job), err
}

func (c *FakeJobs) List(opts api.ListOptions) (result *v1beta1.JobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(jobsResource, c.ns, opts), &v1beta1.JobList{})

	if obj == nil {
		return nil, err
	}

	label := opts.LabelSelector
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.JobList{}
	for _, item := range obj.(*v1beta1.JobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested jobs.
func (c *FakeJobs) Watch(opts api.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(jobsResource, c.ns, opts))

}

// Patch applies the patch and returns the patched job.
func (c *FakeJobs) Patch(name string, pt api.PatchType, data []byte, subresources ...string) (result *v1beta1.Job, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(jobsResource, c.ns, name, data, subresources...), &v1beta1.Job{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Job), err
}