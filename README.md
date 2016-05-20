dockerdemogo
============

This repository contains a pair of applications, a server and a client, to 
demonstrate Docker.

To start the application you need Docker Engine and Docker Compose configured.
If you are using Windows or Mac you may want to try Docker Toolbox from
https://www.docker.com/products/docker-toolbox as it includes all the
components you need. 

You should then be able to run;

    $ docker-compose build
    $ docker-compose up 

to bring up one client and one server. 

You can scale the numbers of clients and servers up and down with

    $ docker-compose scale server=X client=Y

and choose appropriate values for X and Y.

You can see the containers running using;

    $ docker-compose ps

and;

    $ docker ps

You can stop the clients and servers with

    $ docker-compose down
