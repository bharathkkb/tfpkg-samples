module "net-project" {
  source = "terraform-google-modules/project-factory/google"

  name              = var.project_name
  billing_account   = "123"
  folder_id         = "789"
  random_project_id = true
  org_id            = "456"
}

module "prod-net" {
  source = "terraform-google-modules/network/google"

  network_name = "bar"
  project_id   = module.net-project.project_id
  subnets = [{
    subnet_ip     = "10.10.10.0/24"
    subnet_name   = "subnet-01"
    subnet_region = "us-west1"
    }, {
    subnet_ip     = "10.10.20.0/24"
    subnet_name   = "subnet-02"
    subnet_region = "us-east1"
  }]
}
