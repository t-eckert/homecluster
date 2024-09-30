# Setting up Talos

This document will guide you through the process of installing Talos Linux
on a Raspberry Pi and running Homecluster.

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

Use the [Raspberry Pi Imager](https://www.raspberrypi.com/software/) to write an EEPROM update image to an SD card. Select Misc utility images under the "Operating System" tab. Insert the SD card into the Raspberry Pi, power it on, and wait at least 10 seconds. If the update was successful, the green LED light will blink rapidly.

Power off the Raspberry Pi and remove the SD card.

## Installing Talos Linux

Download the Talos Linux image.

```shell
curl -LO https://factory.talos.dev/image/ee21ef4a5ef808a9b7484cc0dda0f25075021691c8c09a276591eedb638ea1f9/v1.8.0/metal-arm64.raw.xz
xz -d metal-arm64.raw.xz
```

Insert your SD card into your computer. On MacOS, use `diskutil` to find the device
name of the SD card.

```shell
distutil list
```

Image the SD card.

```shell
sudo dd if=metal-rpi_generic-arm64.raw of=/dev/<DISK-NAME> conv=fsync bs=4M
```

Insert the SD card into the Raspberry Pi and power it on.

## Configure the Cluster

Set the name of your cluster and the IP address of the Raspberry Pi.

```shell
export CLUSTER_NAME=homecluster
export RASPBERRY_PI_IP=<IP>
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

- [Raspberry Pi Installation Docs](https://www.talos.dev/v1.8/talos-guides/install/single-board-computers/rpi_generic/)
- [Installing Talos Linux on Raspberry Pi 4](https://kubito.dev/posts/talos-linux-raspberry-pi/)
- [Talos Getting Started](https://www.talos.dev/v1.8/introduction/getting-started/)
