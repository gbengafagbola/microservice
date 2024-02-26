const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');
const packageDefinition = protoLoader.loadSync('your_proto_file.proto', { keepCase: true, longs: String, enums: String, defaults: true, oneofs: true });
const hello_proto = grpc.loadPackageDefinition(packageDefinition).hello;

const server = new grpc.Server();

server.addService(hello_proto.Greeter.service, {
  sayHello: (call, callback) => {
    console.log(`Received: ${call.request.name}`);
    const message = `Hello, ${call.request.name}`;
    callback(null, { message });
  },
});

server.bind('0.0.0.0:50052', grpc.ServerCredentials.createInsecure());
console.log('Server running at http://127.0.0.1:50052');
server.start();
