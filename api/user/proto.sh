#!/bin/bash

kratos proto client --proto_path=. --proto_path=$GOPATH/src api/user/v1/common/user.common.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/user/v1/enums/user.enum.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/user/v1/errors/user.error.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/user/v1/resources/user.resource.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/user/v1/services/user.service.v1.proto
