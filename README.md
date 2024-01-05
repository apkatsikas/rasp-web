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

[Unit]
Description=Navidrome Music Server and Streamer compatible with Subsonic/Airsonic
After=remote-fs.target network.target media-ssd.mount
AssertPathExists=/var/lib/navidrome

[Install]
WantedBy=multi-user.target

[Service]
User=navidrome
Group=navidrome
Type=simple
ExecStart=/opt/navidrome/navidrome --configfile "/var/lib/navidrome/navidrome.toml"
WorkingDirectory=/var/lib/navidrome
TimeoutStopSec=20
KillMode=process
Restart=on-failure

# See https://www.freedesktop.org/software/systemd/man/systemd.exec.html
DevicePolicy=closed
NoNewPrivileges=yes
PrivateTmp=yes
PrivateUsers=yes
ProtectControlGroups=yes
ProtectKernelModules=yes
ProtectKernelTunables=yes
RestrictAddressFamilies=AF_UNIX AF_INET AF_INET6
RestrictNamespaces=yes
RestrictRealtime=yes
SystemCallFilter=~@clock @debug @module @mount @obsolete @reboot @setuid @swap
ReadWritePaths=/var/lib/navidrome

# You can uncomment the following line if you're not using the jukebox This
# will prevent navidrome from accessing any real (physical) devices
#PrivateDevices=yes

# You can change the following line to `strict` instead of `full` if you don't
# want navidrome to be able to write anything on your filesystem outside of
# /var/lib/navidrome.
ProtectSystem=full

# You can uncomment the following line if you don't have any media in /home/*.
# This will prevent navidrome from ever reading/writing anything there.
#ProtectHome=true

# You can customize some Navidrome config options by setting environment variables here. Ex:
#Environment=ND_BASEURL="/navidrome"

[Unit]
Description=ddns-route53
Documentation=https://crazymax.dev/ddns-route53/
After=syslog.target
After=network.target

[Service]
RestartSec=2s
Type=simple
User=ddns-route53
Group=ddns-route53
ExecStart=/usr/local/bin/ddns-route53 --config /etc/ddns-route53/ddns-route53.yml
Restart=always
#Environment=TZ=Europe/Paris
#Environment=AWS_ACCESS_KEY_ID=********
#Environment=AWS_SECRET_ACCESS_KEY=********
Environment=SCHEDULE="*/30 * * * *"

[Install]
WantedBy=multi-user.target
