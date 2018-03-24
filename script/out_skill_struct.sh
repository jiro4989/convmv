#!/bin/bash

JSON_FILE=$1

cat "$JSON_FILE" | sed 's@^null,@@g' | awk 'NR==3' | sed 's@,$@@g' | gojson
