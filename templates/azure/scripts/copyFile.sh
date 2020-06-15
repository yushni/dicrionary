#!/bin/bash

ls
pwd
# shellcheck disable=SC2046
# shellcheck disable=SC2005
echo $(caCertificate.secureFilePath)
# shellcheck disable=SC2046
cp $(deployauth.secureFilePath) ./dictionary/sa-key.json

chmod +x  ./dictionary/templates/azure/scripts/do.sh
chmod +x  ./dictionary/templates/azure/scripts/entrypoint.sh