# Basic GoLang server

![Frontend](https://raw.githubusercontent.com/christi4n/hello-go/master/assets/golang-app-hello-go.png)

## Plain HTML response (with already defined route)

    http://localhost:8181/hello

## Plain HTML with template

    http://localhost:8181/

## JSON response

    curl --location --request GET 'http://localhost:8181/api/ping'

## Build a binary

    go build
    ./hello-go

## Build a Docker image

    nerdctl build -t hello-go .
    nerdctl images hello-go

REPOSITORY    TAG       IMAGE ID        CREATED          SIZE
hello-go      latest    e3fcd9cb6639    2 minutes ago    14.9 MiB

## Run the image with nerdctl container runtime

    nerdctl run --rm -p 8181:8181 hello-go

## K8S deployment

Create a new namespace:

    kubectl create ns sample-go-apps

Then, create the deployment

You have two choices:

- declarative - use files in the k8s directory with the **kubectl apply -f** command
- imperative - see below for both deployment and service

### Imperative method

    kubectl create deployment hello-go --image=chrisnt5/hello-go --namespace=sample-go-apps --image-pull-policy=IfNotPresent

Check errors if necessary

    kubectl describe pod hello-go --namespace=sample-go-apps

Check deployment

    kubectl get deployment hello-go --namespace=sample-go-apps

    NAME       READY   UP-TO-DATE   AVAILABLE   AGE
    hello-go   1/1     1            1           22m

Create a service:

    kubectl expose deployment hello-go --type=LoadBalancer --port=8181 --namespace=sample-go-apps

Check the service:

kubectl describe service hello-go --namespace=sample-go-apps
## Snyk analysis for vulnerabilities

    snyk auth
    snyk container test chrisnt5/hello-go:1.0.1 --file-Dockerfile

![Snyk analysis](https://raw.githubusercontent.com/christi4n/hello-go/master/assets/snyk-analysis.png)

## Scan with Kubevious CLI

### Kubevious image

Use the image chrisnt5/kubevious-cli:1.x.x store in Docker hub.

### Command

You can use the following command:

```
   docker run -v $(pwd)/k8s:/src chrisnt5/kubevious-cli:1.0.0 "kubevious guard src --detailed --ignore-unknown"
```
