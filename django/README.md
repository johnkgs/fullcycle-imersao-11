# Projeto DJANGO - Cartola FullCycle.

## Requerimentos

-   [Python](https://www.python.org/) >= 3.10.2
-   [pipenv](https://pipenv.pypa.io/en/latest/) >= 2022.11.30

## Como executar o projeto

Instalar dependências:

```bash
$ pipenv install --dev
```

Popular banco de dados:

```bash
$ pipenv run migrate
$ pipenv run python manage.py loaddata initial_data
```

Executar servidor:

```bash
$ pipenv run server
```

Executar formatador de código:

```bash
$ pipenv run prettify
```
