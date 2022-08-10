/*
Copyright The Kubernetes Authors.

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
	admiralv1 "github.com/istio-ecosystem/admiral/admiral/pkg/apis/admiral/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRoutingPolicies implements RoutingPolicyInterface
type FakeRoutingPolicies struct {
	Fake *FakeAdmiralV1
	ns   string
}

var routingpoliciesResource = schema.GroupVersionResource{Group: "admiral.io", Version: "v1", Resource: "routingpolicies"}

var routingpoliciesKind = schema.GroupVersionKind{Group: "admiral.io", Version: "v1", Kind: "RoutingPolicy"}

// Get takes name of the routingPolicy, and returns the corresponding routingPolicy object, and an error if there is any.
func (c *FakeRoutingPolicies) Get(name string, options v1.GetOptions) (result *admiralv1.RoutingPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(routingpoliciesResource, c.ns, name), &admiralv1.RoutingPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*admiralv1.RoutingPolicy), err
}

// List takes label and field selectors, and returns the list of RoutingPolicies that match those selectors.
func (c *FakeRoutingPolicies) List(opts v1.ListOptions) (result *admiralv1.RoutingPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(routingpoliciesResource, routingpoliciesKind, c.ns, opts), &admiralv1.RoutingPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &admiralv1.RoutingPolicyList{ListMeta: obj.(*admiralv1.RoutingPolicyList).ListMeta}
	for _, item := range obj.(*admiralv1.RoutingPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested routingPolicies.
func (c *FakeRoutingPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(routingpoliciesResource, c.ns, opts))

}

// Create takes the representation of a routingPolicy and creates it.  Returns the server's representation of the routingPolicy, and an error, if there is any.
func (c *FakeRoutingPolicies) Create(routingPolicy *admiralv1.RoutingPolicy) (result *admiralv1.RoutingPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(routingpoliciesResource, c.ns, routingPolicy), &admiralv1.RoutingPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*admiralv1.RoutingPolicy), err
}

// Update takes the representation of a routingPolicy and updates it. Returns the server's representation of the routingPolicy, and an error, if there is any.
func (c *FakeRoutingPolicies) Update(routingPolicy *admiralv1.RoutingPolicy) (result *admiralv1.RoutingPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(routingpoliciesResource, c.ns, routingPolicy), &admiralv1.RoutingPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*admiralv1.RoutingPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRoutingPolicies) UpdateStatus(routingPolicy *admiralv1.RoutingPolicy) (*admiralv1.RoutingPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(routingpoliciesResource, "status", c.ns, routingPolicy), &admiralv1.RoutingPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*admiralv1.RoutingPolicy), err
}

// Delete takes name of the routingPolicy and deletes it. Returns an error if one occurs.
func (c *FakeRoutingPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(routingpoliciesResource, c.ns, name), &admiralv1.RoutingPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRoutingPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(routingpoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &admiralv1.RoutingPolicyList{})
	return err
}

// Patch applies the patch and returns the patched routingPolicy.
func (c *FakeRoutingPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *admiralv1.RoutingPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(routingpoliciesResource, c.ns, name, pt, data, subresources...), &admiralv1.RoutingPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*admiralv1.RoutingPolicy), err
}