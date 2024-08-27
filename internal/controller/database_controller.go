/*
Copyright 2024.

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

package controller

import (
	"context"
	v1 "github.com/wally/my-operator/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	databasev1 "github.com/wallyxjh/my-operator/api/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// DatabaseReconciler reconciles a Database object
type DatabaseReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=database.mydomain.com,resources=databases,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=database.mydomain.com,resources=databases/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=database.mydomain.com,resources=databases/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Database object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.0/pkg/reconcile
func (r *DatabaseReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("database", req.NamespacedName)

	// Fetch the Database instance
	var database v1.Database
	if err := r.Get(ctx, req.NamespacedName, &database); err != nil {
		log.Error(err, "unable to fetch Database")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Here you can add your custom logic to handle the Database resource.
	// For example, create a deployment for the database or allocate storage.
	log.Info("Reconciling Database", "DatabaseName", database.Spec.DatabaseName)

	// Update the status of the Database resource
	database.Status.Conditions = []metav1.Condition{
		{
			Type:    "Available",
			Status:  metav1.ConditionTrue,
			Reason:  "DatabaseCreated",
			Message: "Database instance has been created successfully",
		},
	}
	if err := r.Status().Update(ctx, &database); err != nil {
		log.Error(err, "unable to update Database status")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DatabaseReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&databasev1.Database{}).
		Complete(r)
}
