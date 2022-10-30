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
	user = flag.String("user", "", "create user other than root")
	pass = flag.String("pass", "", "create password of created user")
)

func main() {
	flag.Parse()

	engine := sqle.NewDefault(
		sql.NewDatabaseProvider(
			information_schema.NewInformationSchemaDatabase(),
		))
	engine.Analyzer.Catalog.MySQLDb.AddRootAccount()

	if *user != "" {
		engine.Analyzer.Catalog.MySQLDb.AddSuperUser(*user, *pass)
	}

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
