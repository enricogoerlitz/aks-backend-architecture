# Kubectl commands

```bash
# Get Credentials
$ az aks get-credentials --resource-group <Resource-Group-Name> --name <Cluster-Name>
$ az aks get-credentials --resource-group explore-aks-devops-dev-weu-rg --name explore-aks-devops-dev-weu-aks

# List Kubernetes Worker Nodes
$ kubectl get nodes 
$ kubectl get nodes -o wide

$ kubectl create namespace spark-cluster
$ kubectl apply -f ./kube-manifests

# enter kubernetes cluster in shell -> use nsloopup etc for debugging
$ kubectl run -n default dnsutils --image=busybox:1.28 --rm -it --restart=Never -- sh

$ kubectl get nodes --label-columns agendpool
$ kubectl label node aks-sparkpool001-87494612-vmss00000{0&1&2} nodepool=sparkpool001

```

**Redis check**

```bash
$ kubectl logs job/redis-cluster-init -f

$ kubectl run debug-dns --rm -it --image=alpine -- sh

$ redis-cli -h 9.163.157.93 -p 6379 -a my-password
$ redis-cli -h 9.163.157.93 -p 6379 -a my-password cluster info
$ redis-cli -h 9.163.157.93 -p 6379 -a my-password cluster nodes
$ kubectl exec -it redis-cluster-0 -- redis-cli -a my-password cluster info

# CLuster mode!
$ kubectl exec -it redis-cluster-0 -- redis-cli -c -a my-password
```