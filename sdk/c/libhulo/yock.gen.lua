---@diagnostic disable: undefined-global
::build::
exec({
    "gcc -c ./hulo.c -o libhulo.o",
    "ar rcs libhulo.a libhulo.o",
    "go build -o libhulo.dll -buildmode=c-shared",
    "go build -o libhulo.a -buildmode=c-archive",
    "rm libhulo.o"
})

::clean::
exec({
    "rm cJSON.c",
    "rm cJSON.h",
    "rm libhulo.a",
    "rm libhulo.dll",
    "rm libhulo.h",
    "rm hulo_grpc.pb.go",
    "rm hulo.pb.go",
})

::deploy::
-- deprecated
exec({ "yock ..\\..\\interface\\yock.gen.lua c" })

