---@diagnostic disable: undefined-global
Import({ "cushion", "yock-io" })
local io = require("cushion-io")
io.Exec(
"protoc --go_out=../sdk/c/libhulo --go_opt=paths=source_relative --go-grpc_out=../sdk/c/libhulo --go-grpc_opt=paths=source_relative hulo.proto")
local yock = require("yock-io")
yock.file_replace("../sdk/c/libhulo/hulo.pb.go", "package hulo", "package main")
yock.file_replace("../sdk/c/libhulo/hulo_grpc.pb.go", "package hulo", "package main")
