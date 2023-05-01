---@diagnostic disable: undefined-global
Import({ "yock-io" })
local yock = require("yock-io")
yock.exec({
    "rm cJSON.c",
    "rm cJSON.h",
    "rm libhulo.a",
    "rm libhulo.dll",
    "rm libhulo.h",
    "rm hulo_grpc.pb.go",
    "rm hulo.pb.go",
})
