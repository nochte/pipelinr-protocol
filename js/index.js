module.exports = {
  services: {
    ...require('./protobuf/pipe_grpc_pb'),
    ...require('./protobuf/accountingservice_grpc_pb'),
    ...require('./protobuf/service_grpc_pb'),
  },
  messages: {
    ...require('./protobuf/accountingservice_pb'),
    ...require('./protobuf/apikey_pb'),
    ...require('./protobuf/event_pb'),
    ...require('./protobuf/gogo_pb'),
    ...require('./protobuf/healthcheck_pb'),
    ...require('./protobuf/messageenvelop_pb'),
    ...require('./protobuf/pipe_pb'),
    ...require('./protobuf/service_pb'),
    ...require('./protobuf/user_pb'),    
  }
}

