# healthz-demo

A sample application, simply called `healthz`, used for demonstrating Kubernetes liveness vs readiness probes.

Provides a basic REST API that has three endpoints:

* `/healthz` - always returns 200 OK
* `/readyz` - returns 503 "Service Unavailable," but goes "green" and returns 200 once the app has been running for 20 seconds
* `/unready` - resets the 20 second time used by `/readyz`

## Building and running locally

The included `Dockerfile` is all you need to build and run this app locally.

### Build it

```shell
docker build -t my-healthz .
```

### Run it

```shell
docker run -d -p 8080:8080 my-healthz
```

### Hit it

```shell
# /healthz
curl http://localhost:8080/healthz

# /readyz
curl http://localhost:8080/readyz

# /unready
curl http://localhost:8080/unready
```
