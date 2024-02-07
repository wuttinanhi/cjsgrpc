from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class CreateSandboxRequest(_message.Message):
    __slots__ = ("language", "code", "stdin", "timeout", "docker_image", "compile_command", "compile_timeout", "memory_limit", "cpu_limit", "run_command", "run_timeout")
    LANGUAGE_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    STDIN_FIELD_NUMBER: _ClassVar[int]
    TIMEOUT_FIELD_NUMBER: _ClassVar[int]
    DOCKER_IMAGE_FIELD_NUMBER: _ClassVar[int]
    COMPILE_COMMAND_FIELD_NUMBER: _ClassVar[int]
    COMPILE_TIMEOUT_FIELD_NUMBER: _ClassVar[int]
    MEMORY_LIMIT_FIELD_NUMBER: _ClassVar[int]
    CPU_LIMIT_FIELD_NUMBER: _ClassVar[int]
    RUN_COMMAND_FIELD_NUMBER: _ClassVar[int]
    RUN_TIMEOUT_FIELD_NUMBER: _ClassVar[int]
    language: str
    code: str
    stdin: str
    timeout: int
    docker_image: str
    compile_command: str
    compile_timeout: int
    memory_limit: int
    cpu_limit: int
    run_command: str
    run_timeout: int
    def __init__(self, language: _Optional[str] = ..., code: _Optional[str] = ..., stdin: _Optional[str] = ..., timeout: _Optional[int] = ..., docker_image: _Optional[str] = ..., compile_command: _Optional[str] = ..., compile_timeout: _Optional[int] = ..., memory_limit: _Optional[int] = ..., cpu_limit: _Optional[int] = ..., run_command: _Optional[str] = ..., run_timeout: _Optional[int] = ...) -> None: ...

class CreateSandboxResponse(_message.Message):
    __slots__ = ("id", "error")
    ID_FIELD_NUMBER: _ClassVar[int]
    ERROR_FIELD_NUMBER: _ClassVar[int]
    id: str
    error: str
    def __init__(self, id: _Optional[str] = ..., error: _Optional[str] = ...) -> None: ...

class GetSandboxRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class GetSandboxResponse(_message.Message):
    __slots__ = ("id", "language", "code", "stdin", "stdout", "stderr", "exit_code", "error")
    ID_FIELD_NUMBER: _ClassVar[int]
    LANGUAGE_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    STDIN_FIELD_NUMBER: _ClassVar[int]
    STDOUT_FIELD_NUMBER: _ClassVar[int]
    STDERR_FIELD_NUMBER: _ClassVar[int]
    EXIT_CODE_FIELD_NUMBER: _ClassVar[int]
    ERROR_FIELD_NUMBER: _ClassVar[int]
    id: str
    language: str
    code: str
    stdin: str
    stdout: str
    stderr: str
    exit_code: int
    error: str
    def __init__(self, id: _Optional[str] = ..., language: _Optional[str] = ..., code: _Optional[str] = ..., stdin: _Optional[str] = ..., stdout: _Optional[str] = ..., stderr: _Optional[str] = ..., exit_code: _Optional[int] = ..., error: _Optional[str] = ...) -> None: ...
