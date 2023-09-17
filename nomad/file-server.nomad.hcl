job "collections-ui" {
  // Specifies the datacenter where this job should be run
  // This can be omitted and it will default to ["*"]
  datacenters = ["*"]

  // A group defines a series of tasks that should be co-located
  // on the same client (host). All tasks within a group will be
  // placed on the same host.
  group "servers" {

    // Specifies the number of instances of this group that should be running.
    // Use this to scale or parallelize your job.
    // This can be omitted and it will default to 1.
    count = 1

    network {
      port "grpc" {
        static = 9090
      }
      port "rest" {
        static = 6000
      }
    }

    service {
      provider = "nomad"
      port     = "rest"
    }

    // Tasks are individual units of work that are run by Nomad.
    task "web" {
      // This particular task starts a simple web server within a Docker container
      driver = "docker"

      config {
        image   = "mwitamike/minio:latest"
        ports   = ["grpc","rest"]
      }

      // Specify the maximum resources required to run the task
      resources {
        cpu    = 300
        memory = 300
      }
    }
  }
}
