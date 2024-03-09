// GENERATED CODE -- DO NOT EDIT!

'use strict';
var User_pb = require('./User_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');

function serialize_User_GetUserRequest(arg) {
  if (!(arg instanceof User_pb.GetUserRequest)) {
    throw new Error('Expected argument of type User.GetUserRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_User_GetUserRequest(buffer_arg) {
  return User_pb.GetUserRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_User_User(arg) {
  if (!(arg instanceof User_pb.User)) {
    throw new Error('Expected argument of type User.User');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_User_User(buffer_arg) {
  return User_pb.User.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_User_UserResponse(arg) {
  if (!(arg instanceof User_pb.UserResponse)) {
    throw new Error('Expected argument of type User.UserResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_User_UserResponse(buffer_arg) {
  return User_pb.UserResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_User_UsersResponse(arg) {
  if (!(arg instanceof User_pb.UsersResponse)) {
    throw new Error('Expected argument of type User.UsersResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_User_UsersResponse(buffer_arg) {
  return User_pb.UsersResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_google_protobuf_Empty(arg) {
  if (!(arg instanceof google_protobuf_empty_pb.Empty)) {
    throw new Error('Expected argument of type google.protobuf.Empty');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_google_protobuf_Empty(buffer_arg) {
  return google_protobuf_empty_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}


var UserServiceService = exports['User.UserService'] = {
  getUsers: {
    path: '/User.UserService/GetUsers',
    requestStream: false,
    responseStream: false,
    requestType: google_protobuf_empty_pb.Empty,
    responseType: User_pb.UsersResponse,
    requestSerialize: serialize_google_protobuf_Empty,
    requestDeserialize: deserialize_google_protobuf_Empty,
    responseSerialize: serialize_User_UsersResponse,
    responseDeserialize: deserialize_User_UsersResponse,
  },
  getUser: {
    path: '/User.UserService/GetUser',
    requestStream: false,
    responseStream: false,
    requestType: User_pb.GetUserRequest,
    responseType: User_pb.UserResponse,
    requestSerialize: serialize_User_GetUserRequest,
    requestDeserialize: deserialize_User_GetUserRequest,
    responseSerialize: serialize_User_UserResponse,
    responseDeserialize: deserialize_User_UserResponse,
  },
  createUser: {
    path: '/User.UserService/CreateUser',
    requestStream: false,
    responseStream: false,
    requestType: User_pb.User,
    responseType: User_pb.UserResponse,
    requestSerialize: serialize_User_User,
    requestDeserialize: deserialize_User_User,
    responseSerialize: serialize_User_UserResponse,
    responseDeserialize: deserialize_User_UserResponse,
  },
};

