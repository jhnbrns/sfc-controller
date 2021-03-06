// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/ligato/sfc-controller/plugins/k8scrd/pkg/apis/sfccontroller/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NetworkNodeOverlayLister helps list NetworkNodeOverlays.
type NetworkNodeOverlayLister interface {
	// List lists all NetworkNodeOverlays in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.NetworkNodeOverlay, err error)
	// NetworkNodeOverlays returns an object that can list and get NetworkNodeOverlays.
	NetworkNodeOverlays(namespace string) NetworkNodeOverlayNamespaceLister
	NetworkNodeOverlayListerExpansion
}

// networkNodeOverlayLister implements the NetworkNodeOverlayLister interface.
type networkNodeOverlayLister struct {
	indexer cache.Indexer
}

// NewNetworkNodeOverlayLister returns a new NetworkNodeOverlayLister.
func NewNetworkNodeOverlayLister(indexer cache.Indexer) NetworkNodeOverlayLister {
	return &networkNodeOverlayLister{indexer: indexer}
}

// List lists all NetworkNodeOverlays in the indexer.
func (s *networkNodeOverlayLister) List(selector labels.Selector) (ret []*v1alpha1.NetworkNodeOverlay, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetworkNodeOverlay))
	})
	return ret, err
}

// NetworkNodeOverlays returns an object that can list and get NetworkNodeOverlays.
func (s *networkNodeOverlayLister) NetworkNodeOverlays(namespace string) NetworkNodeOverlayNamespaceLister {
	return networkNodeOverlayNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NetworkNodeOverlayNamespaceLister helps list and get NetworkNodeOverlays.
type NetworkNodeOverlayNamespaceLister interface {
	// List lists all NetworkNodeOverlays in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.NetworkNodeOverlay, err error)
	// Get retrieves the NetworkNodeOverlay from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.NetworkNodeOverlay, error)
	NetworkNodeOverlayNamespaceListerExpansion
}

// networkNodeOverlayNamespaceLister implements the NetworkNodeOverlayNamespaceLister
// interface.
type networkNodeOverlayNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NetworkNodeOverlays in the indexer for a given namespace.
func (s networkNodeOverlayNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NetworkNodeOverlay, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetworkNodeOverlay))
	})
	return ret, err
}

// Get retrieves the NetworkNodeOverlay from the indexer for a given namespace and name.
func (s networkNodeOverlayNamespaceLister) Get(name string) (*v1alpha1.NetworkNodeOverlay, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("networknodeoverlay"), name)
	}
	return obj.(*v1alpha1.NetworkNodeOverlay), nil
}
