# 🚀 Desafio Técnico – Estágio DevOps | Projeto Korp

Este repositório contém a solução para o desafio técnico proposto para a vaga de estágio DevOps, envolvendo desenvolvimento com Golang, conteinerização com Docker, orquestração com Docker Compose e automação com Ansible.

## 📋 Descrição do Desafio

> O desafio consiste em:
- Desenvolver um serviço HTTP em Go que retorna um JSON com nome e horário atual em UTC.
- Empacotar esse serviço em um container Docker.
- Criar um ambiente com dois containers: o serviço em Go e um servidor NGINX configurado como proxy reverso.
- Automatizar toda a configuração com Ansible.
- A resposta ao acessar `http://localhost:80` deve ser o JSON gerado pelo serviço em Go.

---

## 📁 Estrutura do Projeto

```bash
.
├── playbooks/
│   ├── update_packages.yml
│   └── main.yml
├── docker/
│   └── docker-compose.yml
├── proxy/
│   └── http-server-projeto-korp.conf
├── server/
│   ├── cmd
│   │    └── main.go
│   ├── go.mod
│   ├── go.sum
│   └── Dockerfile
└── README.md
```

## 🔧 Tecnologias Utilizadas

- [Golang](https://go.dev/) 
- [Docker](https://docs.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Nginx](https://nginx.org/en/docs/)
- [Ansible](https://docs.ansible.com/ansible/latest/index.html)
- [Linux Ubuntu](https://ubuntu.com/)

## ⚙️ Como Executar o Projeto

Siga os passos abaixo para rodar o projeto localmente.

### Pré-requisitos

Como pré-requisitos desse projeto é preciso um servidor linux ubuntu com git e o ansible instalados.

- [ ] Git
- [ ] Ansible

```bash
sudo apt-get install ansible-core -y
```
```bash
sudo apt-get install git -y
```

### 📦 Clonando o repositório

```bash
git clone https://github.com/wendelsilva/viasoft-korp-devops-test
```

### 🤖 Automação com Ansible

```bash
ansible-playbook viasoft-korp-devops-test/playbooks/main.yml -K
```
> O comando acima irá solicitar a senha de usuário para executar a automação com privilégios de super usuário


A automação realiza a instalação e configuração do Docker e do Go. Em seguida, constrói a imagem do serviço em Go e inicializa os containers (aplicação e NGINX) com Docker Compose. Após o ambiente estar em funcionamento, é executado automaticamente um teste para verificar se o serviço está respondendo corretamente

## 📄 Licença

Este projeto está licenciado sob a MIT License.