package main

import "github.com/cioti/monorepo/cms.api/internal/transport"

func main() {
	builder := transport.NewHttpTransportBuilder()
	builder.Build().Run(":8080")
}
