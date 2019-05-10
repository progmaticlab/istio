minikube stop
sleep 10;
minikube delete
sleep 10;
rm -fr ~/.minikube/
sleep 10;
docker system prune -a -f
sleep 10;
sudo rmmod -f kvm kvm_intel
sleep 10;
sudo modprobe kvm
sudo modprobe kvm_intel
sleep 10;
docker build -t "francoispesce/kalitest:0.9.0" -t francoispesce/kalitest:latest .
sleep 10;
docker push francoispesce/kalitest:latest
sleep 10;
minikube start --vm-driver=kvm2 --bootstrapper=kubeadm --cpus=4 --memory=12288 --disk-size=48g
sleep 10;
kubectl config use-context minikube
sleep 10;
kubectl apply -f $ISTIOPATH/install/kubernetes/istio-demo-auth.yaml
sleep 10
kubectl apply -f $ISTIOPATH/install/kubernetes/istio-demo-auth.yaml
sleep 60
kubectl create -f <(istioctl kube-inject -f ${ISTIOPATH}/samples/bookinfo/platform/kube/bookinfo.yaml)
sleep 600
kubectl apply -f $ISTIOPATH/samples/bookinfo/networking/bookinfo-gateway.yaml
