output "vpc_id" {
  description = "El ID de la VPC"
  value       = aws_vpc.koronet.id 
}

output "public_subnet_id_a" {
  description = "El ID de la subred pública"
  value       = aws_subnet.public_a.id
}

output "public_subnet_id_b" {
  description = "El ID de la subred pública"
  value       = aws_subnet.public_b.id
}

output "private_subnet_id" {
  description = "El ID de la subred privada"
  value       = aws_subnet.private.id
}

output "internet_gateway_id" {
  description = "El ID del Internet Gateway"
  value       = aws_internet_gateway.igw.id
}

output "nat_gateway_id" {
  description = "El ID del NAT Gateway"
  value       = aws_nat_gateway.nat.id
}
