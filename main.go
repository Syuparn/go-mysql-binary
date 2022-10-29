package main

import (
	"flag"
	"fmt"

	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/server"
	"github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/information_schema"
)

var (
	host = flag.String("host", "localhost", "host to listen")
	port = flag.Int("port", 3306, "port to listen")
)

func main() {
	flag.Parse()

	engine := sqle.NewDefault(
		sql.NewDatabaseProvider(
			information_schema.NewInformationSchemaDatabase(),
		))
	engine.Analyzer.Catalog.MySQLDb.AddRootAccount()

	config := server.Config{
		Protocol: "tcp",
		Address:  fmt.Sprintf("%s:%d", *host, *port),
	}

	s, err := server.NewDefaultServer(config, engine)
	if err != nil {
		panic(err)
	}

	s.Start()
}
