#!/bin/bash

kratos proto client --proto_path=. --proto_path=$GOPATH/src api/org/v1/common/org.common.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/org/v1/enums/org.enum.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/org/v1/errors/org.error.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/org/v1/resources/org.resource.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/org/v1/services/org.service.v1.proto
