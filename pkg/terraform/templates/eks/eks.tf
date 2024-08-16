module "eksCluster" {

  source = "github.com/adorigi/localstack/modules/eks"

  for_each = var.eks_configs

  cluster_name = each.value.cluster_name
  role_arn     = each.value.role_arn
  subnet_ids   = each.value.subnet_ids

  #   cluster_name = "testCluster"
  #   role_arn     = "arn:aws:iam::000000000000:role/eks-role"
  #   subnet_ids = [
  #     "subnet-dummysubnet"
  #   ]

  depends_on = [
    module.subnets
  ]


}
