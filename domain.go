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
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type DomainController struct {
	XMLName xml.Name       `xml:"controller"`
	Type    string         `xml:"type,attr"`
	Index   *uint          `xml:"index,attr"`
	Model   string         `xml:"model,attr,omitempty"`
	Address *DomainAddress `xml:"address"`
}

type DomainDiskSecret struct {
	Type  string `xml:"type,attr,omitempty"`
	Usage string `xml:"usage,attr,omitempty"`
	UUID  string `xml:"uuid,attr,omitempty"`
}

type DomainDiskAuth struct {
	Username string            `xml:"username,attr,omitempty"`
	Secret   *DomainDiskSecret `xml:"secret"`
}

type DomainDiskSourceHost struct {
	Transport string `xml:"transport,attr,omitempty"`
	Name      string `xml:"name,attr,omitempty"`
	Port      string `xml:"port,attr,omitempty"`
	Socket    string `xml:"socket,attr,omitempty"`
}

type DomainDiskSource struct {
	File          string                 `xml:"file,attr,omitempty"`
	Device        string                 `xml:"dev,attr,omitempty"`
	Protocol      string                 `xml:"protocol,attr,omitempty"`
	Name          string                 `xml:"name,attr,omitempty"`
	Pool          string                 `xml:"pool,attr,omitempty"`
	Volume        string                 `xml:"volume,attr,omitempty"`
	Hosts         []DomainDiskSourceHost `xml:"host"`
	StartupPolicy string                 `xml:"startupPolicy,attr,omitempty"`
}

type DomainDiskDriver struct {
	Name        string `xml:"name,attr,omitempty"`
	Type        string `xml:"type,attr,omitempty"`
	Cache       string `xml:"cache,attr,omitempty"`
	IO          string `xml:"io,attr,omitempty"`
	ErrorPolicy string `xml:"error_policy,attr,omitempty"`
	Discard     string `xml:"discard,attr,omitempty"`
}

type DomainDiskTarget struct {
	Dev string `xml:"dev,attr,omitempty"`
	Bus string `xml:"bus,attr,omitempty"`
}

type DomainDiskReadOnly struct {
}

type DomainDiskShareable struct {
}

type DomainDiskIOTune struct {
	TotalBytesSec          uint64 `xml:"total_bytes_sec"`
	ReadBytesSec           uint64 `xml:"read_bytes_sec"`
	WriteBytesSec          uint64 `xml:"write_bytes_sec"`
	TotalIopsSec           uint64 `xml:"total_iops_sec"`
	ReadIopsSec            uint64 `xml:"read_iops_sec"`
	WriteIopsSec           uint64 `xml:"write_iops_sec"`
	TotalBytesSecMax       uint64 `xml:"total_bytes_sec_max"`
	ReadBytesSecMax        uint64 `xml:"read_bytes_sec_max"`
	WriteBytesSecMax       uint64 `xml:"write_bytes_sec_max"`
	TotalIopsSecMax        uint64 `xml:"total_iops_sec_max"`
	ReadIopsSecMax         uint64 `xml:"read_iops_sec_max"`
	WriteIopsSecMax        uint64 `xml:"write_iops_sec_max"`
	TotalBytesSecMaxLength uint64 `xml:"total_bytes_sec_max_length"`
	ReadBytesSecMaxLength  uint64 `xml:"read_bytes_sec_max_length"`
	WriteBytesSecMaxLength uint64 `xml:"write_bytes_sec_max_length"`
	TotalIopsSecMaxLength  uint64 `xml:"total_iops_sec_max_length"`
	ReadIopsSecMaxLength   uint64 `xml:"read_iops_sec_max_length"`
	WriteIopsSecMaxLength  uint64 `xml:"write_iops_sec_max_length"`
	SizeIopsSec            uint64 `xml:"size_iops_sec"`
	GroupName              string `xml:"group_name"`
}

type DomainDisk struct {
	XMLName   xml.Name             `xml:"disk"`
	Type      string               `xml:"type,attr"`
	Device    string               `xml:"device,attr"`
	Snapshot  string               `xml:"snapshot,attr,omitempty"`
	Driver    *DomainDiskDriver    `xml:"driver"`
	Auth      *DomainDiskAuth      `xml:"auth"`
	Source    *DomainDiskSource    `xml:"source"`
	Target    *DomainDiskTarget    `xml:"target"`
	IOTune    *DomainDiskIOTune    `xml:"iotune"`
	Serial    string               `xml:"serial,omitempty"`
	ReadOnly  *DomainDiskReadOnly  `xml:"readonly"`
	Shareable *DomainDiskShareable `xml:"shareable"`
	Address   *DomainAddress       `xml:"address"`
	Boot      *DomainDeviceBoot    `xml:"boot"`
	WWN       string               `xml:"wwn,omitempty"`
}

type DomainFilesystemDriver struct {
	Type     string `xml:"type,attr"`
	Name     string `xml:"name,attr,omitempty"`
	WRPolicy string `xml:"wrpolicy,attr,omitempty"`
}

type DomainFilesystemSource struct {
	Dir  string `xml:"dir,attr,omitempty"`
	File string `xml:"file,attr,omitempty"`
}

type DomainFilesystemTarget struct {
	Dir string `xml:"dir,attr"`
}

type DomainFilesystemReadOnly struct {
}

type DomainFilesystemSpaceHardLimit struct {
	Value uint   `xml:",chardata"`
	Unit  string `xml:"unit,attr,omitempty"`
}

type DomainFilesystemSpaceSoftLimit struct {
	Value uint   `xml:",chardata"`
	Unit  string `xml:"unit,attr,omitempty"`
}

