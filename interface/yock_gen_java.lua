---@diagnostic disable: undefined-global
Import({ "cushion", "yock-io" })
local io = require("cushion-io")
io.Exec(
    "protoc --plugin=protoc-gen-grpc-java=D:\\D\\langs\\pkg\\protoc-23.0-rc-1-win64\\bin\\protoc-gen-grpc-java-1.9.1-windows-x86_64.exe  --grpc-java_out=../sdk/java/src/main/java/ --java_out=../sdk/java/src/main/java/ --proto_path=.  hulo.proto")
