# KUBELOG-GO 

A simple web application for getting logs from the pods

# HOW TO USE IT

It`s simple, just choose the environment you are planning to use and configure the env file

1. outside the cluster

you just have to build the dockerfile.out file by docker-compose.yaml 

2. inside the cluster 

build the image by dockerfile.in 

and configure the service account to access the pods information

# ENV files 

env file is located in envs 

in there you configure every environment variable but currently there are just a few.