module "sample" {
  source = "terraform-google-modules/cloud-storage/google//modules/simple_bucket"

  name       = "sample-bucket"
  project_id = "sample-bucket-project"
  location   = "us-east1"
}
