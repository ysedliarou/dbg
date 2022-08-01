# Remote debugging for go-lang applications

## Docker

``` shell
    $ docker run -p 8080:8080 -p 30123:30123 ysedcoupa/dbg:0.1
```

## Kubernetes

``` shell
    $ minikube start

    $ kubectl create -f kub/deployment.yaml

    $ kubectl port-forward <POD> 8080:8080
    $ kubectl port-forward <POD> 30123:30123
    
    $ curl localhost:8080
```


