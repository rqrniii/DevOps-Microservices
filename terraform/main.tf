terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
  required_version = ">= 1.0.0"
}

provider "aws" {
  region = "me-south-1" 
}

module "vpc" {
  source             = "./modules/vpc"
  cidr_block         = "10.0.0.0/16"
  public_subnet_cidr = "10.0.1.0/24"
  az                 = "me-south-1a"
  name               = "k8s-vpc"
}
