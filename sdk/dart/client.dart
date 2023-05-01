import 'package:grpc/service_api.dart';

import 'hulo.dart';
import 'hulo.pbgrpc.dart';

Future<void> main(List<String> args) async {
  var client = HuloClient(9090);
  print(await client.Call(CallRequest()..func = "SayHello"));
  await client.close();
}
