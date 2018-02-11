# Tafalk User Store
This is a user store keeping the application users acting as a [gRPC](https://grpc.io/) server.

![alt text](diagram.png)

### Before all
* Have Go
* Have gRPC
* Have Docker

installed.
### gRPC settings
Everything starts with designing the [protobuf](https://github.com/google/protobuf) file 
(`./proto/user.proto` for this project) and once the messages and services are designed
(or changed later) the following command should be run:
```sh
cd $GOPATH/src/github.com/vahdet/tafalk-user-store
protoc -I proto proto/user.proto --go_out=plugins=grpc:proto
``` 

This command will generate (or recreate if it exists) a `.pb.go` file just beside itself. 
It is `./proto/user.pb.go` here and it allows implementing Go services, repositories, in short any Go code
behind the services defined in the `.proto` file.

For this case, the server implementation is performed in `./grpcserver/server.go` file.

### Datastore
#### Redis in a container
A [Redis](https://redis.io/) instance should be up at the port defined in the config file `./resources/app.yaml`. Unless the port is changed, the default one is `6379` on `localhost`

A sample [Docker]() command to install an official Redis image (with data persistence to disk) is as follows:

```sh
docker container run -d --name redisTest \
-v /opt/docker/data/redis:/data \
-p 6379:6379
redis:latest
``` 

> * See **all running** containers
>
> ```sh
> docker ps
>  ```
> * See **all** containers _(running or not)_
>
> ```sh
> docker ps -a
>  ```
> * Remove Docker image (although it is a bad practice, add `-f` force remove) 
with a few characters of the `CONTAINER_ID` which you can find the one assigned to your
container with the `ps` commands above.
>
> ```sh
> docker rm 
>  ```
> * Bonus: Delete **stopped** containers
>
> ```sh
> docker rm -v $(docker ps -a -q -f status=exited)
>  ```

#### Security
Set a *user/pass* for the redis container.

#### Testing
In this app, data store testing is adopts the approach that a separate redis docker container is created
for test usage. However, there are other ways worth trying:
* [Wrapping for using go-redis way](https://github.com/go-redis/redis/issues/332)

or, if you dare to code through another API anyway:

* [Miniredis](https://github.com/alicebob/miniredis)

### Docker
#### Build
See the `./Dockerfile`

#### Data Persistence
A volume is assigned to redis container by `-v` flag (see previous sections).

There is currently no need for a _volume_ or _bind mount_.

### On Kubernetes work
Once the *container* is build, the logic outside a container (including linking containers) can be built on [Kubernetes](https://kubernetes.io/). 

See:
* [Expose Pod Information to Containers Through Environment Variables](https://kubernetes.io/docs/tasks/inject-data-application/environment-variable-expose-pod-information/)

### Useful links
* [A full, secure gRPC client & server implementation guide](https://medium.com/pantomath/how-we-use-grpc-to-build-a-client-server-system-in-go-dd20045fa1c2)
* GOENV and environment specific settings??
