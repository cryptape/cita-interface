# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: snapshot.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='snapshot.proto',
  package='',
  syntax='proto3',
  serialized_pb=_b('\n\x0esnapshot.proto\"X\n\x0bSnapshotReq\x12\x11\n\x03\x63md\x18\x01 \x01(\x0e\x32\x04.Cmd\x12\x14\n\x0cstart_height\x18\x02 \x01(\x04\x12\x12\n\nend_height\x18\x03 \x01(\x04\x12\x0c\n\x04\x66ile\x18\x04 \x01(\t\"#\n\x0cSnapshotResp\x12\x13\n\x04resp\x18\x01 \x01(\x0e\x32\x05.Resp*?\n\x03\x43md\x12\t\n\x05\x42\x65gin\x10\x00\x12\t\n\x05\x43lear\x10\x01\x12\x0c\n\x08Snapshot\x10\x02\x12\x0b\n\x07Restore\x10\x03\x12\x07\n\x03\x45nd\x10\x04*O\n\x04Resp\x12\x0c\n\x08\x42\x65ginAck\x10\x00\x12\x0c\n\x08\x43learAck\x10\x01\x12\x0f\n\x0bSnapshotAck\x10\x02\x12\x0e\n\nRestoreAck\x10\x03\x12\n\n\x06\x45ndAck\x10\x04\x62\x06proto3')
)

_CMD = _descriptor.EnumDescriptor(
  name='Cmd',
  full_name='Cmd',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='Begin', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Clear', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Snapshot', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Restore', index=3, number=3,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='End', index=4, number=4,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=145,
  serialized_end=208,
)
_sym_db.RegisterEnumDescriptor(_CMD)

Cmd = enum_type_wrapper.EnumTypeWrapper(_CMD)
_RESP = _descriptor.EnumDescriptor(
  name='Resp',
  full_name='Resp',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='BeginAck', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ClearAck', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SnapshotAck', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='RestoreAck', index=3, number=3,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='EndAck', index=4, number=4,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=210,
  serialized_end=289,
)
_sym_db.RegisterEnumDescriptor(_RESP)

Resp = enum_type_wrapper.EnumTypeWrapper(_RESP)
Begin = 0
Clear = 1
Snapshot = 2
Restore = 3
End = 4
BeginAck = 0
ClearAck = 1
SnapshotAck = 2
RestoreAck = 3
EndAck = 4



_SNAPSHOTREQ = _descriptor.Descriptor(
  name='SnapshotReq',
  full_name='SnapshotReq',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='cmd', full_name='SnapshotReq.cmd', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='start_height', full_name='SnapshotReq.start_height', index=1,
      number=2, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='end_height', full_name='SnapshotReq.end_height', index=2,
      number=3, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='file', full_name='SnapshotReq.file', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=18,
  serialized_end=106,
)


_SNAPSHOTRESP = _descriptor.Descriptor(
  name='SnapshotResp',
  full_name='SnapshotResp',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='resp', full_name='SnapshotResp.resp', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=108,
  serialized_end=143,
)

_SNAPSHOTREQ.fields_by_name['cmd'].enum_type = _CMD
_SNAPSHOTRESP.fields_by_name['resp'].enum_type = _RESP
DESCRIPTOR.message_types_by_name['SnapshotReq'] = _SNAPSHOTREQ
DESCRIPTOR.message_types_by_name['SnapshotResp'] = _SNAPSHOTRESP
DESCRIPTOR.enum_types_by_name['Cmd'] = _CMD
DESCRIPTOR.enum_types_by_name['Resp'] = _RESP
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

SnapshotReq = _reflection.GeneratedProtocolMessageType('SnapshotReq', (_message.Message,), dict(
  DESCRIPTOR = _SNAPSHOTREQ,
  __module__ = 'snapshot_pb2'
  # @@protoc_insertion_point(class_scope:SnapshotReq)
  ))
_sym_db.RegisterMessage(SnapshotReq)

SnapshotResp = _reflection.GeneratedProtocolMessageType('SnapshotResp', (_message.Message,), dict(
  DESCRIPTOR = _SNAPSHOTRESP,
  __module__ = 'snapshot_pb2'
  # @@protoc_insertion_point(class_scope:SnapshotResp)
  ))
_sym_db.RegisterMessage(SnapshotResp)


# @@protoc_insertion_point(module_scope)
