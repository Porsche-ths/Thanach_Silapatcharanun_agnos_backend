## How to deploy on local

Clone this repository on your local storage.

Open the terminal and use the command:

```shell
docker-compose up -d
```

## To test a test file

change diretory to the 'app' folder

```shell
cd app
```

and then run the test file using the command:

```shell
go test
```

After finishing the test

```shell
docker-compose down
```
