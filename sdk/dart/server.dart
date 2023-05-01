import 'package:grpc/grpc.dart';
import 'hulo.pb.dart';
import 'hulo.dart';

@Call('SayHello')
Future<CallResponse> doSomething(ServiceCall call, CallRequest request) async {
  print(request);
  return new CallResponse(buf: "I'm Dart");
}

void main(List<String> args) {
  HuloBuilder(args).Start();
}
