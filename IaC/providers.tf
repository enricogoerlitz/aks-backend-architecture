terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "4.26.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.7.1"
    }
  }

#   backend "azurerm" {
#     resource_group_name  = "explore-aks-devops-project-euwest-rg"
#     storage_account_name = "exploreadbeuwestsa"
#     container_name       = "exploreadb-terraform-backend"
#     key                  = "terraform.tfstate"
#   }
}

provider "azurerm" {
  features {}
  subscription_id = "012e925b-f538-41ef-8d23-b0c85e7dbe7b"
}

provider "random" {}
