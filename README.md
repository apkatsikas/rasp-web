`ls -l /dev/disk/by-uuid/` - get uuid

`df -Th` - get filesystem type

ext4

be sure to enable service not just create entry and start

UUID=2d52db9d-3bb2-4b8e-b5de-9764cce7594a /media/ssd/ ext4 defaults,nofail,x-system.automount 0 2

mount drive - `sudo mount /dev/sda /media/ssd/`

make sure everything can boot and work

sudo useradd rasp-web
sudo install -d -o rasp-web -g rasp-web /var/lib/rasp-web
PORTS - sudo setcap CAP_NET_BIND_SERVICE=+eip /opt/rasp-web/rasp-web
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
