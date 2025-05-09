resource "azurerm_kubernetes_cluster" "main" {
  name                = "explore-aks-devops-${terraform.workspace}-weu-aks"
  location            = azurerm_resource_group.main.location
  resource_group_name = azurerm_resource_group.main.name
  dns_prefix          = "eaksdevops${terraform.workspace}weuaks"
  sku_tier            = var.aks_pricing_tier
  node_resource_group = "explore-aks-devops-aks-management-${terraform.workspace}-weu-rg"
  kubernetes_version  = "1.31.7"

  default_node_pool {
    name       = "system"
    node_count = 1
    vm_size    = "Standard_D2ps_v6"
  }

  identity {
    type = "SystemAssigned"
  }

  tags = merge(var.default_tags, {
    "env" = terraform.workspace
  })
}

resource "azurerm_role_assignment" "aks_acr_pull" {
  scope                = data.azurerm_container_registry.cr.id
  role_definition_name = "AcrPull"
  principal_id         = azurerm_kubernetes_cluster.main.identity[0].principal_id
}

resource "azurerm_kubernetes_cluster_node_pool" "application" {
  name                  = "application"
  kubernetes_cluster_id = azurerm_kubernetes_cluster.main.id
  vm_size               = "Standard_D2ps_v6"

  auto_scaling_enabled = true
  min_count            = 1
  max_count            = 1

  tags = merge(var.default_tags, {
    "env" = terraform.workspace
  })
}

resource "azurerm_kubernetes_cluster_node_pool" "database" {
  name                  = "database"
  kubernetes_cluster_id = azurerm_kubernetes_cluster.main.id
  vm_size               = "Standard_D2ps_v6"

  auto_scaling_enabled = true
  min_count            = 1
  max_count            = 1

  tags = merge(var.default_tags, {
    "env" = terraform.workspace
  })
}
