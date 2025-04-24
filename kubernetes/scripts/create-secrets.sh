#!/bin/bash
# chmod +x ./scripts/create-secrets.sh
# ./scripts/create-secrets.sh

set -e

echo "📦 Erstelle Kubernetes Secret aus .env Datei..."

kubectl create secret generic app-secrets \
  --from-env-file=secrets.env \
  --dry-run=client -o yaml | kubectl apply -f -

kubectl create secret generic app-secrets \
  --from-env-file=secrets.env \
  --dry-run=client -o yaml | kubectl apply -f - \
  --namespace=database

kubectl create secret generic app-secrets \
  --from-env-file=secrets.env \
  --dry-run=client -o yaml | kubectl apply -f - \
  --namespace=application

echo "✅ Secret erfolgreich erstellt."
