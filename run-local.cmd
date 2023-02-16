@echo off
setlocal

set kafka_brokers=host.docker.internal:9092
set kafka_messagePack_enabled=true
set server_listenport=8080
build\console.exe 

endlocal
