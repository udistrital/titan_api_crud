#!/bin/sh

set -e
set -u

if [ -n "${PARAMETER_STORE:-}" ]; then
  export TITAN_API_CRUD_DB_USER="$(aws ssm get-parameter --with-decryption --name /${PARAMETER_STORE}/titan_api_crud/db/username --output text --query Parameter.Value)"
  test -n "${TITAN_API_CRUD_DB_USER}"
  export TITAN_API_CRUD_DB_PASS="$(aws ssm get-parameter --with-decryption --name /${PARAMETER_STORE}/titan_api_crud/db/password --output text --query Parameter.Value)"
  test -n "${TITAN_API_CRUD_DB_PASS}"
fi

exec ./main "$@"
