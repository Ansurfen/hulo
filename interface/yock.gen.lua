---@diagnostic disable:  undefined-global, unused-label
local protoc = function(...)
    return exec({ gsub("protoc", ...) })
end
print("start")
local root = "../sdk"
local worksapce = ""
local target = "hulo.proto"
::py::
worksapce = "/python"
exec({ gsub("python -m grpc_tools.protoc -I.",
    "--python_out=" .. root .. worksapce,
    "--pyi_out=" .. root .. worksapce,
    "--grpc_python_out=" .. root .. worksapce,
    target)
})

::__tmp::
protoc("--go_out=" .. root .. worksapce,
    "--go_opt=paths=source_relative",
    "--go-grpc_out=" .. root .. worksapce,
    "--go-grpc_opt=paths=source_relative",
    target)

::go::
worksapce = "/go"
---@link __tmp 

::c::
worksapce = "/c/libhulo"
---@link __tmp
file_replace(root .. worksapce .. "/hulo.pb.go", "package hulo", "package main")
file_replace(root .. worksapce .. "/hulo_grpc.pb.go", "package hulo", "package main")
local repo = "https://raw.githubusercontent.com/DaveGamble/cJSON/master/"
local libs = { "cJSON.c", "cJSON.h" }
for _, lib in ipairs(libs) do
    curl(repo .. lib, root .. worksapce .. "/" .. lib)
end

::java::
local plugin = "D:\\D\\langs\\pkg\\protoc-23.0-rc-1-win64\\bin\\protoc-gen-grpc-java-1.9.1-windows-x86_64.exe"
local worksapce = "/java/src/main/java/"
protoc("--plugin=protoc-gen-grpc-java=" .. plugin,
    "--grpc-java-out" .. root .. worksapce,
    "--java_out=" .. root .. worksapce,
    "--proto_path=.",
    target
)

::dart::
local worksapce = "/dart"
protoc("--dart_out=grpc:" .. root .. worksapce, "-I.", target)
