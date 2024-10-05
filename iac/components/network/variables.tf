variable "vpc_name" {
  type = string
  
}

variable "vpc_cidr" {
  default = "10.0.0.0/16"
  description = "CIDR block para la VPC"
}

variable "public_subnet_cidr_a" {
  type        = string
  description = "CIDR block for public subnet A"
}

variable "public_subnet_cidr_b" {
  type        = string
  description = "CIDR block for public subnet B"
}

variable "private_subnet_cidr" {
  default = "10.0.2.0/24"
  description = "CIDR block para la subred privada"
}

variable "availability_zones" {
  type = list(string)
}
