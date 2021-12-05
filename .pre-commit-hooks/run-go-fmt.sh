#!/bin/bash
set -e

exec 5>&1
output="$(gofmt -l -w "$@" | tee /dev/fd/5)"
[[ -z "$output" ]]
