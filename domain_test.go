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
	"reflect"
	"strings"
	"testing"
)

type PCIAddress struct {
	Domain   uint
	Bus      uint
	Slot     uint
	Function uint
}

type DriveAddress struct {
	Controller uint
	Bus        uint
	Target     uint
	Unit       uint
}

type ISAAddress struct {
	IOBase uint
}

var domainID int = 3

var uhciIndex uint = 0
var uhciAddr = PCIAddress{0, 0, 1, 2}
var pciIndex uint = 0
var pciTargetChassisNr uint = 7
var pciTargetChassis uint = 23
var pciTargetPort uint = 78
var pciTargetBusNr uint = 2
var pciTargetIndex uint = 3
var pciTargetNUMANode uint = 2
var scsiIndex uint = 0
var scsiQueues uint = 3
var scsiCmdPerLUN uint = 8
var scsiMaxSectors uint = 512

var usbHostBus uint = 14
var usbHostDevice uint = 6

var pciHostDomain uint = 0
var pciHostBus uint = 3
var pciHostSlot uint = 14
var pciHostFunction uint = 5

var diskAddr = PCIAddress{0, 0, 3, 0}
var ifaceAddr = PCIAddress{0, 0, 4, 0}
var videoAddr = PCIAddress{0, 0, 5, 0}
var fsAddr = PCIAddress{0, 0, 6, 0}
var balloonAddr = PCIAddress{0, 0, 7, 0}
var panicAddr = ISAAddress{0x505}
var duplexAddr = PCIAddress{0, 0, 8, 0}
var watchdogAddr = PCIAddress{0, 0, 8, 0}
var rngAddr = PCIAddress{0, 0, 9, 0}
var hostdevSCSI = DriveAddress{0, 0, 3, 0}

var serialPort uint = 0
var tabletBus uint = 0
var tabletPort string = "1.1"

var nicAverage int = 1000
var nicBurst int = 10000

var vcpuId0 uint = 0
var vcpuOrder0 uint = 1
var vcpuId1 uint = 1

var memorydevAddressSlot uint = 0
var memorydevAddressBase uint64 = 4294967296

var rebootTimeout uint = 0
var cellID uint = 0

var ipv6Prefix uint = 24

