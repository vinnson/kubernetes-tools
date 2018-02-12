## K8s Installation
## Setup script based on URL below:
## https://medium.com/@SystemMining/setup-kubenetes-cluster-on-ubuntu-16-04-with-kubeadm-336f4061d929
## Command Line
#######################################
sudo apt update && apt install -y apt-transport-https && apt upgrade -y && reboot
sudo curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
sudo cat <<EOF > /etc/apt/sources.list.d/kubernetes.list
deb http://apt.kubernetes.io/ kubernetes-xenial main
EOF
sudo modprobe br_netfilter
sudo sysctl -w net.bridge.bridge-nf-call-iptables=1

sudo apt update && apt install -y docker.io kubelet kubeadm kubectl kubernetes-cni
sudo kubeadm init --pod-network-cidr 10.244.0.0/16

################### Keep the join command from return (for nodes) #######
## kubeadm join --token xxx.xxxxxxx YOUR_MASTER_IP:6443 --discovery-token-ca-cert-hash sha256:xxxxxxxxxxxx
#########################################################################

## No SUDO anymore
$ mkdir -p $HOME/.kube
$ sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
$ sudo chown $(id -u):$(id -g) $HOME/.kube/config
$ kubectl get nodes
## You are the admin user now

## Apply Canal Network ##
$ kubectl apply -f https://raw.githubusercontent.com/projectcalico/canal/master/k8s-install/1.7/rbac.yaml
$ kubectl apply -f https://raw.githubusercontent.com/projectcalico/canal/master/k8s-install/1.7/canal.yaml
