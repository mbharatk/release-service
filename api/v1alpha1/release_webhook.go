/*
Copyright 2022.

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

package v1alpha1

import (
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var releaselog = logf.Log.WithName("release-resource")

func (r *Release) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/validate-appstudio-redhat-com-v1alpha1-release,mutating=false,failurePolicy=fail,sideEffects=None,groups=appstudio.redhat.com,resources=releases,verbs=create;update,versions=v1alpha1,name=vrelease.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Release{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Release) ValidateCreate() error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Release) ValidateUpdate(old runtime.Object) error {
	releaselog.Info("validate update", "name", r.Name)

	return fmt.Errorf("release resources cannot be updated")
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Release) ValidateDelete() error {
	return nil
}
