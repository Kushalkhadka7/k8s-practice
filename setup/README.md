# Setting LXC/LXD

```
$ sudo apt update
```

```
$ sudo apt install lxd && sudo apt install lxc -y
```

# Check the status of lxd and lxc

```
$ sudo systemctl status snap.lxd.daemon && sudo systemctl status lxc
```

# Initialize the lxd enviorment for containers

```
$ lxc init
```

# To lauch sample container

```
$ lxc launch <image>:<version> <container name>
```

## Show all the images in local

```
$ lxc image list
```

# Create profile

```
# Copy defult profile to new prifle.

$ lxc profile copy default <new_profile_name>

```

```
# Save profile for future use.

$ lxc profile show <profile_name> > <profile_name>.yml
```

# Create container using profile

```
# Apply profile to the container

$ lxc launch <image>:<version> <container_name> --profile <profile_name>
```

# Creating a kubernets cluster

```
$ lxc image list
```

# Create master container.
```
$ lxc profile default k8s-master

$ lxc profile edit k8s-master

$ lxc profile show k8s-master > k8s-master.yml

$ lxc launch ubuntu:16.04 --profile k8s-master
```

# Create worker containers (2 worker containers).
```
$ lxc profile default k8s-worker

$ lxc profile edit k8s-worker

$ lxc profile show k8s-worker > k8s-worker.yml

$ lxc launch ubuntu:16.04 --profile k8s-worker

$ lxc exec <master_container_name> bash

$ lxc exec <worker_container_name> bash

$ lxc exec <worker_container_name> bash
```
# Verify installation
```
$ lxc exec kmaster bash

$ kubectl get nodes -o wide

```

# Some useful lxc commands

```
# Help
$ lxc help

# Show storage list
$ lxc storate list

# Search image in remote repo
$ lxc image list image:<image-name>

# Launch container
$ lxc launch <image>:<version> <container name>

# Stop container
$ lxc stop container name

# Delete container
$ lxc delete container name

# Force delete container
$ lxc delete container name --force

# Shell inside the running container
$ lxc exec <container_name> bash

# Ping other lxc container
$ ping <another_container_name>.lxd

# Container info
$ lxc info <container_name>

# Container machine configuration
$ lxc config show <container_name>

# List all profiles
$ lxc profile list

# List profile detail
$ lxc profile show <profile name>

# Copy profile
$ lxc profile copy <souce> <destination>

# Attach profile while creating conatianer
$ lxc launch <image>:<version> <continer_name> --profile <profile>

# Configuer memory limit for container
$ lxc config set <container_name> limits.memory <memory_size> eg.512MB

# Copy file from host to container
$ lxc file push <file_name> <path_inside_container>

# Snapshot container
lxc snapshot <container_name> <snapshot_name>

# Restore container snapshot
lxc restore <container_name> <sapshot_name>
```
