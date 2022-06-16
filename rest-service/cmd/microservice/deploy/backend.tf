terraform {
  cloud {
    organization = "software-engr-full-stack"
    workspaces {
      name = "stackpath-rest-service"
    }
  }
}
