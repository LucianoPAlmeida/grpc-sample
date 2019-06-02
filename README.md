# grpc-sample
A sample grpc client server sample used as demonstration in the grpc talk

* We have here 
  * A go server with a simple service implementation.
  * A iOS App client that will call this server.

## To run the server
* Install Go -> [Here](https://golang.org/doc/install)
* Clone this sample. 
* Go to grpc-sample/go-server
* Take a look on `utils.sh` and run it.
* Then cd to 'src/driver-server' and run `go run main.go`
* Well, if all run well we should have a go gRPC server up and running \o/

## To run the client 
* Go to `grpc-client` folder and run `pod install` 
* Open the workspace in Xcode and then the app.
  
## Licence 

This sample is released under the [MIT License](https://opensource.org/licenses/MIT).
