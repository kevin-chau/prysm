// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: proto/validator/accounts/v2/keymanager.proto

/*
Package ethereum_validator_accounts_v2 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package ethereum_validator_accounts_v2

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = descriptor.ForMessage

func request_RemoteSigner_ListValidatingPublicKeys_0(ctx context.Context, marshaler runtime.Marshaler, client RemoteSignerClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq empty.Empty
	var metadata runtime.ServerMetadata

	msg, err := client.ListValidatingPublicKeys(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_RemoteSigner_ListValidatingPublicKeys_0(ctx context.Context, marshaler runtime.Marshaler, server RemoteSignerServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq empty.Empty
	var metadata runtime.ServerMetadata

	msg, err := server.ListValidatingPublicKeys(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_RemoteSigner_Sign_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_RemoteSigner_Sign_0(ctx context.Context, marshaler runtime.Marshaler, client RemoteSignerClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq SignRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_RemoteSigner_Sign_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Sign(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_RemoteSigner_Sign_0(ctx context.Context, marshaler runtime.Marshaler, server RemoteSignerServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq SignRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_RemoteSigner_Sign_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.Sign(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterRemoteSignerHandlerServer registers the http handlers for service RemoteSigner to "mux".
// UnaryRPC     :call RemoteSignerServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
func RegisterRemoteSignerHandlerServer(ctx context.Context, mux *runtime.ServeMux, server RemoteSignerServer) error {

	mux.Handle("GET", pattern_RemoteSigner_ListValidatingPublicKeys_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_RemoteSigner_ListValidatingPublicKeys_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_RemoteSigner_ListValidatingPublicKeys_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_RemoteSigner_Sign_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_RemoteSigner_Sign_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_RemoteSigner_Sign_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterRemoteSignerHandlerFromEndpoint is same as RegisterRemoteSignerHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterRemoteSignerHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterRemoteSignerHandler(ctx, mux, conn)
}

// RegisterRemoteSignerHandler registers the http handlers for service RemoteSigner to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterRemoteSignerHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterRemoteSignerHandlerClient(ctx, mux, NewRemoteSignerClient(conn))
}

// RegisterRemoteSignerHandlerClient registers the http handlers for service RemoteSigner
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "RemoteSignerClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "RemoteSignerClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "RemoteSignerClient" to call the correct interceptors.
func RegisterRemoteSignerHandlerClient(ctx context.Context, mux *runtime.ServeMux, client RemoteSignerClient) error {

	mux.Handle("GET", pattern_RemoteSigner_ListValidatingPublicKeys_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_RemoteSigner_ListValidatingPublicKeys_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_RemoteSigner_ListValidatingPublicKeys_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_RemoteSigner_Sign_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_RemoteSigner_Sign_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_RemoteSigner_Sign_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_RemoteSigner_ListValidatingPublicKeys_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 0}, []string{"accounts", "v2", "remote"}, "", runtime.AssumeColonVerbOpt(true)))

	pattern_RemoteSigner_Sign_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"accounts", "v2", "remote", "sign"}, "", runtime.AssumeColonVerbOpt(true)))
)

var (
	forward_RemoteSigner_ListValidatingPublicKeys_0 = runtime.ForwardResponseMessage

	forward_RemoteSigner_Sign_0 = runtime.ForwardResponseMessage
)
