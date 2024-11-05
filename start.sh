#!/bin/sh

DATABASE_URL="postgresql://postgres:password@localhost:5432/postgres?sslmode=disable" 
DOMAIN="localhost" 
JWT_KEY="test" 

export DATABASE_URL
export DOMAIN
export JWT_KEY

go run .