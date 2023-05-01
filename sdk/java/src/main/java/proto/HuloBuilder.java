package proto;

import java.io.IOException;
import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;
import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;
import java.util.function.Function;

import io.grpc.Server;
import io.grpc.ServerBuilder;
import proto.Hulo.CallRequest;
import proto.Hulo.CallResponse;

public class HuloBuilder {
    private Server server;

    @Retention(RetentionPolicy.RUNTIME)
    @Target(ElementType.METHOD)
    public @interface Call {
        String name();
    }

    public HuloBuilder(final Class<?> cls, String[] args) {
        int port = 0;
        for (int i = 0; i < args.length; i++) {
            if (args[i].equals("-p")) {
                if (i + 1 < args.length) {
                    try {
                        port = Integer.parseInt(args[i + 1]);
                    } catch (NumberFormatException e) {
                        System.err.println("invalid port");
                        System.exit(1);
                    }
                } else {
                    System.err.println("Missing port number argument after -p");
                    System.exit(1);
                }
            }
        }
        if (port == 0) {
            System.err.println("invalid port");
            System.exit(1);
        }
        HuloInterface hulo = new HuloInterface();
        Method[] methods = cls.getDeclaredMethods();
        for (final Method method : methods) {
            Call annotation = method.getAnnotation(Call.class);
            if (annotation != null) {
                hulo.register(annotation.name(), new Function<CallRequest, CallResponse>() {
                    public CallResponse apply(CallRequest requset) {
                        try {
                            return (CallResponse) method.invoke(cls.getDeclaredConstructor().newInstance(),
                                    requset);
                        } catch (IllegalAccessException | IllegalArgumentException | InvocationTargetException
                                | InstantiationException | NoSuchMethodException | SecurityException e) {
                            System.out.println(e);
                            return CallResponse.newBuilder().setBuf("sys err").build();
                        }
                    }
                });
            }
        }
        try {
            server = ServerBuilder.forPort(port).addService(hulo).build().start();
            Runtime.getRuntime().addShutdownHook(new Thread() {
                @Override
                public void run() {
                    System.out.println("close service...");
                    HuloBuilder.this.stop();
                }
            });
        } catch (IOException e) {
            System.exit(1);
        }
    }

    private void stop() {
        if (server != null)
            server.shutdown();
    }

    private void blockUntilShutdown() throws InterruptedException {
        if (server != null)
            server.awaitTermination();
    }

    public void run() {
        try {
            System.out.println("start service...");
            this.blockUntilShutdown();
        } catch (InterruptedException e) {
        }
    }
}
