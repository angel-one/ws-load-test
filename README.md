# Go Example Project

This project has all the necessary setup done for you to get kick-started with your golang project.

## How to Start?

Clone this repository and move into this directory.

```shell
git clone git@github.com:angel-one/go-example-project.git
cd go-example-project
```

Run the doctor to make sure you have all the pre-requisites installed.
```shell
make doctor
```

Create a new project in GitHub under your organization's namespace and then initialize by providing your organization name and project name.
```shell
make init
```

## What to do when you clone your repository again?

You don't have to call `init` again, it's only once. All you have to do is run the following command.
```shell
make install
```

## What to do when you add a new API or change an existing one?

You will have to provide proper comments to that API as per the documentation mentioned [here](https://github.com/swaggo/swag#general-api-info). Once you have done that, you will have to run the following command to update the documentation.
```shell
make swagger
```
That's it, now just run the application and browse `http://localhost:8080/swagger/index.html` to view your changes.

## How to verify sanctity of your code?

Run the code checker.
```shell
make verify
```

## How to run tests?
```shell
make test
```

## How to run the application?

To run the application, you need to provide the following program arguments.
1. **port** - This is the port number where you have to start the application.
2. **env** - This is the application runtime environment.
3. **base-config-path** - This is the base path that stores all the configurations. You can find the configurations [here](./resources). So the path to this folder has to be provided.

Once, the application is running, the swagger can be accessed at `http://localhost:${port}/swagger/index.html`.
# ws-load-test
