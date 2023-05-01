package cs;

import proto.HuloBuilder;
import proto.Hulo.CallRequest;
import proto.Hulo.CallResponse;
import proto.HuloBuilder.Call;

public class server {
    @Call(name = "SayHello")
    public CallResponse sayHello(CallRequest req) {
        System.out.println(req);
        return CallResponse.newBuilder().setBuf("I'm Java").build();
    }

    public static void main(String[] args) {
        new HuloBuilder(server.class, args).run();
    }
}
