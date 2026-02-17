# Code walkthrough

1. Setup reconcilers (ClusterCache, TCP), webhooks, etc. (`main.go`)
   - How to take advantage of ClusterCache compared to tracker? any other considerations? We can look into it after we sort out general v1beta2 implementation

2. Reconciliation process
   - get owner cluster or requeue if not yet set
   - if paused, requeue
   - if infra not yet ready, requeue
   - init patch helper, requeue if error (PatchHelper is a CAPI resource patching utility)
   - defer status update, run after each reconcile

3. Reconcile external reference (infrastructure template)
   - MachineTemplate is retrieved as Unstructured
   - set OwnerRef on MachineTemplate to Cluster resource

4. Continue reconciliation if ControlPlaneEndpoint is set on Cluster

5. Adopt ControlPlane Machines
   - get ControlPlane Machines for the Cluster (lists Machines in its namespace that have cluster & control-plane label)
   - **ManagementCluster interace for this like in Kubeadm?**

6. Run all reconciliations in sequence
   - EtcdMembers, NodeHealth, Conditions, Kubeconfig, Machines
   - EtcdMembers, healthchecks etcd. does this set healthy condition even if no etcd members are present?
   - Machines actually creates Machines

# Plan

- Implement new conditions without changing too much first
- Modify functions to take v1beta2 API
- Implement ClusterCache, ManagementCluster interface?
