

1. Configuration for database 
2. Create app.go file 
    struct with db and mux objects
    initialize db connection and mux serve in a different package
    new package should be create for all the endpoints
    define models using protocol buffers (proto files)
   create makefile for compiling proto file and inclue in build process



[stage-1]
1. complete couple of basic endpoints to get values from tables
2. create deployment files for the service
3. deploy 3 replicas of the service and should return from which pod the reponse is coming from

[stage-2]
1. Include advanced concepts 
    a. concurrency 
    b. protocol buffers
    c. message queues
2. Another service which handles large number of volumes
3. Create a load testing service
