# Proxmox FTPS File Transfer Server

A self-hosted file-transfer service deployed on a Debian virtual machine running under Proxmox VE. This project documents hands-on practice with Linux service administration, network services, TLS, filesystem permissions, and client-side testing.

## Project status

**Working LAN prototype — security validation still in progress.**

## Objectives

- Deploy a Debian server VM on Proxmox VE.
- Install and configure `vsftpd`.
- Enable explicit FTP over TLS (FTPS).
- Restrict access to a dedicated transfer account.
- Configure a dedicated directory for transferred files.
- Test connectivity and transfers using FileZilla.
- Record security decisions, test evidence, and future improvements.

## Lab architecture

| Component | Details |
|---|---|
| Hypervisor | Proxmox VE |
| Server OS | Debian 13.7 (x86_64) |
| FTP service | vsftpd |
| Protocol | Explicit FTPS |
| Client | FileZilla on Linux |
| Network scope | Local area network (LAN) |
| Server address used during testing | `192.168.1.135` (DHCP; may change) |

The Debian VM hosts the service. A separate client connects over the LAN, authenticates with the dedicated account, and accesses the configured file directory.

## Implemented configuration

- Anonymous FTP access disabled.
- Dedicated account: `ftpsuser`.
- Dedicated transfer directory: `/srv/ftps/ftpsuser/files`.
- Explicit TLS configured using a self-signed certificate.
- TLS required for local-user logins and data connections.
- User-list allowlisting enabled for the transfer account.
- User chroot enabled.
- Passive transfer ports configured as TCP `40000–40009`.
- SSH administration performed with a regular account and `sudo`; root SSH login remains disabled.

A self-signed certificate encrypts the connection but does not provide identity verification by itself. Verify its fingerprint through a trusted channel before accepting it on a client.

## Validation log

| Test | Status | Evidence / notes |
|---|---|---|
| Connect from FileZilla to the server | Passed | Successful client session |
| Authenticate as the dedicated account | Passed | FileZilla status showed login |
| List the remote `/files` directory | Passed | Remote directory visible |
| Upload a test file | Passed | Upload reported successful |
| Download `helloWorld.txt` | Passed | FileZilla reported successful transfer |
| Independently verify TLS certificate and negotiated settings | Pending | Capture and record certificate details |
| Confirm plain, unencrypted FTP is rejected | Pending | Test without weakening the server configuration |
| Test rejected credentials and account restrictions | Pending | Record expected and actual outcomes |
| Test directory isolation and permissions | Pending | Verify access is limited to intended paths |
| Configure and verify LAN-only firewall rules | Pending | Preserve SSH access while testing |
| Reserve a stable DHCP address | Pending | Configure on the router if supported |

## Security scope and limitations

This project is currently intended for LAN use only. It is **not configured for public internet access**. No router port forwarding should be added as part of this initial lab.

The service has been tested for basic connectivity and file transfer, but it should not yet be described as fully hardened. Firewall restrictions, certificate verification, negative authentication tests, directory isolation, and TLS details remain to be validated.

## Repository hygiene

Do not commit passwords, private keys, real credentials, or unredacted configuration dumps. If publishing screenshots, redact passwords and any other information you do not want public. Use sanitized example configuration files rather than copying private key material.

## Planned improvements

1. Inspect the current firewall and listening services.
2. Apply LAN-restricted firewall rules for the FTPS control and passive ports without losing SSH access.
3. Verify the certificate fingerprint and TLS-protected data channel.
4. Test invalid credentials, access restrictions, and directory isolation.
5. Reserve a stable LAN address and update the documentation.
6. Add sanitized configuration examples and troubleshooting notes.
7. Build a separate VPN-server lab as a follow-up project.

## Learning outcomes

This lab provides practical experience with Linux services, virtual machines, TCP ports, TLS certificates, account restrictions, filesystem permissions, troubleshooting, and evidence-based technical documentation.
