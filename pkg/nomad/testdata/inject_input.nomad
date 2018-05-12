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
        image = "redis:3.2"
      }
      meta {
        my-key = "my-value"
      }
    }    
    task "toupdate" {
      driver = "docker"
      config {
        image = "library/redis:3.2@sha256:olddigest"
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