type DomainFilesystem struct {
	XMLName        xml.Name                        `xml:"filesystem"`
	Type           string                          `xml:"type,attr"`
	AccessMode     string                          `xml:"accessmode,attr"`
	Driver         *DomainFilesystemDriver         `xml:"driver"`
	Source         *DomainFilesystemSource         `xml:"source"`
	Target         *DomainFilesystemTarget         `xml:"target"`
	ReadOnly       *DomainFilesystemReadOnly       `xml:"readonly"`
	SpaceHardLimit *DomainFilesystemSpaceHardLimit `xml:"space_hard_limit"`
	SpaceSoftLimit *DomainFilesystemSpaceSoftLimit `xml:"space_soft_limit"`
	Address        *DomainAddress                  `xml:"address"`
}

type DomainInterfaceMAC struct {
	Address string `xml:"address,attr"`
}

type DomainInterfaceModel struct {
	Type string `xml:"type,attr"`
}

type DomainInterfaceSource struct {
	Bridge  string                      `xml:"bridge,attr,omitempty"`
	Dev     string                      `xml:"dev,attr,omitempty"`
	Network string                      `xml:"network,attr,omitempty"`
	Address string                      `xml:"address,attr,omitempty"`
	Type    string                      `xml:"type,attr,omitempty"`
	Path    string                      `xml:"path,attr,omitempty"`
	Mode    string                      `xml:"mode,attr,omitempty"`
	Port    uint                        `xml:"port,attr,omitempty"`
	Service string                      `xml:"service,attr,omitempty"`
	Host    string                      `xml:"host,attr,omitempty"`
	Local   *DomainInterfaceSourceLocal `xml:"local"`
}

type DomainInterfaceSourceLocal struct {
	Address string `xml:"address,attr,omitempty"`
	Port    uint   `xml:"port,attr,omitempty"`
}

type DomainInterfaceTarget struct {
	Dev string `xml:"dev,attr"`
}

type DomainInterfaceAlias struct {
	Name string `xml:"name,attr"`
}

type DomainInterfaceLink struct {
	State string `xml:"state,attr"`
}

type DomainDeviceBoot struct {
	Order uint `xml:"order,attr"`
}

type DomainInterfaceScript struct {
	Path string `xml:"path,attr"`
}

type DomainInterfaceDriver struct {
	Name   string `xml:"name,attr"`
	Queues uint   `xml:"queues,attr,omitempty"`
}

type DomainInterfaceVirtualport struct {
	Type string `xml:"type,attr"`
}

type DomainInterfaceBandwidthParams struct {
	Average *int `xml:"average,attr,omitempty"`
	Peak    *int `xml:"peak,attr,omitempty"`
	Burst   *int `xml:"burst,attr,omitempty"`
	Floor   *int `xml:"floor,attr,omitempty"`
}

type DomainInterfaceBandwidth struct {
	Inbound  *DomainInterfaceBandwidthParams `xml:"inbound"`
	Outbound *DomainInterfaceBandwidthParams `xml:"outbound"`
}

type DomainInterface struct {
	XMLName     xml.Name                    `xml:"interface"`
	Type        string                      `xml:"type,attr"`
	MAC         *DomainInterfaceMAC         `xml:"mac"`
	Model       *DomainInterfaceModel       `xml:"model"`
	Source      *DomainInterfaceSource      `xml:"source"`
	Target      *DomainInterfaceTarget      `xml:"target"`
	Alias       *DomainInterfaceAlias       `xml:"alias"`
	Link        *DomainInterfaceLink        `xml:"link"`
	Boot        *DomainDeviceBoot           `xml:"boot"`
	Script      *DomainInterfaceScript      `xml:"script"`
	Driver      *DomainInterfaceDriver      `xml:"driver"`
	Virtualport *DomainInterfaceVirtualport `xml:"virtualport"`
	Bandwidth   *DomainInterfaceBandwidth   `xml:"bandwidth"`
	Address     *DomainAddress              `xml:"address"`
}

type DomainChardevSource struct {
	Mode    string `xml:"mode,attr,omitempty"`
	Path    string `xml:"path,attr,omitempty"`
	Append  string `xml:"append,attr,omitempty"`
	Host    string `xml:"host,attr,omitempty"`
	Service string `xml:"service,attr,omitempty"`
	TLS     string `xml:"tls,attr,omitempty"`
}

type DomainChardevTarget struct {
	Type  string `xml:"type,attr,omitempty"`
	Name  string `xml:"name,attr,omitempty"`
	State string `xml:"state,attr,omitempty"` // is guest agent connected?
	Port  *uint  `xml:"port,attr"`
}

type DomainConsoleTarget struct {
	Type string `xml:"type,attr,omitempty"`
	Port *uint  `xml:"port,attr"`
}

type DomainSerialTarget struct {
	Type string `xml:"type,attr,omitempty"`
	Port *uint  `xml:"port,attr"`
}

type DomainChannelTarget struct {
	Type  string `xml:"type,attr,omitempty"`
	Name  string `xml:"name,attr,omitempty"`
	State string `xml:"state,attr,omitempty"` // is guest agent connected?
}

type DomainAlias struct {
	Name string `xml:"name,attr"`
}

type DomainAddressPCI struct {
	Domain        *uint  `xml:"domain,attr"`
	Bus           *uint  `xml:"bus,attr"`
	Slot          *uint  `xml:"slot,attr"`
	Function      *uint  `xml:"function,attr"`
	MultiFunction string `xml:"multifunction,attr,omitempty"`
}

type DomainAddressUSB struct {
	Bus  *uint  `xml:"bus,attr"`
	Port string `xml:"port,attr,omitempty"`
}

type DomainAddressDrive struct {
	Controller *uint `xml:"controller,attr"`
	Bus        *uint `xml:"bus,attr"`
	Target     *uint `xml:"target,attr"`
	Unit       *uint `xml:"unit,attr"`
}

type DomainAddressDIMM struct {
	Slot *uint   `xml:"slot,attr"`
	Base *uint64 `xml:"base,attr"`
}

type DomainAddressISA struct {
	IOBase *uint `xml:"iobase,attr"`
	IRQ    *uint `xml:"irq,attr"`
}

