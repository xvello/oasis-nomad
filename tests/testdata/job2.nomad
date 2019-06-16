job "nginx" {
  datacenters = ["dc1"]
  group "nginx" {
    count = 1
    task "nginx" {
      driver = "docker"
      config {
        image = "nginx"
      }
    }
  }
}

