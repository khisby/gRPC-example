import { UserServiceClient } from "./grpc/User/User_grpc_pb";
// import { UsersResponse } from "./grpc/User/User_pb";
import { credentials } from "@grpc/grpc-js";


const client = new UserServiceClient("localhost:9090", credentials.createInsecure());
console.log(client);


// client.getUsers({}, (err, response: UsersResponse) => {
//     if (err) {
//         console.error(err);
//     } else {
//         console.log(response.toObject());
//     }
// });