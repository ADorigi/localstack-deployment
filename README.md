# localstack-deployment
Deploying localstack resources with terraform

# Plan


## configuration 

- one_eks.tfvars.json - this config creates:

    - vpc - 1
    - subnets - 4
    - iam role - 1
    - iam policy attachment - 1 
    - eks cluster - 1


<!-- 
    done - terraform for eks cluster 

     5 eks clusters, 5 node groups, 20 nodes
done 1 eks clusters
done 2 eks clusters
done 
    seperate iam role and policy attachment
    add slog to this repo
    add logs for verbosity in atleast controller, config, jobs and terraform modules
 -->