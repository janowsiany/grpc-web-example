import * as grpc from "@improbable-eng/grpc-web";

import {Greeter} from "./protobuf/service_pb_service";
import {HelloRequest} from "./protobuf/service_pb";

const helloRequest = new HelloRequest();
helloRequest.setName("Bob");

grpc.grpc.unary(Greeter.SayHello, {
  request: helloRequest,
  host: "http://localhost:8080",
  onEnd: console.log
});
