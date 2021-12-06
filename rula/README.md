# Rula

A user service.

# How to run :
1.  prepare redis server
2.  adjust config (.env) values for database and redis
3.  run migrate :
    go run main.go migrate up or ./rula migrate up
4.  run serve :
    go run main.go server or ./rula server
5.  open hostname/swagger/index.html on browser to select apis

# How to generate docs and api client
1. prepare annotations above handlers
2. run go generate ./cmd_tools_imports.go

# How to run blackbox (bdd style tests)
1. run go generate ./blackbox