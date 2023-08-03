# dapr-docker-compose

Example project using Dapr in docker-compose.

## Start

To build client and server images, run:

    docker-compose build

To start everything in the background, run:

    docker-compose up -d

To see client logs, run:

    docker-compose logs client

You should be able to see executed gRPC method calls.
