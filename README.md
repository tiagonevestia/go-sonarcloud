# SonarQube

É uma ferramenta para garantir a qualidade do código-fonte em desenvolvimento.

Ele executa algumas análises durante a compilação do aplicativo, como detecção:

- Snippets de código-fonte que podem gerar erros;
- Repetição de linha de comando para evitar que instruções desnecessárias sejam repetidas;
- Segurança;

![Clarity Hugo Theme](./images/dev-cycle.png)

## Instalando uma instância local do SonarQube

Vamos utilizar uma imagem docker para esse processo.

```bash
$ docker run -d --name sonarqube -e SONAR_ES_BOOTSTRAP_CHECKS_DISABLE=true -p 9000:9000 sonarqube:latest
```

Assim que sua instância docker estiver funcionando, faça login em [http://localhost:9000](http://localhost:9000) usando as credenciais de administrador do sistema:

- login: _admin_
- senha: _admin_

## Visão geral

### Rules

SonarQube executa regras no código-fonte para causar problemas. Existem quatro tipos de regras:

- Code Smell
- Bug
- Vulnerability
- Security Hotspot

### Quality Gates

Quality Gates implementa uma política de qualidade para sua organização respondendo à seguinte pergunta: Meu projeto está pronto para ser lançado?

Para responder a esta pergunta, você definiu um conjunto de condições nas quais o projeto é avaliado. Por exemplo:

- Sem novos problemas de bloqueador
- A cobertura do código no novo código é superior a 80%

### SonarScanner

SonarScanner é um método universal para avaliar o código do seu software. Usado em combinação com SonarQube pode produzir resultados excelentes.

- Baixe o arquivo .zip no seguinte [URL](https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-4.6.0.2311-macosx.zip);
- Salve e descompacte o arquivo em um diretório facilmente acessível de sua escolha.
- Como prática recomendada, renomeie o diretório para sonar-scanner.
- Configure o arquivo sonar-scanner.properties. Remova o hashtag inicial da linha sonar.host.url.
- O último passo é definir as variáveis de ambiente para que o comando Sonar Scanner tenha efeito no seu terminal, pois ali você executará o comando responsável por iniciar a verificação do projeto.

Agora vamos executar os seguintes comandos no terminal:

```bash
sonar-scanner \
  -Dsonar.projectKey=NOME_DO_PROJETO \
  -Dsonar.sources=. \
  -Dsonar.host.url=http://localhost:9000 \
  -Dsonar.login=TOKEN
```

### SonarCloud

SonarCloud é um serviço em nuvem baseado em SonarQube fornecido pela SonarSource. SonarQube é uma plataforma de código aberto amplamente usada que pode verificar continuamente a qualidade do código-fonte e detectar erros, vulnerabilidades e "code smells" em mais de 20 linguagens de programação.
SonarCloud lê as métricas do código-fonte do aplicativo enviado ao serviço para montar um painel para expor as métricas do código e diz a você através do mínimo aceitável (Quality Gates) que pode haver erros, "code smells" e outras variáveis.

## Fontes

- [Curso FullCycle](https://portal.code.education/)
- [SonarQube Documentation](https://docs.sonarqube.org/)
