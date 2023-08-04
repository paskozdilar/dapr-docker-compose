# dapr-docker-compose

Example project using Dapr in docker-compose.

## Start

**NOTE:** If you have initialized Dapr in your host machine with `dapr init`,
please uninitialize it and stop the Dapr containers:

    dapr uninstall
    docker rm --force dapr_redis dapr_zipkin

To build client and server images, run:

    docker-compose build

To start everything in the background, run:

    docker-compose up -d

To see client logs, run:

    docker-compose logs client

You should be able to see executed gRPC method calls.
