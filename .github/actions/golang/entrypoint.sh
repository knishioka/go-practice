#!/bin/bash

set -euxo pipefile

golangci-lint run
