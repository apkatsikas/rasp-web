First temporarily port forward
5030 TCP to the pi
50300 BOTH to the pi - perma?

temporarily firewall?

sudo firewall-cmd --add-port=5030/tcp --permanent
sudo firewall-cmd --add-port=50300/tcp --permanent
sudo firewall-cmd --add-port=50300/udp --permanent



wget https://github.com/slskd/slskd/releases/download/0.19.5/slskd-0.19.5-linux-arm64.zip

sudo useradd -m slskd
sudo install -d -o slskd -g slskd /var/lib/slskd
sudo install -d -o slskd -g slskd /opt/slskd
sudo chown -R slskd:slskd /media/ssd/music

unzip in windows, FTP over
sudo chmod +x /opt/slskd/slskd

SLSKD_SLSK_USERNAME=white_mage SLSKD_SLSK_PASSWORD=uld68Po6yOkCFEOlQ9WY SLSKD_SHARED_DIR=/media/ssd/music/king_cop SLSKD_DOWNLOADS_DIR=/media/ssd/music APP_DIR=/var/lib/slskd  ./slskd

sudo vim /etc/systemd/system/slskd.service

[Unit]
Description=slskd music collection
After=remote-fs.target network.target media-ssd.mount navidrome.service
AssertPathExists=/var/lib/slskd

[Install]
WantedBy=multi-user.target

[Service]
User=slskd
Group=slskd
Type=simple
ExecStart=/opt/slskd/slskd
WorkingDirectory=/var/lib/slskd
TimeoutStopSec=20
KillMode=control-group
Restart=on-failure
Environment=SLSKD_SLSK_USERNAME=white_mage SLSKD_SLSK_PASSWORD=uld68Po6yOkCFEOlQ9WY SLSKD_SHARED_DIR=/media/ssd/music/king_cop SLSKD_DOWNLOADS_DIR=/media/ssd/music APP_DIR=/var/lib/slskd

sudo systemctl daemon-reload
sudo systemctl start

