---
title: 基于k8s搭建个人混合云数据中心：高效利用边缘计算与云服务
---

<!-- Header Section -->
# 基于k3s搭建个人混合云数据中心：高效利用边缘计算与云服务

<!-- Navigation Bar -->
## 目录
1. [背景与问题](#背景与问题)
2. [实现方案](#实现方案)
   1. [方案1: 集群混合部署：公有云计算节点和私有云计算节点混合部署](#方案1-集群混合部署公有云计算节点和私有云计算节点混合部署)
   2. [方案2: 集群部署在私有计算节点上，使用公有云节点当作ingress](#方案2-集群部署在私有计算节点上使用公有云节点当作ingress)
3. [测试](#测试)

<!-- Main Content Section -->
## 背景与问题
随着云计算和边缘计算技术的不断发展，个人用户在使用云主机时面临着一个不可忽视的问题：随着计算需求的增加，云主机的费用也随之增长。这对于那些有特定计算需求，但又不希望一直支付高额云费用的个人用户来说，构建一个混合云架构显得尤为重要。

在传统的云计算模式下，用户的数据和计算任务会被完全转移到云数据中心，随着计算量和存储需求的增加，云计算的费用也会显著提升。而对于许多家庭用户来说，往往有较强的本地计算能力——例如，个人闲置的家用计算机通常配备了较高的硬件配置（如多核CPU、大量内存和硬盘存储），但是由于这些设备并未得到充分利用，造成了资源的浪费。

因此，将边缘计算能力与云服务结合，搭建一个个人混合云数据中心，成为解决这一问题的理想方案。通过这种方式，用户可以充分利用本地计算机的闲置资源，将一些计算密集型任务和存储任务转移到本地边缘节点上，而将需要更高可用性、灵活性或规模的任务交给公有云来处理。

### 这种架构的优势：
1. **降低云计算成本**：通过将部分计算和存储任务转移到本地边缘节点，减少了对云服务的依赖，进而减少了云计算的费用。
2. **高效利用闲置资源**：充分利用家用计算机等本地设备的计算和存储能力，避免了硬件资源的浪费。
3. **降低网络延迟**：将边缘计算部署在离数据源更近的地方，减少了网络延迟，尤其适用于延迟敏感型应用。
4. **灵活性与可扩展性**：通过公有云和本地计算机的结合，用户可以灵活调度和扩展资源，既能保证高效性，也能提供类似公有云的管理体验。

然而，在实现个人混合云数据中心时，仍然面临一些挑战，特别是在网络连接、同步以及资源的调度与管理方面。例如，如何有效地连接本地设备与云服务，保证数据的安全与同步？如何在网络延迟较高的情况下，确保系统稳定运行？

因此，搭建一个既能减少云计算成本，又能高效利用本地硬件资源的个人混合云数据中心，变成了一个有意思的事情。

## 实现方案
这里我们尝试两种方案
1. 集群混合部署：公有云计算节点和私有云计算节点混合部署  
2. 集群部署在私有计算节点上，使用公有云节点当作ingress

### 方案1: 集群混合部署：公有云计算节点和私有云计算节点混合部署 
传统的方案通常将私有云和公有云独立架构，然后使用常用中间件来打通数据流和控制流，这种方式往往增加了运维复杂度，尤其是在数据交互和安全管理方面的挑战。特别是在省钱省事的前提下，成为了安全和运维的难点。本文尝试一种新的解决方案：通过K8s实现公有云和私有计算中心的统一，利用K8s的原生能力来解决网络连接、延迟等带来的数据同步，安全运维等问题。这个方案通过横向扩展两个网络，尝试实现更加高效和简化的云数据中心架构，提供了一个实用的实践方向。 

![high level arch](high-level-arch-hybrid-cluster.png)

#### 打通网络
第一步应该是打通本地主机到云主机的网络，如上图中虚线所示，我们使用wireguard (被Linus成为“the best VPN protocol”)

1. 配置wireguard server，我们标记为A，执行下面脚本 — 注意注释掉的部分

    <details>

    ```bash
    #!/bin/bash

    # wireguard server A
    mkdir -p /wireguard
    cd /wireguard
    wg genkey | tee server_A_private_key | wg pubkey > server_A_public_key

    if [ ! -f /wireguard/client_B_public_key ]; then
        echo "file client public key does not exist"
        exit 1
    fi

    cat <<EOF > /wireguard/wg0.conf
    [Interface]
    PrivateKey = `cat /wireguard/server_A_private_key` # private key of server A
    Address = 10.4.0.1/24      # virtual server IP of A node
    ListenPort = 51820
    SaveConfig = true

    [Peer]
    PublicKey = `cat /wireguard/client_B_public_key`
    AllowedIPs = 10.4.0.2/32   # virtual IP of B node
    EOF

    cp /wireguard/wg0.conf /etc/wireguard/wg0.conf

    sudo wg-quick up wg0
    #sudo systemctl enable wg-quick@wg0

    #sudo ufw allow 51820/udp
    #sudo ufw enable

    #希望 B 设备通过 A 访问 Internet 或其他网络资源，你需要在 A 上配置 IP 转发
    # 启用 IP 转发
    #echo 1 | sudo tee /proc/sys/net/ipv4/ip_forward

    # 修改 sysctl 配置以保持设置
    #sudo sysctl -w net.ipv4.ip_forward=1
    #sudo sysctl -p
    #
    #sudo iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE
    #sudo iptables-save
    ```

    </details>

2. 配置peer client 节点，我们标记为B（同理可以创建C节点），执行下面脚本 — 注意注释掉的部分

    <details>

    ```bash
    #!/bin/bash

    # wireguard client B
    mkdir -p /wireguard
    cd /wireguard
    wg genkey | tee client_B_private_key | wg pubkey > client_B_public_key

    if [ ! -f /wireguard/server_A_public_key ]; then
        echo "file server public key does not exist"
        exit 1
    fi

    cat <<EOF > /etc/wireguard/wg0.conf
    [Interface]
    PrivateKey = `cat /wireguard/client_B_private_key` # private key of B
    Address = 10.4.0.2/32      # virtual server IP of B

    [Peer]
    PublicKey = `cat /wireguard/server_A_public_key`
    Endpoint = 74.121.149.207:51820 # public IP and listening port
    AllowedIPs = 10.4.0.0/24   # virtual IP of B node
    PersistentKeepalive = 25
    EOF

    cp /wireguard/wg0.conf /etc/wireguard/wg0.conf

    sudo wg-quick up wg0
    #sudo systemctl enable wg-quick@wg0 # start on system on
    ```

    </details>

3. 测试链接，ping 成功搞定

    ```bash
    ping 10.4.0.1
    PING 10.4.0.1 (10.4.0.1): 56 data bytes
    64 bytes from 10.4.0.1: icmp_seq=0 ttl=64 time=300.125 ms
    64 bytes from 10.4.0.1: icmp_seq=1 ttl=64 time=320.622 ms
    ```

#### 搭建k3s control plane on cloud
登录到A节点，启动controlplane
```bash
curl -sfL https://get.k3s.io | sh -
```
根据个人环境优化参数，如·网络丢包率为 12%，延时约 250ms·，建议优化启动参数.
也可以使用默认参数，目前默认行为k3s
1. 启动sqlite存储  
2. Kubelet默认同步node节点数据频率为1分钟一次
3. 默认API Server 等待状态更新的最大容忍时间为5m
```/etc/systemd/system/k3s.service
ExecStart=/usr/local/bin/k3s server --tls-san axpz.local

...

--kubelet-arg=--node-status-update-frequency=60s
--kubelet-arg=--node-status-update-grace-period=5m
```
```
sudo systemctl daemon-reload
sudo systemctl restart k3s
```
查看日志
    
<details>
    
```bash  
# 使用以下命令来过滤 K3s 日志中的 Node Token 和 TLS-SAN 信息
# journalctl -u k3s | grep 'node token\|tls-san'

Jan 07 15:30:36 calm-baud-1.localdomain k3s[3094620]: time="2025-01-07T15:30:36Z" level=info msg="Server node token is available at /var/lib/rancher/k3s/server/token"
Jan 07 15:30:36 calm-baud-1.localdomain k3s[3094620]: time="2025-01-07T15:30:36Z" level=info msg="Agent node token is available at /var/lib/rancher/k3s/server/agent-token"
Jan 07 15:30:47 calm-baud-1.localdomain k3s[3094779]: time="2025-01-07T15:30:47Z" level=info msg="Server node token is available at /var/lib/rancher/k3s/server/token"
Jan 07 15:30:47 calm-baud-1.localdomain k3s[3094779]: time="2025-01-07T15:30:47Z" level=info msg="Agent node token is available at /var/lib/rancher/k3s/server/agent-token"
Jan 07 15:30:54 calm-baud-1.localdomain k3s[3094779]: I0107 15:30:54.121292 3094779 kube.go:636] List of node(calm-baud-1.localdomain) annotations: map[string]string{"alpha.kubernetes.io/provided-node-ip":"74.121.149.207", "flannel.alpha.coreos.com/backend-data":"{\"VNI\":1,\"VtepMAC\":\"3e:a5:8d:f1:50:b4\"}", "flannel.alpha.coreos.com/backend-type":"vxlan", "flannel.alpha.coreos.com/kube-subnet-manager":"true", "flannel.alpha.coreos.com/public-ip":"74.121.149.207", "k3s.io/hostname":"calm-baud-1.localdomain", "k3s.io/internal-ip":"74.121.149.207", "k3s.io/node-args":"[\"server\",\"--tls-san\",\"axpz.local\"]", "k3s.io/node-config-hash":"EW5CHGDY6WQFH5IMOMMGEEEOFXADTGZZPMW5RZGHVQBJ3QHTZUHA====", "k3s.io/node-env":"{}", "node.alpha.kubernetes.io/ttl":"0", "volumes.kubernetes.io/controller-managed-attach-detach":"true"}
```

</details>

#### 启动traefik dashboard
edit /var/lib/rancher/k3s/server/manifests/traefik.yaml with
```
apiVersion: helm.cattle.io/v1
kind: HelmChart
metadata:
  name: traefik
  namespace: kube-system
spec:
  chart: https://%{KUBERNETES_API}%/static/charts/traefik-27.0.201+up27.0.2.tgz
  set:
    global.systemDefaultRegistry: ""
    dashboard.enabled: "true"               <<<
    dashboard.domain: "traefik.internal"    <<<
```
kubectl apply -f /var/lib/rancher/k3s/server/manifests/traefik.yaml

#### 搭建k3s worker on `macOS`
首先启动一个轻量级虚拟机
```bash
multipass launch --name k3s-worker --memory 4G --disk 40G --mount ~/go/src/github.com:/mnt/github.com
multipass shell k3s-worker
curl -sfL https://get.k3s.io | K3S_URL=https://axpz.local:6443 K3S_TOKEN=K10...d0eb56c9 sh -
```

#### 系统搭建完成，注意下面几点

1. 云主机的公网IP会自动被treafik的serviceLB创建为LoadBalancer IP，即默认的ingress，你可以通过配置ingressRoute自动分流到到不同的微服务中。
2. 如果各个节点硬件的architecture不一样，需要注意kubernetes的image兼容性问题，比如controlplane是x86-64, 而工作节点是arm64，可能需要额外的工作。
3. 优化： 
    - 优化启动参数：比如上面提到的网络丢包率为 12%，延时约 250ms，需要优化数据同步频率，以及存储等, 比如：`ExecStart=/usr/local/bin/k3s server --tls-san axpz.local --flannel-backend=none --disable-network-policy --no-deploy local-storage --no-deploy traefik`
    - 减少主节点资源消耗：给主节点打污点 `kubectl taint nodes calm-baud-1.localdomain node-role.kubernetes.io/master:NoSchedule`
    - 网络优化，如果作为生产级应用，建议搭建网络专线
    - 其他...

### 方案2: 集群部署在私有计算节点上，使用公有云节点当作ingress

![high level arch](high-level-arch-private-cluster.png)

#### 打通网络、搭建集群

参考方案1 

#### 配置集群负载均衡为wg0

k3s 使用serviceLB来提供负载均衡服务，当有external-ip 时，serviceLB会自动使用其作为负载均衡，当没有的时候会自动使用internal ip作为负载均衡，由于我们配置了vpn连通外网，所以我们配置wg0作为外网ip，具体配置添加启动参数
--node-external-ip 10.4.0.3
所以启动service变为ExecStart=/usr/local/bin/k3s server --node-external-ip 10.4.0.3
```bash
systemctl daemon-reload
systemctl stop k3s
systemctl start k3s
```

## 测试

1. 创建nginx应用
    <details>

    ``` yaml
    apiVersion: apps/v1 
    kind: Deployment 
    metadata:
    name: nginx-deployment
    namespace: default
    spec:
    replicas: 1
    selector:
        matchLabels:
        app: nginx
    template:
        metadata:
        labels:
            app: nginx
        spec:
        containers:
        - name: nginx
            image: nginx:latest
            ports:
            - containerPort: 80
    ---
    apiVersion: apps/v1
    kind: Deployment
    metadata:
    name: nginx-deployment2
    namespace: default
    spec:
    replicas: 1
    selector:
        matchLabels:
        app: nginx2
    template:
        metadata:
        labels:
            app: nginx2
        spec:
        containers:
        - name: nginx
            image: nginx:latest
            ports:
            - containerPort: 80
    ```

    </details>
    

2. 暴露服务
    <details>

    ``` yaml
    apiVersion: v1
    kind: Service
    metadata:
    name: nginx-service
    namespace: default
    spec:
    selector:
        app: nginx
    ports:
        - protocol: TCP
        port: 80
        targetPort: 80
    # 这个服务类型可以改成 ClusterIP 或者 NodePort
    type: ClusterIP
    ```

    </details>

3. 配置IngressRoute
    <details>

    ``` yaml
    apiVersion: traefik.containo.us/v1alpha1
    kind: IngressRoute
    metadata:
    name: nginx-ingressroute
    namespace: default
    spec:
    entryPoints:
        - web   # 确保 Traefik 配置了 `web` entry point
    routes:
        - match: Host(`k3s.local`)  # 使用你自己的域名，或者可以用外部 IP
        kind: Rule
        services:
            - name: nginx-service
            port: 80
        - match: Host(`k3s.local`) && PathPrefix(`/n2`)
        kind: Rule
        services:
            - name: nginx-service2
            port: 80
    ```
    </details>

从互联网访问，..., 搞定, scale up/down pod， ... 搞定

<!-- Footer Section -->
<div class="footer">
    <p>© 2025 基于k3s搭建个人混合云数据中心项目</p>
</div>

