## How to deploy on local

Clone this repository on your local storage.

Then, open the Docker application.

After the Docker is already run, open the terminal and use the command:

```shell
docker-compose up -d
```

## To run a unit test

Change the directory to the `app` folder

```shell
cd app
```

Run the test file using the command:

```shell
go test
```

After finishing the test

```shell
cd ..
docker-compose down
```
