

source "hci-iso" "example" {
  endpoint    = var.endpoint
  tenant     = var.tenant
  user       = var.user
  password   = var.password
  image_name = var.image_name
  boot_command = local.boot_command
  boot_wait     = local.boot_wait
  communicator     = local.communicator
  ssh_username       = var.ssh_username
  ssh_password       = var.ssh_password
  ssh_port         = var.ssh_port
  ssh_timeout      = var.ssh_timeout

  http_directory   = local.http_directory
  target_tenant = var.target_tenant
  specification = var.specification
  repository_name  =  var.repository_name
  iso_name = var.iso_name
  disk_size = var.disk_size
  instance_cidr = var.instance_cidr
}

source "hci-vmx" "example" {
  endpoint    = var.endpoint
  tenant     = var.tenant
  user       = var.user
  password   = var.password
  image_name = var.image_name
  boot_command = local.boot_command
  boot_wait     = local.boot_wait
  communicator     = local.communicator
  ssh_username       = var.ssh_username
  ssh_password       = var.ssh_password
  ssh_port         = var.ssh_port
  ssh_timeout      = var.ssh_timeout

  http_directory   = local.http_directory
  target_tenant = var.target_tenant
  specification = var.specification
  repository_name  =  var.repository_name
  source_image = var.source_image
  disk_size = var.disk_size
  instance_cidr = var.instance_cidr
}

build {
  sources = [
    //"source.hci-iso.example",
    "source.hci-vmx.example",

  ]

  provisioner "shell" {
    inline = local.inline_custom_image_scripts
    execute_command= "echo 'vagrant' | {{ .Vars }} sudo -S -E bash -eux '{{ .Path }}'"
  }

  #post-processor "manifest" {
  # strip_path = true
  # output     = "packer-manifest.json"
  #}
}
