// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package policy

import (
	"yunion.io/x/onecloud/pkg/util/rbacutils"
)

var (
	predefinedDefaultPolicies = []rbacutils.SRbacPolicy{
		{
			Auth:  true,
			Scope: rbacutils.ScopeSystem,
			Rules: []rbacutils.SRbacRule{
				{
					Resource: "tasks",
					Action:   PolicyActionPerform,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "yunionagent",
					Resource: "notices",
					Action:   PolicyActionList,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "yunionagent",
					Resource: "readmarks",
					Action:   PolicyActionList,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "yunionagent",
					Resource: "readmarks",
					Action:   PolicyActionCreate,
					Result:   rbacutils.Allow,
				},
			},
		},
		{
			Auth:  true,
			Scope: rbacutils.ScopeProject,
			Rules: []rbacutils.SRbacRule{
				{
					Resource: "tasks",
					Action:   PolicyActionPerform,
					Result:   rbacutils.Allow,
				},
				{
					// usages for any services
					// Service:  "compute",
					Resource: "usages",
					Action:   PolicyActionGet,
					Result:   rbacutils.Allow,
				},
			},
		},
		{
			// meta服务 dbinstance_skus列表不需要认证
			Auth:  false,
			Scope: rbacutils.ScopeSystem,
			Rules: []rbacutils.SRbacRule{
				{
					Service:  "yunionmeta",
					Resource: "dbinstance_skus",
					Action:   PolicyActionGet,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "yunionmeta",
					Resource: "dbinstance_skus",
					Action:   PolicyActionList,
					Result:   rbacutils.Allow,
				},
			},
		},
		{
			// for domain
			Auth:  true,
			Scope: rbacutils.ScopeDomain,
			Rules: []rbacutils.SRbacRule{
				{
					// usages for any services
					// Service:  "compute",
					Resource: "usages",
					Action:   PolicyActionGet,
					Result:   rbacutils.Allow,
				},
			},
		},
	}
)

func AppendDefaultPolicies(policies []rbacutils.SRbacPolicy) {
	predefinedDefaultPolicies = append(predefinedDefaultPolicies, policies...)
}
