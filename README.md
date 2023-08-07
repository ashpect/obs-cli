# Cli for interacting with the Open Broadcaster Software (OBS)'s' API
This is a simple cli tool for interacting with the OBS API. It is written in Go and leverages the use of cobracli.

## Initialize project
```
cp config.sample.toml config.toml
go mod tidy

```
Update the config.toml file with your creds for OBS 
and run the following commands to build and run the project.
```
go build .
./obs-cli
```
In case you just require a go server of your own to interact with obs-api, and write your own cli, you can find a very basicccc server here : https://github.com/ashpect/sig-obs 