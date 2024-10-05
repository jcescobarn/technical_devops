output "eks_id"{
    value = aws_eks_cluster.koronet_cluster.id
    description = "EKS cluster name"
}

output "eks_arn" {
    value = aws_eks_cluster.koronet_cluster
    description = "EKS cluster ARN"
}

output "eks_network_config" {
    value = aws_eks_cluster.koronet_cluster.kubernetes_network_config
    description = "EKS cluster network configuration"
}