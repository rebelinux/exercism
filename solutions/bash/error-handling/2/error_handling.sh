#!/usr/bin/env bash

case "$#" in
  0 | 2*)
    printf "Usage: error_handling.sh <person>"; exit 1
  ;;
  1)
    printf "Hello, $1"
  ;;
esac