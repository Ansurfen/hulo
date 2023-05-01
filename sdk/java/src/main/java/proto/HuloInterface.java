package proto;

import java.util.Map;
import java.util.TreeMap;
import java.util.function.Function;

import io.grpc.stub.StreamObserver;
import proto.Hulo.CallRequest;
import proto.Hulo.CallResponse;
import proto.Hulo.CompletionRequest;
import proto.Hulo.CompletionResponse;
import proto.HuloInterfaceGrpc.HuloInterfaceImplBase;

public class HuloInterface extends HuloInterfaceImplBase {
    private Map<String, Function<CallRequest, CallResponse>> dict;

    public HuloInterface() {
        this.dict = new TreeMap<String, Function<CallRequest, CallResponse>>();
    }

    public void register(String name, Function<CallRequest, CallResponse> cb) {
        this.dict.put(name, cb);
    }

    @Override
    public void call(CallRequest request, StreamObserver<CallResponse> responseObserver) {
        CallResponse res;
        if (this.dict.containsKey(request.getFunc())) {
            res = this.dict.get(request.getFunc()).apply(request);
        } else {
            res = CallResponse.newBuilder().setBuf("unknown").build();
        }
        responseObserver.onNext(res);
        responseObserver.onCompleted();
    }

    @Override
    public void completion(CompletionRequest request, StreamObserver<CompletionResponse> responseObserver) {
        super.completion(request, responseObserver);
    }
}
