---@diagnostic disable: undefined-global
Import({ "yock-io", "cushion" })
local yock = require("yock-io")
local vm = require("cushion-vm")
local path = require("cushion-path")
yock.copy(path.Join(vm.Workdir(), ".hulo", "sdk", "go"), ".")
yock.move("go", "hulo")
local io = require("cushion-io")
io.Exec("go mod tidy")
