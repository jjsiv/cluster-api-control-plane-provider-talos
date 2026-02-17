package controllers

import (
	"context"

	"github.com/pkg/errors"
	controlplanev1 "github.com/siderolabs/cluster-api-control-plane-provider-talos/api/v1beta1"
	"github.com/siderolabs/talos/pkg/machinery/constants"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/utils/ptr"
	clusterv1 "sigs.k8s.io/cluster-api/api/core/v1beta2"
	"sigs.k8s.io/cluster-api/util"
	"sigs.k8s.io/cluster-api/util/collections"
	"sigs.k8s.io/cluster-api/util/conditions"
	v1beta1conditions "sigs.k8s.io/cluster-api/util/conditions/deprecated/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *TalosControlPlaneReconciler) updateV1Beta1Status(ctx context.Context, tcp *controlplanev1.TalosControlPlane, cluster *clusterv1.Cluster) error {
	clusterSelector := &metav1.LabelSelector{
		MatchLabels: map[string]string{
			clusterv1.ClusterNameLabel:         cluster.Name,
			clusterv1.MachineControlPlaneLabel: "",
		},
	}

	selector, err := metav1.LabelSelectorAsSelector(clusterSelector)
	if err != nil {
		// Since we are building up the LabelSelector above, this should not fail
		return errors.Wrap(err, "failed to parse label selector")
	}
	// Copy label selector to its status counterpart in string format.
	// This is necessary for CRDs including scale subresources.
	tcp.Status.Selector = selector.String()

	ownedMachines, err := r.getControlPlaneMachinesForCluster(ctx, util.ObjectKey(cluster))
	if err != nil {
		return err
	}

	replicas := int32(len(ownedMachines.Items))

	// set basic data that does not require interacting with the workload cluster
	tcp.Status.Deprecated.V1Beta1.ReadyReplicas = 0
	tcp.Status.Deprecated.V1Beta1.UnavailableReplicas = replicas

	// Return early if the deletion timestamp is set, we don't want to try to connect to the workload cluster.
	if !tcp.DeletionTimestamp.IsZero() {
		return nil
	}

	lowestVersion := collections.FromMachineList(&ownedMachines).LowestVersion()
	if lowestVersion != "" {
		tcp.Status.Version = lowestVersion
	}

	c, err := r.ClusterCache.GetClient(ctx, util.ObjectKey(cluster))
	if err != nil {
		r.Log.Info("failed to get kubeconfig for the cluster", "error", err)

		return nil
	}

	nodeSelector := labels.NewSelector()
	req, err := labels.NewRequirement(constants.LabelNodeRoleControlPlane, selection.Exists, []string{})
	if err != nil {
		return err
	}

	var nodes v1.NodeList

	err = c.List(ctx, &nodes, &client.ListOptions{
		LabelSelector: nodeSelector.Add(*req),
	})

	if err != nil {
		r.Log.Info("failed to list controlplane nodes", "error", err)

		return nil
	}

	// if we were able to fetch some resources via control plane endpoint,
	// workload cluster control plane endpoint is available
	tcp.Status.Initialization.ControlPlaneInitialized = ptr.To(true)
	v1beta1conditions.MarkTrue(tcp, controlplanev1.AvailableV1Beta1Condition)

	for _, node := range nodes.Items {
		if util.IsNodeReady(&node) {
			tcp.Status.Deprecated.V1Beta1.ReadyReplicas++
		}
	}

	// fix the case then some Node objects are still visible which were deleted
	if tcp.Status.Deprecated.V1Beta1.ReadyReplicas > *tcp.Status.Replicas {
		tcp.Status.Deprecated.V1Beta1.ReadyReplicas = replicas
	}

	tcp.Status.Deprecated.V1Beta1.UnavailableReplicas = replicas - tcp.Status.Deprecated.V1Beta1.ReadyReplicas

	r.Log.Info("ready replicas", "count", tcp.Status.ReadyReplicas)

	return nil

}

