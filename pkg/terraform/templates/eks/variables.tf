variable "vpc_configs" {
  description = "Map of VPC configurations"
  type = map(object({
    name   = string
    cidr   = string
    region = string
  }))
}

variable "subnet_configs" {
  description = "Map of subnet configurations"
  type = map(object({
    name              = string
    cidr              = string
    availability_zone = string
    vpc_name          = string
  }))
}

variable "iam_configs" {
  description = "Map of iam configurations"
  type = map(object({
    role_name   = string
    role_policy = string
    policy_arn  = string
  }))
}

variable "eks_configs" {
  description = "Map of eks configurations"
  type = map(object({
    cluster_name = string
    role_arn     = string
    subnet_ids   = list(string)
    cluster_tags = map(string)
  }))
}

variable "nodegroup_configs" {
  description = "Map of node group configurations"
  type = map(object({
    cluster_name    = string
    node_group_name = string
    node_role_arn   = string
    subnet_ids      = list(string)
    desired_size    = number
    max_size        = number
    min_size        = number
  }))
}
