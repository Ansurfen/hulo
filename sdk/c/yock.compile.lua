---@diagnostic disable: undefined-global
Import({ "yock-io" })

local yock = require("yock-io")
yock.exec({"gcc ./hulo_test.c ./libhulo/cJSON.c -L ./libhulo -lhulo -o hulo"})