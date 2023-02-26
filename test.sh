#!/bin/bash

set -e

git ls-files | grep '_test.go$' | xargs dirname | awk '{print "./"$1}' | xargs go test 
