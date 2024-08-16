module "vpc" {

  source = "github.com/adorigi/localstack/modules/ec2/vpc"

  for_each = var.vpc_configs

  vpc_name   = each.value.name
  cidr_block = each.value.cidr
  region     = each.value.region

}

module "subnets" {

  source = "github.com/adorigi/localstack/modules/ec2/subnet"

  for_each = var.subnet_configs

  subnet_name       = each.value.name
  cidr_block        = each.value.cidr
  availability_zone = each.value.availability_zone
  vpc_id            = module.vpc[each.value.vpc_name].vpc_id

}
