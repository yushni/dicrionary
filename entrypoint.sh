#!/bin/bash
set -e

cd /pipeline

echo "Authenticating with SA key"
echo ls 

gcloud auth activate-service-account --key-file ./sa-key.json


exec "$@"