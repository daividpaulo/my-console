package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func NewClient() *cli.App {

	client := cli.NewApp()
	client.Name = "My console app in GO Lang"
	client.Usage = "Cli utilirios de Rede"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "teste.com.br",
		},
	}

	client.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Obtem IPs de um host na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "ns",
			Usage:  "Obtem nomes de servidores para um host na internet",
			Flags:  flags,
			Action: buscarNs,
		},
	}

	return client
}

func buscarNs(c *cli.Context) {

	host := c.String("host")
	ns, erro := net.LookupNS(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range ns {
		fmt.Println(server.Host)
	}
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
