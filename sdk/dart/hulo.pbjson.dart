///
//  Generated code. Do not modify.
//  source: hulo.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use completionRequestDescriptor instead')
const CompletionRequest$json = const {
  '1': 'CompletionRequest',
  '2': const [
    const {'1': 'str', '3': 1, '4': 1, '5': 9, '10': 'str'},
  ],
};

/// Descriptor for `CompletionRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List completionRequestDescriptor = $convert.base64Decode('ChFDb21wbGV0aW9uUmVxdWVzdBIQCgNzdHIYASABKAlSA3N0cg==');
@$core.Deprecated('Use completionResponseDescriptor instead')
const CompletionResponse$json = const {
  '1': 'CompletionResponse',
  '2': const [
    const {'1': 'suggests', '3': 1, '4': 3, '5': 11, '6': '.Hulo.Suggest', '10': 'suggests'},
  ],
};

/// Descriptor for `CompletionResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List completionResponseDescriptor = $convert.base64Decode('ChJDb21wbGV0aW9uUmVzcG9uc2USKQoIc3VnZ2VzdHMYASADKAsyDS5IdWxvLlN1Z2dlc3RSCHN1Z2dlc3Rz');
@$core.Deprecated('Use suggestDescriptor instead')
const Suggest$json = const {
  '1': 'Suggest',
  '2': const [
    const {'1': 'Text', '3': 1, '4': 1, '5': 9, '10': 'Text'},
    const {'1': 'Description', '3': 2, '4': 1, '5': 9, '10': 'Description'},
    const {'1': 'Comment', '3': 3, '4': 1, '5': 8, '10': 'Comment'},
  ],
};

/// Descriptor for `Suggest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List suggestDescriptor = $convert.base64Decode('CgdTdWdnZXN0EhIKBFRleHQYASABKAlSBFRleHQSIAoLRGVzY3JpcHRpb24YAiABKAlSC0Rlc2NyaXB0aW9uEhgKB0NvbW1lbnQYAyABKAhSB0NvbW1lbnQ=');
@$core.Deprecated('Use callRequestDescriptor instead')
const CallRequest$json = const {
  '1': 'CallRequest',
  '2': const [
    const {'1': 'Func', '3': 1, '4': 1, '5': 9, '10': 'Func'},
    const {'1': 'Arg', '3': 2, '4': 1, '5': 9, '10': 'Arg'},
  ],
};

/// Descriptor for `CallRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List callRequestDescriptor = $convert.base64Decode('CgtDYWxsUmVxdWVzdBISCgRGdW5jGAEgASgJUgRGdW5jEhAKA0FyZxgCIAEoCVIDQXJn');
@$core.Deprecated('Use callResponseDescriptor instead')
const CallResponse$json = const {
  '1': 'CallResponse',
  '2': const [
    const {'1': 'Ok', '3': 1, '4': 1, '5': 8, '10': 'Ok'},
    const {'1': 'Buf', '3': 2, '4': 1, '5': 9, '10': 'Buf'},
  ],
};

/// Descriptor for `CallResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List callResponseDescriptor = $convert.base64Decode('CgxDYWxsUmVzcG9uc2USDgoCT2sYASABKAhSAk9rEhAKA0J1ZhgCIAEoCVIDQnVm');
