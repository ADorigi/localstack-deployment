module "eksCluster" {

  source       = "github.com/adorigi/localstack/modules/eks"
  cluster_name = "testCluster"
  role_arn     = "arn:aws:iam::000000000000:role/eks-role"
  subnet_ids = [
    "subnet-dummysubnet"
  ]

  depends_on = [
    module.subnets
  ]


}
