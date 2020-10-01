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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	jenkinsiov1 "github.com/jenkins-x/jx-api/v3/pkg/apis/jenkins.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePipelineStructures implements PipelineStructureInterface
type FakePipelineStructures struct {
	Fake *FakeJenkinsV1
	ns   string
}

var pipelinestructuresResource = schema.GroupVersionResource{Group: "jenkins.io", Version: "v1", Resource: "pipelinestructures"}

var pipelinestructuresKind = schema.GroupVersionKind{Group: "jenkins.io", Version: "v1", Kind: "PipelineStructure"}

// Get takes name of the pipelineStructure, and returns the corresponding pipelineStructure object, and an error if there is any.
func (c *FakePipelineStructures) Get(ctx context.Context, name string, options v1.GetOptions) (result *jenkinsiov1.PipelineStructure, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pipelinestructuresResource, c.ns, name), &jenkinsiov1.PipelineStructure{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.PipelineStructure), err
}

// List takes label and field selectors, and returns the list of PipelineStructures that match those selectors.
func (c *FakePipelineStructures) List(ctx context.Context, opts v1.ListOptions) (result *jenkinsiov1.PipelineStructureList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pipelinestructuresResource, pipelinestructuresKind, c.ns, opts), &jenkinsiov1.PipelineStructureList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &jenkinsiov1.PipelineStructureList{ListMeta: obj.(*jenkinsiov1.PipelineStructureList).ListMeta}
	for _, item := range obj.(*jenkinsiov1.PipelineStructureList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pipelineStructures.
func (c *FakePipelineStructures) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pipelinestructuresResource, c.ns, opts))

}

// Create takes the representation of a pipelineStructure and creates it.  Returns the server's representation of the pipelineStructure, and an error, if there is any.
func (c *FakePipelineStructures) Create(ctx context.Context, pipelineStructure *jenkinsiov1.PipelineStructure, opts v1.CreateOptions) (result *jenkinsiov1.PipelineStructure, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pipelinestructuresResource, c.ns, pipelineStructure), &jenkinsiov1.PipelineStructure{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.PipelineStructure), err
}

// Update takes the representation of a pipelineStructure and updates it. Returns the server's representation of the pipelineStructure, and an error, if there is any.
func (c *FakePipelineStructures) Update(ctx context.Context, pipelineStructure *jenkinsiov1.PipelineStructure, opts v1.UpdateOptions) (result *jenkinsiov1.PipelineStructure, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pipelinestructuresResource, c.ns, pipelineStructure), &jenkinsiov1.PipelineStructure{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.PipelineStructure), err
}

// Delete takes name of the pipelineStructure and deletes it. Returns an error if one occurs.
func (c *FakePipelineStructures) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pipelinestructuresResource, c.ns, name), &jenkinsiov1.PipelineStructure{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePipelineStructures) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pipelinestructuresResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &jenkinsiov1.PipelineStructureList{})
	return err
}

// Patch applies the patch and returns the patched pipelineStructure.
func (c *FakePipelineStructures) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *jenkinsiov1.PipelineStructure, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pipelinestructuresResource, c.ns, name, pt, data, subresources...), &jenkinsiov1.PipelineStructure{})

	if obj == nil {
		return nil, err
	}
	return obj.(*jenkinsiov1.PipelineStructure), err
}
