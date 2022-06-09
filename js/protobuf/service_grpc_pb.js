// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var pipe_pb = require('./pipe_pb.js');
var apikey_pb = require('./apikey_pb.js');
var user_pb = require('./user_pb.js');

function serialize_accounts_ConfirmSignupRequest(arg) {
  if (!(arg instanceof user_pb.ConfirmSignupRequest)) {
    throw new Error('Expected argument of type accounts.ConfirmSignupRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_accounts_ConfirmSignupRequest(buffer_arg) {
  return user_pb.ConfirmSignupRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_accounts_CreateAPIKeyRequest(arg) {
  if (!(arg instanceof apikey_pb.CreateAPIKeyRequest)) {
    throw new Error('Expected argument of type accounts.CreateAPIKeyRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_accounts_CreateAPIKeyRequest(buffer_arg) {
  return apikey_pb.CreateAPIKeyRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_accounts_CreateAPIKeyResponse(arg) {
  if (!(arg instanceof apikey_pb.CreateAPIKeyResponse)) {
    throw new Error('Expected argument of type accounts.CreateAPIKeyResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_accounts_CreateAPIKeyResponse(buffer_arg) {
  return apikey_pb.CreateAPIKeyResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_accounts_JWTValidationRequest(arg) {
  if (!(arg instanceof user_pb.JWTValidationRequest)) {
    throw new Error('Expected argument of type accounts.JWTValidationRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_accounts_JWTValidationRequest(buffer_arg) {
  return user_pb.JWTValidationRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_accounts_ListAPIKeyResponse(arg) {
  if (!(arg instanceof apikey_pb.ListAPIKeyResponse)) {
    throw new Error('Expected argument of type accounts.ListAPIKeyResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_accounts_ListAPIKeyResponse(buffer_arg) {
  return apikey_pb.ListAPIKeyResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_accounts_LoginResponse(arg) {
  if (!(arg instanceof user_pb.LoginResponse)) {
    throw new Error('Expected argument of type accounts.LoginResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_accounts_LoginResponse(buffer_arg) {
  return user_pb.LoginResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_accounts_SignupRequest(arg) {
  if (!(arg instanceof user_pb.SignupRequest)) {
    throw new Error('Expected argument of type accounts.SignupRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_accounts_SignupRequest(buffer_arg) {
  return user_pb.SignupRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_accounts_UsernameLoginRequest(arg) {
  if (!(arg instanceof user_pb.UsernameLoginRequest)) {
    throw new Error('Expected argument of type accounts.UsernameLoginRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_accounts_UsernameLoginRequest(buffer_arg) {
  return user_pb.UsernameLoginRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_accounts_ValidateAPIKeyRequestResponse(arg) {
  if (!(arg instanceof user_pb.ValidateAPIKeyRequestResponse)) {
    throw new Error('Expected argument of type accounts.ValidateAPIKeyRequestResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_accounts_ValidateAPIKeyRequestResponse(buffer_arg) {
  return user_pb.ValidateAPIKeyRequestResponse.deserializeBinary(new Uint8Array(buffer_arg));
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

function serialize_pipes_Null(arg) {
  if (!(arg instanceof pipe_pb.Null)) {
    throw new Error('Expected argument of type pipes.Null');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_pipes_Null(buffer_arg) {
  return pipe_pb.Null.deserializeBinary(new Uint8Array(buffer_arg));
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


var AccountsService = exports.AccountsService = {
  // These are unauthenticated, wide-open
signup: {
    path: '/accounts.Accounts/Signup',
    requestStream: false,
    responseStream: false,
    requestType: user_pb.SignupRequest,
    responseType: pipe_pb.GenericResponse,
    requestSerialize: serialize_accounts_SignupRequest,
    requestDeserialize: deserialize_accounts_SignupRequest,
    responseSerialize: serialize_pipes_GenericResponse,
    responseDeserialize: deserialize_pipes_GenericResponse,
  },
  confirmSignup: {
    path: '/accounts.Accounts/ConfirmSignup',
    requestStream: false,
    responseStream: false,
    requestType: user_pb.ConfirmSignupRequest,
    responseType: user_pb.LoginResponse,
    requestSerialize: serialize_accounts_ConfirmSignupRequest,
    requestDeserialize: deserialize_accounts_ConfirmSignupRequest,
    responseSerialize: serialize_accounts_LoginResponse,
    responseDeserialize: deserialize_accounts_LoginResponse,
  },
  login: {
    path: '/accounts.Accounts/Login',
    requestStream: false,
    responseStream: false,
    requestType: user_pb.UsernameLoginRequest,
    responseType: user_pb.LoginResponse,
    requestSerialize: serialize_accounts_UsernameLoginRequest,
    requestDeserialize: deserialize_accounts_UsernameLoginRequest,
    responseSerialize: serialize_accounts_LoginResponse,
    responseDeserialize: deserialize_accounts_LoginResponse,
  },
  forgotPassword: {
    path: '/accounts.Accounts/ForgotPassword',
    requestStream: false,
    responseStream: false,
    requestType: user_pb.UsernameLoginRequest,
    responseType: pipe_pb.GenericResponse,
    requestSerialize: serialize_accounts_UsernameLoginRequest,
    requestDeserialize: deserialize_accounts_UsernameLoginRequest,
    responseSerialize: serialize_pipes_GenericResponse,
    responseDeserialize: deserialize_pipes_GenericResponse,
  },
  confirmForgotPassword: {
    path: '/accounts.Accounts/ConfirmForgotPassword',
    requestStream: false,
    responseStream: false,
    requestType: user_pb.ConfirmSignupRequest,
    responseType: user_pb.LoginResponse,
    requestSerialize: serialize_accounts_ConfirmSignupRequest,
    requestDeserialize: deserialize_accounts_ConfirmSignupRequest,
    responseSerialize: serialize_accounts_LoginResponse,
    responseDeserialize: deserialize_accounts_LoginResponse,
  },
  validateJWT: {
    path: '/accounts.Accounts/ValidateJWT',
    requestStream: false,
    responseStream: false,
    requestType: user_pb.JWTValidationRequest,
    responseType: user_pb.LoginResponse,
    requestSerialize: serialize_accounts_JWTValidationRequest,
    requestDeserialize: deserialize_accounts_JWTValidationRequest,
    responseSerialize: serialize_accounts_LoginResponse,
    responseDeserialize: deserialize_accounts_LoginResponse,
  },
  createAPIKey: {
    path: '/accounts.Accounts/CreateAPIKey',
    requestStream: false,
    responseStream: false,
    requestType: apikey_pb.CreateAPIKeyRequest,
    responseType: apikey_pb.CreateAPIKeyResponse,
    requestSerialize: serialize_accounts_CreateAPIKeyRequest,
    requestDeserialize: deserialize_accounts_CreateAPIKeyRequest,
    responseSerialize: serialize_accounts_CreateAPIKeyResponse,
    responseDeserialize: deserialize_accounts_CreateAPIKeyResponse,
  },
  // NOTE: ListAPIKeys will not return the actual Key value
listAPIKeys: {
    path: '/accounts.Accounts/ListAPIKeys',
    requestStream: false,
    responseStream: false,
    requestType: pipe_pb.Null,
    responseType: apikey_pb.ListAPIKeyResponse,
    requestSerialize: serialize_pipes_Null,
    requestDeserialize: deserialize_pipes_Null,
    responseSerialize: serialize_accounts_ListAPIKeyResponse,
    responseDeserialize: deserialize_accounts_ListAPIKeyResponse,
  },
  deleteAPIKey: {
    path: '/accounts.Accounts/DeleteAPIKey',
    requestStream: false,
    responseStream: false,
    requestType: pipe_pb.Xid,
    responseType: pipe_pb.GenericResponse,
    requestSerialize: serialize_pipes_Xid,
    requestDeserialize: deserialize_pipes_Xid,
    responseSerialize: serialize_pipes_GenericResponse,
    responseDeserialize: deserialize_pipes_GenericResponse,
  },
  validateAPIKey: {
    path: '/accounts.Accounts/ValidateAPIKey',
    requestStream: false,
    responseStream: false,
    requestType: user_pb.ValidateAPIKeyRequestResponse,
    responseType: user_pb.ValidateAPIKeyRequestResponse,
    requestSerialize: serialize_accounts_ValidateAPIKeyRequestResponse,
    requestDeserialize: deserialize_accounts_ValidateAPIKeyRequestResponse,
    responseSerialize: serialize_accounts_ValidateAPIKeyRequestResponse,
    responseDeserialize: deserialize_accounts_ValidateAPIKeyRequestResponse,
  },
};

exports.AccountsClient = grpc.makeGenericClientConstructor(AccountsService);
