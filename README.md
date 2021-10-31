# firecracker-ctl

A CLI to control [firecracker](https://github.com/firecracker-microvm/firecracker)

## Short introduction

### Starting a VM

_Note: x86 guest_

```shell
# Download Kernel + RootFS
wget https://s3.amazonaws.com/spec.ccfc.min/img/quickstart_guide/x86_64/kernels/vmlinux.bin
wget https://s3.amazonaws.com/spec.ccfc.min/img/quickstart_guide/x86_64/rootfs/bionic.rootfs.ext4

# Start Firecracker (insecure mode)
firecracker --api-sock /tmp/firecracker.sock
```

Now, in another terminal (same `$PWD`), give the following `firecracker-ctl` commands:

```shell
firecracker-ctl boot-source \
  --boot-args "console=ttyS0 reboot=k panic=1 pci=off" \
  "$PWD/vmlinux.bin"
  
firecracker-ctl drive \
  --id rootfs \
  -r=false \
  -R=true \
  -p "$PWD/bionic.rootfs.ext4"
  
firecracker-ctl action start
```