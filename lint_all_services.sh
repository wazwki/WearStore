#!/bin/bash
set -e

SERVICES=("api-gateway" "cart-service" "product-service" "user-service")

for SERVICE in "${SERVICES[@]}"; do
  echo "Running golangci-lint for $SERVICE..."
  pushd "$SERVICE" > /dev/null
  golangci-lint run
  popd > /dev/null
done

echo "Linting completed for all services!"
