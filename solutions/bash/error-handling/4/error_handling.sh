#!/usr/bin/env bash

case $# in
  1)
    printf 'Hello, %s\n' "$1";;
  *)
    printf 'Usage: error_handling.sh <person>\n'; exit 1;;
esac