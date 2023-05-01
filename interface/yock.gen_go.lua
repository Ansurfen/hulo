---@diagnostic disable: undefined-global
Import({ "cushion", "yock-io" })
local io = require("cushion-io")
io.Exec(
    "protoc --go_out=../sdk/go --go_opt=paths=source_relative --go-grpc_out=../sdk/go --go-grpc_opt=paths=source_relative hulo.proto")
