# CI/CD Commands

## Login into Azure

```bash
# Login into az
$ az login -u <username> -p <password>
```

## Build and Deploy Docker images

```bash
# Login into azure acr
$ az acr login --name <acr-name>

# Build and Deploy Flyway-DB-Migration-Service
$ az acr build \
  --registry <acr-name> \
  --resource-group <resource-group-name> \
  --file docker/backend/database/migration-service/flyway-migration-service.Dockerfile \
  --image aks-backend-flyway-migration-service:latest \
  --platform linux/arm64 \
  backend/database/migration-service

# Build and Deploy RESTAPI Image
$ az acr build \
  --registry eaksprojectweucr \
  --resource-group <resource-group-name> \
  --file docker/backend/app/restapi.Dockerfile \
  --image aks-backend-restapi-service:latest \
  --platform linux/arm64 \
  backend/app
```

## Build, Deploy and Create/Update Cloud Infrastructure

```bash
$ terraform init
$ terraform workspace select <env>
$ terraform apply -var-file=tfvars.$(terraform workspace show).tfvars -auto-approve
```

## Build, Deploy and Create/Update Kubernetes Infrastructure

```bash
$ az aks get-credentials --resource-group <resource-group-name> --name <Cluster-Name>
$ kubectl apply -f kubernetes/kube-manifests/cache/
$ kubectl apply -f kubernetes/kube-manifests/database/
$ kubectl apply -f kubernetes/kube-manifests/backend/database/migration/
$ kubectl apply -f kubernetes/kube-manifests/backend/app
$ kubectl apply -f kubernetes/kube-manifests/frontend/
```
