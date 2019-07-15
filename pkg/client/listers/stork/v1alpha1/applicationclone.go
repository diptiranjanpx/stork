/*
Copyright 2018 Openstorage.org

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

package v1alpha1

import (
	v1alpha1 "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ApplicationCloneLister helps list ApplicationClones.
type ApplicationCloneLister interface {
	// List lists all ApplicationClones in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ApplicationClone, err error)
	// ApplicationClones returns an object that can list and get ApplicationClones.
	ApplicationClones(namespace string) ApplicationCloneNamespaceLister
	ApplicationCloneListerExpansion
}

// applicationCloneLister implements the ApplicationCloneLister interface.
type applicationCloneLister struct {
	indexer cache.Indexer
}

// NewApplicationCloneLister returns a new ApplicationCloneLister.
func NewApplicationCloneLister(indexer cache.Indexer) ApplicationCloneLister {
	return &applicationCloneLister{indexer: indexer}
}

// List lists all ApplicationClones in the indexer.
func (s *applicationCloneLister) List(selector labels.Selector) (ret []*v1alpha1.ApplicationClone, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApplicationClone))
	})
	return ret, err
}

// ApplicationClones returns an object that can list and get ApplicationClones.
func (s *applicationCloneLister) ApplicationClones(namespace string) ApplicationCloneNamespaceLister {
	return applicationCloneNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApplicationCloneNamespaceLister helps list and get ApplicationClones.
type ApplicationCloneNamespaceLister interface {
	// List lists all ApplicationClones in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ApplicationClone, err error)
	// Get retrieves the ApplicationClone from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ApplicationClone, error)
	ApplicationCloneNamespaceListerExpansion
}

// applicationCloneNamespaceLister implements the ApplicationCloneNamespaceLister
// interface.
type applicationCloneNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApplicationClones in the indexer for a given namespace.
func (s applicationCloneNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApplicationClone, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApplicationClone))
	})
	return ret, err
}

// Get retrieves the ApplicationClone from the indexer for a given namespace and name.
func (s applicationCloneNamespaceLister) Get(name string) (*v1alpha1.ApplicationClone, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("applicationclone"), name)
	}
	return obj.(*v1alpha1.ApplicationClone), nil
}