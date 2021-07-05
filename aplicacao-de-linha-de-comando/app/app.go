package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Applicação de linha de comando"
	app.Usage = "Busca de IPs e nomes de servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca ips",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "busca dos servidores",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servers, erro := net.LookupNS(host) //name server

	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
