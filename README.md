## Spaced - spaced repetition in the terminal
Like Anki, but in terminal for you to plug it into your scripts.

A simple CLI for spaced repetition - a learning technique involving reviewing learned material at gradually increasing intervals improving retention.

## Requirements
- Go >=1.26

## Installation

### Installation script

Install the app with systemd service using a bash script:
```bash
curl https://raw.githubusercontent.com/d1gitale/spaced/refs/heads/main/scripts/install_update_linux.sh | bash
```

## Usage

### In the terminal

Call `spaced` in your terminal to see a help message. Basic flow goes like so:

1. Add a card titled by a concept you want to remember via `spaced add <card name>`
2. Next day you will receive a notification on your desktop
3. Go to some place to review the concept
4. Call `spaced list --due` to list due cards and copy the ID of the checked card
5. Check the card with `spaced check --id <id> --score <1-5>` with 1 being "completely forgotten" and 5 being "remember perfectly"
6. From now review intervals will grow gradually

## Other features

1. Delete cards with `spaced delete --id <id>`
2. List cards with `spaced list [--format <json/plain> --due]`
3. Rename cards with `spaced rename --id <id>`
