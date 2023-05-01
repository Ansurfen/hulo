---@diagnostic disable: undefined-global
Import({ "yock-io" })
local yock = require("yock-io")
yock.curl("https://raw.githubusercontent.com/DaveGamble/cJSON/master/cJSON.c", "cJSON.c")
yock.curl("https://raw.githubusercontent.com/DaveGamble/cJSON/master/cJSON.h", "cJSON.h")