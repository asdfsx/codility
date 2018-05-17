使用docker 搭建etcd

使用etcd来做测试环境
```
docker pull fishead/quay.io.coreos.etcd:v3.2.4
docker run -b \
-e "ETCD_LISTEN_CLIENT_URLS=http://0.0.0.0:2379" \
-e "ETCD_ADVERTISE_CLIENT_URLS=http://192.168.99.100:2379" \
-p 2379:2379 fishead/quay.io.coreos.etcd:v3.2.4
```