module "eks_cluster_role" {

    source = "github.com/adorigi/localstack/modules/iam"
  
    for_each = var.iam_configs

    role_name = each.value.role_name
    assume_role_policy = each.value.role_policy
    policy_arn = each.value.policy_arn
  
}