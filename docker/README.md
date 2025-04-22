

```bash
# Go to specific directory
$ cd /Users/enricogoerlitz/Developer/azure/aks-backend-architecture/docker

# Login into azure acr
$ az acr login --name eaksprojectweucr

# Build and Deploy Flyway-DB-Migration-Service
$ az acr build \
  --registry eaksprojectweucr \
  --resource-group explore-aks-devops-project-weu-rg \
  --file ./backend/database/migration-service/flyway-migration-service.Dockerfile \
  --image aks-backend-flyway-migration-service:latest \
  ../backend/database/migration-service

# Build and Deploy RESTAPI Image
$ az acr build \
  --registry eaksprojectweucr \
  --resource-group explore-aks-devops-project-weu-rg \
  --file ./backend/app/restapi.Dockerfile \
  --image aks-backend-restapi-service:latest \
  ../backend/app
```