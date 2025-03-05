[![Build Status](http://ec2-54-171-246-10.eu-west-1.compute.amazonaws.com/buildStatus/icon?job=git-pipeline)](http://ec2-54-171-246-10.eu-west-1.compute.amazonaws.com/job/git-pipeline/) 
## rest-client
This is my first project written in Go. It was fun to learn a new language and implement a client library.

This rest client library, implemented in Go, accesses a dummy account API provided as Docker container in the file docker-compose.yaml.

Create, Fetch, and Delete operations on the accounts resource have been implemented. List operation is also implemented in addition.

Using docker command `docker-compose up` would run all test cases.

## Future enhancements
* Implement Exponential back-off retry algorithm.
* Publish and distribute the module.
* Integrate with jenkins pipeline
