/*
 * Copyright 2022 Singularity Data
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package webhook

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"

	risingwavev1alpha1 "github.com/singularity-data/risingwave-operator/apis/risingwave/v1alpha1"
)

type RisingWavePodTemplateMutatingWebhook struct {
}

func (w *RisingWavePodTemplateMutatingWebhook) setDefault(ctx context.Context, obj *risingwavev1alpha1.RisingWavePodTemplate) error {
	return nil
}

func (w *RisingWavePodTemplateMutatingWebhook) Default(ctx context.Context, obj runtime.Object) error {
	return w.setDefault(ctx, obj.(*risingwavev1alpha1.RisingWavePodTemplate))
}

func NewRisingWavePodTemplateMutatingWebhook() webhook.CustomDefaulter {
	return &RisingWavePodTemplateMutatingWebhook{}
}