var domainTestData = []struct {
	Object   Document
	Expected []string
}{
	{
		Object: &Domain{
			Type: "kvm",
			Name: "test",
			ID:   &domainID,
		},
		Expected: []string{
			`<domain type="kvm" id="3">`,
			`  <name>test</name>`,
			`</domain>`,
		},
	},
	{
		Object: &Domain{
			Type:        "kvm",
			Name:        "test",
			Title:       "Test",
			Description: "A test guest config",
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
						Serial: "fishfood",
						Boot: &DomainDeviceBoot{
							Order: 1,
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
						Address: &DomainAddress{
							PCI: &DomainAddressPCI{
								Domain:   &diskAddr.Domain,
								Bus:      &diskAddr.Bus,
								Slot:     &diskAddr.Slot,
								Function: &diskAddr.Function,
							},
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
						Shareable: &DomainDiskShareable{},
					},
					DomainDisk{
						Type:   "volume",
						Device: "cdrom",
						Driver: &DomainDiskDriver{
							Cache:       "none",
							IO:          "native",
							ErrorPolicy: "stop",
						},
						Source: &DomainDiskSource{
							Pool:   "default",
							Volume: "myvolume",
						},
						Target: &DomainDiskTarget{
							Dev: "vde",
							Bus: "virtio",
						},
						ReadOnly: &DomainDiskReadOnly{},
					},
				},
			},
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`  <title>Test</title>`,
			`  <description>A test guest config</description>`,
			`  <devices>`,
			`    <disk type="file" device="cdrom">`,
			`      <driver name="qemu" type="qcow2"></driver>`,
			`      <source file="/var/lib/libvirt/images/demo.qcow2"></source>`,
			`      <target dev="vda" bus="virtio"></target>`,
			`      <serial>fishfood</serial>`,
			`      <boot order="1"></boot>`,
			`    </disk>`,
			`    <disk type="block" device="disk">`,
			`      <driver name="qemu" type="raw"></driver>`,
			`      <source dev="/dev/sda1"></source>`,
			`      <target dev="vdb" bus="virtio"></target>`,
			`      <address type="pci" domain="0x0000" bus="0x00" slot="0x03" function="0x0"></address>`,
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
			`      <shareable></shareable>`,
			`    </disk>`,
			`    <disk type="volume" device="cdrom">`,
			`      <driver cache="none" io="native" error_policy="stop"></driver>`,
			`      <source pool="default" volume="myvolume"></source>`,
			`      <target dev="vde" bus="virtio"></target>`,
			`      <readonly></readonly>`,
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
						Address: &DomainAddress{
							USB: &DomainAddressUSB{
								Bus:  &tabletBus,
								Port: tabletPort,
							},
						},
					},
					DomainInput{
						Type: "keyboard",
						Bus:  "ps2",
					},
				},
				Videos: []DomainVideo{
					DomainVideo{
						Model: DomainVideoModel{
							Type:   "cirrus",
							Heads:  1,
							Ram:    4096,
							VRam:   8192,
							VGAMem: 256,
						},
						Address: &DomainAddress{
							PCI: &DomainAddressPCI{
								Domain:   &videoAddr.Domain,
								Bus:      &videoAddr.Bus,
								Slot:     &videoAddr.Slot,
								Function: &videoAddr.Function,
							},
						},
					},
				},
				Graphics: []DomainGraphic{
					DomainGraphic{
						Type: "vnc",
					},
				},
				MemBalloon: &DomainMemBalloon{
					Model: "virtio",
					Address: &DomainAddress{
						PCI: &DomainAddressPCI{
							Domain:   &balloonAddr.Domain,
							Bus:      &balloonAddr.Bus,
							Slot:     &balloonAddr.Slot,
							Function: &balloonAddr.Function,
						},
					},
				},
				Panics: []DomainPanic{
					DomainPanic{
						Model: "hyperv",
					},
					DomainPanic{
						Model: "isa",
						Address: &DomainAddress{
							ISA: &DomainAddressISA{
								IOBase: &panicAddr.IOBase,
							},
						},
					},
				},
				Consoles: []DomainConsole{
					DomainConsole{
						Type: "pty",
						Target: &DomainConsoleTarget{
							Type: "virtio",
							Port: &serialPort,
						},
					},
				},
				Serials: []DomainSerial{
					DomainSerial{
						Type: "pty",
						Target: &DomainSerialTarget{
							Type: "isa",
							Port: &serialPort,
						},
					},
					DomainSerial{
						Type: "file",
						Source: &DomainChardevSource{
							Path:   "/tmp/serial.log",
							Append: "off",
						},
						Target: &DomainSerialTarget{
							Port: &serialPort,
						},
					},
					DomainSerial{
						Type: "tcp",
						Source: &DomainChardevSource{
							Mode:    "bind",
							Host:    "127.0.0.1",
							Service: "1234",
							TLS:     "yes",
						},
						Protocol: &DomainSerialProtocol{
							Type: "telnet",
						},
						Target: &DomainSerialTarget{
							Port: &serialPort,
						},
					},
				},
				Channels: []DomainChannel{
					DomainChannel{
						Type: "pty",
						Target: &DomainChannelTarget{
							Type:  "virtio",
							Name:  "org.redhat.spice",
							State: "connected",
						},
					},
				},
				Sounds: []DomainSound{
					DomainSound{
						Model: "ich6",
						Codec: &DomainSoundCodec{
							Type: "duplex",
						},
						Address: &DomainAddress{
							PCI: &DomainAddressPCI{
								Domain:   &duplexAddr.Domain,
								Bus:      &duplexAddr.Bus,
								Slot:     &duplexAddr.Slot,
								Function: &duplexAddr.Function,
							},
						},
					},
				},
				RNGs: []DomainRNG{
					DomainRNG{
						Model: "virtio",
						Rate: &DomainRNGRate{
							Period: 2000,
							Bytes:  1234,
						},
						Backend: &DomainRNGBackend{
							EGD: &DomainRNGBackendEGD{
								Type: "udp",
								Sources: []DomainInterfaceSource{
									DomainInterfaceSource{
										Mode:    "bind",
										Service: "1234",
									},
									DomainInterfaceSource{
										Mode:    "connect",
										Host:    "1.2.3.4",
										Service: "1234",
									},
								},
								Protocol: &DomainRNGProtocol{
									Type: "raw",
								},
							},
						},
					},
				},
				Memorydevs: []DomainMemorydev{
					DomainMemorydev{
						Model:  "dimm",
						Access: "private",
						Target: &DomainMemorydevTarget{
							Size: &DomainMemorydevTargetSize{
								Value: 1,
								Unit:  "GiB",
							},
							Node: &DomainMemorydevTargetNode{
								Value: 0,
							},
						},
						Address: &DomainAddress{
							DIMM: &DomainAddressDIMM{
								Slot: &memorydevAddressSlot,
								Base: &memorydevAddressBase,
							},
						},
					},
				},
			},
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`  <devices>`,
			`    <serial type="pty">`,
			`      <target type="isa" port="0"></target>`,
			`    </serial>`,
			`    <serial type="file">`,
			`      <source path="/tmp/serial.log" append="off"></source>`,
			`      <target port="0"></target>`,
			`    </serial>`,
			`    <serial type="tcp">`,
			`      <source mode="bind" host="127.0.0.1" service="1234" tls="yes"></source>`,
			`      <protocol type="telnet"></protocol>`,
			`      <target port="0"></target>`,
			`    </serial>`,
			`    <console type="pty">`,
			`      <target type="virtio" port="0"></target>`,
			`    </console>`,
			`    <input type="tablet" bus="usb">`,
			`      <address type="usb" bus="0" port="1.1"></address>`,
			`    </input>`,
			`    <input type="keyboard" bus="ps2"></input>`,
			`    <graphics type="vnc"></graphics>`,
			`    <video>`,
			`      <model type="cirrus" heads="1" ram="4096" vram="8192" vgamem="256"></model>`,
			`      <address type="pci" domain="0x0000" bus="0x00" slot="0x05" function="0x0"></address>`,
			`    </video>`,
			`    <channel type="pty">`,
			`      <target type="virtio" name="org.redhat.spice" state="connected"></target>`,
			`    </channel>`,
			`    <memballoon model="virtio">`,
			`      <address type="pci" domain="0x0000" bus="0x00" slot="0x07" function="0x0"></address>`,
			`    </memballoon>`,
			`    <panic model="hyperv"></panic>`,
			`    <panic model="isa">`,
			`      <address type="isa" iobase="0x505"></address>`,
			`    </panic>`,
			`    <sound model="ich6">`,
			`      <codec type="duplex"></codec>`,
			`      <address type="pci" domain="0x0000" bus="0x00" slot="0x08" function="0x0"></address>`,
			`    </sound>`,
			`    <rng model="virtio">`,
			`      <rate bytes="1234" period="2000"></rate>`,
			`      <backend model="egd" type="udp">`,
			`        <source mode="bind" service="1234"></source>`,
			`        <source mode="connect" service="1234" host="1.2.3.4"></source>`,
			`        <protocol type="raw"></protocol>`,
			`      </backend>`,
			`    </rng>`,
			`    <memory model="dimm" access="private">`,
			`      <target>`,
			`        <size unit="GiB">1</size>`,
			`        <node>0</node>`,
			`      </target>`,
			`      <address type="dimm" slot="0" base="0x100000000"></address>`,
			`    </memory>`,
			`  </devices>`,
			`</domain>`,
		},
	},
	{
		Object: &Domain{
			Type: "kvm",
			Name: "test",
			Memory: &DomainMemory{
				Unit:     "KiB",
				Value:    8192,
				DumpCore: "yes",
			},
			CurrentMemory: &DomainCurrentMemory{
				Unit:  "KiB",
				Value: 4096,
			},
			MaximumMemory: &DomainMaxMemory{
				Unit:  "KiB",
				Value: 16384,
				Slots: 2,
			},
			MemoryBacking: &DomainMemoryBacking{
				MemoryHugePages: &DomainMemoryHugepages{
					Hugepages: []DomainMemoryHugepage{
						{
							Size:    1,
							Unit:    "G",
							Nodeset: "0-3,5",
						},
						{
							Size:    2,
							Unit:    "M",
							Nodeset: "4",
						},
					},
				},
				MemoryNosharepages: &DomainMemoryNosharepages{},
				MemoryLocked:       &DomainMemoryLocked{},
				MemorySource: &DomainMemorySource{
					Type: "file",
				},
				MemoryAccess: &DomainMemoryAccess{
					Mode: "shared",
				},
				MemoryAllocation: &DomainMemoryAllocation{
					Mode: "immediate",
				},
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
				DTB: "/some/path",
				ACPI: &DomainACPI{
					Tables: []DomainACPITable{
						DomainACPITable{
							Type: "slic",
							Path: "/some/data",
						},
					},
				},
				SMBios: &DomainSMBios{
					Mode: "sysinfo",
				},
				BIOS: &DomainBIOS{
					UseSerial:     "yes",
					RebootTimeout: &rebootTimeout,
				},
				Init: "/bin/systemd",
				InitArgs: []string{
					"--unit",
					"emergency.service",
				},
				InitEnv: []DomainOSInitEnv{
					DomainOSInitEnv{
						Name:  "HOME",
						Value: "/home/fred",
					},
					DomainOSInitEnv{
						Name:  "USER",
						Value: "fred",
					},
				},
				InitUser:  "fred",
				InitGroup: "fred",
				InitDir:   "/home/fred",
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
			Clock: &DomainClock{
				Offset:     "variable",
				Basis:      "utc",
				Adjustment: 28794,
				TimeZone:   "Europe/Paris",
				Timer: []DomainTimer{
					DomainTimer{
						Name:       "rtc",
						Track:      "boot",
						TickPolicy: "catchup",
						CatchUp: &DomainTimerCatchUp{
							Threshold: 123,
							Slew:      120,
							Limit:     10000,
						},
						Frequency: 120,
						Mode:      "auto",
					},
				},
			},
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`  <maxMemory unit="KiB" slots="2">16384</maxMemory>`,
			`  <memory unit="KiB" dumpCore="yes">8192</memory>`,
			`  <currentMemory unit="KiB">4096</currentMemory>`,
			`  <memoryBacking>`,
			`    <hugepages>`,
			`      <page size="1" unit="G" nodeset="0-3,5"></page>`,
			`      <page size="2" unit="M" nodeset="4"></page>`,
			`    </hugepages>`,
			`    <nosharepages></nosharepages>`,
			`    <locked></locked>`,
			`    <source type="file"></source>`,
			`    <access mode="shared"></access>`,
			`    <allocation mode="immediate"></allocation>`,
			`  </memoryBacking>`,
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
			`  <os>`,
			`    <type arch="x86_64" machine="pc">hvm</type>`,
			`    <init>/bin/systemd</init>`,
			`    <initarg>--unit</initarg>`,
			`    <initarg>emergency.service</initarg>`,
			`    <initenv name="HOME">/home/fred</initenv>`,
			`    <initenv name="USER">fred</initenv>`,
			`    <initdir>/home/fred</initdir>`,
			`    <inituser>fred</inituser>`,
			`    <initgroup>fred</initgroup>`,
			`    <loader readonly="yes" secure="no" type="rom">/loader</loader>`,
			`    <dtb>/some/path</dtb>`,
			`    <acpi>`,
			`      <table type="slic">/some/data</table>`,
			`    </acpi>`,
			`    <boot dev="hd"></boot>`,
			`    <bios useserial="yes" rebootTimeout="0"></bios>`,
			`    <smbios mode="sysinfo"></smbios>`,
			`  </os>`,
			`  <clock offset="variable" basis="utc" adjustment="28794" timezone="Europe/Paris">`,
			`    <timer name="rtc" track="boot" tickpolicy="catchup" frequency="120" mode="auto">`,
			`      <catchup threshold="123" slew="120" limit="10000"></catchup>`,
			`    </timer>`,
			`  </clock>`,
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
					Enable:  "yes",
					Timeout: "3000",
				},
			},
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`  <os>`,
			`    <nvram template="/t.fd">/vars.fd</nvram>`,
			`    <bootmenu enable="yes" timeout="3000"></bootmenu>`,
			`  </os>`,
			`</domain>`,
		},
	},
	{
		Object: &Domain{
			Type: "kvm",
			Name: "test",
			BlockIOTune: &DomainBlockIOTune{
				Weight: 900,
				Device: []DomainBlockIOTuneDevice{
					DomainBlockIOTuneDevice{
						Path:          "/dev/sda",
						Weight:        500,
						ReadIopsSec:   300,
						WriteIopsSec:  200,
						ReadBytesSec:  3000,
						WriteBytesSec: 2000,
					},
					DomainBlockIOTuneDevice{
						Path:          "/dev/sdb",
						Weight:        600,
						ReadIopsSec:   100,
						WriteIopsSec:  40,
						ReadBytesSec:  1000,
						WriteBytesSec: 400,
					},
				},
			},
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`  <blkiotune>`,
			`    <weight>900</weight>`,
			`    <device>`,
			`      <path>/dev/sda</path>`,
			`      <weight>500</weight>`,
			`      <read_iops_sec>300</read_iops_sec>`,
			`      <write_iops_sec>200</write_iops_sec>`,
			`      <read_bytes_sec>3000</read_bytes_sec>`,
			`      <write_bytes_sec>2000</write_bytes_sec>`,
			`    </device>`,
			`    <device>`,
			`      <path>/dev/sdb</path>`,
			`      <weight>600</weight>`,
			`      <read_iops_sec>100</read_iops_sec>`,
			`      <write_iops_sec>40</write_iops_sec>`,
			`      <read_bytes_sec>1000</read_bytes_sec>`,
			`      <write_bytes_sec>400</write_bytes_sec>`,
			`    </device>`,
			`  </blkiotune>`,
			`</domain>`,
		},
	},
	{
		Object: &Domain{
			Type: "kvm",
			Name: "test",
			MemoryTune: &DomainMemoryTune{
				HardLimit: &DomainMemoryTuneLimit{
					Value: 1024,
					Unit:  "MiB",
				},
				SoftLimit: &DomainMemoryTuneLimit{
					Value: 1024,
				},
				MinGuarantee: &DomainMemoryTuneLimit{
					Value: 1024,
				},
				SwapHardLimit: &DomainMemoryTuneLimit{
					Value: 1024,
				},
			},
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`  <memtune>`,
			`    <hard_limit unit="MiB">1024</hard_limit>`,
			`    <soft_limit>1024</soft_limit>`,
			`    <min_guarantee>1024</min_guarantee>`,
			`    <swap_hard_limit>1024</swap_hard_limit>`,
			`  </memtune>`,
			`</domain>`,
		},
	},
	{
		Object: &Domain{
			Type: "kvm",
			Name: "test",
			PM: &DomainPM{
				SuspendToMem: &DomainPMPolicy{
					Enabled: "no",
				},
				SuspendToDisk: &DomainPMPolicy{
					Enabled: "yes",
				},
			},
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`  <pm>`,
			`    <suspend-to-mem enabled="no"></suspend-to-mem>`,
			`    <suspend-to-disk enabled="yes"></suspend-to-disk>`,
			`  </pm>`,
			`</domain>`,
		},
	},
	{
		Object: &Domain{
			Type: "kvm",
			Name: "test",
			SecLabel: []DomainSecLabel{
				DomainSecLabel{
					Type:       "dynamic",
					Model:      "selinux",
					Relabel:    "yes",
					Label:      "system_u:system_r:svirt_t:s0:c143,c762",
					ImageLabel: "system_u:object_r:svirt_image_t:s0:c143,c762",
					BaseLabel:  "system_u:system_r:svirt_t:s0",
				},
				DomainSecLabel{
					Type:    "dynamic",
					Model:   "dac",
					Relabel: "no",
				},
			},
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`  <seclabel type="dynamic" model="selinux" relabel="yes">`,
			`    <label>system_u:system_r:svirt_t:s0:c143,c762</label>`,
			`    <imagelabel>system_u:object_r:svirt_image_t:s0:c143,c762</imagelabel>`,
			`    <baselabel>system_u:system_r:svirt_t:s0</baselabel>`,
			`  </seclabel>`,
			`  <seclabel type="dynamic" model="dac" relabel="no"></seclabel>`,
			`</domain>`,
		},
	},
	{
		Object: &Domain{
			Type: "kvm",
			Name: "test",
			OS: &DomainOS{
				Kernel:  "/vmlinuz",
				Initrd:  "/initrd",
				Cmdline: "arg",
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
			VCPUs: &DomainVCPUs{
				VCPU: []DomainVCPUsVCPU{
					DomainVCPUsVCPU{
						Id:           &vcpuId0,
						Enabled:      "yes",
						Hotpluggable: "no",
						Order:        &vcpuOrder0,
					},
					DomainVCPUsVCPU{
						Id:           &vcpuId1,
						Enabled:      "no",
						Hotpluggable: "yes",
						Order:        nil,
					},
				},
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
						Virtualport: &DomainInterfaceVirtualport{
							Type: "openvswitch",
						},
					},
				},
			},
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`  <vcpu placement="static" cpuset="1-4,^3,6" current="1">2</vcpu>`,
			`  <vcpus>`,
			`    <vcpu id="0" enabled="yes" hotpluggable="no" order="1"></vcpu>`,
			`    <vcpu id="1" enabled="no" hotpluggable="yes"></vcpu>`,
			`  </vcpus>`,
			`  <devices>`,
			`    <interface type="network">`,
			`      <mac address="00:11:22:33:44:55"></mac>`,
			`      <model type="virtio"></model>`,
			`      <virtualport type="openvswitch"></virtualport>`,
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
				Check: "none",
				Model: &DomainCPUModel{
					Fallback: "allow",
					Value:    "core2duo",
					VendorID: "LibvirtQEMU",
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
				Cache: &DomainCPUCache{
					Level: 1,
					Mode:  "emulate",
				},
				Numa: &DomainNuma{
					[]DomainCell{
						{ID: &cellID, CPUs: "0-3", Memory: "512000", Unit: "KiB", MemAccess: "private"},
					},
				},
			},
			Devices: &DomainDeviceList{
				Emulator: "/bin/qemu-kvm",
			},
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`  <cpu match="exact" check="none">`,
			`    <model fallback="allow" vendor_id="LibvirtQEMU">core2duo</model>`,
			`    <vendor>Intel</vendor>`,
			`    <topology sockets="1" cores="2" threads="1"></topology>`,
			`    <cache level="1" mode="emulate"></cache>`,
			`    <feature policy="disable" name="lahf_lm"></feature>`,
			`    <numa>`,
			`      <cell id="0" cpus="0-3" memory="512000" unit="KiB" memAccess="private"></cell>`,
			`    </numa>`,
			`  </cpu>`,
			`  <devices>`,
			`    <emulator>/bin/qemu-kvm</emulator>`,
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
						Type: "udp",
						MAC: &DomainInterfaceMAC{
							Address: "52:54:00:39:97:ac",
						},
						Model: &DomainInterfaceModel{
							Type: "virtio",
						},
						Source: &DomainInterfaceSource{
							Address: "127.0.0.1",
							Port:    1234,
							Local: &DomainInterfaceSourceLocal{
								Address: "127.0.0.1",
								Port:    1235,
							},
						},
					},
				},
			},
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`  <devices>`,
			`    <interface type="udp">`,
			`      <mac address="52:54:00:39:97:ac"></mac>`,
			`      <model type="virtio"></model>`,
			`      <source address="127.0.0.1" port="1234">`,
			`        <local address="127.0.0.1" port="1235"></local>`,
			`      </source>`,
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
						Type: "direct",
						MAC: &DomainInterfaceMAC{
							Address: "52:54:00:39:97:ac",
						},
						Model: &DomainInterfaceModel{
							Type: "e1000",
						},
						Source: &DomainInterfaceSource{
							Dev:  "eth0",
							Mode: "bridge",
						},
					},
				},
			},
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`  <devices>`,
			`    <interface type="direct">`,
			`      <mac address="52:54:00:39:97:ac"></mac>`,
			`      <model type="e1000"></model>`,
			`      <source dev="eth0" mode="bridge"></source>`,
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
						Boot: &DomainDeviceBoot{
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
						Address: &DomainAddress{
							PCI: &DomainAddressPCI{
								Domain:   &ifaceAddr.Domain,
								Bus:      &ifaceAddr.Bus,
								Slot:     &ifaceAddr.Slot,
								Function: &ifaceAddr.Function,
							},
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
			`      <address type="pci" domain="0x0000" bus="0x00" slot="0x04" function="0x0"></address>`,
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
				Interfaces: []DomainInterface{
					DomainInterface{
						Type: "vhostuser",
						MAC: &DomainInterfaceMAC{
							Address: "52:54:00:39:97:ac",
						},
						Model: &DomainInterfaceModel{
							Type: "virtio",
						},
						Bandwidth: &DomainInterfaceBandwidth{
							Inbound: &DomainInterfaceBandwidthParams{
								Average: &nicAverage,
								Burst:   &nicBurst,
							},
							Outbound: &DomainInterfaceBandwidthParams{
								Average: new(int),
								Burst:   new(int),
							},
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
			`      <bandwidth>`,
			`        <inbound average="1000" burst="10000"></inbound>`,
			`        <outbound average="0" burst="0"></outbound>`,
			`      </bandwidth>`,
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
						Address: &DomainAddress{
							PCI: &DomainAddressPCI{
								Domain:   &fsAddr.Domain,
								Bus:      &fsAddr.Bus,
								Slot:     &fsAddr.Slot,
								Function: &fsAddr.Function,
							},
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
			`      <address type="pci" domain="0x0000" bus="0x00" slot="0x06" function="0x0"></address>`,
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
	{
		Object: &Domain{
			Type: "kvm",
			Name: "test",
			Features: &DomainFeatureList{
				PAE:     &DomainFeature{},
				ACPI:    &DomainFeature{},
				APIC:    &DomainFeatureAPIC{},
				HAP:     &DomainFeatureState{},
				PrivNet: &DomainFeature{},
				HyperV: &DomainFeatureHyperV{
					Relaxed: &DomainFeatureState{State: "on"},
					VAPIC:   &DomainFeatureState{State: "on"},
					Spinlocks: &DomainFeatureHyperVSpinlocks{
						DomainFeatureState{State: "on"}, 4096,
					},
					VPIndex: &DomainFeatureState{State: "on"},
					Runtime: &DomainFeatureState{State: "on"},
					Synic:   &DomainFeatureState{State: "on"},
					Reset:   &DomainFeatureState{State: "on"},
					VendorId: &DomainFeatureHyperVVendorId{
						DomainFeatureState{State: "on"}, "KVM Hv",
					},
				},
				KVM: &DomainFeatureKVM{
					Hidden: &DomainFeatureState{State: "on"},
				},
				PVSpinlock: &DomainFeatureState{State: "on"},
				GIC:        &DomainFeatureGIC{Version: "2"},
			},
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`  <features>`,
			`    <pae></pae>`,
			`    <acpi></acpi>`,
			`    <apic></apic>`,
			`    <hap></hap>`,
			`    <privnet></privnet>`,
			`    <hyperv>`,
			`      <relaxed state="on"></relaxed>`,
			`      <vapic state="on"></vapic>`,
			`      <spinlocks state="on" retries="4096"></spinlocks>`,
			`      <vpindex state="on"></vpindex>`,
			`      <runtime state="on"></runtime>`,
			`      <synic state="on"></synic>`,
			`      <reset state="on"></reset>`,
			`      <vendor_id state="on" value="KVM Hv"></vendor_id>`,
			`    </hyperv>`,
			`    <kvm>`,
			`      <hidden state="on"></hidden>`,
			`    </kvm>`,
			`    <pvspinlock state="on"></pvspinlock>`,
			`    <gic version="2"></gic>`,
			`  </features>`,
			`</domain>`,
		},
	},
	{
		Object: &Domain{
			Type: "kvm",
			Name: "test",
			Devices: &DomainDeviceList{
				Controllers: []DomainController{
					DomainController{
						Type:  "usb",
						Index: &uhciIndex,
						Model: "piix3-uhci",
						Address: &DomainAddress{
							PCI: &DomainAddressPCI{
								Domain:   &uhciAddr.Domain,
								Bus:      &uhciAddr.Bus,
								Slot:     &uhciAddr.Slot,
								Function: &uhciAddr.Function,
							},
						},
					},
					DomainController{
						Type:  "usb",
						Index: nil,
						Model: "ehci",
						USB: &DomainControllerUSB{
							Master: &DomainControllerUSBMaster{
								StartPort: 0,
							},
						},
					},
				},
			},
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`  <devices>`,
			`    <controller type="usb" index="0" model="piix3-uhci">`,
			`      <address type="pci" domain="0x0000" bus="0x00" slot="0x01" function="0x2"></address>`,
			`    </controller>`,
			`    <controller type="usb" model="ehci">`,
			`      <master startport="0"></master>`,
			`    </controller>`,
			`  </devices>`,
			`</domain>`,
		},
	},
	{
		Object: &Domain{
			Type: "kvm",
			Name: "test",
			Devices: &DomainDeviceList{
				Controllers: []DomainController{
					DomainController{
						Type:  "pci",
						Index: &pciIndex,
						Model: "pci-expander-bus",
						PCI: &DomainControllerPCI{
							Model: &DomainControllerPCIModel{
								Name: "pxb",
							},
							Target: &DomainControllerPCITarget{
								ChassisNr: &pciTargetChassisNr,
								Chassis:   &pciTargetChassis,
								Port:      &pciTargetPort,
								BusNr:     &pciTargetBusNr,
								Index:     &pciTargetIndex,
								NUMANode:  &pciTargetNUMANode,
							},
							Hole64: &DomainControllerPCIHole64{
								Size: 1024,
							},
						},
					},
				},
			},
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`  <devices>`,
			`    <controller type="pci" index="0" model="pci-expander-bus">`,
			`      <model name="pxb"></model>`,
			`      <target chassisNr="7" chassis="23" port="78" busNr="2" index="3">`,
			`        <node>2</node>`,
			`      </target>`,
			`      <pcihole64>1024</pcihole64>`,
			`    </controller>`,
			`  </devices>`,
			`</domain>`,
		},
	},
	{
		Object: &Domain{
			Type: "kvm",
			Name: "test",
			Devices: &DomainDeviceList{
				Controllers: []DomainController{
					DomainController{
						Type:  "scsi",
						Index: &scsiIndex,
						Driver: &DomainControllerDriver{
							Queues:     &scsiQueues,
							CmdPerLUN:  &scsiCmdPerLUN,
							MaxSectors: &scsiMaxSectors,
							IOEventFD:  "yes",
							IOThread:   3,
							IOMMU:      "yes",
							ATS:        "no",
						},
					},
				},
			},
		},
		Expected: []string{
			`<domain type="kvm">`,
			`  <name>test</name>`,
			`  <devices>`,
			`    <controller type="scsi" index="0">`,
			`      <driver queues="3" cmd_per_lun="8" max_sectors="512" ioeventfd="yes" iothread="3" iommu="yes" ats="no"></driver>`,
			`    </controller>`,
			`  </devices>`,
			`</domain>`,
		},
	},
	{
		Object: &Domain{
			Type: "qemu",
			Name: "test",
			QEMUCommandline: &DomainQEMUCommandline{
				Args: []DomainQEMUCommandlineArg{
					DomainQEMUCommandlineArg{Value: "-newarg"},
					DomainQEMUCommandlineArg{Value: "-oldarg"},
				},
				Envs: []DomainQEMUCommandlineEnv{
					DomainQEMUCommandlineEnv{Name: "QEMU_ENV", Value: "VAL"},
					DomainQEMUCommandlineEnv{Name: "QEMU_VAR", Value: "VAR"},
				},
			},
		},
		Expected: []string{
			`<domain type="qemu">`,
			`  <name>test</name>`,
			`  <commandline xmlns="http://libvirt.org/schemas/domain/qemu/1.0">`,
			`    <arg value="-newarg"></arg>`,
			`    <arg value="-oldarg"></arg>`,
			`    <env name="QEMU_ENV" value="VAL"></env>`,
			`    <env name="QEMU_VAR" value="VAR"></env>`,
			`  </commandline>`,
			`</domain>`,
		},
	},
	{
		Object: &Domain{
			Name:      "test",
			IOThreads: 4,
			IOThreadIDs: &DomainIOThreadIDs{
				IOThreads: []DomainIOThread{
					DomainIOThread{
						ID: 0,
					},
					DomainIOThread{
						ID: 1,
					},
					DomainIOThread{
						ID: 2,
					},
					DomainIOThread{
						ID: 3,
					},
				},
			},
			CPUTune: &DomainCPUTune{
				Shares: &DomainCPUTuneShares{Value: 1024},
				Period: &DomainCPUTunePeriod{Value: 500000},
				Quota:  &DomainCPUTuneQuota{Value: -1},
			},
		},
		Expected: []string{
			`<domain>`,
			`  <name>test</name>`,
			`  <iothreads>4</iothreads>`,
			`  <iothreadids>`,
			`    <iothread id="0"></iothread>`,
			`    <iothread id="1"></iothread>`,
			`    <iothread id="2"></iothread>`,
			`    <iothread id="3"></iothread>`,
			`  </iothreadids>`,
			`  <cputune>`,
			`    <shares>1024</shares>`,
			`    <period>500000</period>`,
			`    <quota>-1</quota>`,
			`  </cputune>`,
			`</domain>`,
		},
	},
	{
		Object: &Domain{
			Name: "test",
			KeyWrap: &DomainKeyWrap{
				Ciphers: []DomainKeyWrapCipher{
					DomainKeyWrapCipher{
						Name:  "aes",
						State: "on",
					},
					DomainKeyWrapCipher{
						Name:  "dea",
						State: "off",
					},
				},
			},
		},
		Expected: []string{
			`<domain>`,
			`  <name>test</name>`,
			`  <keywrap>`,
			`    <cipher name="aes" state="on"></cipher>`,
			`    <cipher name="dea" state="off"></cipher>`,
			`  </keywrap>`,
			`</domain>`,
		},
	},
	{
		Object: &Domain{
			Name: "test",
			IDMap: &DomainIDMap{
				UIDs: []DomainIDMapRange{
					DomainIDMapRange{
						Start:  0,
						Target: 1000,
						Count:  50,
					},
					DomainIDMapRange{
						Start:  1000,
						Target: 5000,
						Count:  5000,
					},
				},
				GIDs: []DomainIDMapRange{
					DomainIDMapRange{
						Start:  0,
						Target: 1000,
						Count:  50,
					},
					DomainIDMapRange{
						Start:  1000,
						Target: 5000,
						Count:  5000,
					},
				},
			},
		},
		Expected: []string{
			`<domain>`,
			`  <name>test</name>`,
			`  <idmap>`,
			`    <uid start="0" target="1000" count="50"></uid>`,
			`    <uid start="1000" target="5000" count="5000"></uid>`,
			`    <gid start="0" target="1000" count="50"></gid>`,
			`    <gid start="1000" target="5000" count="5000"></gid>`,
			`  </idmap>`,
			`</domain>`,
		},
	},
	{
		Object: &Domain{
			Name: "test",
			NUMATune: &DomainNUMATune{
				Memory: &DomainNUMATuneMemory{
					Mode:      "strict",
					Nodeset:   "2-3",
					Placement: "static",
				},
				MemNodes: []DomainNUMATuneMemNode{
					DomainNUMATuneMemNode{
						CellID:  0,
						Mode:    "strict",
						Nodeset: "2",
					},
					DomainNUMATuneMemNode{
						CellID:  1,
						Mode:    "strict",
						Nodeset: "3",
					},
				},
			},
		},
		Expected: []string{
			`<domain>`,
			`  <name>test</name>`,
			`  <numatune>`,
			`    <memory mode="strict" nodeset="2-3" placement="static"></memory>`,
			`    <memnode cellid="0" mode="strict" nodeset="2"></memnode>`,
			`    <memnode cellid="1" mode="strict" nodeset="3"></memnode>`,

			`  </numatune>`,
			`</domain>`,
		},
	},

	/* Tests for sub-documents that can be hotplugged */
	{
		Object: &DomainController{
			Type:  "usb",
			Index: &uhciIndex,
			Model: "piix3-uhci",
			Address: &DomainAddress{
				PCI: &DomainAddressPCI{
					Domain:   &uhciAddr.Domain,
					Bus:      &uhciAddr.Bus,
					Slot:     &uhciAddr.Slot,
					Function: &uhciAddr.Function,
				},
			},
		},
		Expected: []string{
			`<controller type="usb" index="0" model="piix3-uhci">`,
			`  <address type="pci" domain="0x0000" bus="0x00" slot="0x01" function="0x2"></address>`,
			`</controller>`,
		},
	},
	{
		Object: &DomainDisk{
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
			Serial: "fishfood",
			WWN:    "0123456789abcdef",
		},
		Expected: []string{
			`<disk type="file" device="cdrom">`,
			`  <driver name="qemu" type="qcow2"></driver>`,
			`  <source file="/var/lib/libvirt/images/demo.qcow2"></source>`,
			`  <target dev="vda" bus="virtio"></target>`,
			`  <serial>fishfood</serial>`,
			`  <wwn>0123456789abcdef</wwn>`,
			`</disk>`,
		},
	},
	{
		Object: &DomainFilesystem{
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
			Address: &DomainAddress{
				PCI: &DomainAddressPCI{
					Domain:   &fsAddr.Domain,
					Bus:      &fsAddr.Bus,
					Slot:     &fsAddr.Slot,
					Function: &fsAddr.Function,
				},
			},
		},

		Expected: []string{
			`<filesystem type="mount" accessmode="mapped">`,
			`  <driver type="path" wrpolicy="immediate"></driver>`,
			`  <source dir="/home/user/test"></source>`,
			`  <target dir="user-test-mount"></target>`,
			`  <address type="pci" domain="0x0000" bus="0x00" slot="0x06" function="0x0"></address>`,
			`</filesystem>`,
		},
	},
	{
		Object: &DomainInterface{
			Type: "network",
			MAC: &DomainInterfaceMAC{
				Address: "00:11:22:33:44:55",
			},
			Model: &DomainInterfaceModel{
				Type: "virtio",
			},
		},
		Expected: []string{
			`<interface type="network">`,
			`  <mac address="00:11:22:33:44:55"></mac>`,
			`  <model type="virtio"></model>`,
			`</interface>`,
		},
	},
	{
		Object: &DomainSerial{
			Type: "pty",
			Target: &DomainSerialTarget{
				Type: "isa",
				Port: &serialPort,
			},
			Log: &DomainChardevLog{
				File:   "/some/path",
				Append: "on",
			},
		},

		Expected: []string{
			`<serial type="pty">`,
			`  <target type="isa" port="0"></target>`,
			`  <log file="/some/path" append="on"></log>`,
			`</serial>`,
		},
	},
	{
		Object: &DomainConsole{
			Type: "pty",
			Target: &DomainConsoleTarget{
				Type: "virtio",
				Port: &serialPort,
			},
		},

		Expected: []string{
			`<console type="pty">`,
			`  <target type="virtio" port="0"></target>`,
			`</console>`,
		},
	},
	{
		Object: &DomainInput{
			Type: "tablet",
			Bus:  "usb",
			Address: &DomainAddress{
				USB: &DomainAddressUSB{
					Bus:  &tabletBus,
					Port: tabletPort,
				},
			},
		},

		Expected: []string{
			`<input type="tablet" bus="usb">`,
			`  <address type="usb" bus="0" port="1.1"></address>`,
			`</input>`,
		},
	},
	{
		Object: &DomainVideo{
			Model: DomainVideoModel{
				Type:    "cirrus",
				Heads:   1,
				Ram:     4096,
				VRam:    8192,
				VRam64:  8192,
				VGAMem:  256,
				Primary: "yes",
				Accel: &DomainVideoAccel{
					Accel3D: "yes",
				},
			},
			Address: &DomainAddress{
				PCI: &DomainAddressPCI{
					Domain:        &videoAddr.Domain,
					Bus:           &videoAddr.Bus,
					Slot:          &videoAddr.Slot,
					Function:      &videoAddr.Function,
					MultiFunction: "on",
				},
			},
		},

		Expected: []string{
			`<video>`,
			`  <model type="cirrus" heads="1" ram="4096" vram="8192" vram64="8192" vgamem="256" primary="yes">`,
			`    <acceleration accel3d="yes"></acceleration>`,
			`  </model>`,
			`  <address type="pci" domain="0x0000" bus="0x00" slot="0x05" function="0x0" multifunction="on"></address>`,
			`</video>`,
		},
	},
	{
		Object: &DomainChannel{
			Type: "pty",
			Target: &DomainChannelTarget{
				Type:  "virtio",
				Name:  "org.redhat.spice",
				State: "connected",
			},
		},

		Expected: []string{
			`<channel type="pty">`,
			`  <target type="virtio" name="org.redhat.spice" state="connected"></target>`,
			`</channel>`,
		},
	},
	{
		Object: &DomainMemBalloon{
			Model:       "virtio",
			AutoDeflate: "on",
			Stats: &DomainMemBalloonStats{
				Period: 10,
			},
			Address: &DomainAddress{
				PCI: &DomainAddressPCI{
					Domain:   &balloonAddr.Domain,
					Bus:      &balloonAddr.Bus,
					Slot:     &balloonAddr.Slot,
					Function: &balloonAddr.Function,
				},
			},
		},

		Expected: []string{
			`<memballoon model="virtio" autodeflate="on">`,
			`  <stats period="10"></stats>`,
			`  <address type="pci" domain="0x0000" bus="0x00" slot="0x07" function="0x0"></address>`,
			`</memballoon>`,
		},
	},
	{
		Object: &DomainWatchdog{
			Model:  "ib700",
			Action: "inject-nmi",
			Address: &DomainAddress{
				PCI: &DomainAddressPCI{
					Domain:   &watchdogAddr.Domain,
					Bus:      &watchdogAddr.Bus,
					Slot:     &watchdogAddr.Slot,
					Function: &watchdogAddr.Function,
				},
			},
		},

		Expected: []string{
			`<watchdog model="ib700" action="inject-nmi">`,
			`  <address type="pci" domain="0x0000" bus="0x00" slot="0x08" function="0x0"></address>`,
			`</watchdog>`,
		},
	},
	{
		Object: &DomainSound{
			Model: "ich6",
			Codec: &DomainSoundCodec{
				Type: "duplex",
			},
			Address: &DomainAddress{
				PCI: &DomainAddressPCI{
					Domain:   &duplexAddr.Domain,
					Bus:      &duplexAddr.Bus,
					Slot:     &duplexAddr.Slot,
					Function: &duplexAddr.Function,
				},
			},
		},

		Expected: []string{
			`<sound model="ich6">`,
			`  <codec type="duplex"></codec>`,
			`  <address type="pci" domain="0x0000" bus="0x00" slot="0x08" function="0x0"></address>`,
			`</sound>`,
		},
	},
	{
		Object: &DomainRNG{
			Model: "virtio",
			Rate: &DomainRNGRate{
				Period: 2000,
				Bytes:  1234,
			},
			Backend: &DomainRNGBackend{
				Random: &DomainRNGBackendRandom{
					Device: "/dev/random",
				},
			},
			Address: &DomainAddress{
				PCI: &DomainAddressPCI{
					Domain:   &rngAddr.Domain,
					Bus:      &rngAddr.Bus,
					Slot:     &rngAddr.Slot,
					Function: &rngAddr.Function,
				},
			},
		},

		Expected: []string{
			`<rng model="virtio">`,
			`  <rate bytes="1234" period="2000"></rate>`,
			`  <backend model="random">/dev/random</backend>`,
			`  <address type="pci" domain="0x0000" bus="0x00" slot="0x09" function="0x0"></address>`,
			`</rng>`,
		},
	},
	{
		Object: &DomainRNG{
			Model: "virtio",
			Rate: &DomainRNGRate{
				Period: 2000,
				Bytes:  1234,
			},
			Backend: &DomainRNGBackend{
				EGD: &DomainRNGBackendEGD{
					Type: "udp",
					Sources: []DomainInterfaceSource{
						DomainInterfaceSource{
							Mode:    "bind",
							Service: "1234",
						},
						DomainInterfaceSource{
							Mode:    "connect",
							Host:    "1.2.3.4",
							Service: "1234",
						},
					},
				},
			},
		},

		Expected: []string{
			`<rng model="virtio">`,
			`  <rate bytes="1234" period="2000"></rate>`,
			`  <backend model="egd" type="udp">`,
			`    <source mode="bind" service="1234"></source>`,
			`    <source mode="connect" service="1234" host="1.2.3.4"></source>`,
			`  </backend>`,
			`</rng>`,
		},
	},
	{
		Object: &DomainHostdev{
			SubsysSCSI: &DomainHostdevSubsysSCSI{
				SGIO:  "unfiltered",
				RawIO: "yes",
				Source: &DomainHostdevSubsysSCSISource{
					Host: &DomainHostdevSubsysSCSISourceHost{
						Adapter: &DomainHostdevSubsysSCSIAdapter{
							Name: "scsi_host0",
						},
						Address: &DomainAddressDrive{
							Bus:    &hostdevSCSI.Bus,
							Target: &hostdevSCSI.Target,
							Unit:   &hostdevSCSI.Unit,
						},
					},
				},
			},
			Address: &DomainAddress{
				Drive: &DomainAddressDrive{
					Controller: &hostdevSCSI.Controller,
					Bus:        &hostdevSCSI.Bus,
					Target:     &hostdevSCSI.Target,
					Unit:       &hostdevSCSI.Unit,
				},
			},
		},

		Expected: []string{
			`<hostdev mode="subsystem" type="scsi" sgio="unfiltered" rawio="yes">`,
			`  <source>`,
			`    <adapter name="scsi_host0"></adapter>`,
			`    <address bus="0" target="3" unit="0"></address>`,
			`  </source>`,
			`  <address type="drive" controller="0" bus="0" target="3" unit="0"></address>`,
			`</hostdev>`,
		},
	},
	{
		Object: &DomainHostdev{
			SubsysSCSI: &DomainHostdevSubsysSCSI{
				SGIO:  "unfiltered",
				RawIO: "yes",
				Source: &DomainHostdevSubsysSCSISource{
					ISCSI: &DomainHostdevSubsysSCSISourceISCSI{
						Name: "iqn.1992-01.com.example:storage/1",
						Host: []DomainDiskSourceHost{
							DomainDiskSourceHost{
								Name: "example.org",
								Port: "3260",
							},
						},
						Auth: &DomainDiskAuth{
							Username: "myname",
							Secret: &DomainDiskSecret{
								Type:  "iscsi",
								Usage: "mycluster_myname",
							},
						},
					},
				},
			},
			Address: &DomainAddress{
				Drive: &DomainAddressDrive{
					Controller: &hostdevSCSI.Controller,
					Bus:        &hostdevSCSI.Bus,
					Target:     &hostdevSCSI.Target,
					Unit:       &hostdevSCSI.Unit,
				},
			},
		},

		Expected: []string{
			`<hostdev mode="subsystem" type="scsi" sgio="unfiltered" rawio="yes">`,
			`  <source protocol="iscsi" name="iqn.1992-01.com.example:storage/1">`,
			`    <host name="example.org" port="3260"></host>`,
			`    <auth username="myname">`,
			`      <secret type="iscsi" usage="mycluster_myname"></secret>`,
			`    </auth>`,
			`  </source>`,
			`  <address type="drive" controller="0" bus="0" target="3" unit="0"></address>`,
			`</hostdev>`,
		},
	},
	{
		Object: &DomainHostdev{
			SubsysSCSIHost: &DomainHostdevSubsysSCSIHost{
				Source: &DomainHostdevSubsysSCSIHostSource{
					Protocol: "vhost",
					WWPN:     "naa.5123456789abcde0",
				},
			},
		},

		Expected: []string{
			`<hostdev mode="subsystem" type="scsi_host">`,
			`  <source protocol="vhost" wwpn="naa.5123456789abcde0"></source>`,
			`</hostdev>`,
		},
	},
	{
		Object: &DomainHostdev{
			SubsysUSB: &DomainHostdevSubsysUSB{
				Source: &DomainHostdevSubsysUSBSource{
					Address: &DomainAddressUSB{
						Bus:    &usbHostBus,
						Device: &usbHostDevice,
					},
				},
			},
		},

		Expected: []string{
			`<hostdev mode="subsystem" type="usb">`,
			`  <source>`,
			`    <address bus="14" device="6"></address>`,
			`  </source>`,
			`</hostdev>`,
		},
	},
	{
		Object: &DomainHostdev{
			Managed: "yes",
			SubsysPCI: &DomainHostdevSubsysPCI{
				Source: &DomainHostdevSubsysPCISource{
					Address: &DomainAddressPCI{
						Domain:   &pciHostDomain,
						Bus:      &pciHostBus,
						Slot:     &pciHostSlot,
						Function: &pciHostFunction,
					},
				},
			},
		},

		Expected: []string{
			`<hostdev mode="subsystem" type="pci" managed="yes">`,
			`  <source>`,
			`    <address domain="0x0000" bus="0x03" slot="0x0e" function="0x5"></address>`,
			`  </source>`,
			`</hostdev>`,
		},
	},
	{
		Object: &DomainHostdev{
			SubsysMDev: &DomainHostdevSubsysMDev{
				Model: "vfio-pci",
				Source: &DomainHostdevSubsysMDevSource{
					Address: &DomainAddressMDev{
						UUID: "53764d0e-85a0-42b4-af5c-2046b460b1dc",
					},
				},
			},
		},

		Expected: []string{
			`<hostdev mode="subsystem" type="mdev" model="vfio-pci">`,
			`  <source>`,
			`    <address uuid="53764d0e-85a0-42b4-af5c-2046b460b1dc"></address>`,
			`  </source>`,
			`</hostdev>`,
		},
	},
	{
		Object: &DomainHostdev{
			CapsStorage: &DomainHostdevCapsStorage{
				Source: &DomainHostdevCapsStorageSource{
					Block: "/dev/sda",
				},
			},
		},

		Expected: []string{
			`<hostdev mode="capabilities" type="storage">`,
			`  <source>`,
			`    <block>/dev/sda</block>`,
			`  </source>`,
			`</hostdev>`,
		},
	},
	{
		Object: &DomainHostdev{
			CapsMisc: &DomainHostdevCapsMisc{
				Source: &DomainHostdevCapsMiscSource{
					Char: "/dev/kvm",
				},
			},
		},

		Expected: []string{
			`<hostdev mode="capabilities" type="misc">`,
			`  <source>`,
			`    <char>/dev/kvm</char>`,
			`  </source>`,
			`</hostdev>`,
		},
	},
	{
		Object: &DomainHostdev{
			CapsNet: &DomainHostdevCapsNet{
				Source: &DomainHostdevCapsNetSource{
					Interface: "eth0",
				},
				IP: []DomainIP{
					DomainIP{
						Address: "192.168.122.2",
						Family:  "ipv4",
					},
					DomainIP{
						Address: "2003:db8:1:0:214:1234:fe0b:3596",
						Family:  "ipv6",
						Prefix:  &ipv6Prefix,
					},
				},
				Route: []DomainRoute{
					DomainRoute{
						Family:  "ipv4",
						Address: "0.0.0.0",
						Gateway: "192.168.122.1",
					},
					DomainRoute{
						Family:  "ipv6",
						Address: "::",
						Gateway: "2003:db8:1:0:214:1234:fe0b:3595",
					},
				},
			},
		},

		Expected: []string{
			`<hostdev mode="capabilities" type="net">`,
			`  <source>`,
			`    <interface>eth0</interface>`,
			`  </source>`,
			`  <ip address="192.168.122.2" family="ipv4"></ip>`,
			`  <ip address="2003:db8:1:0:214:1234:fe0b:3596" family="ipv6" prefix="24"></ip>`,
			`  <route family="ipv4" address="0.0.0.0" gateway="192.168.122.1"></route>`,
			`  <route family="ipv6" address="::" gateway="2003:db8:1:0:214:1234:fe0b:3595"></route>`,
			`</hostdev>`,
		},
	},
	{
		Object: &DomainMemorydev{
			Model:  "dimm",
			Access: "private",
			Target: &DomainMemorydevTarget{
				Size: &DomainMemorydevTargetSize{
					Value: 1,
					Unit:  "GiB",
				},
				Node: &DomainMemorydevTargetNode{
					Value: 0,
				},
			},
		},

		Expected: []string{
			`<memory model="dimm" access="private">`,
			`  <target>`,
			`    <size unit="GiB">1</size>`,
			`    <node>0</node>`,
			`  </target>`,
			`</memory>`,
		},
	},
	/* Host Bootloader -- bhyve, Xen */
	{
		Object: &Domain{
			Type:           "bhyve",
			Name:           "test",
			Bootloader:     "/usr/local/sbin/grub-bhyve",
			BootloaderArgs: "-r cd0 -m /tmp/test-device.map -M 1024M linuxguest",
		},
		Expected: []string{
			`<domain type="bhyve">`,
			`  <name>test</name>`,
			`  <bootloader>/usr/local/sbin/grub-bhyve</bootloader>`,
			`  <bootloader_args>-r cd0 -m /tmp/test-device.map -M 1024M linuxguest</bootloader_args>`,
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
			t.Fatal("Bad initial xml:\n", string(doc), "\n does not match\n", expect, "\n")
		}

		typ := reflect.ValueOf(test.Object).Elem().Type()

		newobj := reflect.New(typ)

		newdocobj, ok := newobj.Interface().(Document)
		if !ok {
			t.Fatal("Could not clone %s", newobj.Interface())
		}

		err = newdocobj.Unmarshal(expect)
		if err != nil {
			t.Fatal(err)
		}

		doc, err = newdocobj.Marshal()
		if err != nil {
			t.Fatal(err)
		}

		if doc != expect {
			t.Fatal("Bad roundtrip xml:\n", string(doc), "\n does not match\n", expect, "\n")
		}
	}
}
