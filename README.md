# 🔍 NetTools CLI

> Ferramentas de linha de comando escritas em Go para análise de rede, resolução de domínios, testes de latência, varredura de portas e consultas WHOIS.

---

## 📦 Sobre o projeto

**NetTools CLI** é um utilitário de linha de comando construído em Go, que oferece uma série de comandos úteis para diagnosticar e investigar informações de rede, como IPs, servidores DNS, registros MX, consultas WHOIS, varredura de portas TCP e mais.

Este projeto também faz parte de um estudo prático sobre a biblioteca `net` da linguagem Go, com foco no aprofundamento dos recursos de rede e no treinamento dos fundamentos de testes automatizados, incluindo uso de mocks e injeção de dependência.

O projeto é modular, testável, desacoplado por meio de injeção de dependência, com suporte a testes automatizados e pronto para produção.

---

## 🚀 Funcionalidades

- 🔍 `ip` – Busca os IPs de um domínio
- 🌐 `servidores` – Lista os servidores de nomes (NS)
- ✉️ `mx` – Lista os servidores de e-mail (MX)
- 🔗 `cname` – Exibe o CNAME de um domínio
- 🔁 `reverso` – Realiza lookup reverso de IP para hostname
- 🧠 `whois` – Executa consulta WHOIS do domínio
- 📶 `ping` – Mede a latência de resolução do host
- 🔓 `tcp-scan` – Escaneia portas abertas em um host

---

## 📥 Instalação

### Com Go instalado:

git clone https://github.com/ZecaSouza/nettools-cli.git

cd nettools-cli

go run main.go --help

git clone https://github.com/ZecaSouza/nettools-cli.git

cd nettools-cli

go run main.go --help

---

## Compilar binário:

go build -o nettools main.go

./nettools ip --host=google.com

---

## 🧪 Testes

Este projeto é 100% testável com mocks e segue boas práticas de injeção de dependência.

### 🚀 Executar todos os testes:

go test ./...

---

## 💻 Exemplos de uso

go run main.go ip --host=google.com

go run main.go servidores --host=example.com

go run main.go mx --host=example.com

go run main.go whois --host=example.com

go run main.go tcp-scan --host=example.com --inicio=20 --fim=25

---

🧱 Estrutura de diretórios
/

├── app/

│   ├── app.go              # Configuração da CLI

│   ├── resolver.go         # Interface Resolver

│   ├── netresolver.go      # Implementação real

│   ├── mockresolver.go     # Mock para testes

│   ├── app_test.go         # Testes automatizados

├── main.go                 # Entrada principal

├── go.mod                  # Módulo Go

└── README.md

---

🧰 Tecnologias utilizadas:

° Go (Golang)

° urfave/cli.v2

° DNS padrão (net)

° WHOIS via conexão TCP simples

---

🤝 Contribuições

1 - Contribuições são bem-vindas!

2 - Faça um fork do projeto

3 - Crie sua feature branch (git checkout -b minha-feature)

4 - Commit suas alterações (git commit -m 'feat: nova feature')

5 - Push na branch (git push origin minha-feature)

7 - Abra um Pull Request 🚀

---

📌 Autor

Zeca Souza

Desenvolvedor, entusiasta de software livre e ferramentas de rede.

https://github.com/ZecaSouza
