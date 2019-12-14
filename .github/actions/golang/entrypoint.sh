#!/bin/bash

set -euo pipefile

golangci-lint run --enable-all
