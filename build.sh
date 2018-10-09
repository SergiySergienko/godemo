#!/usr/bin/env bash
# Stops the process if something fails
set -xe

# All of the dependencies needed/fetched for your project.
# This is what actually fixes the problem so that EB can find your dependencies. 
# FOR EXAMPLE:
go get "github.com/aws/aws-sdk-go"
go get "github.com/aws/aws-sdk-go/aws/"
go get "github.com/aws/aws-sdk-go/service/"

# create the application binary that eb uses
GOOS=linux GOARCH=amd64 go build -o bin/application -ldflags="-s -w"