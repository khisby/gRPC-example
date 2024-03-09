// package: User
// file: User.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as User_pb from "./User_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";

interface IUserServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    getUsers: IUserServiceService_IGetUsers;
    getUser: IUserServiceService_IGetUser;
    createUser: IUserServiceService_ICreateUser;
}

interface IUserServiceService_IGetUsers extends grpc.MethodDefinition<google_protobuf_empty_pb.Empty, User_pb.UsersResponse> {
    path: "/User.UserService/GetUsers";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    requestDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
    responseSerialize: grpc.serialize<User_pb.UsersResponse>;
    responseDeserialize: grpc.deserialize<User_pb.UsersResponse>;
}
interface IUserServiceService_IGetUser extends grpc.MethodDefinition<User_pb.GetUserRequest, User_pb.UserResponse> {
    path: "/User.UserService/GetUser";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<User_pb.GetUserRequest>;
    requestDeserialize: grpc.deserialize<User_pb.GetUserRequest>;
    responseSerialize: grpc.serialize<User_pb.UserResponse>;
    responseDeserialize: grpc.deserialize<User_pb.UserResponse>;
}
interface IUserServiceService_ICreateUser extends grpc.MethodDefinition<User_pb.User, User_pb.UserResponse> {
    path: "/User.UserService/CreateUser";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<User_pb.User>;
    requestDeserialize: grpc.deserialize<User_pb.User>;
    responseSerialize: grpc.serialize<User_pb.UserResponse>;
    responseDeserialize: grpc.deserialize<User_pb.UserResponse>;
}

export const UserServiceService: IUserServiceService;

export interface IUserServiceServer extends grpc.UntypedServiceImplementation {
    getUsers: grpc.handleUnaryCall<google_protobuf_empty_pb.Empty, User_pb.UsersResponse>;
    getUser: grpc.handleUnaryCall<User_pb.GetUserRequest, User_pb.UserResponse>;
    createUser: grpc.handleUnaryCall<User_pb.User, User_pb.UserResponse>;
}

export interface IUserServiceClient {
    getUsers(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: User_pb.UsersResponse) => void): grpc.ClientUnaryCall;
    getUsers(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: User_pb.UsersResponse) => void): grpc.ClientUnaryCall;
    getUsers(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: User_pb.UsersResponse) => void): grpc.ClientUnaryCall;
    getUser(request: User_pb.GetUserRequest, callback: (error: grpc.ServiceError | null, response: User_pb.UserResponse) => void): grpc.ClientUnaryCall;
    getUser(request: User_pb.GetUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: User_pb.UserResponse) => void): grpc.ClientUnaryCall;
    getUser(request: User_pb.GetUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: User_pb.UserResponse) => void): grpc.ClientUnaryCall;
    createUser(request: User_pb.User, callback: (error: grpc.ServiceError | null, response: User_pb.UserResponse) => void): grpc.ClientUnaryCall;
    createUser(request: User_pb.User, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: User_pb.UserResponse) => void): grpc.ClientUnaryCall;
    createUser(request: User_pb.User, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: User_pb.UserResponse) => void): grpc.ClientUnaryCall;
}

export class UserServiceClient extends grpc.Client implements IUserServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public getUsers(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: User_pb.UsersResponse) => void): grpc.ClientUnaryCall;
    public getUsers(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: User_pb.UsersResponse) => void): grpc.ClientUnaryCall;
    public getUsers(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: User_pb.UsersResponse) => void): grpc.ClientUnaryCall;
    public getUser(request: User_pb.GetUserRequest, callback: (error: grpc.ServiceError | null, response: User_pb.UserResponse) => void): grpc.ClientUnaryCall;
    public getUser(request: User_pb.GetUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: User_pb.UserResponse) => void): grpc.ClientUnaryCall;
    public getUser(request: User_pb.GetUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: User_pb.UserResponse) => void): grpc.ClientUnaryCall;
    public createUser(request: User_pb.User, callback: (error: grpc.ServiceError | null, response: User_pb.UserResponse) => void): grpc.ClientUnaryCall;
    public createUser(request: User_pb.User, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: User_pb.UserResponse) => void): grpc.ClientUnaryCall;
    public createUser(request: User_pb.User, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: User_pb.UserResponse) => void): grpc.ClientUnaryCall;
}
