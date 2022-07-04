/*
Copyright 2022 The KCP Authors.

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

package shared

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/kcp-dev/logicalcluster"
	"github.com/martinlindhe/base36"

	"k8s.io/apimachinery/pkg/types"
)

const (
	NamespaceLocatorAnnotation = "kcp.dev/namespace-locator"
)

// NamespaceLocator stores a logical cluster and namespace and is used
// as the source for the mapped namespace name in a physical cluster.
type NamespaceLocator struct {
	SyncTarget SyncTargetLocator   `json:"syncTarget"`
	Workspace  logicalcluster.Name `json:"workspace,omitempty"`
	Namespace  string              `json:"namespace"`
}

type SyncTargetLocator struct {
	Path logicalcluster.Name `json:"path"`
	Name string              `json:"name"`
	UID  types.UID           `json:"uid"`
}

func NewNamespaceLocator(workspace, workloadClusterWorkspace logicalcluster.Name, workloadClusterUID types.UID, workloadLogicalClusterName, upstreamNamespace string) NamespaceLocator {
	return NamespaceLocator{
		SyncTarget: SyncTargetLocator{
			Path: syncTargetWorkspace,
			Name: workloadLogicalClusterName,
			UID:  syncTargetUID,
		},
		Workspace: workspace,
		Namespace: upstreamNamespace,
	}
}

func LocatorFromAnnotations(annotations map[string]string) (*NamespaceLocator, error) {
	annotation := annotations[NamespaceLocatorAnnotation]
	if len(annotation) == 0 {
		return nil, nil
	}
	var locator NamespaceLocator
	if err := json.Unmarshal([]byte(annotation), &locator); err != nil {
		return nil, err
	}
	return &locator, nil
}

// PhysicalClusterNamespaceName encodes the NamespaceLocator into a new
// namespace name for use on a physical cluster. The encoding is repeatable.
func PhysicalClusterNamespaceName(l NamespaceLocator) (string, error) {
	b, err := json.Marshal(l)
	if err != nil {
		return "", err
	}
	// hash the marshalled locator
	hash := sha256.Sum224(b[:])
	// convert the hash to base36 (alphanumeric) to decrease collision probabilities
	base36hash := strings.ToLower(base36.EncodeBytes(hash[:]))
	// use 12 chars of the base36hash, should be enough to avoid collisions and
	// keep the namespaces short enough.
	return fmt.Sprintf("kcp-%s", base36hash[:12]), nil
}
