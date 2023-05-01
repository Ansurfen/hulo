const hulo_proto = require('./proto')
const grpc = require('@grpc/grpc-js')

/**
 * @typedef {object} CallRequest
 * @property {String} Func function name
 * @property {String} Arg function argument
 */

/**
 * @typedef {object} CallResponse
 * @property {Boolean} Ok
 * @property {String} Buf
 */

class HuloClient {
    client

    /**
     * 
     * @param {number} port 
     */
    constructor(port) {
        this.client = new hulo_proto.HuloInterface(`localhost:${port}`, grpc.credentials.createInsecure())
    }

    /**
     * 
     * @param {CallRequest} request 
     * @param {Function} handle
     */
    Call(request, handle) {
        if (this.client == null) {
            handle({ Ok: false, Buf: "sys err" })
        }
        this.client.Call(request, function (err, response) {
            if (err) {
                handle({ Ok: false, Buf: "sys err" })
            }
            handle(response)
        })
    }
}

class HuloBuilder {
    server
    dict

    constructor() {
        this.server = new grpc.Server()
        this.dict = new Map()
    }

    callHandle(context, callback) {
        console.log()
        let huloCall = this.dict.get(context.request.Func)
        if (huloCall !== undefined) {
            return callback(null, huloCall(context.request))
        }
        return callback(null, { Ok: false, Buf: "unknown" })
    }

    /**
     * 
     * @param {String} func 
     * @param {Function} huloCall 
     */
    Call(func, huloCall) {
        this.dict.set(func, huloCall)
    }

    run() {
        this.server.addService(hulo_proto.HuloInterface.service, { Call: this.callHandle.bind(this) })
        const args = process.argv.slice(2);
        let port = 0;
        for (let i = 0; i < args.length; i++) {
            if (args[i] === '-p') {
                port = parseInt(args[i + 1], 10);
                break;
            }
        }
        if (port === 0) {
            console.log("invalid port")
            process.exit(1);
        }
        this.server.bindAsync(`localhost:${port}`, grpc.ServerCredentials.createInsecure(), () => {
            this.server.start()
            console.log('grpc server started')
        })
    }
}

module.exports = {
    HuloBuilder: HuloBuilder,
    HuloClient: HuloClient
}