type DomainAddressVirtioMMIO struct {
}

type DomainAddressCCW struct {
	Cssid *uint `xml:"cssid,attr"`
	Ssid  *uint `xml:"ssid,attr"`
	DevNo *uint `xml:"devno,attr"`
}

type DomainAddressVirtioSerial struct {
	Controller *uint `xml:"controller,attr"`
	Bus        *uint `xml:"bus,attr"`
	Port       *uint `xml:"port,attr"`
}

type DomainAddressSpaprVIO struct {
	Reg *uint64 `xml:"reg,attr"`
}

type DomainAddress struct {
	USB          *DomainAddressUSB
	PCI          *DomainAddressPCI
	Drive        *DomainAddressDrive
	DIMM         *DomainAddressDIMM
	ISA          *DomainAddressISA
	VirtioMMIO   *DomainAddressVirtioMMIO
	CCW          *DomainAddressCCW
	VirtioSerial *DomainAddressVirtioSerial
	SpaprVIO     *DomainAddressSpaprVIO
}

type DomainChardevLog struct {
	File   string `xml:"file,attr"`
	Append string `xml:"append,attr,omitempty"`
}

type DomainConsole struct {
	XMLName xml.Name             `xml:"console"`
	Type    string               `xml:"type,attr"`
	Source  *DomainChardevSource `xml:"source"`
	Target  *DomainConsoleTarget `xml:"target"`
	Log     *DomainChardevLog    `xml:"log"`
	Alias   *DomainAlias         `xml:"alias"`
	Address *DomainAddress       `xml:"address"`
}

type DomainSerial struct {
	XMLName  xml.Name              `xml:"serial"`
	Type     string                `xml:"type,attr"`
	Source   *DomainChardevSource  `xml:"source"`
	Protocol *DomainSerialProtocol `xml:"protocol"`
	Target   *DomainSerialTarget   `xml:"target"`
	Log      *DomainChardevLog     `xml:"log"`
	Alias    *DomainAlias          `xml:"alias"`
	Address  *DomainAddress        `xml:"address"`
}

type DomainSerialProtocol struct {
	Type string `xml:"type,attr"`
}

type DomainChannel struct {
	XMLName xml.Name             `xml:"channel"`
	Type    string               `xml:"type,attr"`
	Source  *DomainChardevSource `xml:"source"`
	Target  *DomainChannelTarget `xml:"target"`
	Log     *DomainChardevLog    `xml:"log"`
	Alias   *DomainAlias         `xml:"alias"`
	Address *DomainAddress       `xml:"address"`
}

type DomainInput struct {
	XMLName xml.Name       `xml:"input"`
	Type    string         `xml:"type,attr"`
	Bus     string         `xml:"bus,attr"`
	Address *DomainAddress `xml:"address"`
}

type DomainGraphicListener struct {
	Type    string `xml:"type,attr"`
	Address string `xml:"address,attr,omitempty"`
	Network string `xml:"network,attr,omitempty"`
	Socket  string `xml:"socket,attr,omitempty"`
}

type DomainGraphic struct {
	XMLName       xml.Name                `xml:"graphics"`
	Type          string                  `xml:"type,attr"`
	AutoPort      string                  `xml:"autoport,attr,omitempty"`
	Port          int                     `xml:"port,attr,omitempty"`
	TLSPort       int                     `xml:"tlsPort,attr,omitempty"`
	WebSocket     int                     `xml:"websocket,attr,omitempty"`
	Listen        string                  `xml:"listen,attr,omitempty"`
	Socket        string                  `xml:"socket,attr,omitempty"`
	Keymap        string                  `xml:"keymap,attr,omitempty"`
	Passwd        string                  `xml:"passwd,attr,omitempty"`
	PasswdValidTo string                  `xml:"passwdValidTo,attr,omitempty"`
	Connected     string                  `xml:"connected,attr,omitempty"`
	SharePolicy   string                  `xml:"sharePolicy,attr,omitempty"`
	DefaultMode   string                  `xml:"defaultMode,attr,omitempty"`
	Display       string                  `xml:"display,attr,omitempty"`
	XAuth         string                  `xml:"xauth,attr,omitempty"`
	FullScreen    string                  `xml:"fullscreen,attr,omitempty"`
	ReplaceUser   string                  `xml:"replaceUser,attr,omitempty"`
	MultiUser     string                  `xml:"multiUser,attr,omitempty"`
	Listeners     []DomainGraphicListener `xml:"listen"`
}

type DomainVideoAccel struct {
	Accel3D string `xml:"accel3d,attr,omitempty"`
}

type DomainVideoModel struct {
	Type    string            `xml:"type,attr"`
	Heads   uint              `xml:"heads,attr,omitempty"`
	Ram     uint              `xml:"ram,attr,omitempty"`
	VRam    uint              `xml:"vram,attr,omitempty"`
	VRam64  uint              `xml:"vram64,attr,omitempty"`
	VGAMem  uint              `xml:"vgamem,attr,omitempty"`
	Primary string            `xml:"primary,attr,omitempty"`
	Accel   *DomainVideoAccel `xml:"acceleration"`
}

type DomainVideo struct {
	XMLName xml.Name         `xml:"video"`
	Model   DomainVideoModel `xml:"model"`
	Address *DomainAddress   `xml:"address"`
}

type DomainMemBalloonStats struct {
	Period uint `xml:"period,attr"`
}

type DomainMemBalloon struct {
	XMLName     xml.Name               `xml:"memballoon"`
	Model       string                 `xml:"model,attr"`
	AutoDeflate string                 `xml:"autodeflate,attr,omitempty"`
	Stats       *DomainMemBalloonStats `xml:"stats"`
	Address     *DomainAddress         `xml:"address"`
}

type DomainPanic struct {
	XMLName xml.Name       `xml:"panic"`
	Model   string         `xml:"model,attr"`
	Address *DomainAddress `xml:"address"`
}

