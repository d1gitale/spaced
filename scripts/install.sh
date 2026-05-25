#!/bin/bash
set -e

APP_NAME="spaced"
BINARY_NAME="spaced"
REPO_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

echo "🔨 Building $APP_NAME..."
CGO_ENABLED=0 go build -o "$BINARY_NAME" .

echo "📦 Installing to /usr/local/bin..."
sudo cp "$BINARY_NAME" /usr/local/bin/$BINARY_NAME
sudo chmod +x /usr/local/bin/$BINARY_NAME

echo "⚙️ Setting up systemd user service for notifications..."

mkdir -p ~/.config/systemd/user/

cat > ~/.config/systemd/user/$APP_NAME-boot.service <<EOF
[Unit]
Description=Send SRS Notifications on Boot
After=graphical-session.target
Wants=graphical-session.target

[Service]
Type=oneshot
ExecStart=/usr/local/bin/$BINARY_NAME notify --check
Environment=DBUS_SESSION_BUS_ADDRESS=unix:path=/run/user/%U/bus

[Install]
WantedBy=default.target
EOF

systemctl --user daemon-reload
systemctl --user enable $APP_NAME-boot.service

echo "✅ Installed successfully!"
echo "Run '$BINARY_NAME' from anywhere."
echo "Notifications will trigger on next login."
