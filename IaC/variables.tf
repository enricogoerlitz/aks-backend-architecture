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