type DomainSoundCodec struct {
	Type string `xml:"type,attr"`
}

type DomainSound struct {
	XMLName xml.Name          `xml:"sound"`
	Model   string            `xml:"model,attr"`
	Codec   *DomainSoundCodec `xml:"codec"`
	Address *DomainAddress    `xml:"address"`
}

type DomainRNGRate struct {
	Bytes  uint `xml:"bytes,attr"`
	Period uint `xml:"period,attr,omitempty"`
}

type DomainRNGProtocol struct {
	Type string `xml:"type,attr"`
}

type DomainRNGBackend struct {
	Device   string                  `xml:",chardata"`
	Model    string                  `xml:"model,attr"`
	Type     string                  `xml:"type,attr,omitempty"`
	Sources  []DomainInterfaceSource `xml:"source"`
	Protocol *DomainRNGProtocol      `xml:"protocol"`
}

type DomainRNG struct {
	XMLName xml.Name          `xml:"rng"`
	Model   string            `xml:"model,attr"`
	Rate    *DomainRNGRate    `xml:"rate"`
	Backend *DomainRNGBackend `xml:"backend"`
	Address *DomainAddress    `xml:"address"`
}

type DomainHostdevAdapter struct {
	Name string `xml:"name,attr,omitempty"`
}

type DomainHostdevSource struct {
	Protocol string                `xml:"protocol,attr,omitempty"`
	Name     string                `xml:"name,attr,omitempty"`
	WWPN     string                `xml:"wwpn,attr,omitempty"`
	Adapter  *DomainHostdevAdapter `xml:"adapter"`
	Address  *DomainAddress        `xml:"address"`
}

type DomainHostdev struct {
	XMLName xml.Name             `xml:"hostdev"`
	Mode    string               `xml:"mode,attr"`
	Type    string               `xml:"type,attr"`
	SGIO    string               `xml:"sgio,attr,omitempty"`
	RawIO   string               `xml:"rawio,attr,omitempty"`
	Managed string               `xml:"managed,attr,omitempty"`
	Source  *DomainHostdevSource `xml:"source"`
	Address *DomainAddress       `xml:"address"`
}

type DomainMemorydevTargetNode struct {
	Value uint `xml:",chardata"`
}

type DomainMemorydevTarget struct {
	Size *DomainMemory              `xml:"size"`
	Node *DomainMemorydevTargetNode `xml:"node"`
}

type DomainMemorydev struct {
	XMLName xml.Name               `xml:"memory"`
	Model   string                 `xml:"model,attr"`
	Access  string                 `xml:"access,attr"`
	Target  *DomainMemorydevTarget `xml:"target"`
	Address *DomainAddress         `xml:"address"`
}

type DomainWatchdog struct {
	XMLName xml.Name       `xml:"watchdog"`
	Model   string         `xml:"model,attr"`
	Action  string         `xml:"action,attr,omitempty"`
	Address *DomainAddress `xml:"address"`
}

type DomainDeviceList struct {
	Emulator    string             `xml:"emulator,omitempty"`
	Controllers []DomainController `xml:"controller"`
	Disks       []DomainDisk       `xml:"disk"`
	Filesystems []DomainFilesystem `xml:"filesystem"`
	Interfaces  []DomainInterface  `xml:"interface"`
	Serials     []DomainSerial     `xml:"serial"`
	Consoles    []DomainConsole    `xml:"console"`
	Inputs      []DomainInput      `xml:"input"`
	Graphics    []DomainGraphic    `xml:"graphics"`
	Videos      []DomainVideo      `xml:"video"`
	Channels    []DomainChannel    `xml:"channel"`
	MemBalloon  *DomainMemBalloon  `xml:"memballoon"`
	Panics      []DomainPanic      `xml:"panic"`
	Sounds      []DomainSound      `xml:"sound"`
	RNGs        []DomainRNG        `xml:"rng"`
	Hostdevs    []DomainHostdev    `xml:"hostdev"`
	Memorydevs  []DomainMemorydev  `xml:"memory"`
	Watchdog    *DomainWatchdog    `xml:"watchdog"`
}

type DomainMemory struct {
	Value uint   `xml:",chardata"`
	Unit  string `xml:"unit,attr,omitempty"`
}

type DomainMaxMemory struct {
	Value uint   `xml:",chardata"`
	Unit  string `xml:"unit,attr,omitempty"`
	Slots uint   `xml:"slots,attr,omitempty"`
}

type DomainMemoryHugepage struct {
	Size    uint   `xml:"size,attr"`
	Unit    string `xml:"unit,attr,omitempty"`
	Nodeset string `xml:"nodeset,attr,omitempty"`
}

type DomainMemoryHugepages struct {
	Hugepages []DomainMemoryHugepage `xml:"page"`
}

type DomainMemoryNosharepages struct {
}

type DomainMemoryLocked struct {
}

type DomainMemorySource struct {
	Type string `xml:"type,attr,omitempty"`
}

type DomainMemoryAccess struct {
	Mode string `xml:"mode,attr,omitempty"`
}

type DomainMemoryAllocation struct {
	Mode string `xml:"mode,attr,omitempty"`
}

type DomainMemoryBacking struct {
	MemoryHugePages    *DomainMemoryHugepages    `xml:"hugepages"`
	MemoryNosharepages *DomainMemoryNosharepages `xml:"nosharepages"`
	MemoryLocked       *DomainMemoryLocked       `xml:"locked"`
	MemorySource       *DomainMemorySource       `xml:"source"`
	MemoryAccess       *DomainMemoryAccess       `xml:"access"`
	MemoryAllocation   *DomainMemoryAllocation   `xml:"allocation"`
}

type DomainOSType struct {
	Arch    string `xml:"arch,attr,omitempty"`
	Machine string `xml:"machine,attr,omitempty"`
	Type    string `xml:",chardata"`
}

