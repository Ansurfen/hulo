///
//  Generated code. Do not modify.
//  source: hulo.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'hulo.pb.dart' as $0;
export 'hulo.pb.dart';

class HuloInterfaceClient extends $grpc.Client {
  static final _$completion =
      $grpc.ClientMethod<$0.CompletionRequest, $0.CompletionResponse>(
          '/Hulo.HuloInterface/Completion',
          ($0.CompletionRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.CompletionResponse.fromBuffer(value));
  static final _$call = $grpc.ClientMethod<$0.CallRequest, $0.CallResponse>(
      '/Hulo.HuloInterface/Call',
      ($0.CallRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.CallResponse.fromBuffer(value));

  HuloInterfaceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.CompletionResponse> completion(
      $0.CompletionRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$completion, request, options: options);
  }

  $grpc.ResponseFuture<$0.CallResponse> call($0.CallRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$call, request, options: options);
  }
}

abstract class HuloInterfaceServiceBase extends $grpc.Service {
  $core.String get $name => 'Hulo.HuloInterface';

  HuloInterfaceServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.CompletionRequest, $0.CompletionResponse>(
        'Completion',
        completion_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.CompletionRequest.fromBuffer(value),
        ($0.CompletionResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.CallRequest, $0.CallResponse>(
        'Call',
        call_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.CallRequest.fromBuffer(value),
        ($0.CallResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.CompletionResponse> completion_Pre($grpc.ServiceCall call,
      $async.Future<$0.CompletionRequest> request) async {
    return completion(call, await request);
  }

  $async.Future<$0.CallResponse> call_Pre(
      $grpc.ServiceCall _call, $async.Future<$0.CallRequest> request) async {
    return call(_call, await request);
  }

  $async.Future<$0.CompletionResponse> completion(
      $grpc.ServiceCall call, $0.CompletionRequest request);
  $async.Future<$0.CallResponse> call(
      $grpc.ServiceCall call, $0.CallRequest request);
}
