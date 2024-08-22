module "eksCluster" {

  source = "github.com/adorigi/localstack/modules/eks"

  for_each = var.eks_configs

  cluster_name = format("%s-%s", each.value.cluster_name, var.region)
  role_arn     = each.value.role_arn
  subnet_ids   = each.value.subnet_ids

  depends_on = [
    module.subnets
  ]
}

