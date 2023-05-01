from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CallRequest(_message.Message):
    __slots__ = ["Arg", "Func"]
    ARG_FIELD_NUMBER: _ClassVar[int]
    Arg: str
    FUNC_FIELD_NUMBER: _ClassVar[int]
    Func: str
    def __init__(self, Func: _Optional[str] = ..., Arg: _Optional[str] = ...) -> None: ...

class CallResponse(_message.Message):
    __slots__ = ["Buf", "Ok"]
    BUF_FIELD_NUMBER: _ClassVar[int]
    Buf: str
    OK_FIELD_NUMBER: _ClassVar[int]
    Ok: bool
    def __init__(self, Ok: bool = ..., Buf: _Optional[str] = ...) -> None: ...

class CompletionRequest(_message.Message):
    __slots__ = ["str"]
    STR_FIELD_NUMBER: _ClassVar[int]
    str: str
    def __init__(self, str: _Optional[str] = ...) -> None: ...

class CompletionResponse(_message.Message):
    __slots__ = ["suggests"]
    SUGGESTS_FIELD_NUMBER: _ClassVar[int]
    suggests: _containers.RepeatedCompositeFieldContainer[Suggest]
    def __init__(self, suggests: _Optional[_Iterable[_Union[Suggest, _Mapping]]] = ...) -> None: ...

class Suggest(_message.Message):
    __slots__ = ["Comment", "Description", "Text"]
    COMMENT_FIELD_NUMBER: _ClassVar[int]
    Comment: bool
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    Description: str
    TEXT_FIELD_NUMBER: _ClassVar[int]
    Text: str
    def __init__(self, Text: _Optional[str] = ..., Description: _Optional[str] = ..., Comment: bool = ...) -> None: ...
