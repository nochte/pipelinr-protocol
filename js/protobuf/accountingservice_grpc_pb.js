// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var accountingservice_pb = require('./accountingservice_pb.js');
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');

function serialize_accounting_GetReportRequest(arg) {
  if (!(arg instanceof accountingservice_pb.GetReportRequest)) {
    throw new Error('Expected argument of type accounting.GetReportRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_accounting_GetReportRequest(buffer_arg) {
  return accountingservice_pb.GetReportRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_accounting_GetReportResponse(arg) {
  if (!(arg instanceof accountingservice_pb.GetReportResponse)) {
    throw new Error('Expected argument of type accounting.GetReportResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_accounting_GetReportResponse(buffer_arg) {
  return accountingservice_pb.GetReportResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var AccountingService = exports.AccountingService = {
  getReport: {
    path: '/accounting.Accounting/GetReport',
    requestStream: false,
    responseStream: false,
    requestType: accountingservice_pb.GetReportRequest,
    responseType: accountingservice_pb.GetReportResponse,
    requestSerialize: serialize_accounting_GetReportRequest,
    requestDeserialize: deserialize_accounting_GetReportRequest,
    responseSerialize: serialize_accounting_GetReportResponse,
    responseDeserialize: deserialize_accounting_GetReportResponse,
  },
};

exports.AccountingClient = grpc.makeGenericClientConstructor(AccountingService);
