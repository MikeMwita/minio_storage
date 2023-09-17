data_dir  = "/home/mike/Desktop/comm/dockerconfigs"

bind_addr = "0.0.0.0" # the default

advertise {
  # Defaults to the first private IP address.
  http = "127.0.0.1:6000"
  rpc  = "127.0.0.1:9090"
  serf = "127.0.0.1:5648" # non-default ports may be specified

}

server {
  enabled          = true
  bootstrap_expect = 5
}

client {
  enabled       = true
}

plugin "raw_exec" {
  config {
    enabled = true
  }
}

plugin "dns" {
  config {
    enabled = true
  }
}
consul {
  address = "127.0.0.1:8500"
}