type DomainSMBios struct {
	Mode string `xml:"mode,attr"`
}

type DomainNVRam struct {
	NVRam    string `xml:",chardata"`
	Template string `xml:"template,attr,omitempty"`
}

type DomainBootDevice struct {
	Dev string `xml:"dev,attr"`
}

type DomainBootMenu struct {
	Enable  string `xml:"enable,attr,omitempty"`
	Timeout string `xml:"timeout,attr,omitempty"`
}

type DomainSysInfo struct {
	Type      string               `xml:"type,attr"`
	System    []DomainSysInfoEntry `xml:"system>entry"`
	BIOS      []DomainSysInfoEntry `xml:"bios>entry"`
	BaseBoard []DomainSysInfoEntry `xml:"baseBoard>entry"`
}

type DomainSysInfoEntry struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

type DomainBIOS struct {
	UseSerial     string `xml:"useserial,attr"`
	RebootTimeout string `xml:"rebootTimeout,attr"`
}

type DomainLoader struct {
	Path     string `xml:",chardata"`
	Readonly string `xml:"readonly,attr,omitempty"`
	Secure   string `xml:"secure,attr,omitempty"`
	Type     string `xml:"type,attr,omitempty"`
}

type DomainACPI struct {
	Tables []DomainACPITable `xml:"table"`
}

type DomainACPITable struct {
	Type string `xml:"type,attr"`
	Path string `xml:",chardata"`
}

type DomainOS struct {
	Type        *DomainOSType      `xml:"type"`
	Init        string             `xml:"init,omitempty"`
	InitArgs    []string           `xml:"initarg"`
	Loader      *DomainLoader      `xml:"loader"`
	NVRam       *DomainNVRam       `xml:"nvram"`
	Kernel      string             `xml:"kernel,omitempty"`
	Initrd      string             `xml:"initrd,omitempty"`
	Cmdline     string             `xml:"cmdline,omitempty"`
	DTB         string             `xml:"dtb,omitempty"`
	ACPI        *DomainACPI        `xml:"acpi"`
	BootDevices []DomainBootDevice `xml:"boot"`
	BootMenu    *DomainBootMenu    `xml:"bootmenu"`
	BIOS        *DomainBIOS        `xml:"bios"`
	SMBios      *DomainSMBios      `xml:"smbios"`
}

type DomainResource struct {
	Partition string `xml:"partition,omitempty"`
}

type DomainVCPU struct {
	Placement string `xml:"placement,attr,omitempty"`
	CPUSet    string `xml:"cpuset,attr,omitempty"`
	Current   string `xml:"current,attr,omitempty"`
	Value     int    `xml:",chardata"`
}

type DomainVCPUsVCPU struct {
	Id           *uint  `xml:"id,attr,omitempty"`
	Enabled      string `xml:"enabled,attr,omitempty"`
	Hotpluggable string `xml:"hotpluggable,attr,omitempty"`
	Order        *uint  `xml:"order,attr,omitempty"`
}

type DomainVCPUs struct {
	VCPU []DomainVCPUsVCPU `xml:"vcpu"`
}

type DomainCPUModel struct {
	Fallback string `xml:"fallback,attr,omitempty"`
	Value    string `xml:",chardata"`
}

type DomainCPUTopology struct {
	Sockets int `xml:"sockets,attr,omitempty"`
	Cores   int `xml:"cores,attr,omitempty"`
	Threads int `xml:"threads,attr,omitempty"`
}

type DomainCPUFeature struct {
	Policy string `xml:"policy,attr,omitempty"`
	Name   string `xml:"name,attr,omitempty"`
}

type DomainCPU struct {
	Match    string             `xml:"match,attr,omitempty"`
	Mode     string             `xml:"mode,attr,omitempty"`
	Check    string             `xml:"check,attr,omitempty"`
	Model    *DomainCPUModel    `xml:"model"`
	Vendor   string             `xml:"vendor,omitempty"`
	Topology *DomainCPUTopology `xml:"topology"`
	Features []DomainCPUFeature `xml:"feature"`
	Numa     *DomainNuma        `xml:"numa,omitempty"`
}

type DomainNuma struct {
	Cell []DomainCell `xml:"cell"`
}

type DomainCell struct {
	ID     string `xml:"id,attr"`
	CPUs   string `xml:"cpus,attr"`
	Memory string `xml:"memory,attr"`
	Unit   string `xml:"unit,attr"`
}

type DomainClock struct {
	Offset     string        `xml:"offset,attr,omitempty"`
	Basis      string        `xml:"basis,attr,omitempty"`
	Adjustment int           `xml:"adjustment,attr,omitempty"`
	Timer      []DomainTimer `xml:"timer,omitempty"`
}

type DomainTimer struct {
	Name       string              `xml:"name,attr"`
	Track      string              `xml:"track,attr,omitempty"`
	TickPolicy string              `xml:"tickpolicy,attr,omitempty"`
	CatchUp    *DomainTimerCatchUp `xml:"catchup,omitempty"`
	Frequency  uint32              `xml:"frequency,attr,omitempty"`
	Mode       string              `xml:"mode,attr,omitempty"`
	Present    string              `xml:"present,attr,omitempty"`
}

type DomainTimerCatchUp struct {
	Threshold uint `xml:"threshold,attr,omitempty"`
	Slew      uint `xml:"slew,attr,omitempty"`
	Limit     uint `xml:"limit,attr,omitempty"`
}

type DomainFeature struct {
}

type DomainFeatureState struct {
	State string `xml:"state,attr,omitempty"`
}

type DomainFeatureAPIC struct {
	EOI string `xml:"eio,attr,omitempty"`
}

type DomainFeatureHyperVVendorId struct {
	DomainFeatureState
	Value string `xml:"value,attr,omitempty"`
}

type DomainFeatureHyperVSpinlocks struct {
	DomainFeatureState
	Retries uint `xml:"retries,attr,omitempty"`
}

