# Bookinfo Sample

See <https://istio.io/docs/examples/bookinfo/>

## Build docker images without pushing

```bash
src/build-services.sh <version>
```

The bookinfo versions are different from Istio versions since the sample should work with any version of Istio.

## Update docker images in the yaml files

```bash
sed -i "s/\(istio\/examples-bookinfo-.*\):[[:digit:]]\.[[:digit:]]\.[[:digit:]]/<your docker image with tag>/g" */bookinfo*.yaml
```

## Push docker images to docker hub

One script to build the docker images, push them to docker hub and to update the yaml files

```bash
build_push_update_images.sh <version>
```

## Tests

Bookinfo is tested by e2e smoke test on every PR. The Bookinfo e2e test is in [tests/e2e/tests/bookinfo](https://github.com/istio/istio/tree/master/tests/e2e/tests/bookinfo), make target `e2e_bookinfo`.

The reference productpage HTML files are in [tests/apps/bookinfo/output](https://github.com/istio/istio/tree/master/tests/apps/bookinfo/output). If the productpage HTML produced by the app is changed, remember to regenerate the reference HTML files and commit them with the same PR.

# NextGen deployment

Nodes

Create 3 nodes, name them node1, node2, node3

Install kubernetes

git clone http://github.com/tungstenfabric/tf-devstack
tf-devstack/k8s_manifests/run.sh platform

Install istio

see istio guide

Get bookinfo

git clone http://github.com/alexandrelevine/istio

Configure on all nodes

cat /etc/docker/daemon.json
{
    "insecure-registries": ["node1:5000"]
}

Build bookinfo

sudo docker run -d -p 5000:5000 --restart=always --name registry registry:2
cd istio/samples/bookinfo
sudo src/build-services.sh 1.1 "node1:5000/istio"

