#!/bin/bash
set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

APP_NAME="spaced"
REPO_URL="https://github.com/d1gitale/spaced"
BRANCH="main"

echo -e "${GREEN}🚀 Starting installation for ${APP_NAME}...${NC}"

if ! command -v go &> /dev/null; then
    echo -e "${RED}❌ Go is not installed.${NC}"
    exit 1
fi

if ! command -v git &> /dev/null; then
    echo -e "${RED}❌ Git is not installed.${NC}"
    exit 1
fi

TEMP_DIR=$(mktemp -d)
trap "rm -rf $TEMP_DIR" EXIT

git clone --depth 1 --branch $BRANCH $REPO_URL "$TEMP_DIR"
cd "$TEMP_DIR/cmd/app"

CGO_ENABLED=0 go build -ldflags="-s -w" -o "$APP_NAME" .

if [ ! -f "$APP_NAME" ]; then
    echo -e "${RED}❌ Build failed. Binary not found.${NC}"
    exit 1
fi

sudo cp "$APP_NAME" /usr/local/bin/$APP_NAME
sudo chmod +x /usr/local/bin/$APP_NAME

echo -e "${YELLOW}⚙️ Configuring systemd user service for notifications...${NC}"

SERVICE_DIR="$HOME/.config/systemd/user"
mkdir -p "$SERVICE_DIR"

cat > "$SERVICE_DIR/${APP_NAME}-boot.service" <<EOF
[Unit]
Description=Spaced Repetition System Notification Check
After=graphical-session.target
Wants=graphical-session.target

[Service]
Type=oneshot
ExecStart=/usr/local/bin/$APP_NAME notify --check
Environment=DBUS_SESSION_BUS_ADDRESS=unix:path=/run/user/%U/bus

[Install]
WantedBy=default.target
EOF

systemctl --user daemon-reload
systemctl --user enable "${APP_NAME}-boot.service"

echo -e "${GREEN}✅ Installation complete!${NC}"
echo -e "Run '${GREEN}${APP_NAME}${NC}' from anywhere."
echo -e "Notifications will start working after next login/reboot."
