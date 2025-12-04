//go:build tools
// +build tools

package main

import (
	_ "google.golang.org/genproto/googleapis/rpc/status"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
