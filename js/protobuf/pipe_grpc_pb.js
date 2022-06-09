// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var pipe_pb = require('./pipe_pb.js');
var messageenvelop_pb = require('./messageenvelop_pb.js');
var event_pb = require('./event_pb.js');
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');

function serialize_messages_Event(arg) {
  if (!(arg instanceof event_pb.Event)) {
    throw new Error('Expected argument of type messages.Event');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_messages_Event(buffer_arg) {
  return event_pb.Event.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_messages_Events(arg) {
  if (!(arg instanceof event_pb.Events)) {
    throw new Error('Expected argument of type messages.Events');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_messages_Events(buffer_arg) {
  return event_pb.Events.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_messages_MessageEnvelop(arg) {
  if (!(arg instanceof messageenvelop_pb.MessageEnvelop)) {
    throw new Error('Expected argument of type messages.MessageEnvelop');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_messages_MessageEnvelop(buffer_arg) {
  return messageenvelop_pb.MessageEnvelop.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_pipes_AddStepsRequest(arg) {
  if (!(arg instanceof pipe_pb.AddStepsRequest)) {
    throw new Error('Expected argument of type pipes.AddStepsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pipes_AddStepsRequest(buffer_arg) {
  return pipe_pb.AddStepsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_pipes_CompleteRequest(arg) {
  if (!(arg instanceof pipe_pb.CompleteRequest)) {
    throw new Error('Expected argument of type pipes.CompleteRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pipes_CompleteRequest(buffer_arg) {
  return pipe_pb.CompleteRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_pipes_Decorations(arg) {
  if (!(arg instanceof pipe_pb.Decorations)) {
    throw new Error('Expected argument of type pipes.Decorations');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pipes_Decorations(buffer_arg) {
  return pipe_pb.Decorations.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_pipes_GenericResponse(arg) {
  if (!(arg instanceof pipe_pb.GenericResponse)) {
    throw new Error('Expected argument of type pipes.GenericResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pipes_GenericResponse(buffer_arg) {
  return pipe_pb.GenericResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_pipes_GenericResponses(arg) {
  if (!(arg instanceof pipe_pb.GenericResponses)) {
    throw new Error('Expected argument of type pipes.GenericResponses');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pipes_GenericResponses(buffer_arg) {
  return pipe_pb.GenericResponses.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_pipes_ListRequest(arg) {
  if (!(arg instanceof pipe_pb.ListRequest)) {
    throw new Error('Expected argument of type pipes.ListRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pipes_ListRequest(buffer_arg) {
  return pipe_pb.ListRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_pipes_ReceiveOptions(arg) {
  if (!(arg instanceof pipe_pb.ReceiveOptions)) {
    throw new Error('Expected argument of type pipes.ReceiveOptions');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pipes_ReceiveOptions(buffer_arg) {
  return pipe_pb.ReceiveOptions.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_pipes_RouteLogRequest(arg) {
  if (!(arg instanceof pipe_pb.RouteLogRequest)) {
    throw new Error('Expected argument of type pipes.RouteLogRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pipes_RouteLogRequest(buffer_arg) {
  return pipe_pb.RouteLogRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_pipes_Xid(arg) {
  if (!(arg instanceof pipe_pb.Xid)) {
    throw new Error('Expected argument of type pipes.Xid');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pipes_Xid(buffer_arg) {
  return pipe_pb.Xid.deserializeBinary(new Uint8Array(buffer_arg));
}


var PipeService = exports.PipeService = {
  streamSend: {
    path: '/pipes.Pipe/StreamSend',
    requestStream: true,
    responseStream: true,
    requestType: messageenvelop_pb.MessageEnvelop,
    responseType: pipe_pb.Xid,
    requestSerialize: serialize_messages_MessageEnvelop,
    requestDeserialize: deserialize_messages_MessageEnvelop,
    responseSerialize: serialize_pipes_Xid,
    responseDeserialize: deserialize_pipes_Xid,
  },
  // Send will send up a message envelop, and will return an event id, error if invalid for any reason
send: {
    path: '/pipes.Pipe/Send',
    requestStream: false,
    responseStream: false,
    requestType: messageenvelop_pb.MessageEnvelop,
    responseType: pipe_pb.Xid,
    requestSerialize: serialize_messages_MessageEnvelop,
    requestDeserialize: deserialize_messages_MessageEnvelop,
    responseSerialize: serialize_pipes_Xid,
    responseDeserialize: deserialize_pipes_Xid,
  },
  // Recv will request to receive with options. Defaults to {
//     AutoAck = false,
//     Block = true,
//     Count = 1,
//     Timeout = inf
// }
recv: {
    path: '/pipes.Pipe/Recv',
    requestStream: false,
    responseStream: false,
    requestType: pipe_pb.ReceiveOptions,
    responseType: event_pb.Events,
    requestSerialize: serialize_pipes_ReceiveOptions,
    requestDeserialize: deserialize_pipes_ReceiveOptions,
    responseSerialize: serialize_messages_Events,
    responseDeserialize: deserialize_messages_Events,
  },
  streamRecv: {
    path: '/pipes.Pipe/StreamRecv',
    requestStream: false,
    responseStream: true,
    requestType: pipe_pb.ReceiveOptions,
    responseType: event_pb.Event,
    requestSerialize: serialize_pipes_ReceiveOptions,
    requestDeserialize: deserialize_pipes_ReceiveOptions,
    responseSerialize: serialize_messages_Event,
    responseDeserialize: deserialize_messages_Event,
  },
  // Ack acknowledges that a message by id was received and can be discarded from the re-enqueue queue queue
ack: {
    path: '/pipes.Pipe/Ack',
    requestStream: false,
    responseStream: false,
    requestType: pipe_pb.CompleteRequest,
    responseType: pipe_pb.GenericResponse,
    requestSerialize: serialize_pipes_CompleteRequest,
    requestDeserialize: deserialize_pipes_CompleteRequest,
    responseSerialize: serialize_pipes_GenericResponse,
    responseDeserialize: deserialize_pipes_GenericResponse,
  },
  // Complete takes a Xid and step, marking the step as complete (to be enqueued into the next pipe, if needed)
complete: {
    path: '/pipes.Pipe/Complete',
    requestStream: false,
    responseStream: false,
    requestType: pipe_pb.CompleteRequest,
    responseType: pipe_pb.GenericResponse,
    requestSerialize: serialize_pipes_CompleteRequest,
    requestDeserialize: deserialize_pipes_CompleteRequest,
    responseSerialize: serialize_pipes_GenericResponse,
    responseDeserialize: deserialize_pipes_GenericResponse,
  },
  // AppendLog takes a routelog and adds it to the message. If step is not given, assumes current step
appendLog: {
    path: '/pipes.Pipe/AppendLog',
    requestStream: false,
    responseStream: false,
    requestType: pipe_pb.RouteLogRequest,
    responseType: pipe_pb.GenericResponse,
    requestSerialize: serialize_pipes_RouteLogRequest,
    requestDeserialize: deserialize_pipes_RouteLogRequest,
    responseSerialize: serialize_pipes_GenericResponse,
    responseDeserialize: deserialize_pipes_GenericResponse,
  },
  // AddSteps adds steps to the route. If After is not given, assumes after current step
addSteps: {
    path: '/pipes.Pipe/AddSteps',
    requestStream: false,
    responseStream: false,
    requestType: pipe_pb.AddStepsRequest,
    responseType: pipe_pb.GenericResponse,
    requestSerialize: serialize_pipes_AddStepsRequest,
    requestDeserialize: deserialize_pipes_AddStepsRequest,
    responseSerialize: serialize_pipes_GenericResponse,
    responseDeserialize: deserialize_pipes_GenericResponse,
  },
  // Decorate takes a set of decorations and applies them to the message
decorate: {
    path: '/pipes.Pipe/Decorate',
    requestStream: false,
    responseStream: false,
    requestType: pipe_pb.Decorations,
    responseType: pipe_pb.GenericResponses,
    requestSerialize: serialize_pipes_Decorations,
    requestDeserialize: deserialize_pipes_Decorations,
    responseSerialize: serialize_pipes_GenericResponses,
    responseDeserialize: deserialize_pipes_GenericResponses,
  },
};

exports.PipeClient = grpc.makeGenericClientConstructor(PipeService);
var StoreService = exports.StoreService = {
  get: {
    path: '/pipes.Store/Get',
    requestStream: false,
    responseStream: false,
    requestType: pipe_pb.Xid,
    responseType: event_pb.Event,
    requestSerialize: serialize_pipes_Xid,
    requestDeserialize: deserialize_pipes_Xid,
    responseSerialize: serialize_messages_Event,
    responseDeserialize: deserialize_messages_Event,
  },
  del: {
    path: '/pipes.Store/Del',
    requestStream: false,
    responseStream: false,
    requestType: pipe_pb.Xid,
    responseType: pipe_pb.GenericResponse,
    requestSerialize: serialize_pipes_Xid,
    requestDeserialize: deserialize_pipes_Xid,
    responseSerialize: serialize_pipes_GenericResponse,
    responseDeserialize: deserialize_pipes_GenericResponse,
  },
  list: {
    path: '/pipes.Store/List',
    requestStream: false,
    responseStream: false,
    requestType: pipe_pb.ListRequest,
    responseType: event_pb.Events,
    requestSerialize: serialize_pipes_ListRequest,
    requestDeserialize: deserialize_pipes_ListRequest,
    responseSerialize: serialize_messages_Events,
    responseDeserialize: deserialize_messages_Events,
  },
};

exports.StoreClient = grpc.makeGenericClientConstructor(StoreService);
