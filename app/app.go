// app/app.go
package app

import (
	"fmt"
	"github.com/urfave/cli"
)

func Gerar(res Resolver) *cli.App {
	app := cli.NewApp()
	app.Name = "NetTools CLI"
	app.Usage = "Ferramentas para análise de rede"
	app.Version = "2.1.0"

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de um domínio",
			Flags:  flags,
			Action: func(c *cli.Context) error {
				return buscarIps(res, c)
			},
		},
		{
			Name:   "servidores",
			Usage:  "Busca os servidores de nomes",
			Flags:  flags,
			Action: func(c *cli.Context) error {
				return buscarServidores(res, c)
			},
		},
		{
			Name:   "mx",
			Usage:  "Busca servidores de e-mail (MX)",
			Flags:  flags,
			Action: func(c *cli.Context) error {
				return buscarMX(res, c)
			},
		},
		{
			Name:   "cname",
			Usage:  "Busca o CNAME de um domínio",
			Flags:  flags,
			Action: func(c *cli.Context) error {
				return buscarCNAME(res, c)
			},
		},
		{
			Name:   "whois",
			Usage:  "Executa consulta WHOIS",
			Flags:  flags,
			Action: func(c *cli.Context) error {
				return whoisConsulta(res, c)
			},
		},
		{
			Name:   "reverso",
			Usage:  "Busca reversa de IP",
			Flags:  flags,
			Action: func(c *cli.Context) error {
				return buscarReverso(res, c)
			},
		},
		{
			Name:   "ping",
			Usage:  "Verifica latência de um host",
			Flags:  flags,
			Action: func(c *cli.Context) error {
				return executarPing(res, c)
			},
		},
		{
			Name:   "tcp-scan",
			Usage:  "Escaneia portas TCP",
			Flags: append(flags,
				&cli.IntFlag{
					Name:  "inicio",
					Value: 1,
				},
				&cli.IntFlag{
					Name:  "fim",
					Value: 1024,
				}),
			Action: func(c *cli.Context) error {
				return escanearPortasTCP(res, c)
			},
		},
	}

	return app
}


func buscarIps(res Resolver, c *cli.Context) error {
	host := c.String("host")
	ips, err := res.LookupIP(host)
	if err != nil {
		return fmt.Errorf("erro ao buscar IPs: %v", err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
	return nil
}


func escanearPortasTCP(res Resolver, c *cli.Context) error {
	host := c.String("host")
	inicio := c.Int("inicio")
	fim := c.Int("fim")

	for port := inicio; port <= fim; port++ {
		open, err := res.TCPScan(host, port)
		if err != nil {
			fmt.Printf("Erro na porta %d: %v\n", port, err)
			continue
		}
		if open {
			fmt.Println("Porta aberta:", port)
		}
	}
	return nil
}

func buscarServidores(res Resolver, c *cli.Context) error {
	host := c.String("host")
	servidores, err := res.LookupNS(host)
	if err != nil {
		return fmt.Errorf("erro ao buscar servidores de nomes: %v", err)
	}
	for _, s := range servidores {
		fmt.Println(s)
	}
	return nil
}

func buscarMX(res Resolver, c *cli.Context) error {
	host := c.String("host")
	mxs, err := res.LookupMX(host)
	if err != nil {
		return fmt.Errorf("erro ao buscar registros MX: %v", err)
	}
	for _, mx := range mxs {
		fmt.Printf("%s (prioridade %d)\n", mx.Host, mx.Pref)
	}
	return nil
}

func buscarCNAME(res Resolver, c *cli.Context) error {
	host := c.String("host")
	cname, err := res.LookupCNAME(host)
	if err != nil {
		return fmt.Errorf("erro ao buscar CNAME: %v", err)
	}
	fmt.Println(cname)
	return nil
}

func buscarReverso(res Resolver, c *cli.Context) error {
	host := c.String("host")
	ips, err := res.LookupIP(host)
	if err != nil {
		return fmt.Errorf("erro ao buscar IPs para reverso: %v", err)
	}
	for _, ip := range ips {
		ptrs, err := res.LookupAddr(ip)
		if err != nil {
			fmt.Printf("Erro ao buscar reverso para %s: %v\n", ip, err)
			continue
		}
		for _, ptr := range ptrs {
			fmt.Println(ptr)
		}
	}
	return nil
}

func whoisConsulta(res Resolver, c *cli.Context) error {
	host := c.String("host")
	output, err := res.Whois(host)
	if err != nil {
		return fmt.Errorf("erro ao consultar WHOIS: %v", err)
	}
	fmt.Println(output)
	return nil
}

func executarPing(res Resolver, c *cli.Context) error {
	host := c.String("host")
	latency, err := res.Ping(host)
	if err != nil {
		return fmt.Errorf("erro ao pingar o host: %v", err)
	}
	fmt.Printf("Latência: %dms\n", latency)
	return nil
}