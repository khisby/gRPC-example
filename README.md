# Description  
GRPC Example golang and JS

# How to run  
Need to Setup:  
1. install protoc, protoc-gen-go, grpc_tools_node_protoc_plugin  
2. run `cd js_grpc && npm install`  
  
Command:  
run in root project  
`make gen-proto` : generate proto file to `go_grpc/grpc` and `js_grpc/grpc`  
`make clean` : remove folder grpc generated  
`make run-go`: run grpc server go  
`make run-js`: run grpc client js  
  