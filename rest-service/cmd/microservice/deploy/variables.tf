variable "aws_region" {
  type    = string
  default = "us-west-1"
}

variable "name" {
  type = string
  description = "Base name of deployment."
}

variable "path" {
  type = string
  description = "The main path."
}

variable "source_file" {
  type = string
  description = "Path to the handler bin."
}

variable "output_path" {
  type = string
  description = "Path to the output zip file."
}
