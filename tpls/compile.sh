#!/usr/bin/env sh

go-bindata -o tpls.go dist/...

sed -i "" "s/package main/package tpls/g" tpls.go
