# TMUX RANDOM session name generator.

Tmux random name session generator, inspired by Docker.
This is small plugin that will generate random session name every time you open new Terminal session or it will reconnect your previously disconnected sessions.

## Installation

Script is written in Golang.

iTerm2 Specific.

1. Clone this repository somewhere on you drive.
2. Preferences -> Profiles -> General
   * Under **Command Section:**
    * **Send test at start:** Add tmux.py

This will start tmux.py everytime you open new window.

## Usage

Open New terminal session.

## Inspired

If Docker can generate random names every time it creates container so can tmux when creating new session.

## Credits

Inspired by Docker.
