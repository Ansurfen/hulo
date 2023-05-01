import 'dart:async';
import 'dart:io';
import 'dart:mirrors';
import 'package:args/args.dart';
import 'package:grpc/grpc.dart';
import 'hulo.pbgrpc.dart';

typedef HuloCall = Future<CallResponse> Function(ServiceCall, CallRequest);

class HuloInterface extends HuloInterfaceServiceBase {
  late Map<String, HuloCall> dict;

  HuloInterface() {
    this.dict = new Map();
  }

  @override
  Future<CallResponse> call(ServiceCall call, CallRequest request) async {
    if (this.dict.containsKey(request.func)) {
      print(this.dict);
      return this.dict[request.func]!(call, request);
    }
    return CallResponse(buf: "unknown");
  }

  @override
  Future<CompletionResponse> completion(
      ServiceCall call, CompletionRequest request) {
    throw UnimplementedError();
  }

  void Register(String name, HuloCall cb) {
    this.dict[name] = cb;
  }
}

class Call {
  final String name;
  const Call(this.name);
}

class HuloBuilder {
  late final Server server;
  late int port;

  HuloBuilder(List<String> args) {
    HuloInterface hulo = new HuloInterface();
    final parser = ArgParser();
    parser.addOption('port', abbr: 'p');
    ArgResults result = parser.parse(args);
    try {
      port = int.parse(result['port']);
    } catch (e) {
      print("invalid port");
      exit(1);
    }
    currentMirrorSystem().libraries.forEach((_, lib) {
      lib.declarations.forEach((s, decl) {
        decl.metadata.where((m) => m.reflectee is Call).forEach((m) {
          var anno = m.reflectee as Call;
          if (decl is MethodMirror) {
            hulo.Register(anno.name, (ServiceCall call, CallRequest req) async {
              return ((decl).owner as LibraryMirror)
                  .invoke(s, [call, req]).reflectee;
            });
          }
          ;
        });
      });
    });
    server = Server([hulo]);
  }

  Future<void> Start() async {
    print("service start...");
    await this.server.serve(port: this.port);
  }
}

class HuloClient {
  late final ClientChannel channel;
  late final HuloInterfaceClient cli;

  HuloClient(int port) {
    this.channel = ClientChannel('localhost',
        port: port,
        options:
            const ChannelOptions(credentials: ChannelCredentials.insecure()));
    this.cli = HuloInterfaceClient(channel);
  }

  Future<CallResponse> Call(CallRequest request) {
    return this.cli.call(request);
  }

  Future<void> close() async {
    this.channel.shutdown();
  }
}
