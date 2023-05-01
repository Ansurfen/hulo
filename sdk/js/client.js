const HuloClient = require("./hulo").HuloClient
var cli = new HuloClient(9090)
cli.Call({ Func: "SayHello" }, (request) => {
    console.log(request)
})