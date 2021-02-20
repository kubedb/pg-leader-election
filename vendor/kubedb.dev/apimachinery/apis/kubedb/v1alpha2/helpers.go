/*
Copyright AppsCode Inc. and Contributors

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

package v1alpha2

import (
	"fmt"

	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/labels"
	appslister "k8s.io/client-go/listers/apps/v1"
	apps_util "kmodules.xyz/client-go/apps/v1"
	ofst "kmodules.xyz/offshoot-api/api/v1"
)

func checkReplicas(lister appslister.StatefulSetNamespaceLister, selector labels.Selector, expectedItems int) (bool, string, error) {
	items, err := lister.List(selector)
	if err != nil {
		return false, "", err
	}
	if len(items) < expectedItems {
		return false, fmt.Sprintf("All StatefulSets are not available. Desire number of StatefulSet: %d, Available: %d", expectedItems, len(items)), nil
	}

	// return isReplicasReady, message, error
	ready, msg := apps_util.StatefulSetsAreReady(items)
	return ready, msg, nil
}

// HasServiceTemplate returns "true" if the desired serviceTemplate provided in "aliaS" is present in the serviceTemplate list.
// Otherwise, it returns "false".
func HasServiceTemplate(templates []NamedServiceTemplateSpec, alias ServiceAlias) bool {
	for i := range templates {
		if templates[i].Alias == alias {
			return true
		}
	}
	return false
}

// GetServiceTemplate returns a pointer to the desired serviceTemplate referred by "aliaS". Otherwise, it returns nil.
func GetServiceTemplate(templates []NamedServiceTemplateSpec, alias ServiceAlias) ofst.ServiceTemplateSpec {
	for i := range templates {
		c := templates[i]
		if c.Alias == alias {
			return c.ServiceTemplateSpec
		}
	}
	return ofst.ServiceTemplateSpec{}
}

func SetDefaultResourceLimits(req *core.ResourceRequirements, defaultLimits core.ResourceList) {
	// if request is set,
	//		- limit set:
	//			- return max(limit,request)
	// else if limit set:
	//		- return limit
	// else
	//		- return default
	fn := func(name core.ResourceName, defaultValue resource.Quantity) resource.Quantity {
		if r, ok := req.Requests[name]; ok {
			// l is greater than r == 1.
			if l, exist := req.Limits[name]; exist && l.Cmp(r) == 1 {
				return l
			}
			return r
		}
		if l, ok := req.Limits[name]; ok {
			return l
		}
		return defaultValue
	}

	if req.Limits == nil {
		req.Limits = core.ResourceList{}
	}
	if req.Requests == nil {
		req.Requests = core.ResourceList{}
	}

	// for: cpu & memory
	//		- calculate limit
	//		- if request not set, set to limit
	for resourceName := range defaultLimits {
		req.Limits[resourceName] = fn(resourceName, defaultLimits[resourceName])

		if _, ok := req.Requests[resourceName]; !ok {
			req.Requests[resourceName] = req.Limits[resourceName]
		}
	}
}
