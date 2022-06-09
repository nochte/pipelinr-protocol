// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var pipe_pb = require('./pipe_pb.js');

function serialize_pipes_GenericResponse(arg) {
  if (!(arg instanceof pipe_pb.GenericResponse)) {
    throw new Error('Expected argument of type pipes.GenericResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pipes_GenericResponse(buffer_arg) {
  return pipe_pb.GenericResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_pipes_Null(arg) {
  if (!(arg instanceof pipe_pb.Null)) {
    throw new Error('Expected argument of type pipes.Null');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pipes_Null(buffer_arg) {
  return pipe_pb.Null.deserializeBinary(new Uint8Array(buffer_arg));
}


var HealthcheckService = exports.HealthcheckService = {
  isReady: {
    path: '/healthcheck.Healthcheck/IsReady',
    requestStream: false,
    responseStream: false,
    requestType: pipe_pb.Null,
    responseType: pipe_pb.GenericResponse,
    requestSerialize: serialize_pipes_Null,
    requestDeserialize: deserialize_pipes_Null,
    responseSerialize: serialize_pipes_GenericResponse,
    responseDeserialize: deserialize_pipes_GenericResponse,
  },
  isHealthy: {
    path: '/healthcheck.Healthcheck/IsHealthy',
    requestStream: false,
    responseStream: false,
    requestType: pipe_pb.Null,
    responseType: pipe_pb.GenericResponse,
    requestSerialize: serialize_pipes_Null,
    requestDeserialize: deserialize_pipes_Null,
    responseSerialize: serialize_pipes_GenericResponse,
    responseDeserialize: deserialize_pipes_GenericResponse,
  },
};

exports.HealthcheckClient = grpc.makeGenericClientConstructor(HealthcheckService);
