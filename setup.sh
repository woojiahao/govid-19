#!/usr/bin/env bash

read -rp "Database user >>> " user
read -rp "Database password >>> " password
read -rp "Database name >>> " name
read -rp "Host >>> " host

{
 echo "POSTGRES_DB=$name"
 echo "POSTGRES_USER=$user"
 echo "POSTGRES_PASSWORD=$password"
 echo "HOST=$host"
} >> .env
