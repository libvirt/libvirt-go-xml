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
						FileSource: &DomainDiskFileSource{
							File: "/var/lib/libvirt/images/demo.qcow2",
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
			`  <devices>`,
			`    <interface type="network">`,
			`      <mac address="00:11:22:33:44:55"></mac>`,
			`      <model type="virtio"></model>`,
			`    </interface>`,
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
