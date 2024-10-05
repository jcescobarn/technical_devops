variable "eks_cluster_name" {
    type = string
    description = "EKS cluster name"
}

variable "eks_cluster_version" {
    type = string
    description = "EKS cluster version"
    default = ""
}

variable "eks_service_ipv4_cidr" {
    type = string
}

variable "name_prefix" {
    type = string
}

variable "eks_security_group_ids" {
    type = list(string)
}

variable "eks_subnet_ids" {
  type = list(string)
}

variable "creator_iam_user_arn" {
    type = string
}

variable "vpc_id" {
    type = string
}