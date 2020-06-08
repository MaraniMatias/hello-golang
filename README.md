# hello-golang

Jugando con Go

## Weather file if change it run "go run ." command

Nodemon pero para go

```sh
npm install -g go-mon
```

Using

```sh
 gomon .
```

## Git houck

Crear un archivo ./.git/hooks/pre-commit

```sh
#!/usr/bin/env bash

git diff HEAD --name-only --diff-filter=ACM | grep '.go$' | go test
```

Para saltar la verification use `git commit -a --no-verify`

> Nota:
> Para usar los archivos use el argumento 'xarg'
> git diff HEAD --name-only --diff-filter=ACM | grep '.go\$' | xargs echo