// TODO: separate functions for setting status on v1beta1 and v1alpha3 APIs. Move to a separate file
func (r *TalosControlPlaneReconciler) updateStatus(ctx context.Context, tcp *controlplanev1.TalosControlPlane, cluster *clusterv1.Cluster) error {
	clusterSelector := &metav1.LabelSelector{
		MatchLabels: map[string]string{
			clusterv1.ClusterNameLabel:         cluster.Name,
			clusterv1.MachineControlPlaneLabel: "",
		},
	}

	selector, err := metav1.LabelSelectorAsSelector(clusterSelector)
	if err != nil {
		// Since we are building up the LabelSelector above, this should not fail
		return errors.Wrap(err, "failed to parse label selector")
	}
	// Copy label selector to its status counterpart in string format.
	// This is necessary for CRDs including scale subresources.
	tcp.Status.Selector = selector.String()

	ownedMachines, err := r.getControlPlaneMachinesForCluster(ctx, util.ObjectKey(cluster))
	if err != nil {
		return err
	}

	// TODO: return a collection from getControlPlaneMachinesForCluster
	machineCollection := collections.FromMachineList(&ownedMachines)
	var readyReplicas, availableReplicas, upToDateReplicas int32
	for _, machine := range machineCollection {
		if conditions.IsTrue(machine, clusterv1.MachineReadyCondition) {
			readyReplicas++
		}
		if conditions.IsTrue(machine, clusterv1.MachineAvailableCondition) {
			availableReplicas++
		}
		if conditions.IsTrue(machine, clusterv1.MachineUpToDateCondition) {
			upToDateReplicas++
		}
	}

	replicas := int32(len(ownedMachines.Items))
	tcp.Status.Initialization.ControlPlaneInitialized = ptr.To(false)
	tcp.Status.Replicas = &replicas
	tcp.Status.ReadyReplicas = ptr.To(readyReplicas)
	tcp.Status.AvailableReplicas = ptr.To(availableReplicas)
	tcp.Status.UpToDateReplicas = ptr.To(upToDateReplicas)

	// set basic data that does not require interacting with the workload cluster
	tcp.Status.Deprecated.V1Beta1.ReadyReplicas = 0
	tcp.Status.Deprecated.V1Beta1.UnavailableReplicas = replicas

	// Return early if the deletion timestamp is set, we don't want to try to connect to the workload cluster.
	if !tcp.DeletionTimestamp.IsZero() {
		return nil
	}

	lowestVersion := collections.FromMachineList(&ownedMachines).LowestVersion()
	if lowestVersion != "" {
		tcp.Status.Version = lowestVersion
	}

	c, err := r.ClusterCache.GetClient(ctx, util.ObjectKey(cluster))
	if err != nil {
		r.Log.Info("failed to get kubeconfig for the cluster", "error", err)

		return nil
	}

	nodeSelector := labels.NewSelector()
	req, err := labels.NewRequirement(constants.LabelNodeRoleControlPlane, selection.Exists, []string{})
	if err != nil {
		return err
	}

	var nodes v1.NodeList

	err = c.List(ctx, &nodes, &client.ListOptions{
		LabelSelector: nodeSelector.Add(*req),
	})

	if err != nil {
		r.Log.Info("failed to list controlplane nodes", "error", err)

		return nil
	}

	// if we were able to fetch some resources via control plane endpoint,
	// workload cluster control plane endpoint is available
	tcp.Status.Initialization.ControlPlaneInitialized = ptr.To(true)
	v1beta1conditions.MarkTrue(tcp, controlplanev1.AvailableV1Beta1Condition)

	for _, node := range nodes.Items {
		if util.IsNodeReady(&node) {
			tcp.Status.Deprecated.V1Beta1.ReadyReplicas++
		}
	}

	// fix the case then some Node objects are still visible which were deleted
	if tcp.Status.Deprecated.V1Beta1.ReadyReplicas > *tcp.Status.Replicas {
		tcp.Status.Deprecated.V1Beta1.ReadyReplicas = replicas
	}

	tcp.Status.Deprecated.V1Beta1.UnavailableReplicas = replicas - tcp.Status.Deprecated.V1Beta1.ReadyReplicas

	// TODO: set status condition instead
	// if *tcp.Status.ReadyReplicas > 0 {
	// 	tcp.Status.Ready = true
	// }

	r.Log.Info("ready replicas", "count", tcp.Status.ReadyReplicas)

	return nil
}
