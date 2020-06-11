#!/bin/bash
set -e

cd /pipeline

echo "Authenticating with SA key"
gcloud auth activate-service-account --key-file ./sa-key.json

exec "$@"