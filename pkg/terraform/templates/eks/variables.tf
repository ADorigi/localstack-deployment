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
