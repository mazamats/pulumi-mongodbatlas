#!/bin/bash
PROVIDER=mongodbatlas

make prepare NAME="$PROVIDER" REPOSITORY=github.com/pulumi/pulumi-"$PROVIDER"
go get github.com/pulumi/scripts/gomod-doccopy
GO111MODULE=on go get github.com/pulumi/pulumi-terraform@master
GO111MODULE=on go get github.com/terraform-providers/terraform-provider-"$PROVIDER"
GO111MODULE=on go mod vendor
make ensure
