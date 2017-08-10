# Docker Engine Playground

Simple application to test the Docker Engine Go 1.3 API

### What it do:
Will output all processes for each container running on a host

### How you do it:
You can run this one of two ways.

Standalone:
```
$ go build
$ ./hello-container
```

Inside another container:

```
$ docker build -t my-golang-app .
$ docker run -v /var/run/docker.sock:/var/run/docker.sock -it --rm --name my-running-app my-golang-app
```


### Sample Output:
```
5d7a39a05189734e5e580c4eede1d839782c82c8206f724de66386e330b6276f
PID USER TIME COMMAND
4308 root 3:29 /clair -config /config/config.yaml
cf70aa382c7c56357522426f4913f01327947d6fb29ccb46fce351c0a1720210
PID USER TIME COMMAND
4175 999 0:00 postgres
4351 999 0:02 postgres: checkpointer process
4352 999 0:00 postgres: writer process
4353 999 0:00 postgres: wal writer process
4354 999 0:00 postgres: autovacuum launcher process
4355 999 0:00 postgres: stats collector process
4369 999 3:09 postgres: postgres postgres 172.18.0.3(37576) idle
4375 999 2:24 postgres: postgres postgres 172.18.0.3(37580) idle
```
