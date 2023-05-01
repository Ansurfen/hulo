package cs;

import proto.HuloClient;
import proto.Hulo.CallRequest;

public class client {
    public static void main(String[] args) {
        HuloClient cli = new HuloClient(9090);
        System.out.println(cli.Call(CallRequest.newBuilder().setFunc("SayHello").build()));
    }
}
