// Copyright 2017 uSwitch
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
package k8s

import (
	"context"
	"k8s.io/api/core/v1"
)

type RoleFinder interface {
	// Finds a uncompleted pod from its IP address
	FindRoleFromIP(ctx context.Context, ip string) (string, error)
}

type PodAnnouncer interface {
	// Will receive a Pod whenever there's a change/addition for a Pod with a role.
	Pods() <-chan *v1.Pod
	// Return whether there are still uncompleted pods in the specified role
	IsActivePodsForRole(role string) (bool, error)
}

type NamespaceFinder interface {
	FindNamespace(ctx context.Context, name string) (*v1.Namespace, error)
}