type DomainFeatureHyperV struct {
	DomainFeature
	Relaxed   *DomainFeatureState           `xml:"relaxed"`
	VAPIC     *DomainFeatureState           `xml:"vapic"`
	Spinlocks *DomainFeatureHyperVSpinlocks `xml:"spinlocks"`
	VPIndex   *DomainFeatureState           `xml:"vpindex"`
	Runtime   *DomainFeatureState           `xml:"runtime"`
	Synic     *DomainFeatureState           `xml:"synic"`
	STimer    *DomainFeatureState           `xml:"stimer"`
	Reset     *DomainFeatureState           `xml:"reset"`
	VendorId  *DomainFeatureHyperVVendorId  `xml:"vendor_id"`
}

type DomainFeatureKVM struct {
	Hidden *DomainFeatureState `xml:"hidden"`
}

type DomainFeatureGIC struct {
	Version string `xml:"version,attr,omitempty"`
}

type DomainFeatureList struct {
	PAE        *DomainFeature       `xml:"pae"`
	ACPI       *DomainFeature       `xml:"acpi"`
	APIC       *DomainFeatureAPIC   `xml:"apic"`
	HAP        *DomainFeatureState  `xml:"hap"`
	Viridian   *DomainFeature       `xml:"viridian"`
	PrivNet    *DomainFeature       `xml:"privnet"`
	HyperV     *DomainFeatureHyperV `xml:"hyperv"`
	KVM        *DomainFeatureKVM    `xml:"kvm"`
	PVSpinlock *DomainFeatureState  `xml:"pvspinlock"`
	PMU        *DomainFeatureState  `xml:"pmu"`
	VMPort     *DomainFeatureState  `xml:"vmport"`
	GIC        *DomainFeatureGIC    `xml:"gic"`
	SMM        *DomainFeatureState  `xml:"smm"`
}

type DomainCPUTuneShares struct {
	Value uint `xml:",chardata"`
}

type DomainCPUTunePeriod struct {
	Value uint64 `xml:",chardata"`
}

type DomainCPUTuneQuota struct {
	Value int64 `xml:",chardata"`
}

type DomainCPUTune struct {
	Shares *DomainCPUTuneShares `xml:"shares"`
	Period *DomainCPUTunePeriod `xml:"period"`
	Quota  *DomainCPUTuneQuota  `xml:"quota"`
}

type DomainQEMUCommandlineArg struct {
	Value string `xml:"value,attr"`
}

type DomainQEMUCommandlineEnv struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr,omitempty"`
}

type DomainQEMUCommandline struct {
	XMLName xml.Name                   `xml:"http://libvirt.org/schemas/domain/qemu/1.0 commandline"`
	Args    []DomainQEMUCommandlineArg `xml:"arg"`
	Envs    []DomainQEMUCommandlineEnv `xml:"env"`
}

type DomainBlockIOTune struct {
	Weight uint                      `xml:"weight,omitempty"`
	Device []DomainBlockIOTuneDevice `xml:"device"`
}

type DomainBlockIOTuneDevice struct {
	Path          string `xml:"path"`
	Weight        uint   `xml:"weight,omitempty"`
	ReadIopsSec   uint   `xml:"read_iops_sec"`
	WriteIopsSec  uint   `xml:"write_iops_sec"`
	ReadBytesSec  uint   `xml:"read_bytes_sec"`
	WriteBytesSec uint   `xml:"write_bytes_sec"`
}

type DomainPM struct {
	SuspendToMem  *DomainPMPolicy `xml:"suspend-to-mem"`
	SuspendToDisk *DomainPMPolicy `xml:"suspend-to-disk"`
}

type DomainPMPolicy struct {
	Enabled string `xml:"enabled,attr"`
}

type DomainSecLabel struct {
	Type       string `xml:"type,attr,omitempty"`
	Model      string `xml:"model,attr,omitempty"`
	Relabel    string `xml:"relabel,attr,omitempty"`
	Label      string `xml:"label,omitempty"`
	ImageLabel string `xml:"imagelabel,omitempty"`
	BaseLabel  string `xml:"baselabel,omitempty"`
}

type DomainNUMATune struct {
	Memory   *DomainNUMATuneMemory   `xml:"memory"`
	MemNodes []DomainNUMATuneMemNode `xml:"memnode"`
}

type DomainNUMATuneMemory struct {
	Mode      string `xml:"mode,attr"`
	Nodeset   string `xml:"nodeset,attr,omitempty"`
	Placement string `xml:"placement,attr,omitempty"`
}

type DomainNUMATuneMemNode struct {
	CellID  uint   `xml:"cellid,attr"`
	Mode    string `xml:"mode,attr"`
	Nodeset string `xml:"nodeset,attr"`
}

type DomainIOThreadIDs struct {
	IOThreads []DomainIOThread `xml:"iothread"`
}

type DomainIOThread struct {
	ID uint `xml:"id,attr"`
}

type DomainKeyWrap struct {
	Ciphers []DomainKeyWrapCipher `xml:"cipher"`
}

type DomainKeyWrapCipher struct {
	Name  string `xml:"name,attr"`
	State string `xml:"state,attr"`
}

type DomainIDMap struct {
	UIDs []DomainIDMapRange `xml:"uid"`
	GIDs []DomainIDMapRange `xml:"gid"`
}

type DomainIDMapRange struct {
	Start  uint `xml:"start,attr"`
	Target uint `xml:"target,attr"`
	Count  uint `xml:"count,attr"`
}

type DomainMemoryTuneLimit struct {
	Value uint64 `xml:",chardata"`
	Unit  string `xml:"unit,attr,omitempty"`
}

type DomainMemoryTune struct {
	HardLimit     *DomainMemoryTuneLimit `xml:"hard_limit"`
	SoftLimit     *DomainMemoryTuneLimit `xml:"soft_limit"`
	MinGuarantee  *DomainMemoryTuneLimit `xml:"min_guarantee"`
	SwapHardLimit *DomainMemoryTuneLimit `xml:"swap_hard_limit"`
}

