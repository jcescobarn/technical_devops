provider "aws" {
    alias = "virginia"
    region = "us-east-1"

    default_tags {
      tags = ["koronet"] 
    }

    
}

terraform {
        required_providers {
            aws = "5.64.0"
        }
    }