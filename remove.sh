
echo Stopped loanapp container
docker stop loanapp

echo Removed loanapp container
docker rm loanapp

docker container prune --filter until=15m

docker system prune -a