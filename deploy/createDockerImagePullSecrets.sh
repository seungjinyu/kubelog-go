kubectl create secret docker-registry registry-secret --namespace=logger \
        --docker-server=${REGISTRY} \
        --docker-username=${USERNAME} \
        --docker-password=${PASSWORD} \
        --docker-email=${USEREMAIL