/*
Copyright 2019 The Seldon Authors.

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

package v1alpha3

import (
	v1alpha3 "github.com/seldonio/seldon-core/operator/apis/machinelearning/v1alpha3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SeldonDeploymentLister helps list SeldonDeployments.
type SeldonDeploymentLister interface {
	// List lists all SeldonDeployments in the indexer.
	List(selector labels.Selector) (ret []*v1alpha3.SeldonDeployment, err error)
	// SeldonDeployments returns an object that can list and get SeldonDeployments.
	SeldonDeployments(namespace string) SeldonDeploymentNamespaceLister
	SeldonDeploymentListerExpansion
}

// seldonDeploymentLister implements the SeldonDeploymentLister interface.
type seldonDeploymentLister struct {
	indexer cache.Indexer
}

// NewSeldonDeploymentLister returns a new SeldonDeploymentLister.
func NewSeldonDeploymentLister(indexer cache.Indexer) SeldonDeploymentLister {
	return &seldonDeploymentLister{indexer: indexer}
}

// List lists all SeldonDeployments in the indexer.
func (s *seldonDeploymentLister) List(selector labels.Selector) (ret []*v1alpha3.SeldonDeployment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha3.SeldonDeployment))
	})
	return ret, err
}

// SeldonDeployments returns an object that can list and get SeldonDeployments.
func (s *seldonDeploymentLister) SeldonDeployments(namespace string) SeldonDeploymentNamespaceLister {
	return seldonDeploymentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SeldonDeploymentNamespaceLister helps list and get SeldonDeployments.
type SeldonDeploymentNamespaceLister interface {
	// List lists all SeldonDeployments in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha3.SeldonDeployment, err error)
	// Get retrieves the SeldonDeployment from the indexer for a given namespace and name.
	Get(name string) (*v1alpha3.SeldonDeployment, error)
	SeldonDeploymentNamespaceListerExpansion
}

// seldonDeploymentNamespaceLister implements the SeldonDeploymentNamespaceLister
// interface.
type seldonDeploymentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SeldonDeployments in the indexer for a given namespace.
func (s seldonDeploymentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha3.SeldonDeployment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha3.SeldonDeployment))
	})
	return ret, err
}

// Get retrieves the SeldonDeployment from the indexer for a given namespace and name.
func (s seldonDeploymentNamespaceLister) Get(name string) (*v1alpha3.SeldonDeployment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha3.Resource("seldondeployment"), name)
	}
	return obj.(*v1alpha3.SeldonDeployment), nil
}
