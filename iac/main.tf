 locals {
    account_id = var.account_id 
 }


module "network" {
    source = "./components/network"
    providers = {
      aws.virginia = aws.virginia
    }
    vpc_name = var.vpc_name
    vpc_cidr = "10.0.0.0/24"
    public_subnet_cidr_a = "10.0.0.0/26"
    public_subnet_cidr_b = "10.0.0.64/26"
    private_subnet_cidr = "10.0.0.128/26"
    availability_zones = ["us-east-1a", "us-east-1b"]
}

 module "backend" {
    source ="./components/backend"

    providers = {
      aws.virginia = aws.virginia
    }
    eks_cluster_name = "koronet_interview"
    eks_cluster_version = "1.30"
    eks_service_ipv4_cidr = "172.20.0.0/16"
    name_prefix = "kokornet"
    eks_security_group_ids = []
    eks_subnet_ids = [module.network.public_subnet_id_a,module.network.public_subnet_id_b]
    creator_iam_user_arn = var.eks_user_arn 
    vpc_id = module.network.vpc_id 

 }

