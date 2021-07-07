#!/bin/bash
## set version `git describe --tags $(git rev-list --tags --max-count=1)`
VERSION=`git describe --tags $(git rev-list --tags --max-count=1)`
## set build `git log -1 --pretty=format:%h`
BUILD=`git log -1 --pretty=format:%h`

## go build
go build -ldflags "-X main.version=${VERSION}@${BUILD}" 

## test help
./enumable_builder -h

## 테스트 
./enumable_builder -P hello h e l l o > hello.go
cat hello.go
## hello 파일 지우기
rm hello.go