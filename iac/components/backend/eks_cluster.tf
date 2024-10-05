resource "aws_eks_cluster" "koronet_cluster" {
  name     = var.eks_cluster_name
  version  = var.eks_cluster_version
  role_arn = aws_iam_role.eks_cluster_role.arn

  vpc_config {
    subnet_ids              = var.eks_subnet_ids
    security_group_ids      = var.eks_security_group_ids
  }

  kubernetes_network_config {
    service_ipv4_cidr = var.eks_service_ipv4_cidr
  }

  access_config {
    authentication_mode = "API_AND_CONFIG_MAP"
  }

  # Configuración de logs en CloudWatch
  enabled_cluster_log_types = ["api", "audit", "authenticator", "controllerManager", "scheduler"]

  depends_on = [
    aws_iam_role_policy_attachment.eks_cluster_policy
  ]
}

resource "aws_cloudwatch_log_group" "eks_cluster_log_group" {
  name              = "/aws/eks/${var.name_prefix}/cluster"
  retention_in_days = 7
}

# IAM Role para EKS Cluster
resource "aws_iam_role" "eks_cluster_role" {
  name = "eks-cluster-role"

  assume_role_policy = <<EOF
  {
    "Version": "2012-10-17",
    "Statement": [
      {
        "Action": "sts:AssumeRole",
        "Principal": {
          "Service": "eks.amazonaws.com"
        },
        "Effect": "Allow",
        "Sid": ""
      }
    ]
  }
  EOF

  managed_policy_arns = [
    "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy",
    "arn:aws:iam::aws:policy/AmazonEKSServicePolicy",
    "arn:aws:iam::aws:policy/AmazonEKSVPCResourceController"
  ]
}

# Asociar políticas IAM con el rol del cluster
resource "aws_iam_role_policy_attachment" "eks_cluster_policy" {
  role       = aws_iam_role.eks_cluster_role.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"
}

# Configuración de acceso y permisos del creador del cluster
resource "aws_eks_access_entry" "cluster_creator_access" {
  cluster_name   = aws_eks_cluster.koronet_cluster.name
  principal_arn       = var.creator_iam_user_arn
  user_name       = "cluster-admin"
  type = "STANDARD"

  depends_on = [
    aws_eks_cluster.koronet_cluster
  ]
}

resource "aws_security_group" "eks_security_group" {
  vpc_id = var.vpc_id

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "${var.eks_cluster_name}-eks-sg"
  }
}

resource "aws_eks_node_group" "koronet_nodes" {
  cluster_name    = aws_eks_cluster.koronet_cluster.name
  node_group_name = "${var.name_prefix}-eks-nodes"
  node_role_arn   = aws_iam_role.eks_node_role.arn
  subnet_ids      = var.eks_subnet_ids
  scaling_config {
    desired_size = 0 
    max_size     = 1
    min_size     = 0
  }

  # Aquí defines el tipo de instancias que usarán los nodos.
  instance_types = ["t2.micro"]

  depends_on = [
    aws_eks_cluster.koronet_cluster
  ]
}

resource "aws_iam_role" "eks_node_role" {
  name = "eks-node-role"

  assume_role_policy = <<EOF
  {
    "Version": "2012-10-17",
    "Statement": [
      {
        "Action": "sts:AssumeRole",
        "Principal": {
          "Service": "ec2.amazonaws.com"
        },
        "Effect": "Allow",
        "Sid": ""
      }
    ]
  }
  EOF

  managed_policy_arns = [
    "arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy",
    "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly",
    "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"
  ]
}

