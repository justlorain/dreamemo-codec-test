package main

import (
	"context"
	"dreamemo-codec-test/protobuf"
	thrift2 "dreamemo-codec-test/thrift"
	"github.com/apache/thrift/lib/go/thrift"
	"google.golang.org/protobuf/proto"
	"net/http"
)

func main() {
	http.HandleFunc("/thrift", func(w http.ResponseWriter, req *http.Request) {
		serializer := thrift.NewTSerializer()
		body, _ := serializer.Write(context.Background(), &thrift2.GetResponse{Value: []byte("hello")})
		w.Header().Set("Content-Type", "application/octet-stream")
		_, _ = w.Write(body)
	})
	http.HandleFunc("/protobuf", func(w http.ResponseWriter, req *http.Request) {
		body, _ := proto.Marshal(&protobuf.GetResponse{Value: []byte("hello")})
		w.Header().Set("Content-Type", "application/octet-stream")
		_, _ = w.Write(body)
	})
	_ = http.ListenAndServe(":8080", nil)
}
