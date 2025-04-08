from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class RosTypeGenType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    MESSAGE: _ClassVar[RosTypeGenType]
    SERVICE: _ClassVar[RosTypeGenType]
MESSAGE: RosTypeGenType
SERVICE: RosTypeGenType

class Empty(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class LinkCallReq(_message.Message):
    __slots__ = ("cid", "method", "data")
    CID_FIELD_NUMBER: _ClassVar[int]
    METHOD_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    cid: str
    method: str
    data: str
    def __init__(self, cid: _Optional[str] = ..., method: _Optional[str] = ..., data: _Optional[str] = ...) -> None: ...

class LinkCallRsp(_message.Message):
    __slots__ = ("data",)
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: str
    def __init__(self, data: _Optional[str] = ...) -> None: ...

class LinkCommandDownstream(_message.Message):
    __slots__ = ("id", "payload_ros_exec", "payload_ros_type", "payload_ros_list")
    ID_FIELD_NUMBER: _ClassVar[int]
    PAYLOAD_ROS_EXEC_FIELD_NUMBER: _ClassVar[int]
    PAYLOAD_ROS_TYPE_FIELD_NUMBER: _ClassVar[int]
    PAYLOAD_ROS_LIST_FIELD_NUMBER: _ClassVar[int]
    id: str
    payload_ros_exec: LinkCommandPayloadRosExec
    payload_ros_type: LinkCommandPayloadRosType
    payload_ros_list: LinkCommandPayloadRosList
    def __init__(self, id: _Optional[str] = ..., payload_ros_exec: _Optional[_Union[LinkCommandPayloadRosExec, _Mapping]] = ..., payload_ros_type: _Optional[_Union[LinkCommandPayloadRosType, _Mapping]] = ..., payload_ros_list: _Optional[_Union[LinkCommandPayloadRosList, _Mapping]] = ...) -> None: ...

class LinkCommandPayloadRosExec(_message.Message):
    __slots__ = ("action", "ros_topic", "data")
    ACTION_FIELD_NUMBER: _ClassVar[int]
    ROS_TOPIC_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    action: str
    ros_topic: str
    data: str
    def __init__(self, action: _Optional[str] = ..., ros_topic: _Optional[str] = ..., data: _Optional[str] = ...) -> None: ...

class LinkCommandPayloadRosType(_message.Message):
    __slots__ = ("type", "ros_topic")
    TYPE_FIELD_NUMBER: _ClassVar[int]
    ROS_TOPIC_FIELD_NUMBER: _ClassVar[int]
    type: RosTypeGenType
    ros_topic: str
    def __init__(self, type: _Optional[_Union[RosTypeGenType, str]] = ..., ros_topic: _Optional[str] = ...) -> None: ...

class LinkCommandPayloadRosList(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class LinkCommandUpstream(_message.Message):
    __slots__ = ("id", "ok", "cdr_data", "err_msg", "cid", "type_gen_result", "type_list_result")
    ID_FIELD_NUMBER: _ClassVar[int]
    OK_FIELD_NUMBER: _ClassVar[int]
    CDR_DATA_FIELD_NUMBER: _ClassVar[int]
    ERR_MSG_FIELD_NUMBER: _ClassVar[int]
    CID_FIELD_NUMBER: _ClassVar[int]
    TYPE_GEN_RESULT_FIELD_NUMBER: _ClassVar[int]
    TYPE_LIST_RESULT_FIELD_NUMBER: _ClassVar[int]
    id: str
    ok: bool
    cdr_data: str
    err_msg: str
    cid: str
    type_gen_result: LinkTypeGenResult
    type_list_result: TypeListRsp
    def __init__(self, id: _Optional[str] = ..., ok: bool = ..., cdr_data: _Optional[str] = ..., err_msg: _Optional[str] = ..., cid: _Optional[str] = ..., type_gen_result: _Optional[_Union[LinkTypeGenResult, _Mapping]] = ..., type_list_result: _Optional[_Union[TypeListRsp, _Mapping]] = ...) -> None: ...

class LinkTypeGenResult(_message.Message):
    __slots__ = ("name", "req", "rsp")
    NAME_FIELD_NUMBER: _ClassVar[int]
    REQ_FIELD_NUMBER: _ClassVar[int]
    RSP_FIELD_NUMBER: _ClassVar[int]
    name: str
    req: str
    rsp: str
    def __init__(self, name: _Optional[str] = ..., req: _Optional[str] = ..., rsp: _Optional[str] = ...) -> None: ...

class TypeListReq(_message.Message):
    __slots__ = ("cid",)
    CID_FIELD_NUMBER: _ClassVar[int]
    cid: str
    def __init__(self, cid: _Optional[str] = ...) -> None: ...

class TypeListRsp(_message.Message):
    __slots__ = ("messages", "services")
    MESSAGES_FIELD_NUMBER: _ClassVar[int]
    SERVICES_FIELD_NUMBER: _ClassVar[int]
    messages: _containers.RepeatedScalarFieldContainer[str]
    services: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, messages: _Optional[_Iterable[str]] = ..., services: _Optional[_Iterable[str]] = ...) -> None: ...

class TypeGenReq(_message.Message):
    __slots__ = ("cid", "topic_regexp")
    CID_FIELD_NUMBER: _ClassVar[int]
    TOPIC_REGEXP_FIELD_NUMBER: _ClassVar[int]
    cid: str
    topic_regexp: str
    def __init__(self, cid: _Optional[str] = ..., topic_regexp: _Optional[str] = ...) -> None: ...

class TypeGenRsp(_message.Message):
    __slots__ = ("pb",)
    class PbEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    PB_FIELD_NUMBER: _ClassVar[int]
    pb: _containers.ScalarMap[str, str]
    def __init__(self, pb: _Optional[_Mapping[str, str]] = ...) -> None: ...

class TypeSchema(_message.Message):
    __slots__ = ("ros_schema", "pb_schema", "pb_type_name")
    ROS_SCHEMA_FIELD_NUMBER: _ClassVar[int]
    PB_SCHEMA_FIELD_NUMBER: _ClassVar[int]
    PB_TYPE_NAME_FIELD_NUMBER: _ClassVar[int]
    ros_schema: str
    pb_schema: str
    pb_type_name: str
    def __init__(self, ros_schema: _Optional[str] = ..., pb_schema: _Optional[str] = ..., pb_type_name: _Optional[str] = ...) -> None: ...

class AgentListRsp(_message.Message):
    __slots__ = ("agents",)
    class AgentsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    AGENTS_FIELD_NUMBER: _ClassVar[int]
    agents: _containers.ScalarMap[str, str]
    def __init__(self, agents: _Optional[_Mapping[str, str]] = ...) -> None: ...

class AgentListenReq(_message.Message):
    __slots__ = ("cid",)
    CID_FIELD_NUMBER: _ClassVar[int]
    cid: str
    def __init__(self, cid: _Optional[str] = ...) -> None: ...

class AgentListenRsp(_message.Message):
    __slots__ = ("is_upstream", "topic", "json_data")
    IS_UPSTREAM_FIELD_NUMBER: _ClassVar[int]
    TOPIC_FIELD_NUMBER: _ClassVar[int]
    JSON_DATA_FIELD_NUMBER: _ClassVar[int]
    is_upstream: bool
    topic: str
    json_data: str
    def __init__(self, is_upstream: bool = ..., topic: _Optional[str] = ..., json_data: _Optional[str] = ...) -> None: ...
