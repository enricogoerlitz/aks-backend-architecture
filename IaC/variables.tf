variable "project_resourcegroup_name" {
  description = "The name of the resource group for the project"
  type        = string
  default     = "explore-aks-devops-project-weu-rg"
}

variable "default_tags" {
  description = "The default tags for all resources"
  type        = map(string)
  default = {
    "project" = "explore-aks-devops"
  }
}

variable "westeurope_location" {
  description = "The location for the resources"
  type        = string
  default     = "westeurope"
}

variable "aks_pricing_tier" {
  description = "The pricing tier for the AKS cluster"
  type        = string
}


variable "project_acr_name" {
  description = "The name of the Azure Container Registry for the project"
  type        = string
  default     = "eaksprojectweucr"
}

variable "project_acr_url" {
  description = "The URL of the Azure Container Registry for the project"
  type        = string
  default     = "eaksprojectweucr.azurecr.io"
}
