#!/usr/bin/env bash
make
sls create_domain
serverless deploy --verbose