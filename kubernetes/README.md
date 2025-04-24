# AKS Kubernetes Manifests

## Commands

**Setup**

```bash
# login into aks cluster
$ az aks get-credentials --resource-group explore-aks-devops-dev-weu-rg --name explore-aks-devops-dev-weu-aks

# List Kubernetes Worker Nodes
$ kubectl get nodes 
$ kubectl get nodes -o wide

# Setup secrets
$ ./scripts/create-secrets.sh
```

**Debugging**

```bash
# Get logs
$ kubectl logs redis-cluster-init -f -n database

# Shell in cluster
$ kubectl run debug-dns --rm -it --image=alpine -n <namespace> -- sh

# Apply / Delete
$ kubectl apply -f kube-manifests/database/
$ kubectl delete -f kube-manifests/database/

# Redis CLI
$ kubectl exec -it redis-cluster-0 -n database -- redis-cli -a my-password
$ cluster info
$ kubectl exec -it redis-cluster-0 -n database -- redis-cli -a my-password cluster info
```