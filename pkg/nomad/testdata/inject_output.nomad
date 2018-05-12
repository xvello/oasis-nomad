job "testjob" {
  datacenters = ["dc1"]
  group "nodocker" {
    count = 1
    task "task1" {
      driver = "raw_exec"
      config {
        command = "date"
      }
    }
  }
  group "dockerized" {
    count = 1
    task "knownimage" {
      driver = "docker"
      config {
        image = "library/redis:3.2@sha256:6ff2a3a2ddb62378e778180ead0acaf5b44f6e719e42a1ae8c261dd969a16f2a"
      }
      meta {
        my-key = "my-value"
      }
    }
    task "toupdate" {
      driver = "docker"
      config {
        image = "library/redis:3.2@sha256:6ff2a3a2ddb62378e778180ead0acaf5b44f6e719e42a1ae8c261dd969a16f2a"
      }
    }  
    task "unknownimage" {
      driver = "docker"
      config {
        image = "alpine:3.7"
      }
    }   
    task "noimage" {
      driver = "docker"
      config {
      }
    }   
  }
}