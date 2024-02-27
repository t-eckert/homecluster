# Homecluster Setup

Homecluster runs on Kubernetes running on a Raspberry Pi using Talos Linux.
This document will guide you through the process of installing Talos Linux
and running Homecluster.

## Prerequisites

- `talosctl`
- an SD card

## `talosctl`

Install `talosctl`.

```shell
curl -sL https://talos.dev/install | sh
```

This tool is needed to interact with the Talos API. Talos Linux does not expose
SSH access.

## Update the EEPROM on the Raspberry Pi

Use the [Raspberry Pi Imager](https://www.raspberrypi.com/software/)
to write an EEPROM update image to an SD card. Select Misc utility images under
the "Operating System" tab. Insert the SD card into the Raspberr Pi, power it on,
and wait at least 10 seconds. If the update was successful, the green LED light
will blink rapidly.

Power off the Raspberry Pi and remove the SD card.

## Set the IP Address of the Pi

It would be a good idea if you set up the static IPs of each of the Raspberry
Pi boards via their MAC address on your router before beginning. This is so you
don’t have to look for the dynamically assigned IP when you boot them up. If you
don’t know how to find out the MAC address, check this
[guide](https://kubito.dev/posts/getting-pi-mac-address/)

## Installing Talos Linux

Download the Talos Linux image.

```shell
curl -LO https://github.com/siderolabs/talos/releases/download/v1.6.4/metal-rpi_generic-arm64.raw.xz
xz -d metal-rpi_generic-arm64.raw.xz
```

Image the SD card.

```shell
sudo dd if=metal-rpi_generic-arm64.raw of=/dev/mmcblk0 conv=fsync bs=4M
```

Bootstrap the cluster.

```shell
talosctl apply-config --insecure --mode=interactive --nodes <node IP or DNS name>
```

Retrieve the Kubeconfig.

```shell
talosctl kubeconfig
```

## Reference

- [Raspberry Pi Installation Docs](https://www.talos.dev/v1.6/talos-guides/install/single-board-computers/rpi_generic/)
- [Installing Talos Linux on Raspberry Pi 4](https://kubito.dev/posts/talos-linux-raspberry-pi/)
