---@diagnostic disable: undefined-global

local gcc = function(str)
    return "gcc " + str
end
local target = { "libhulo.o", "hulo.c" }
Yock({
    ["*"] = function()
        -- 默认执行的
    end,
    clean = function()
        -- yock.exec({
        --     "rm cJSON.c",
        --     "rm cJSON.h",
        --     "rm libhulo.a",
        --     "rm libhulo.dll",
        --     "rm libhulo.h",
        -- })
        yock.rm("cJSON.c", "cJSON.h", "libhulo.a", "libhulo.dll", "libhulo.h")
    end,
    deploy = function()
        yock.curl("https://raw.githubusercontent.com/DaveGamble/cJSON/master/cJSON.c", "cJSON.c")
        yock.curl("https://raw.githubusercontent.com/DaveGamble/cJSON/master/cJSON.h", "cJSON.h")
    end,
    build = function()
        yock.exec({
            "gcc -c ./hulo.c -o libhulo.o",
            "ar rcs libhulo.a libhulo.o",
            "go build -o libhulo.dll -buildmode=c-shared",
            "go build -o libhulo.a -buildmode=c-archive",
            "rm libhulo.o",
            gcc("-c ./hulo.c -o libhulo.o" + target),
        })
    end
})
