# wolooth

Trigger wake on LAN signals from specific bluetooth devices activity!

## Initial purpose

I'm using this tool to wake my [homemade "Steam Deck"]((https://bazzite.gg/)) with my Xbox bluetooth controller. 

Useful when your system only supports wake on lan and you have a spare raspberry on hand 🙂! (In my example the raspberry is connected via ethernet to my gaming machine)


## What does this do

Wolooth listens for any kind of activity emitted by specific bluetooth devices to trigger a Wake on Lan magic packet to the specified target.

## Requirements

- `go >= 1.24`
- [etherwake](https://www.mkssoftware.com/docs/man1/etherwake.1.asp) installed (to send WOL signals)

## Usage

> ℹ️ Run the script as root because etherwake needs elevated permissions

Run `WOL_TARGET=... BT_DEVICES=...,... go run wolooth.go` on the device that will send WOL signals.

Where `WOL_TARGET` is the interface's MAC address of the device you want to wake and `BT_DEVICES` a list of MAC addresses of bluetooth devices that will be monitored for activity.

You can also use an `.env` file if you wish.