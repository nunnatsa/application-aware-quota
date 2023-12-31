/*
Copyright 2023 The AAQ Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "kubevirt.io/applications-aware-quota/pkg/generated/aaq/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// AAQs returns a AAQInformer.
	AAQs() AAQInformer
	// AAQJobQueueConfigs returns a AAQJobQueueConfigInformer.
	AAQJobQueueConfigs() AAQJobQueueConfigInformer
	// ApplicationsResourceQuotas returns a ApplicationsResourceQuotaInformer.
	ApplicationsResourceQuotas() ApplicationsResourceQuotaInformer
	// ClusterAppsResourceQuotas returns a ClusterAppsResourceQuotaInformer.
	ClusterAppsResourceQuotas() ClusterAppsResourceQuotaInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// AAQs returns a AAQInformer.
func (v *version) AAQs() AAQInformer {
	return &aAQInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// AAQJobQueueConfigs returns a AAQJobQueueConfigInformer.
func (v *version) AAQJobQueueConfigs() AAQJobQueueConfigInformer {
	return &aAQJobQueueConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ApplicationsResourceQuotas returns a ApplicationsResourceQuotaInformer.
func (v *version) ApplicationsResourceQuotas() ApplicationsResourceQuotaInformer {
	return &applicationsResourceQuotaInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ClusterAppsResourceQuotas returns a ClusterAppsResourceQuotaInformer.
func (v *version) ClusterAppsResourceQuotas() ClusterAppsResourceQuotaInformer {
	return &clusterAppsResourceQuotaInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
