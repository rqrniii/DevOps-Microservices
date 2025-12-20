variable "ami_id" {
  description = "Ubuntu 22.04 AMI"
  type        = string
  default     = "ami-0147469985ff98928"
}

variable "key_name" {
  description = "SSH key name for EC2"
  type        = string
  default     = "devops-microservices-key"
}