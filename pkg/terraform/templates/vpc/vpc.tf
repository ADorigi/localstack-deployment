module "vpc" {

    source = "github.com/adorigi/localstack/modules/ec2/vpc"

    for_each = var.vpc_configs

    vpc_name = each.value.name
    cidr_block = each.value.cidr
    region = each.value.region
  
}