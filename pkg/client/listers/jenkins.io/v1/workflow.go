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
// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/jenkins-x/jx-api/v3/pkg/apis/jenkins.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// WorkflowLister helps list Workflows.
// All objects returned here must be treated as read-only.
type WorkflowLister interface {
	// List lists all Workflows in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Workflow, err error)
	// Workflows returns an object that can list and get Workflows.
	Workflows(namespace string) WorkflowNamespaceLister
	WorkflowListerExpansion
}

// workflowLister implements the WorkflowLister interface.
type workflowLister struct {
	indexer cache.Indexer
}

// NewWorkflowLister returns a new WorkflowLister.
func NewWorkflowLister(indexer cache.Indexer) WorkflowLister {
	return &workflowLister{indexer: indexer}
}

// List lists all Workflows in the indexer.
func (s *workflowLister) List(selector labels.Selector) (ret []*v1.Workflow, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Workflow))
	})
	return ret, err
}

// Workflows returns an object that can list and get Workflows.
func (s *workflowLister) Workflows(namespace string) WorkflowNamespaceLister {
	return workflowNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WorkflowNamespaceLister helps list and get Workflows.
// All objects returned here must be treated as read-only.
type WorkflowNamespaceLister interface {
	// List lists all Workflows in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Workflow, err error)
	// Get retrieves the Workflow from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Workflow, error)
	WorkflowNamespaceListerExpansion
}

// workflowNamespaceLister implements the WorkflowNamespaceLister
// interface.
type workflowNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Workflows in the indexer for a given namespace.
func (s workflowNamespaceLister) List(selector labels.Selector) (ret []*v1.Workflow, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Workflow))
	})
	return ret, err
}

// Get retrieves the Workflow from the indexer for a given namespace and name.
func (s workflowNamespaceLister) Get(name string) (*v1.Workflow, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("workflow"), name)
	}
	return obj.(*v1.Workflow), nil
}
