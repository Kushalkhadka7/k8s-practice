## Setup k8s enviornment

- [Setup](setup)

kubectl expose deployemnt <name> --port <port in which app is running in container> --type NodePort
kubectl port-forward <pod name> <local port>:<port in which the app is running>

metallb deployment

- create name space metallb-system
- kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.10.2/manifests/namespace.yaml

- deploy metallb
- kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.10.2/manifests/metallb.yaml

if we want to give a range of ip address to the network the n we can add a config as mentioned in metallb doc layar2 config, if not external ip will be available.
we can access it throug the external ip.

install kubernetes dashboard ui.

- kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.2.0/aio/deploy/recommended.yaml

it will create different resources including namespace, role, clusterrole binding, custer role pods, services etc.
But the service account that is created doesnt have full accces to the dashboard.
so we need to create a cluster role with admin access and cluster role binding.

using /kubernetes/sa/dashboard-admin.yml
it will create a sercive account dashboard-admin with admin access.

it will be deployed on kube-system namespace.

get token to login.

kubectl -n kube-system describe sa dashboard-admin
copy the token secret

kiubectl -n kube-system describe secret <token secret>

it will provide the actual token to loing.

now the sevice kubernetes-dashboard wont be exposed.
so expose it using nodeport, portforwarding or any one method.
use the token to login to the dashboard.

Node Pod selector.

- Attach a label to the node in which we want to deploy the pods.
- kubectl label node <node_name> <key=value> -> label must be on key value pair
- eg. kubectl label node kworker1 env=dev -> here env=dev it the label.
- kubectl get node <node_name> --show-labels

now in pod specification add

nodeSelector:
env: "dev"

- enable pod node selector plugin.
- if we deploy pods ot a namespace these pods should be deployed to specific pod only.
- eg. Deploy pod to dev name space , the this pod should be created to the node where label is env=dev.
- Pod node selector is a plugin need to be enabled from the contorl plane k8s.
- It is the alternative of node pod selector.

how to do it:
in the master node go to cd /etc/kubernetes/manifests/
update the kube-apiserver.yml file
updtate this - --enable-admission-plugins=NodeRestriction,PodNodeSelector
master pod will be restarted and the pulgin will be enabled.

Pod Disruption budget

- Used in maintainance of the cluster.
- set pod distruption budget to 50%, says that 50% of the pods will always be there.
- Prevents accidently deleting the pods while maintainance window actives.

<!-- jobs and corn jobs -->

jobs:

- run to completion one time
- These jobs continues to run as long as they are not sucessfully completed.
- If the are interupted in middle of the run, the run agian until they gets completed.
- we can limit the no of retries for the job.(backoffLimit: 2).
- We can run job as many times as needed by using (completions: 2).
- We can also run jobs in parrallel ( parallism: 2)=> runs 2 jobs in parallel.
- We can also run the jobs forever.

corn jobs:

- run to completion in specific scheduled time.
- we can do much things similar to that of jobs.
- we can suspend the job after that it wont do or create any new jobs from now.
- we can resume it anytime.
- We can even run the jobs concurrently (enabled by default).(more than one job running at new time).
- We can also say complete the job and only run the new one.
  useCases:
  - generally used for automated tasks
  - eg. sending email, db backup
