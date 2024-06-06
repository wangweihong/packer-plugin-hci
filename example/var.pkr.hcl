locals {
  communicator = var.communicator == null ? (
  var.is_windows ? "winrm" : "ssh"
  ) : var.communicator

  http_directory   = var.http_directory == null ? "${path.root}/http" : var.http_directory

  boot_command = ["<enter><wait>",
    "<f6><esc>",
    "<bs><bs><bs><bs><bs><bs><bs><bs><bs><bs>",
    "<bs><bs><bs><bs><bs><bs><bs><bs><bs><bs>",
    "<bs><bs><bs><bs><bs><bs><bs><bs><bs><bs>",
    "<bs><bs><bs><bs><bs><bs><bs><bs><bs><bs>",
    "<bs><bs><bs><bs><bs><bs><bs><bs><bs><bs>",
    "<bs><bs><bs><bs><bs><bs><bs><bs><bs><bs>",
    "<bs><bs><bs><bs><bs><bs><bs><bs><bs><bs>",
    "<bs><bs><bs><bs><bs><bs><bs><bs><bs><bs>",
    "<bs><bs><bs>",
    "/install/vmlinuz ",
    "initrd=/install/initrd.gz ",
    #   "net.ifnames=0 ",
    "auto-install/enable=true ",
    "debconf/priority=critical ",
    "preseed/url=http://{{ .HTTPIP }}:{{ .HTTPPort }}/ubuntu/16.04/preseed.cfg ",
    "<enter>"]

  boot_wait = "5s"

  inline_custom_image_scripts =[
    "env",
  ]

}


variable "shutdown_command" {
  type    = string
  default = null
}

variable "shutdown_timeout" {
  type    = string
  default = "15m"
}

variable "ssh_username" {
  type        = string
  default     = "vagrant"
  description = "通过ssh连接系统的账号密码。如果是通过iso安装, 必须和预设账号密码保持一致"
}

variable "ssh_password" {
  type    = string
  default = "vagrant"
}

variable "ssh_port" {
  type    = number
  default = 22
}

variable "ssh_timeout" {
  type    = string
  default = "30m"
}

variable "endpoint" {
  type    = string
  default = "http://10.30.100.25:9990"
}

variable "tenant" {
  type    = string
  default = "system"
}

variable "user" {
  type    = string
  default = "wwhvw"
}

variable "password" {
  type    = string
  default = "wwhvw"
}

variable "target_tenant" {
  type    = string
  default = "container"
}

variable "specification" {
  type    = string
  default = "高性能虚拟机"
}

variable "repository_name" {
  type = string
  default = "container"
}

variable "iso_name" {
  type = string
  //default = "ubuntu-16.04.6-server-amd64.iso"
  default = "ubuntu-16.04.4-server-amd64.iso"
}

variable "source_image" {
  type = string
  // default = "ubuntu-base-04-27.img"
  default = "ubuntu-base-16.04-20240605.raw"
}

variable "image_name" {
  type    = string
  default = "test-ubuntu-16-04-packer"
  description = "构建的镜像的名称"
}

variable "export_repository" {
  type    = string
  default = "container"
  description = "镜像导出的仓库"
}

variable "disk_size" {
  type  = number
  default = 107374182400
}

variable "communicator" {
  type = string
  default = null
}

variable "is_windows" {
  type= bool
  default =false
}

variable "http_directory" {
  type= string
  default = null
}

variable "instance_cidr" {
  type= string
  default = "10.30.100.222/32"
}