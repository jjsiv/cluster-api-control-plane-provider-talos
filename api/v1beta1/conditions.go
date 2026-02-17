// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// +kubebuilder:object:generate=true
// +groupName=controlplane.cluster.x-k8s.io
package v1beta1

import clusterv1 "sigs.k8s.io/cluster-api/api/core/v1beta2"

// Conditions and condition Reasons for the TalosControlPlane object

// TalosControlPlane's MachinesReady condition and corresponding reasons.
const (
	// TalosControlPlaneMachinesReadyCondition surfaces detail of issues on the controlled machines, if any.
	// Please note this will include also APIServerPodHealthy, ControllerManagerPodHealthy, SchedulerPodHealthy conditions.
	// If not using an external etcd also EtcdPodHealthy, EtcdMemberHealthy conditions are included.
	TalosControlPlaneMachinesReadyCondition = clusterv1.MachinesReadyCondition

	// TalosControlPlaneMachinesReadyReason surfaces when all the controlled machine's Ready conditions are true.
	TalosControlPlaneMachinesReadyReason = clusterv1.ReadyReason

	// TalosControlPlaneMachinesNotReadyReason surfaces when at least one of the controlled machine's Ready conditions is false.
	TalosControlPlaneMachinesNotReadyReason = clusterv1.NotReadyReason

	// TalosControlPlaneMachinesReadyUnknownReason surfaces when at least one of the controlled machine's Ready conditions is unknown
	// and no one of the controlled machine's Ready conditions is false.
	TalosControlPlaneMachinesReadyUnknownReason = clusterv1.ReadyUnknownReason

	// TalosControlPlaneMachinesReadyNoReplicasReason surfaces when no machines exist for the TalosControlPlane.
	TalosControlPlaneMachinesReadyNoReplicasReason = clusterv1.NoReplicasReason

	// TalosControlPlaneMachinesReadyInternalErrorReason surfaces unexpected failures when computing the MachinesReady condition.
	TalosControlPlaneMachinesReadyInternalErrorReason = clusterv1.InternalErrorReason
)

// TalosControlPlane's MachinesUpToDate condition and corresponding reasons.
const (
	// TalosControlPlaneMachinesUpToDateCondition surfaces details of controlled machines not up to date, if any.
	// Note: New machines are considered 10s after machine creation. This gives time to the machine's owner controller to recognize the new machine and add the UpToDate condition.
	TalosControlPlaneMachinesUpToDateCondition = clusterv1.MachinesUpToDateCondition

	// TalosControlPlaneMachinesUpToDateReason surfaces when all the controlled machine's UpToDate conditions are true.
	TalosControlPlaneMachinesUpToDateReason = clusterv1.UpToDateReason

	// TalosControlPlaneMachinesNotUpToDateReason surfaces when at least one of the controlled machine's UpToDate conditions is false.
	TalosControlPlaneMachinesNotUpToDateReason = clusterv1.NotUpToDateReason

	// TalosControlPlaneMachinesUpToDateUnknownReason surfaces when at least one of the controlled machine's UpToDate conditions is unknown
	// and no one of the controlled machine's UpToDate conditions is false.
	TalosControlPlaneMachinesUpToDateUnknownReason = clusterv1.UpToDateUnknownReason

	// TalosControlPlaneMachinesUpToDateNoReplicasReason surfaces when no machines exist for the TalosControlPlane.
	TalosControlPlaneMachinesUpToDateNoReplicasReason = clusterv1.NoReplicasReason

	// TalosControlPlaneMachinesUpToDateInternalErrorReason surfaces unexpected failures when computing the MachinesUpToDate condition.
	TalosControlPlaneMachinesUpToDateInternalErrorReason = clusterv1.InternalErrorReason
)

const (
	// MachinesBootstrappedCondition is tracking control planes bootstrap status.
	MachinesBootstrappedCondition clusterv1.ConditionType = "MachinesBootstrapped"

	// WaitingForMachinesReason (Severity=Info) documents a TalosControlPlane bootstrap is waiting
	// for all control plane nodes to be created.
	WaitingForMachinesReason = "WaitingForMachines"
)

