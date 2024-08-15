variable "vpc_configs" {
  description = "Map of VPC configurations"
  type = map(object({
    name = string
    cidr = string
    region = string
  }))
}