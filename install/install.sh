echo "...... [ Welcome To teleterm installer]"
echo "...... [ Check Golang ]"
if [ -d /usr/local/go ];
then
  echo "Golang Installed"
  echo "...... [ Build ]"
  go build -o ../teleterm ../main.go
  cp ../teleterm /usr/local/bin/
  echo "...... [ Finish :) ]"
else
  echo "Install Golang First"
fi