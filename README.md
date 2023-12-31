sudo useradd rasp-web
sudo install -d -o rasp-web -g rasp-web /var/lib/rasp-web
make it executable?

[Unit]
Description=rasp-web https go web server
After=navidrome.service
After=remote-fs.target
After=network.target

[Install]
WantedBy=multi-user.target

[Service]
RestartSec=2s
Type=exec
User=rasp-web
Group=rasp-web
ExecStart=/opt/rasp-web/rasp-web
Restart=on-failure
KillMode=control-group
TimeoutStopSec=20
WorkingDirectory=/var/lib/rasp-web