// NB, try to keep the order of fields in this struct
// matching the order of XML elements that libvirt
// will generate when dumping XML.
type Domain struct {
	XMLName         xml.Name             `xml:"domain"`
	Type            string               `xml:"type,attr,omitempty"`
	ID              *int                 `xml:"id,attr"`
	Name            string               `xml:"name"`
	UUID            string               `xml:"uuid,omitempty"`
	Title           string               `xml:"title,omitempty"`
	Description     string               `xml:"description,omitempty"`
	MaximumMemory   *DomainMaxMemory     `xml:"maxMemory"`
	Memory          *DomainMemory        `xml:"memory"`
	CurrentMemory   *DomainMemory        `xml:"currentMemory"`
	BlockIOTune     *DomainBlockIOTune   `xml:"blkiotune"`
	MemoryTune      *DomainMemoryTune    `xml:"memtune"`
	MemoryBacking   *DomainMemoryBacking `xml:"memoryBacking"`
	VCPU            *DomainVCPU          `xml:"vcpu"`
	VCPUs           *DomainVCPUs         `xml:"vcpus"`
	IOThreads       uint                 `xml:"iothreads,omitempty"`
	IOThreadIDs     *DomainIOThreadIDs   `xml:"iothreadids"`
	CPUTune         *DomainCPUTune       `xml:"cputune"`
	NUMATune        *DomainNUMATune      `xml:"numatune"`
	Resource        *DomainResource      `xml:"resource"`
	SysInfo         *DomainSysInfo       `xml:"sysinfo"`
	Bootloader      string               `xml:"bootloader,omitempty"`
	BootloaderArgs  string               `xml:"bootloader_args,omitempty"`
	OS              *DomainOS            `xml:"os"`
	IDMap           *DomainIDMap         `xml:"idmap"`
	Features        *DomainFeatureList   `xml:"features"`
	CPU             *DomainCPU           `xml:"cpu"`
	Clock           *DomainClock         `xml:"clock,omitempty"`
	OnPoweroff      string               `xml:"on_poweroff,omitempty"`
	OnReboot        string               `xml:"on_reboot,omitempty"`
	OnCrash         string               `xml:"on_crash,omitempty"`
	PM              *DomainPM            `xml:"pm"`
	Devices         *DomainDeviceList    `xml:"devices"`
	SecLabel        []DomainSecLabel     `xml:"seclabel"`
	QEMUCommandline *DomainQEMUCommandline
	KeyWrap         *DomainKeyWrap `xml:"keywrap"`
}

func (d *Domain) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *Domain) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *DomainController) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainController) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *DomainDisk) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainDisk) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *DomainFilesystem) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainFilesystem) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *DomainInterface) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainInterface) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *DomainConsole) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainConsole) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *DomainSerial) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainSerial) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *DomainInput) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainInput) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *DomainVideo) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainVideo) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *DomainChannel) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainChannel) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *DomainMemBalloon) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainMemBalloon) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *DomainSound) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainSound) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *DomainRNG) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainRNG) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *DomainHostdev) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainHostdev) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *DomainMemorydev) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainMemorydev) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *DomainWatchdog) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainWatchdog) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func marshallUintAttr(start *xml.StartElement, name string, val *uint, format string) {
	if val != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: name}, fmt.Sprintf(format, *val),
		})
	}
}

func marshallUint64Attr(start *xml.StartElement, name string, val *uint64, format string) {
	if val != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: name}, fmt.Sprintf(format, *val),
		})
	}
}

func (a *DomainAddressPCI) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{
		xml.Name{Local: "type"}, "pci",
	})
	marshallUintAttr(&start, "domain", a.Domain, "0x%04x")
	marshallUintAttr(&start, "bus", a.Bus, "0x%02x")
	marshallUintAttr(&start, "slot", a.Slot, "0x%02x")
	marshallUintAttr(&start, "function", a.Function, "0x%x")
	if a.MultiFunction != "" {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "multifunction"}, a.MultiFunction,
		})
	}
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressUSB) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{
		xml.Name{Local: "type"}, "usb",
	})
	marshallUintAttr(&start, "bus", a.Bus, "%d")
	start.Attr = append(start.Attr, xml.Attr{
		xml.Name{Local: "port"}, a.Port,
	})
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressDrive) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{
		xml.Name{Local: "type"}, "drive",
	})
	marshallUintAttr(&start, "controller", a.Controller, "%d")
	marshallUintAttr(&start, "bus", a.Bus, "%d")
	marshallUintAttr(&start, "target", a.Target, "%d")
	marshallUintAttr(&start, "unit", a.Unit, "%d")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressDIMM) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{
		xml.Name{Local: "type"}, "dimm",
	})
	marshallUintAttr(&start, "slot", a.Slot, "%d")
	marshallUint64Attr(&start, "base", a.Base, "0x%x")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressISA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{
		xml.Name{Local: "type"}, "isa",
	})
	marshallUintAttr(&start, "iobase", a.IOBase, "0x%x")
	marshallUintAttr(&start, "irq", a.IRQ, "0x%x")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressVirtioMMIO) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{
		xml.Name{Local: "type"}, "virtio-mmio",
	})
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressCCW) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{
		xml.Name{Local: "type"}, "ccw",
	})
	marshallUintAttr(&start, "cssid", a.Cssid, "0x%x")
	marshallUintAttr(&start, "ssid", a.Ssid, "0x%x")
	marshallUintAttr(&start, "devno", a.DevNo, "0x%04x")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressVirtioSerial) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{
		xml.Name{Local: "type"}, "virtio-serial",
	})
	marshallUintAttr(&start, "controller", a.Controller, "%d")
	marshallUintAttr(&start, "bus", a.Bus, "%d")
	marshallUintAttr(&start, "port", a.Port, "%d")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressSpaprVIO) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{
		xml.Name{Local: "type"}, "spapr-vio",
	})
	marshallUint64Attr(&start, "reg", a.Reg, "%x")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddress) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if a.USB != nil {
		return a.USB.MarshalXML(e, start)
	} else if a.PCI != nil {
		return a.PCI.MarshalXML(e, start)
	} else if a.Drive != nil {
		return a.Drive.MarshalXML(e, start)
	} else if a.DIMM != nil {
		return a.DIMM.MarshalXML(e, start)
	} else if a.ISA != nil {
		return a.ISA.MarshalXML(e, start)
	} else if a.VirtioMMIO != nil {
		return a.VirtioMMIO.MarshalXML(e, start)
	} else if a.CCW != nil {
		return a.CCW.MarshalXML(e, start)
	} else if a.VirtioSerial != nil {
		return a.VirtioSerial.MarshalXML(e, start)
	} else if a.SpaprVIO != nil {
		return a.SpaprVIO.MarshalXML(e, start)
	} else {
		return nil
	}
}

