package proto;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import proto.Hulo.CallRequest;
import proto.Hulo.CallResponse;

public class HuloClient {
    ManagedChannel channel;

    public HuloClient(int port) {
        this.channel = ManagedChannelBuilder.forAddress("localhost", port).usePlaintext().build();
    }

    public CallResponse Call(CallRequest request) {
        try {
            HuloInterfaceGrpc.HuloInterfaceBlockingStub stub = HuloInterfaceGrpc.newBlockingStub(channel);
            return stub.call(request);
        } catch (Exception e) {
            return CallResponse.newBuilder().setBuf("sys err").build();
        }
    }
}
