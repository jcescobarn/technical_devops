terraform {
    backend "s3" {
       bucket = "koronet-interview" 
        key    = "koronet_iac.tfstate"
        region = "us-east-1" 
    }
}