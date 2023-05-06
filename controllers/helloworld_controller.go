package controllers

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	hellov1 "github.com/your-username/helloworld-operator/api/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// HelloWorldReconciler reconcilia um objeto HelloWorld
type HelloWorldReconciler struct {
	client.Client
	Log logr.Logger
}

//+kubebuilder:rbac:groups=hello.example.com,resources=helloworlds,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=hello.example.com,resources=helloworlds/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=hello.example.com,resources=helloworlds/finalizers,verbs=update

func (r *HelloWorldReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("helloworld", req.NamespacedName)

//Obtenha o recurso HelloWorld
var hw hellov1.HelloWorld 
if err := r.Get(ctx, req.NamespacedName, &hw); err != nil { 
  log.Error(err, "unable to fetch HelloWorld") 
  return ctrl.Result{}, client.IgnoreNotFound(err) }

// Imprimir mensagem HelloWorld
fmt.Printf("HelloWorld message: %s\n", hw.Spec.Message)

// Atualize o status HelloWorld
hw.Status.Message = fmt.Sprintf("HelloWorld message: %s", hw.Spec.Message)
if err := r.Status().Update(ctx, &hw); err != nil {	
  log.Error(err, "unable to update HelloWorld status")	
  return ctrl.Result{}, err}return ctrl.Result{}, nil

// SetupWithManager configura o controlador com o Manager.
func (r *HelloWorldReconciler) SetupWithManager(mgr ctrl.Manager) error { 
  return ctrl.NewControllerManagedBy(mgr) 
  For(&hellov1.HelloWorld{}) 
   Complete(r) }}