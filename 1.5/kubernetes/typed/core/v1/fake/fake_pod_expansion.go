/*
Copyright 2014 The Kubernetes Authors.

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
	"github.com/Icelandair/client-go/1.5/pkg/api/v1"
	"github.com/Icelandair/client-go/1.5/pkg/apis/policy/v1alpha1"
	"github.com/Icelandair/client-go/1.5/rest"
	"github.com/Icelandair/client-go/1.5/testing"
)

func (c *FakePods) Bind(binding *v1.Binding) error {
	action := testing.CreateActionImpl{}
	action.Verb = "create"
	action.Resource = podsResource
	action.Subresource = "bindings"
	action.Object = binding

	_, err := c.Fake.Invokes(action, binding)
	return err
}

func (c *FakePods) GetLogs(name string, opts *v1.PodLogOptions) *rest.Request {
	action := testing.GenericActionImpl{}
	action.Verb = "get"
	action.Namespace = c.ns
	action.Resource = podsResource
	action.Subresource = "logs"
	action.Value = opts

	_, _ = c.Fake.Invokes(action, &v1.Pod{})
	return &rest.Request{}
}

func (c *FakePods) Evict(eviction *v1alpha1.Eviction) error {
	action := testing.CreateActionImpl{}
	action.Verb = "create"
	action.Resource = podsResource
	action.Subresource = "eviction"
	action.Object = eviction

	_, err := c.Fake.Invokes(action, eviction)
	return err
}
