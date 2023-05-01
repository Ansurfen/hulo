---@diagnostic disable: undefined-global
Import({ "cushion", "yock-io" })
local io = require("cushion-io")
io.Exec(
    "python -m grpc_tools.protoc -I. --python_out=../sdk/python --pyi_out=../sdk/python --grpc_python_out=../sdk/python hulo.proto")
