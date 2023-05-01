///
//  Generated code. Do not modify.
//  source: hulo.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class CompletionRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'CompletionRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Hulo'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'str')
    ..hasRequiredFields = false
  ;

  CompletionRequest._() : super();
  factory CompletionRequest({
    $core.String? str,
  }) {
    final _result = create();
    if (str != null) {
      _result.str = str;
    }
    return _result;
  }
  factory CompletionRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CompletionRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  CompletionRequest clone() => CompletionRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  CompletionRequest copyWith(void Function(CompletionRequest) updates) => super.copyWith((message) => updates(message as CompletionRequest)) as CompletionRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static CompletionRequest create() => CompletionRequest._();
  CompletionRequest createEmptyInstance() => create();
  static $pb.PbList<CompletionRequest> createRepeated() => $pb.PbList<CompletionRequest>();
  @$core.pragma('dart2js:noInline')
  static CompletionRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CompletionRequest>(create);
  static CompletionRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get str => $_getSZ(0);
  @$pb.TagNumber(1)
  set str($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasStr() => $_has(0);
  @$pb.TagNumber(1)
  void clearStr() => clearField(1);
}

class CompletionResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'CompletionResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Hulo'), createEmptyInstance: create)
    ..pc<Suggest>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'suggests', $pb.PbFieldType.PM, subBuilder: Suggest.create)
    ..hasRequiredFields = false
  ;

  CompletionResponse._() : super();
  factory CompletionResponse({
    $core.Iterable<Suggest>? suggests,
  }) {
    final _result = create();
    if (suggests != null) {
      _result.suggests.addAll(suggests);
    }
    return _result;
  }
  factory CompletionResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CompletionResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  CompletionResponse clone() => CompletionResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  CompletionResponse copyWith(void Function(CompletionResponse) updates) => super.copyWith((message) => updates(message as CompletionResponse)) as CompletionResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static CompletionResponse create() => CompletionResponse._();
  CompletionResponse createEmptyInstance() => create();
  static $pb.PbList<CompletionResponse> createRepeated() => $pb.PbList<CompletionResponse>();
  @$core.pragma('dart2js:noInline')
  static CompletionResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CompletionResponse>(create);
  static CompletionResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Suggest> get suggests => $_getList(0);
}

class Suggest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Suggest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Hulo'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Text', protoName: 'Text')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Description', protoName: 'Description')
    ..aOB(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Comment', protoName: 'Comment')
    ..hasRequiredFields = false
  ;

  Suggest._() : super();
  factory Suggest({
    $core.String? text,
    $core.String? description,
    $core.bool? comment,
  }) {
    final _result = create();
    if (text != null) {
      _result.text = text;
    }
    if (description != null) {
      _result.description = description;
    }
    if (comment != null) {
      _result.comment = comment;
    }
    return _result;
  }
  factory Suggest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Suggest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Suggest clone() => Suggest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Suggest copyWith(void Function(Suggest) updates) => super.copyWith((message) => updates(message as Suggest)) as Suggest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Suggest create() => Suggest._();
  Suggest createEmptyInstance() => create();
  static $pb.PbList<Suggest> createRepeated() => $pb.PbList<Suggest>();
  @$core.pragma('dart2js:noInline')
  static Suggest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Suggest>(create);
  static Suggest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get text => $_getSZ(0);
  @$pb.TagNumber(1)
  set text($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasText() => $_has(0);
  @$pb.TagNumber(1)
  void clearText() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get description => $_getSZ(1);
  @$pb.TagNumber(2)
  set description($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasDescription() => $_has(1);
  @$pb.TagNumber(2)
  void clearDescription() => clearField(2);

  @$pb.TagNumber(3)
  $core.bool get comment => $_getBF(2);
  @$pb.TagNumber(3)
  set comment($core.bool v) { $_setBool(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasComment() => $_has(2);
  @$pb.TagNumber(3)
  void clearComment() => clearField(3);
}

class CallRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'CallRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Hulo'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Func', protoName: 'Func')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Arg', protoName: 'Arg')
    ..hasRequiredFields = false
  ;

  CallRequest._() : super();
  factory CallRequest({
    $core.String? func,
    $core.String? arg,
  }) {
    final _result = create();
    if (func != null) {
      _result.func = func;
    }
    if (arg != null) {
      _result.arg = arg;
    }
    return _result;
  }
  factory CallRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CallRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  CallRequest clone() => CallRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  CallRequest copyWith(void Function(CallRequest) updates) => super.copyWith((message) => updates(message as CallRequest)) as CallRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static CallRequest create() => CallRequest._();
  CallRequest createEmptyInstance() => create();
  static $pb.PbList<CallRequest> createRepeated() => $pb.PbList<CallRequest>();
  @$core.pragma('dart2js:noInline')
  static CallRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CallRequest>(create);
  static CallRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get func => $_getSZ(0);
  @$pb.TagNumber(1)
  set func($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasFunc() => $_has(0);
  @$pb.TagNumber(1)
  void clearFunc() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get arg => $_getSZ(1);
  @$pb.TagNumber(2)
  set arg($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasArg() => $_has(1);
  @$pb.TagNumber(2)
  void clearArg() => clearField(2);
}

class CallResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'CallResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'Hulo'), createEmptyInstance: create)
    ..aOB(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Ok', protoName: 'Ok')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Buf', protoName: 'Buf')
    ..hasRequiredFields = false
  ;

  CallResponse._() : super();
  factory CallResponse({
    $core.bool? ok,
    $core.String? buf,
  }) {
    final _result = create();
    if (ok != null) {
      _result.ok = ok;
    }
    if (buf != null) {
      _result.buf = buf;
    }
    return _result;
  }
  factory CallResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CallResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  CallResponse clone() => CallResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  CallResponse copyWith(void Function(CallResponse) updates) => super.copyWith((message) => updates(message as CallResponse)) as CallResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static CallResponse create() => CallResponse._();
  CallResponse createEmptyInstance() => create();
  static $pb.PbList<CallResponse> createRepeated() => $pb.PbList<CallResponse>();
  @$core.pragma('dart2js:noInline')
  static CallResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CallResponse>(create);
  static CallResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.bool get ok => $_getBF(0);
  @$pb.TagNumber(1)
  set ok($core.bool v) { $_setBool(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasOk() => $_has(0);
  @$pb.TagNumber(1)
  void clearOk() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get buf => $_getSZ(1);
  @$pb.TagNumber(2)
  set buf($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasBuf() => $_has(1);
  @$pb.TagNumber(2)
  void clearBuf() => clearField(2);
}

