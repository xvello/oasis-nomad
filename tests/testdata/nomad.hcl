data_dir = "/tmp/testing/nomad"

ports {
  http = "44646"
  rpc = "44647"
  serf = "44648"
}

server {
  enabled = true
  bootstrap_expect = 1
  num_schedulers = 1
}
