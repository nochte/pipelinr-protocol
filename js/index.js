const fs = require('fs');

const protoLoader = require('@grpc/proto-loader');
const defs = protoLoader.loadSync(
  
)

// const files = fs.readdirSync(__dirname + '/protobuf');

// files.forEach((f) => {
//   require(__dirname + '/protobuf/' + f);
// });