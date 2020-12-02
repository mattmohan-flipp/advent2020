#!/bin/bash
find . -maxdepth 1 -mindepth 1 -type d -iname day\* | xargs go test