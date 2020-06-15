#!/bin/bash

echo ls
echo $(deployauth.secureFilePath)
echo pwd

# shellcheck disable=SC2046
cp $(deployauth.secureFilePath) ./dictionary/sa-key.json

chmod +x  ./dictionary/templates/azure/scripts/do.sh
chmod +x  ./dictionary/templates/azure/scripts/entrypoint.sh