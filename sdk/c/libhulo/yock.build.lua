---@diagnostic disable: undefined-global
Import({ "yock-io" })
local yock = require("yock-io")
yock.exec({
    "gcc -c ./hulo.c -o libhulo.o",
    "ar rcs libhulo.a libhulo.o",
    "go build -o libhulo.dll -buildmode=c-shared",
    "go build -o libhulo.a -buildmode=c-archive",
    "rm libhulo.o"
})
