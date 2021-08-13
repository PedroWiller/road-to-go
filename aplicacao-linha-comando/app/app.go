package app

import (
	"log"
	"net"

	"github.com/urfave/cli"
)

//Gerar vai retornar a aplicacao que sera retornada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicacao de linha de comando comando"
	app.Usage = "Busca ip e nomes na Internet"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca Ips de endereco na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: buscarIp,
		},
	}

	return app
}

func buscarIp(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupHost(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		log.Println(ip)
	}
}
