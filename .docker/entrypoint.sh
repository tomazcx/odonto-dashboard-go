#!/bin/bash

templ generate
go mod tidy
go build -buildvcs=false -o ./temp/main
./temp/main