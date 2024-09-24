# Gofile

`Gofile` is an alternative way to package a Go binary into a minimal container image without using a `Dockerfile`.

This is purely a fun project to help me learn more about BuildKit its features.

> It's pretty much similar to the [gockerfile](https://github.com/po3rin/gockerfile) project but maybe I'll add some extra features in the future, who knows.

## Quick Start

Create a `Gofile.yaml` file indicating the Git repository, the path, and the Git ref where the Go application is located.

```yaml
#syntax=felipecruz/gofile
repo: https://github.com/dockersamples/helloworld-go-demo
path: .
ref: main
scratch: true # Use scratch image as base with CA certs, otherwise defaults to alpine
```

Or to build a Go binary from a local directory:

```yaml
path: ./example/demo
scratch: true
```

Build the image:

```bash
docker build -t felipecruz/gofile-demo -f Gofile.yaml .
```

Check the image size:

```bash
docker image ls felipecruz/gofile-demo
REPOSITORY               TAG       IMAGE ID       CREATED          SIZE
felipecruz/gofile-demo   latest    fe66d2c25ca2   14 minutes ago   7.33MB
```

Run the container:

```bash
docker run --rm -p 8080:8080 felipecruz/gofile-demo /bin/server
```

Make a request:

```bash
curl localhost:8080

          ##         .
    ## ## ##        ==
 ## ## ## ## ##    ===
/"""""""""""""""""\___/ ===
{                       /  ===-
\______ O           __/
 \    \         __/
  \____\_______/


Hello from Docker!
```
