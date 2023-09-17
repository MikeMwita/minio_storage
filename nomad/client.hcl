client {
  enabled = true

  network_interface = "{{ GetPrivateInterfaces | include \"network\" \"10.0.0.0/8\" | attr \"name\" }}"

  # Add support for cni
  cni_path = "/opt/cni/bin"

  # Add support for cni config
  #cni_config_dir = "/etc/cni/net.d"

  # Managing volume mounts

  host_volume "nocodb_data" {
    path      = "/opt/data/nocodb"
    read_only = false
  }

  host_volume "nocodb_db" {
    path      = "/opt/data/nocodb/db"
    read_only = false
  }
}
