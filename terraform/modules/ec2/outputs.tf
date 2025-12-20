output "master_public_ips" {
  value = aws_instance.master[*].public_ip
}

output "worker_public_ips" {
  value = aws_instance.worker[*].public_ip
}

output "security_group_id" {
  value = aws_security_group.k8s.id
}