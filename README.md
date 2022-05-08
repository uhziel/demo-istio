# demo-istio

试验 istio 的一些功能。主要侧重 tcp。

# 操作

```
$ kubectl create namespace tcp-test
$ kubectl label namespaces tcp-test istio-injection=enabled
$ kubectl apply -f manifests
```
