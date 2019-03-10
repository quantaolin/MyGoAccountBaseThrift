package demoserver

import (
	"demo"
    "strings"
    "github.com/apache/thrift/lib/go/thrift"
    "fmt"
    "log"
    "context"
)

type FormatDataImpl struct {}

func (fdi *FormatDataImpl) DoFormat(ctx context.Context, data *demo.Data) (r *demo.Data, err error){
    var rData demo.Data
    rData.Text = strings.ToUpper(data.Text)

    return &rData, nil
}

const (
    HOST = "localhost"
    PORT = "8080"
)

func main() {

    handler := &FormatDataImpl{}
    processor := demo.NewFormatDataProcessor(handler)
    serverTransport, err := thrift.NewTServerSocket(HOST + ":" + PORT)
    if err != nil {
        log.Fatalln("Error:", err)
    }
    transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
    protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

    server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
    fmt.Println("Running at:", HOST + ":" + PORT)
    server.Serve()
}