func unmarshallUint64Attr(valstr string, valptr **uint64, base int) error {
	if base == 16 {
		valstr = strings.TrimPrefix(valstr, "0x")
	}
	val, err := strconv.ParseUint(valstr, base, 64)
	if err != nil {
		return err
	}
	*valptr = &val
	return nil
}

func unmarshallUintAttr(valstr string, valptr **uint, base int) error {
	if base == 16 {
		valstr = strings.TrimPrefix(valstr, "0x")
	}
	val, err := strconv.ParseUint(valstr, base, 64)
	if err != nil {
		return err
	}
	vali := uint(val)
	*valptr = &vali
	return nil
}

func (a *DomainAddressUSB) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "bus" {
			if err := unmarshallUintAttr(attr.Value, &a.Bus, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "port" {
			a.Port = attr.Value
		}
	}
	return nil
}

func (a *DomainAddressPCI) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "domain" {
			if err := unmarshallUintAttr(attr.Value, &a.Domain, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "bus" {
			if err := unmarshallUintAttr(attr.Value, &a.Bus, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "slot" {
			if err := unmarshallUintAttr(attr.Value, &a.Slot, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "function" {
			if err := unmarshallUintAttr(attr.Value, &a.Function, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "multifunction" {
			a.MultiFunction = attr.Value
		}
	}
	return nil
}

func (a *DomainAddressDrive) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "controller" {
			if err := unmarshallUintAttr(attr.Value, &a.Controller, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "bus" {
			if err := unmarshallUintAttr(attr.Value, &a.Bus, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "target" {
			if err := unmarshallUintAttr(attr.Value, &a.Target, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "unit" {
			if err := unmarshallUintAttr(attr.Value, &a.Unit, 10); err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *DomainAddressDIMM) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "slot" {
			if err := unmarshallUintAttr(attr.Value, &a.Slot, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "base" {
			if err := unmarshallUint64Attr(attr.Value, &a.Base, 16); err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *DomainAddressISA) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "iobase" {
			if err := unmarshallUintAttr(attr.Value, &a.IOBase, 16); err != nil {
				return err
			}
		} else if attr.Name.Local == "irq" {
			if err := unmarshallUintAttr(attr.Value, &a.IRQ, 16); err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *DomainAddressVirtioMMIO) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return nil
}

func (a *DomainAddressCCW) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "cssid" {
			if err := unmarshallUintAttr(attr.Value, &a.Cssid, 16); err != nil {
				return err
			}
		} else if attr.Name.Local == "ssid" {
			if err := unmarshallUintAttr(attr.Value, &a.Ssid, 16); err != nil {
				return err
			}
		} else if attr.Name.Local == "devno" {
			if err := unmarshallUintAttr(attr.Value, &a.DevNo, 16); err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *DomainAddressVirtioSerial) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "controller" {
			if err := unmarshallUintAttr(attr.Value, &a.Controller, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "bus" {
			if err := unmarshallUintAttr(attr.Value, &a.Bus, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "port" {
			if err := unmarshallUintAttr(attr.Value, &a.Port, 10); err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *DomainAddressSpaprVIO) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "reg" {
			if err := unmarshallUint64Attr(attr.Value, &a.Reg, 16); err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *DomainAddress) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var typ string
	d.Skip()
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			typ = attr.Value
			break
		}
	}
	if typ == "" {
		return nil
	}

	if typ == "usb" {
		a.USB = &DomainAddressUSB{}
		return a.USB.UnmarshalXML(d, start)
	} else if typ == "pci" {
		a.PCI = &DomainAddressPCI{}
		return a.PCI.UnmarshalXML(d, start)
	} else if typ == "drive" {
		a.Drive = &DomainAddressDrive{}
		return a.Drive.UnmarshalXML(d, start)
	} else if typ == "dimm" {
		a.DIMM = &DomainAddressDIMM{}
		return a.DIMM.UnmarshalXML(d, start)
	} else if typ == "isa" {
		a.ISA = &DomainAddressISA{}
		return a.ISA.UnmarshalXML(d, start)
	} else if typ == "virtio-mmio" {
		a.VirtioMMIO = &DomainAddressVirtioMMIO{}
		return a.VirtioMMIO.UnmarshalXML(d, start)
	} else if typ == "ccw" {
		a.CCW = &DomainAddressCCW{}
		return a.CCW.UnmarshalXML(d, start)
	} else if typ == "virtio-serial" {
		a.VirtioSerial = &DomainAddressVirtioSerial{}
		return a.VirtioSerial.UnmarshalXML(d, start)
	} else if typ == "spapr-vio" {
		a.SpaprVIO = &DomainAddressSpaprVIO{}
		return a.SpaprVIO.UnmarshalXML(d, start)
	}

	return nil
}
