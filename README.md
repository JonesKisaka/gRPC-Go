# gRPC-Go
gRPC Server written in Go.

#issue
In server.go, the &s in "chat.RegisterChatServiceServer(grpcServer, &s)" pulls an error when compiling. 
the error: cannot use &s (value of type *chat.Server) as chat.ChatServiceServer value in argument to chat.RegisterChatServiceServer: *chat.Server does not implement chat.ChatServiceServer (missing method mustEmbedUnimplementedChatServiceServer)

Any help to this will be accepted.
