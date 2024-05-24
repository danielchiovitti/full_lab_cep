## Serviço para exibir o clima atual através do CEP

### Iniciar serviço localmente
````bash
docker compose up
````

### Rodar testes automatizados
```bash
go test -v ./...
```

### CURLS para chamar o endpoint localmente e para o Google Cloud Run

#### CEP válido
````
curl --location 'http://localhost:3500/cep/05330999'
curl --location 'https://full-lab-cep-f6b7bmiviq-uc.a.run.app/cep/05330-011'
````

#### CEP não encontrado
````
curl --location 'http://localhost:3500/cep/05330-999'
curl --location 'https://full-lab-cep-f6b7bmiviq-uc.a.run.app/cep/05330-999'
````

#### CEP inválido
````
curl --location 'http://localhost:3500/cep/05330999'
curl --location 'https://full-lab-cep-f6b7bmiviq-uc.a.run.app/cep/05330999'
````