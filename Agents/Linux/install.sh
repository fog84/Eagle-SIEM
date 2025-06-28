mkdir /etc/eagle_agent
mkdir /etc/eagle_agent/conf
cp -r conf/* /etc/eagle_agent/conf/
mkdir /etc/eagle_agent/save_last_lines
go build -o eagle_agent main.go
cp eagle_agent /usr/local/bin
cp eagle_agent.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl start eagle_agent.service
sudo systemctl status eagle_agent.service