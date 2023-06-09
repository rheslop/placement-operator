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

// Generated by:
//
// operator-sdk create webhook --group placement --version v1beta1 --kind PlacementAPI --programmatic-validation --defaulting
//

package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// PlacementAPIDefaults -
type PlacementAPIDefaults struct {
	ContainerImageURL string
}

var placementAPIDefaults PlacementAPIDefaults

// log is for logging in this package.
var placementapilog = logf.Log.WithName("placementapi-resource")

// SetupPlacementAPIDefaults - initialize PlacementAPI spec defaults for use with either internal or external webhooks
func SetupPlacementAPIDefaults(defaults PlacementAPIDefaults) {
	placementAPIDefaults = defaults
	placementapilog.Info("PlacementAPI defaults initialized", "defaults", defaults)
}

// SetupWebhookWithManager sets up the webhook with the Manager
func (r *PlacementAPI) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-placement-openstack-org-v1beta1-placementapi,mutating=true,failurePolicy=fail,sideEffects=None,groups=placement.openstack.org,resources=placementapis,verbs=create;update,versions=v1beta1,name=mplacementapi.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &PlacementAPI{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *PlacementAPI) Default() {
	placementapilog.Info("default", "name", r.Name)

	r.Spec.Default()
}

// Default - set defaults for this PlacementAPI spec
func (spec *PlacementAPISpec) Default() {
	if spec.ContainerImage == "" {
		spec.ContainerImage = placementAPIDefaults.ContainerImageURL
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-placement-openstack-org-v1beta1-placementapi,mutating=false,failurePolicy=fail,sideEffects=None,groups=placement.openstack.org,resources=placementapis,verbs=create;update,versions=v1beta1,name=vplacementapi.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &PlacementAPI{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *PlacementAPI) ValidateCreate() error {
	placementapilog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *PlacementAPI) ValidateUpdate(old runtime.Object) error {
	placementapilog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *PlacementAPI) ValidateDelete() error {
	placementapilog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
