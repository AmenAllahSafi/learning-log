# Proxmox Homelab

Documentation for my homelab, running on Proxmox VE.

## Overview

- **Purpose:** The purpose of this lab is straightforward as a solution to my need of constant VM for various reasons such as testing, hosting, or purely personal needs (such as my FTPS server).
- **Proxmox VE version:** 9.2
- **Last updated:** May 21, 2026 
- **Location:** My room.

## Hardware

| Component | Details |
|---|---|
| Host | HPProDesk 600 Series G5|
| CPU | i5-9400F |
| Memory | DDR4 12 GB |
| Storage | 256 NVMe SSD |
| Networking | LAN via UTP cable, one NIC. |

## Network

| Network/VLAN | Subnet | Purpose |
|---|---|---|
| [LAN] | `[192.168.1.0/24` | [Clients and management] |


## Virtual Machines and Containers

| ID | Name | Type | OS | vCPU | RAM | Storage | Purpose |
|---:|---|---|---|---:|---:|---:|---|
| [100] | kalimachine | VM | kali linux 2026.3 | 2 | 2 GB | 40 GB | Penetration testing  |
 | [101] | testingmachine | VM | Debian 13.0 | 1 | 2 GB | 40 GB | software testing |
| [102] | FTPSserver | VM | Debian 13.0  | 2 | 4 GB | 60 GB | FTPS server |

## Backups and Recovery

- **Backup tool/location:** external SSD disk
- **Schedule:** Whenever major changes happen, prior to major updates also.

## Services

| Service | Runs on | Purpose | Access |
|---|---|---|---|
| FTPS | VM ID 102 | files retention and backup | LAN only |


## Maintenance

- Updates only happen on stable tested updates prior to each update I backup my configurations and important files to external physical medium. 
- Configurations are tracked post each change to a local markdown file.