const (
	// AvailableCondition documents that the first control plane instance has completed Talos boot sequence
	// and so the control plane is available and an API server instance is ready for processing requests.
	AvailableCondition = "Available"

	// WaitingForTalosBootReason (Severity=Info) documents a TalosControlPlane object waiting for the first
	// control plane instance to complete Talos boot sequence.
	WaitingForTalosBootReason = "WaitingForTalosBoot"

	// InvalidControlPlaneConfigReason (Severity=Error) documents that controlplane config is invalid and the provider
	// can not proceed with the bootstrap.
	InvalidControlPlaneConfigReason = "InvalidControlPlaneConfig"
)

const (
	// MachinesSpecUpToDateCondition documents that the spec of the machines controlled by the TalosControlPlane
	// is up to date. When this condition is false, the TalosControlPlane is executing a rolling upgrade.
	MachinesSpecUpToDateCondition clusterv1.ConditionType = "MachinesSpecUpToDate"

	// RollingUpdateInProgressReason (Severity=Warning) documents a TalosControlPlane object executing a
	// rolling upgrade for aligning the machines spec to the desired state.
	RollingUpdateInProgressReason = "RollingUpdateInProgress"
)

const (
	// ResizedCondition documents a TalosControlPlane that is resizing the set of controlled machines.
	ResizedCondition clusterv1.ConditionType = "Resized"

	// ScalingUpReason (Severity=Info) documents a TalosControlPlane that is increasing the number of replicas.
	ScalingUpReason = "ScalingUp"

	// ScalingDownReason (Severity=Info) documents a TalosControlPlane that is decreasing the number of replicas.
	ScalingDownReason = "ScalingDown"
)

const (
	// ControlPlaneComponentsHealthyCondition reports the overall status of control plane components
	// implemented as static pods generated by Talos including kube-api-server, kube-controller manager,
	// kube-scheduler and etcd.
	ControlPlaneComponentsHealthyCondition clusterv1.ConditionType = "ControlPlaneComponentsHealthy"

	// ControlPlaneComponentsUnhealthyReason (Severity=Error) documents a control plane component not healthy.
	ControlPlaneComponentsUnhealthyReason = "ControlPlaneComponentsUnhealthy"

	// ControlPlaneComponentsInspectionFailedReason documents a failure in inspecting the control plane component status.
	ControlPlaneComponentsInspectionFailedReason = "ControlPlaneComponentsInspectionFailed"
)

const (
	// EtcdClusterHealthyCondition documents the overall etcd cluster's health.
	EtcdClusterHealthyCondition clusterv1.ConditionType = "EtcdClusterHealthyCondition"

	// EtcdClusterUnhealthyReason (Severity=Error) is set when the etcd cluster is unhealthy.
	EtcdClusterUnhealthyReason = "EtcdClusterUnhealthy"
)

const (
	// MachinesCreatedCondition documents that the machines controlled by the TalosControlPlane are created.
	// When this condition is false, it indicates that there was an error when cloning the infrastructure/bootstrap template or
	// when generating the machine object.
	MachinesCreatedCondition clusterv1.ConditionType = "MachinesCreated"

	// InfrastructureTemplateCloningFailedReason (Severity=Error) documents a TalosControlPlane failing to
	// clone the infrastructure template.
	InfrastructureTemplateCloningFailedReason = "InfrastructureTemplateCloningFailed"

	// BootstrapTemplateCloningFailedReason (Severity=Error) documents a TalosControlPlane failing to
	// clone the bootstrap template.
	BootstrapTemplateCloningFailedReason = "BootstrapTemplateCloningFailed"

	// MachineGenerationFailedReason (Severity=Error) documents a TalosControlPlane failing to
	// generate a machine object.
	MachineGenerationFailedReason = "MachineGenerationFailed"
)
