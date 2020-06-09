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

## Go unit

https://github.com/hexdigest/gounit-vim

```bash
go get -u github.com/hexdigest/gounit/cmd/gounit
```

## Library

- [ENV](https://github.com/ilyakaznacheev/cleanenv)
- [MongoDB](https://github.com/kamva/mgm)
- [Gorgonia - Deep learning in Go](https://gorgonia.org/)

## Articles

- [Deep learning in Go](https://medium.com/@hackintoshrao/deep-learning-in-go-f13e586f7d8a)
- [A Beginner’s Journey into Machine Learning: CNN with Gorgonia](https://medium.com/@msdewittdev/a-beginners-journey-into-machine-learning-cnn-with-gorgonia-b1d78e626aea)
