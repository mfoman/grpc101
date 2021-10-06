# Setup of repository

1. make ``go.mod`` file with:

    ``$ go mod init``

2. make a ``.proto`` file in a subfolder, for example ``time/time.proto``.
3. run command:

    ``$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative time/time.proto``

    which should create the two ``pb.go`` files.
4. run command:

    ``$ go mod tidy``

    to install dependencies and create the ``go.sum`` file.
5. create a ``client\client.go`` file with the ``client_template.txt`` as template.
6. create a ``server\server.go`` file with the ``server_template.txt`` as template.
7. switch out the "myPackage" with your actual package.
8. switch our the method names with actual method names.
9. add more methods to the ``client.go`` file, so that there's a method for each request in the ``.proto`` file.
10. when everything is compilable, open a terminal, change directory to the ``server`` folder, and run the command:

    ``$ go run .``

    this will start your server.
11. open a new terminal, change directory to the ``client`` folder and run the command:

    ``$ go run.``

    this will run the requests listed in the ``main`` method.
