const HuloBuilder = require("./hulo").HuloBuilder

var builder = new HuloBuilder()

builder.Call("SayHello", function (request) {
    console.log(request)
    return { Buf: "I'm nodejs" }
})
builder.run()
