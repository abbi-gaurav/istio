# Overview

Install a custom Istio adapter to a Kyma cluster.

## References

* [Mixer Out of Process Adapter Walk through](https://github.com/istio/istio/wiki/Mixer-Out-Of-Process-Adapter-Walkthrough)
* [A reference guide](https://medium.com/google-cloud/simple-istio-mixer-out-of-process-authorization-adapter-5f9363cd9bbc)

## Customize

* Logic resides in [adapter code](./mygrpcadapter.go).
* Make any required changes.

## Build

* Builds and pushes the docker image

  ```shell script
    DOCKER_ACCOUNT=<your-docker-account> make push-image
    ```

## Install on a Kyma cluster
* Change the image value in [deployment](./deployment/deployment.yaml).

* Deploy adapter as a k8s service.

    ```shell script
    kubectl apply -f deployment/deployment.yaml
    ```

* Setup the attributes maps and deploy

  ```shell script
    kubectl apply -f testdata/attributes.yaml
    ```

* Create template

    ```shell script
    kubectl apply -f testdata/template.yaml
    ```

* Apply adapter configuration

  ```shell script
    kubectl apply -f testdata/mygrpcadapter.yaml
    ``` 

* Create the operator

  ```shell script
    kubectl apply -f testdata/sample_operator_cfg.yaml
    ```

* Check the mixer logs

  ```shell script
    kubectl -n istio-system logs $(kubectl -n istio-system get pods -lchart=mixer -o jsonpath='{.items[0].metadata.name}') -c mixer
    ``` 

    ```shell script
    # a sample
    2019-11-18T12:24:50.239542Z     info    ccResolverWrapper: sending update to cc: {[{mygrpcadapter:44225 0  <nil>}] <nil>}
    2019-11-18T12:24:50.239567Z     info    grpcAdapter     Connected to: mygrpcadapter:44225
    2019-11-18T12:24:50.239799Z     info    base.baseBalancer: got new ClientConn state: {{[{mygrpcadapter:44225 0  <nil>}] <nil>} <nil>}
    ```

* Send a HTTP request

* Check the adapter logs

    ```shell script
    kubectl -n istio-system logs -l app=mygrpcadapter
    ```