# test-flow

## Starting multipass

I have changed the memory to 8GB in case 4GB is not enough and increased the disk space.

multipass launch --cpus 4 --disk 20G --memory 8G --name direktiv --cloud-init https://raw.githubusercontent.com/direktiv/direktiv/all-in-one/build/docker/all/multipass/init.yaml

Mount this repository to a namespace `test-flow`. 

## basic

The flow `basic` shows how to transition between 3 different states and merge the results to a meaningful response for the service. 

```
curl http://<IP>/api/namespaces/test-flow/tree/basic?op=wait
```
## async 

Simple example how to call a function in async mode. These functions return no value.

```
curl http://<IP>/api/namespaces/test-flow/tree/async?op=wait
```

## python

This is a three state example with python. In this release it is based on naming conventions and this will be changed in the next one. If names are like `workflowname.yaml.mything.txt` it will be created as workflow variable. These vraiables can be referenced in the `files` section in the workflow and they are available in the function.

```
curl "http://<IP>/api/namespaces/test-flow/tree/python?op=wait&input.currency=USD
```

## parallel

This is a simple parallel workflow. Functions usually return their data in `return` to the state data. If it is parallel it is an arry and need to be access with `return[X]`.

```
http://<IP>/api/namespaces/test-flow/tree/python?op=wait
```

## custom

The folder `custom-container` contains an example how to build a simple custom function including a `Dockerfile`. It is used in the workflow `custom.yaml`. You can write all kinds of custom functions/containers. As long as they listen to port 8080 and returning JSON. 

```
curl http://<IP>/api/namespaces/test-flow/tree/custom?op=wait
```

### Call With No Payload

```
curl "http://<IP>/api/namespaces/test-flow/tree/custom?op=wait"
```

### Call With Content

Payloads for flows can be POST or GET requests. In GET requests the payload can be passed with a prefix `input.`.

```
curl "http://<IP>/api/namespaces/test-flow/tree/custom?op=wait&input.name=WORLD"
```