# BPCopServer

## MySQL
BPCopServer will not boot if the local MySQL instance is not running. Before starting the program make sure to run `mysql.server start`

Access MySQL locally by running `mysql --user=root`. Once the client is open, here are some useful commands:
- `show databases`: lists all available databases 
- `use blood_pressure`: accesses all the 