package proto;

import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ClientCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ClientCalls.asyncClientStreamingCall;
import static io.grpc.stub.ClientCalls.asyncServerStreamingCall;
import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.blockingServerStreamingCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.stub.ServerCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ServerCalls.asyncClientStreamingCall;
import static io.grpc.stub.ServerCalls.asyncServerStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.9.1)",
    comments = "Source: hulo.proto")
public final class HuloInterfaceGrpc {

  private HuloInterfaceGrpc() {}

  public static final String SERVICE_NAME = "Hulo.HuloInterface";

  // Static method descriptors that strictly reflect the proto.
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  @java.lang.Deprecated // Use {@link #getCompletionMethod()} instead. 
  public static final io.grpc.MethodDescriptor<proto.Hulo.CompletionRequest,
      proto.Hulo.CompletionResponse> METHOD_COMPLETION = getCompletionMethod();

  private static volatile io.grpc.MethodDescriptor<proto.Hulo.CompletionRequest,
      proto.Hulo.CompletionResponse> getCompletionMethod;

  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static io.grpc.MethodDescriptor<proto.Hulo.CompletionRequest,
      proto.Hulo.CompletionResponse> getCompletionMethod() {
    io.grpc.MethodDescriptor<proto.Hulo.CompletionRequest, proto.Hulo.CompletionResponse> getCompletionMethod;
    if ((getCompletionMethod = HuloInterfaceGrpc.getCompletionMethod) == null) {
      synchronized (HuloInterfaceGrpc.class) {
        if ((getCompletionMethod = HuloInterfaceGrpc.getCompletionMethod) == null) {
          HuloInterfaceGrpc.getCompletionMethod = getCompletionMethod = 
              io.grpc.MethodDescriptor.<proto.Hulo.CompletionRequest, proto.Hulo.CompletionResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "Hulo.HuloInterface", "Completion"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.Hulo.CompletionRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.Hulo.CompletionResponse.getDefaultInstance()))
                  .setSchemaDescriptor(new HuloInterfaceMethodDescriptorSupplier("Completion"))
                  .build();
          }
        }
     }
     return getCompletionMethod;
  }
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  @java.lang.Deprecated // Use {@link #getCallMethod()} instead. 
  public static final io.grpc.MethodDescriptor<proto.Hulo.CallRequest,
      proto.Hulo.CallResponse> METHOD_CALL = getCallMethod();

  private static volatile io.grpc.MethodDescriptor<proto.Hulo.CallRequest,
      proto.Hulo.CallResponse> getCallMethod;

  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static io.grpc.MethodDescriptor<proto.Hulo.CallRequest,
      proto.Hulo.CallResponse> getCallMethod() {
    io.grpc.MethodDescriptor<proto.Hulo.CallRequest, proto.Hulo.CallResponse> getCallMethod;
    if ((getCallMethod = HuloInterfaceGrpc.getCallMethod) == null) {
      synchronized (HuloInterfaceGrpc.class) {
        if ((getCallMethod = HuloInterfaceGrpc.getCallMethod) == null) {
          HuloInterfaceGrpc.getCallMethod = getCallMethod = 
              io.grpc.MethodDescriptor.<proto.Hulo.CallRequest, proto.Hulo.CallResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(
                  "Hulo.HuloInterface", "Call"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.Hulo.CallRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  proto.Hulo.CallResponse.getDefaultInstance()))
                  .setSchemaDescriptor(new HuloInterfaceMethodDescriptorSupplier("Call"))
                  .build();
          }
        }
     }
     return getCallMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static HuloInterfaceStub newStub(io.grpc.Channel channel) {
    return new HuloInterfaceStub(channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static HuloInterfaceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    return new HuloInterfaceBlockingStub(channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static HuloInterfaceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    return new HuloInterfaceFutureStub(channel);
  }

  /**
   */
  public static abstract class HuloInterfaceImplBase implements io.grpc.BindableService {

    /**
     */
    public void completion(proto.Hulo.CompletionRequest request,
        io.grpc.stub.StreamObserver<proto.Hulo.CompletionResponse> responseObserver) {
      asyncUnimplementedUnaryCall(getCompletionMethod(), responseObserver);
    }

    /**
     */
    public void call(proto.Hulo.CallRequest request,
        io.grpc.stub.StreamObserver<proto.Hulo.CallResponse> responseObserver) {
      asyncUnimplementedUnaryCall(getCallMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getCompletionMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                proto.Hulo.CompletionRequest,
                proto.Hulo.CompletionResponse>(
                  this, METHODID_COMPLETION)))
          .addMethod(
            getCallMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                proto.Hulo.CallRequest,
                proto.Hulo.CallResponse>(
                  this, METHODID_CALL)))
          .build();
    }
  }

  /**
   */
  public static final class HuloInterfaceStub extends io.grpc.stub.AbstractStub<HuloInterfaceStub> {
    private HuloInterfaceStub(io.grpc.Channel channel) {
      super(channel);
    }

    private HuloInterfaceStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected HuloInterfaceStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new HuloInterfaceStub(channel, callOptions);
    }

    /**
     */
    public void completion(proto.Hulo.CompletionRequest request,
        io.grpc.stub.StreamObserver<proto.Hulo.CompletionResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getCompletionMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void call(proto.Hulo.CallRequest request,
        io.grpc.stub.StreamObserver<proto.Hulo.CallResponse> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getCallMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class HuloInterfaceBlockingStub extends io.grpc.stub.AbstractStub<HuloInterfaceBlockingStub> {
    private HuloInterfaceBlockingStub(io.grpc.Channel channel) {
      super(channel);
    }

    private HuloInterfaceBlockingStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected HuloInterfaceBlockingStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new HuloInterfaceBlockingStub(channel, callOptions);
    }

    /**
     */
    public proto.Hulo.CompletionResponse completion(proto.Hulo.CompletionRequest request) {
      return blockingUnaryCall(
          getChannel(), getCompletionMethod(), getCallOptions(), request);
    }

    /**
     */
    public proto.Hulo.CallResponse call(proto.Hulo.CallRequest request) {
      return blockingUnaryCall(
          getChannel(), getCallMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class HuloInterfaceFutureStub extends io.grpc.stub.AbstractStub<HuloInterfaceFutureStub> {
    private HuloInterfaceFutureStub(io.grpc.Channel channel) {
      super(channel);
    }

    private HuloInterfaceFutureStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected HuloInterfaceFutureStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new HuloInterfaceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<proto.Hulo.CompletionResponse> completion(
        proto.Hulo.CompletionRequest request) {
      return futureUnaryCall(
          getChannel().newCall(getCompletionMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<proto.Hulo.CallResponse> call(
        proto.Hulo.CallRequest request) {
      return futureUnaryCall(
          getChannel().newCall(getCallMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_COMPLETION = 0;
  private static final int METHODID_CALL = 1;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final HuloInterfaceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(HuloInterfaceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_COMPLETION:
          serviceImpl.completion((proto.Hulo.CompletionRequest) request,
              (io.grpc.stub.StreamObserver<proto.Hulo.CompletionResponse>) responseObserver);
          break;
        case METHODID_CALL:
          serviceImpl.call((proto.Hulo.CallRequest) request,
              (io.grpc.stub.StreamObserver<proto.Hulo.CallResponse>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class HuloInterfaceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    HuloInterfaceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return proto.Hulo.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("HuloInterface");
    }
  }

  private static final class HuloInterfaceFileDescriptorSupplier
      extends HuloInterfaceBaseDescriptorSupplier {
    HuloInterfaceFileDescriptorSupplier() {}
  }

  private static final class HuloInterfaceMethodDescriptorSupplier
      extends HuloInterfaceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    HuloInterfaceMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (HuloInterfaceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new HuloInterfaceFileDescriptorSupplier())
              .addMethod(getCompletionMethod())
              .addMethod(getCallMethod())
              .build();
        }
      }
    }
    return result;
  }
}
