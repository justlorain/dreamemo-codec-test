package main

import (
	"bytes"
	"context"
	"dreamemo-codec-test/protobuf"
	thrift2 "dreamemo-codec-test/thrift"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
)

func main() {
	thriftRequest()
	protobufRequest()
}

func thriftRequest() {
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/thrift", nil)
	tpt := http.DefaultTransport
	resp, _ := tpt.RoundTrip(req)

	b := &bytes.Buffer{}
	_, _ = io.Copy(b, resp.Body)

	deserializer := thrift.NewTDeserializer()
	out := &thrift2.GetResponse{}
	_ = deserializer.Read(context.Background(), out, b.Bytes())

	fmt.Println(string(out.GetValue()))
}

func protobufRequest() {
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/protobuf", nil)
	tpt := http.DefaultTransport
	resp, _ := tpt.RoundTrip(req)

	b := &bytes.Buffer{}
	_, _ = io.Copy(b, resp.Body)

	out := &protobuf.GetResponse{}
	_ = proto.Unmarshal(b.Bytes(), out)

	fmt.Println(string(out.GetValue()))
}
