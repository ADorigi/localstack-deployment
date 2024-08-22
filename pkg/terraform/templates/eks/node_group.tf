module "nodegroup" {
  source = "github.com/adorigi/localstack/modules/eks/nodegroup"

  for_each = var.nodegroup_configs

  cluster_name    = format("%s-%s", each.value.cluster_name, var.region)
  node_group_name = format("%s-%s", each.value.node_group_name, var.region)
  node_role_arn   = each.value.node_role_arn
  subnet_ids      = each.value.subnet_ids
  desired_size    = each.value.desired_size
  max_size        = each.value.max_size
  min_size        = each.value.min_size


  depends_on = [module.eksCluster]
}