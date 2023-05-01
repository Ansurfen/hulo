---@diagnostic disable: undefined-global
Import({ "cushion", "yock-io" })
local io = require("cushion-io")
io.Exec(
    "protoc --dart_out=grpc:../sdk/dart -I. hulo.proto")
