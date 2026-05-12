#!/usr/bin/env bash

case "$#" in
  0 | [2-9])
    printf "Usage: error_handling.sh <person>\n"; exit 1;;
  1)
    printf 'Hello, %s\n' "$1";;
esac