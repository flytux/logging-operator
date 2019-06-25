/*
 * Copyright © 2019 Banzai Cloud
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package fluentd

import (
	"context"

	"github.com/banzaicloud/logging-operator/pkg/resources/templates"
	"github.com/banzaicloud/logging-operator/pkg/util"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

func (r *Reconciler) appconfigMap() runtime.Object {
	current := &corev1.ConfigMap{}
	err := r.Client.Get(context.TODO(), types.NamespacedName{Namespace: r.Fluentd.Namespace, Name: appConfigMapName}, current)
	if err != nil {
		return &corev1.ConfigMap{
			ObjectMeta: templates.FluentdObjectMeta(appConfigMapName, util.MergeLabels(r.Fluentd.Labels, labelSelector), r.Fluentd),
			Data:       map[string]string{},
		}
	}
	return current
}
