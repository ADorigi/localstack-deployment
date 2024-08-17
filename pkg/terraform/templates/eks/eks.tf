module "eksCluster" {

  source = "github.com/adorigi/localstack/modules/eks"

  for_each = var.eks_configs

  cluster_name = each.value.cluster_name
  role_arn     = each.value.role_arn
  subnet_ids   = each.value.subnet_ids

  cluster_tags = each.value.cluster_tags

  depends_on = [
    module.subnets
  ]
}
