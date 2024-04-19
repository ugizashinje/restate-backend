#!/bin/bash
set -euo pipefail

cd ./docker/db/postgres/postgres-initdb


openssl req -new -text -passout pass:abcd -subj /CN=localhost -out server.req -keyout privkey.pem
openssl rsa -in privkey.pem -passin pass:abcd -out server.key
openssl req -x509 -in server.req -text -key server.key -out server.crt
chmod 600 server.key
# test $(uname -s) == Linux && chown 999 server.key