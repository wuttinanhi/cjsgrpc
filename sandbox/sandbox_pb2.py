# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: sandbox.proto
# Protobuf Python Version: 4.25.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\rsandbox.proto\x12\x07sandbox\"\xf1\x01\n\x14\x43reateSandboxRequest\x12\x10\n\x08language\x18\x01 \x01(\t\x12\x0c\n\x04\x63ode\x18\x02 \x01(\t\x12\r\n\x05stdin\x18\x03 \x01(\t\x12\x0f\n\x07timeout\x18\x04 \x01(\r\x12\x14\n\x0c\x64ocker_image\x18\x05 \x01(\t\x12\x17\n\x0f\x63ompile_command\x18\x06 \x01(\t\x12\x17\n\x0f\x63ompile_timeout\x18\x07 \x01(\r\x12\x14\n\x0cmemory_limit\x18\x08 \x01(\r\x12\x11\n\tcpu_limit\x18\t \x01(\r\x12\x13\n\x0brun_command\x18\n \x01(\t\x12\x13\n\x0brun_timeout\x18\x0b \x01(\r\"2\n\x15\x43reateSandboxResponse\x12\n\n\x02id\x18\x01 \x01(\t\x12\r\n\x05\x65rror\x18\x02 \x01(\t\"\x1f\n\x11GetSandboxRequest\x12\n\n\x02id\x18\x01 \x01(\t\"\x91\x01\n\x12GetSandboxResponse\x12\n\n\x02id\x18\x01 \x01(\t\x12\x10\n\x08language\x18\x02 \x01(\t\x12\x0c\n\x04\x63ode\x18\x03 \x01(\t\x12\r\n\x05stdin\x18\x04 \x01(\t\x12\x0e\n\x06stdout\x18\x05 \x01(\t\x12\x0e\n\x06stderr\x18\x06 \x01(\t\x12\x11\n\texit_code\x18\x08 \x01(\r\x12\r\n\x05\x65rror\x18\x07 \x01(\t2\xa0\x01\n\x07Sandbox\x12N\n\rCreateSandbox\x12\x1d.sandbox.CreateSandboxRequest\x1a\x1e.sandbox.CreateSandboxResponse\x12\x45\n\nGetSandbox\x12\x1a.sandbox.GetSandboxRequest\x1a\x1b.sandbox.GetSandboxResponseB2Z0github.com/wuttinanhi/cjsgrpc/sandbox/go/sandboxb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'sandbox_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z0github.com/wuttinanhi/cjsgrpc/sandbox/go/sandbox'
  _globals['_CREATESANDBOXREQUEST']._serialized_start=27
  _globals['_CREATESANDBOXREQUEST']._serialized_end=268
  _globals['_CREATESANDBOXRESPONSE']._serialized_start=270
  _globals['_CREATESANDBOXRESPONSE']._serialized_end=320
  _globals['_GETSANDBOXREQUEST']._serialized_start=322
  _globals['_GETSANDBOXREQUEST']._serialized_end=353
  _globals['_GETSANDBOXRESPONSE']._serialized_start=356
  _globals['_GETSANDBOXRESPONSE']._serialized_end=501
  _globals['_SANDBOX']._serialized_start=504
  _globals['_SANDBOX']._serialized_end=664
# @@protoc_insertion_point(module_scope)
