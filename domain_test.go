/*
 * This file is part of the libvirt-go-xml project
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 *
 * Copyright (C) 2016 Red Hat, Inc.
 *
 */

package libvirtxml

import (
	"strings"
	"testing"
)

var domainTestData = []struct {
	Object   *Domain
	Expected []string
}{
	{
		Object: &Domain{
			Type: "kvm",
			Name: "test",
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`</domain>`,
		},
	},
	{
		Object: &Domain{
			Type: "kvm",
			Name: "test",
			Devices: &DomainDeviceList{
				Disks: []DomainDisk{
					DomainDisk{
						Type:   "file",
						Device: "cdrom",
						Driver: &DomainDiskDriver{
							Name: "qemu",
							Type: "qcow2",
						},
						Source: &DomainDiskSource{
							File: "/var/lib/libvirt/images/demo.qcow2",
						},
						Target: &DomainDiskTarget{
							Dev: "vda",
							Bus: "virtio",
						},
					},
					DomainDisk{
						Type:   "block",
						Device: "disk",
						Driver: &DomainDiskDriver{
							Name: "qemu",
							Type: "raw",
						},
						Source: &DomainDiskSource{
							Device: "/dev/sda1",
						},
						Target: &DomainDiskTarget{
							Dev: "vdb",
							Bus: "virtio",
						},
					},
					DomainDisk{
						Type:   "network",
						Device: "disk",
						Auth: &DomainDiskAuth{
							Username: "fred",
							Secret: &DomainDiskSecret{
								Type: "ceph",
								UUID: "e49f09c9-119e-43fd-b5a9-000d41e65493",
							},
						},
						Source: &DomainDiskSource{
							Protocol: "rbd",
							Name:     "somepool/somevol",
							Hosts: []DomainDiskSourceHost{
								DomainDiskSourceHost{
									Transport: "tcp",
									Name:      "rbd1.example.com",
									Port:      "3000",
								},
								DomainDiskSourceHost{
									Transport: "tcp",
									Name:      "rbd2.example.com",
									Port:      "3000",
								},
							},
						},
						Target: &DomainDiskTarget{
							Dev: "vdc",
							Bus: "virtio",
						},
					},
					DomainDisk{
						Type:   "network",
						Device: "disk",
						Source: &DomainDiskSource{
							Protocol: "nbd",
							Hosts: []DomainDiskSourceHost{
								DomainDiskSourceHost{
									Transport: "unix",
									Socket:    "/var/run/nbd.sock",
								},
							},
						},
						Target: &DomainDiskTarget{
							Dev: "vdd",
							Bus: "virtio",
						},
					},
					DomainDisk{
						Type: "volume",
						Device: "cdrom",
						Source: &DomainDiskSource{
							Pool: "default",
							Volume: "myvolume",
						},
						Target: &DomainDiskTarget{
							Dev: "vde",
							Bus: "virtio",
						},
					},
				},
			},
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`  <devices>`,
			`    <disk type="file" device="cdrom">`,
			`      <driver name="qemu" type="qcow2"></driver>`,
			`      <source file="/var/lib/libvirt/images/demo.qcow2"></source>`,
			`      <target dev="vda" bus="virtio"></target>`,
			`    </disk>`,
			`    <disk type="block" device="disk">`,
			`      <driver name="qemu" type="raw"></driver>`,
			`      <source dev="/dev/sda1"></source>`,
			`      <target dev="vdb" bus="virtio"></target>`,
			`    </disk>`,
			`    <disk type="network" device="disk">`,
			`      <auth username="fred">`,
			`        <secret type="ceph" uuid="e49f09c9-119e-43fd-b5a9-000d41e65493"></secret>`,
			`      </auth>`,
			`      <source protocol="rbd" name="somepool/somevol">`,
			`        <host transport="tcp" name="rbd1.example.com" port="3000"></host>`,
			`        <host transport="tcp" name="rbd2.example.com" port="3000"></host>`,
			`      </source>`,
			`      <target dev="vdc" bus="virtio"></target>`,
			`    </disk>`,
			`    <disk type="network" device="disk">`,
			`      <source protocol="nbd">`,
			`        <host transport="unix" socket="/var/run/nbd.sock"></host>`,
			`      </source>`,
			`      <target dev="vdd" bus="virtio"></target>`,
			`    </disk>`,
			`    <disk type="volume" device="cdrom">`,
			`      <source pool="default" volume="myvolume"></source>`,
			`      <target dev="vde" bus="virtio"></target>`,
			`    </disk>`,
			`  </devices>`,
			`</domain>`,
		},
	},
	{
		Object: &Domain{
			Type: "kvm",
			Name: "test",
			Devices: &DomainDeviceList{
				Inputs: []DomainInput{
					DomainInput{
						Type: "tablet",
						Bus:  "usb",
					},
					DomainInput{
						Type: "keyboard",
						Bus:  "ps2",
					},
				},
				Videos: []DomainVideo{
					DomainVideo{
						Model: DomainVideoModel{
							Type: "cirrus",
						},
					},
				},
				Graphics: []DomainGraphic{
					DomainGraphic{
						Type: "vnc",
					},
				},
			},
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`  <devices>`,
			`    <input type="tablet" bus="usb"></input>`,
			`    <input type="keyboard" bus="ps2"></input>`,
			`    <graphics type="vnc"></graphics>`,
			`    <video>`,
			`      <model type="cirrus"></model>`,
			`    </video>`,
			`  </devices>`,
			`</domain>`,
		},
	},
	{
		Object: &Domain{
			Type: "kvm",
			Name: "test",
			Memory: &DomainMemory{
				Unit:  "KiB",
				Value: 8192,
			},
			CurrentMemory: &DomainMemory{
				Unit:  "KiB",
				Value: 4096,
			},
			MaximumMemory: &DomainMaxMemory{
				Unit:  "KiB",
				Value: 16384,
				Slots: 2,
			},
			OS: &DomainOS{
				Type: &DomainOSType{
					Arch:    "x86_64",
					Machine: "pc",
					Type:    "hvm",
				},
				BootDevices: []DomainBootDevice{
					DomainBootDevice{
						Dev: "hd",
					},
				},
				Loader: &DomainLoader{
					Readonly: "yes",
					Secure:   "no",
					Type:     "rom",
					Path:     "/loader",
				},
				SMBios: &DomainSMBios{
					Mode: "sysinfo",
				},
				BIOS: &DomainBIOS{
					UseSerial:     "yes",
					RebootTimeout: "0",
				},
				Init: "/bin/systemd",
				InitArgs: []string{
					"--unit",
					"emergency.service",
				},
			},
			SysInfo: &DomainSysInfo{
				Type: "smbios",
				BIOS: []DomainSysInfoEntry{
					DomainSysInfoEntry{
						Name:  "vendor",
						Value: "vendor",
					},
				},
				System: []DomainSysInfoEntry{
					DomainSysInfoEntry{
						Name:  "manufacturer",
						Value: "manufacturer",
					},
					DomainSysInfoEntry{
						Name:  "product",
						Value: "product",
					},
					DomainSysInfoEntry{
						Name:  "version",
						Value: "version",
					},
				},
				BaseBoard: []DomainSysInfoEntry{
					DomainSysInfoEntry{
						Name:  "manufacturer",
						Value: "manufacturer",
					},
					DomainSysInfoEntry{
						Name:  "product",
						Value: "product",
					},
					DomainSysInfoEntry{
						Name:  "version",
						Value: "version",
					},
					DomainSysInfoEntry{
						Name:  "serial",
						Value: "serial",
					},
				},
			},
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`  <memory unit="KiB">8192</memory>`,
			`  <currentMemory unit="KiB">4096</currentMemory>`,
			`  <maxMemory unit="KiB" slots="2">16384</maxMemory>`,
			`  <os>`,
			`    <type arch="x86_64" machine="pc">hvm</type>`,
			`    <loader readonly="yes" secure="no" type="rom">/loader</loader>`,
			`    <boot dev="hd"></boot>`,
			`    <smbios mode="sysinfo"></smbios>`,
			`    <bios useserial="yes" rebootTimeout="0"></bios>`,
			`    <init>/bin/systemd</init>`,
			`    <initarg>--unit</initarg>`,
			`    <initarg>emergency.service</initarg>`,
			`  </os>`,
			`  <sysinfo type="smbios">`,
			`    <system>`,
			`      <entry name="manufacturer">manufacturer</entry>`,
			`      <entry name="product">product</entry>`,
			`      <entry name="version">version</entry>`,
			`    </system>`,
			`    <bios>`,
			`      <entry name="vendor">vendor</entry>`,
			`    </bios>`,
			`    <baseBoard>`,
			`      <entry name="manufacturer">manufacturer</entry>`,
			`      <entry name="product">product</entry>`,
			`      <entry name="version">version</entry>`,
			`      <entry name="serial">serial</entry>`,
			`    </baseBoard>`,
			`  </sysinfo>`,
			`</domain>`,
		},
	},
	{
		Object: &Domain{
			Type: "kvm",
			Name: "test",
			OS: &DomainOS{
				NVRam: &DomainNVRam{
					Template: "/t.fd",
					NVRam:    "/vars.fd",
				},
				BootMenu: &DomainBootMenu{
					Enabled: "yes",
					Timeout: "3000",
				},
			},
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`  <os>`,
			`    <nvram template="/t.fd">/vars.fd</nvram>`,
			`    <bootmenu enabled="yes" timeout="3000"></bootmenu>`,
			`  </os>`,
			`</domain>`,
		},
	},
	{
		Object: &Domain{
			Type: "kvm",
			Name: "test",
			OS: &DomainOS{
				Kernel:     "/vmlinuz",
				Initrd:     "/initrd",
				KernelArgs: "arg",
			},
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`  <os>`,
			`    <kernel>/vmlinuz</kernel>`,
			`    <initrd>/initrd</initrd>`,
			`    <cmdline>arg</cmdline>`,
			`  </os>`,
			`</domain>`,
		},
	},
	{
		Object: &Domain{
			Type: "kvm",
			Name: "test",
			Resource: &DomainResource{
				Partition: "/machines/production",
			},
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`  <resource>`,
			`    <partition>/machines/production</partition>`,
			`  </resource>`,
			`</domain>`,
		},
	},
	{
		Object: &Domain{
			Type: "kvm",
			Name: "test",
			VCPU: &DomainVCPU{
				Placement: "static",
				CPUSet:    "1-4,^3,6",
				Current:   "1",
				Value:     2,
			},
			Devices: &DomainDeviceList{
				Interfaces: []DomainInterface{
					DomainInterface{
						Type: "network",
						MAC: &DomainInterfaceMAC{
							Address: "00:11:22:33:44:55",
						},
						Model: &DomainInterfaceModel{
							Type: "virtio",
						},
					},
				},
			},
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`  <vcpu placement="static" cpuset="1-4,^3,6" current="1">2</vcpu>`,
			`  <devices>`,
			`    <interface type="network">`,
			`      <mac address="00:11:22:33:44:55"></mac>`,
			`      <model type="virtio"></model>`,
			`    </interface>`,
			`  </devices>`,
			`</domain>`,
		},
	},
	{
		Object: &Domain{
			Type: "kvm",
			Name: "test",
			CPU: &DomainCPU{
				Match: "exact",
				Model: &DomainCPUModel{
					Fallback: "allow",
					Value:    "core2duo",
				},
				Vendor: "Intel",
				Topology: &DomainCPUTopology{
					Sockets: 1,
					Cores:   2,
					Threads: 1,
				},
				Features: []DomainCPUFeature{
					DomainCPUFeature{Policy: "disable", Name: "lahf_lm"},
				},
			},
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`  <cpu match="exact">`,
			`    <model fallback="allow">core2duo</model>`,
			`    <vendor>Intel</vendor>`,
			`    <topology sockets="1" cores="2" threads="1"></topology>`,
			`    <feature policy="disable" name="lahf_lm"></feature>`,
			`  </cpu>`,
			`</domain>`,
		},
	},
	{
		Object: &Domain{
			Type: "kvm",
			Name: "test",
			Devices: &DomainDeviceList{
				Interfaces: []DomainInterface{
					DomainInterface{
						Type: "bridge",
						MAC: &DomainInterfaceMAC{
							Address: "06:39:b4:00:00:46",
						},
						Model: &DomainInterfaceModel{
							Type: "virtio",
						},
						Source: &DomainInterfaceSource{
							Bridge: "private",
						},
						Target: &DomainInterfaceTarget{
							Dev: "vnet3",
						},
						Alias: &DomainInterfaceAlias{
							Name: "net1",
						},
					},
				},
			},
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`  <devices>`,
			`    <interface type="bridge">`,
			`      <mac address="06:39:b4:00:00:46"></mac>`,
			`      <model type="virtio"></model>`,
			`      <source bridge="private"></source>`,
			`      <target dev="vnet3"></target>`,
			`      <alias name="net1"></alias>`,
			`    </interface>`,
			`  </devices>`,
			`</domain>`,
		},
	},
	{
		Object: &Domain{
			Type: "kvm",
			Name: "test",
			Devices: &DomainDeviceList{
				Interfaces: []DomainInterface{
					DomainInterface{
						Type: "network",
						MAC: &DomainInterfaceMAC{
							Address: "52:54:00:39:97:ac",
						},
						Model: &DomainInterfaceModel{
							Type: "e1000",
						},
						Source: &DomainInterfaceSource{
							Network: "default",
						},
					},
				},
			},
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`  <devices>`,
			`    <interface type="network">`,
			`      <mac address="52:54:00:39:97:ac"></mac>`,
			`      <model type="e1000"></model>`,
			`      <source network="default"></source>`,
			`    </interface>`,
			`  </devices>`,
			`</domain>`,
		},
	},
	{
		Object: &Domain{
			Type: "kvm",
			Name: "test",
			Devices: &DomainDeviceList{
				Interfaces: []DomainInterface{
					DomainInterface{
						Type: "user",
						MAC: &DomainInterfaceMAC{
							Address: "52:54:00:39:97:ac",
						},
						Model: &DomainInterfaceModel{
							Type: "virtio",
						},
						Link: &DomainInterfaceLink{
							State: "up",
						},
						Boot: &DomainInterfaceBoot{
							Order: 1,
						},
						Driver: &DomainInterfaceDriver{
							Name:   "vhost",
							Queues: 5,
						},
					},
				},
			},
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`  <devices>`,
			`    <interface type="user">`,
			`      <mac address="52:54:00:39:97:ac"></mac>`,
			`      <model type="virtio"></model>`,
			`      <link state="up"></link>`,
			`      <boot order="1"></boot>`,
			`      <driver name="vhost" queues="5"></driver>`,
			`    </interface>`,
			`  </devices>`,
			`</domain>`,
		},
	},
	{
		Object: &Domain{
			Type: "kvm",
			Name: "test",
			Devices: &DomainDeviceList{
				Interfaces: []DomainInterface{
					DomainInterface{
						Type: "server",
						MAC: &DomainInterfaceMAC{
							Address: "52:54:00:39:97:ac",
						},
						Model: &DomainInterfaceModel{
							Type: "virtio",
						},
						Source: &DomainInterfaceSource{
							Address: "127.0.0.1",
							Port:    1234,
						},
					},
				},
			},
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`  <devices>`,
			`    <interface type="server">`,
			`      <mac address="52:54:00:39:97:ac"></mac>`,
			`      <model type="virtio"></model>`,
			`      <source address="127.0.0.1" port="1234"></source>`,
			`    </interface>`,
			`  </devices>`,
			`</domain>`,
		},
	},
	{
		Object: &Domain{
			Type: "kvm",
			Name: "test",
			Devices: &DomainDeviceList{
				Interfaces: []DomainInterface{
					DomainInterface{
						Type: "ethernet",
						MAC: &DomainInterfaceMAC{
							Address: "52:54:00:39:97:ac",
						},
						Model: &DomainInterfaceModel{
							Type: "virtio",
						},
						Script: &DomainInterfaceScript{
							Path: "/etc/qemu-ifup",
						},
					},
				},
			},
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`  <devices>`,
			`    <interface type="ethernet">`,
			`      <mac address="52:54:00:39:97:ac"></mac>`,
			`      <model type="virtio"></model>`,
			`      <script path="/etc/qemu-ifup"></script>`,
			`    </interface>`,
			`  </devices>`,
			`</domain>`,
		},
	},
	{
		Object: &Domain{
			Type: "kvm",
			Name: "test",
			Devices: &DomainDeviceList{
				Interfaces: []DomainInterface{
					DomainInterface{
						Type: "vhostuser",
						MAC: &DomainInterfaceMAC{
							Address: "52:54:00:39:97:ac",
						},
						Model: &DomainInterfaceModel{
							Type: "virtio",
						},
						Source: &DomainInterfaceSource{
							Type: "unix",
							Path: "/tmp/vhost0.sock",
							Mode: "server",
						},
					},
				},
			},
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`  <devices>`,
			`    <interface type="vhostuser">`,
			`      <mac address="52:54:00:39:97:ac"></mac>`,
			`      <model type="virtio"></model>`,
			`      <source type="unix" path="/tmp/vhost0.sock" mode="server"></source>`,
			`    </interface>`,
			`  </devices>`,
			`</domain>`,
		},
	},
	{
		Object: &Domain{
			Type: "kvm",
			Name: "test",
			Devices: &DomainDeviceList{
				Filesystems: []DomainFilesystem{
					DomainFilesystem{
						Type:       "mount",
						AccessMode: "mapped",
						Driver: &DomainFilesystemDriver{
							Type:     "path",
							WRPolicy: "immediate",
						},
						Source: &DomainFilesystemSource{
							Dir: "/home/user/test",
						},
						Target: &DomainFilesystemTarget{
							Dir: "user-test-mount",
						},
					},
					DomainFilesystem{
						Type:       "file",
						AccessMode: "passthrough",
						Driver: &DomainFilesystemDriver{
							Name: "loop",
							Type: "raw",
						},
						Source: &DomainFilesystemSource{
							File: "/home/user/test.img",
						},
						Target: &DomainFilesystemTarget{
							Dir: "user-file-test-mount",
						},
					},
				},
			},
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`  <devices>`,
			`    <filesystem type="mount" accessmode="mapped">`,
			`      <driver type="path" wrpolicy="immediate"></driver>`,
			`      <source dir="/home/user/test"></source>`,
			`      <target dir="user-test-mount"></target>`,
			`    </filesystem>`,
			`    <filesystem type="file" accessmode="passthrough">`,
			`      <driver type="raw" name="loop"></driver>`,
			`      <source file="/home/user/test.img"></source>`,
			`      <target dir="user-file-test-mount"></target>`,
			`    </filesystem>`,
			`  </devices>`,
			`</domain>`,
		},
	},
}

func TestDomain(t *testing.T) {
	for _, test := range domainTestData {
		doc, err := test.Object.Marshal()
		if err != nil {
			t.Fatal(err)
		}

		expect := strings.Join(test.Expected, "\n")

		if doc != expect {
			t.Fatal("Bad xml:\n", string(doc), "\n does not match\n", expect, "\n")
		}
	}
}
