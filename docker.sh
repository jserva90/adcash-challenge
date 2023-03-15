clear
docker image build -t loanapp .

echo
echo 'Server running at https://localhost:8080/'
echo
echo Use \"sh remove.sh\" to stop and remove the container.
echo

docker container run -p 8080:8080 --detach --name loanapp loanapp

docker ps
