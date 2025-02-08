#!/bin/bash

set -x
set -o errexit
set -o pipefail
set -o nounset

helm repo add spiderpool https://spidernet-io.github.io/spiderpool
helm repo update spiderpool

VERSION=v0.9.7

echo "load images to kind cluster ${E2E_CLUSTER_NAME}"
IMAGES=$( helm template spiderpool spiderpool/spiderpool --version ${VERSION} --set global.imageRegistryOverride=ghcr.m.daocloud.io | grep "image:"  | awk '{print $2}' | sort | tr -d '"' | uniq )
echo "IMAGES"
echo "${IMAGES}"
for IMAGE in $IMAGES; do
    echo "loading $IMAGE"
    docker inspect $IMAGE &>/dev/null || docker pull $IMAGE 
    kind load docker-image $IMAGE --name ${E2E_CLUSTER_NAME}
done

echo "install spiderpool"
helm uninstall spiderpool -n  spiderpool || true 
helm install spiderpool spiderpool/spiderpool \
  --wait \
  --version ${VERSION} \
  --namespace spiderpool \
  --create-namespace \
  --set global.imageRegistryOverride=ghcr.m.daocloud.io \
  --set plugins.installCNI=true


# INTERFACE=eth0
# cat <<EOF | kubectl apply -f -
# apiVersion: spiderpool.spidernet.io/v2beta1
# kind: SpiderMultusConfig
# metadata:
#   name: ${INTERFACE}-macvlan
#   namespace: spiderpool
# spec:
#   cniType: macvlan
#   disableIPAM: true
#   macvlan:
#     master: ["${INTERFACE}"]
# EOF

INTERFACE=eth1
cat <<EOF | kubectl apply -f -
apiVersion: spiderpool.spidernet.io/v2beta1
kind: SpiderIPPool
metadata:
  name: ${INTERFACE}-pool
spec:
  #gateway: 192.168.1.1
  subnet: 192.168.1.0/24
  ips:
    - 192.168.1.200-192.168.1.250
---
apiVersion: spiderpool.spidernet.io/v2beta1
kind: SpiderMultusConfig
metadata:
  name: ${INTERFACE}-macvlan
  namespace: spiderpool
spec:
  cniType: macvlan
  macvlan:
    master: ["${INTERFACE}"]
    ippools:
      ipv4: ["${INTERFACE}-pool"]
EOF

INTERFACE=eth1
VLAN_ID=10
cat <<EOF | kubectl apply -f -
apiVersion: spiderpool.spidernet.io/v2beta1
kind: SpiderIPPool
metadata:
  name: ${INTERFACE}-vlan${VLAN_ID}-pool
spec:
  #gateway: 192.168.10.1
  subnet: 192.168.10.0/24
  ips:
    - 192.168.10.200-192.168.10.250
---
apiVersion: spiderpool.spidernet.io/v2beta1
kind: SpiderMultusConfig
metadata:
  name: ${INTERFACE}-vlan${VLAN_ID}-macvlan
  namespace: spiderpool
spec:
  cniType: macvlan
  macvlan:
    master: ["${INTERFACE}"]
    vlanID: ${VLAN_ID}
    ippools:
      ipv4: ["${INTERFACE}-vlan${VLAN_ID}-pool"]
EOF

echo "finish installing spiderpool"

