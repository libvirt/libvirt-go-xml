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
	"io"
	"strconv"
	"strings"
)

type DomainControllerPCIHole64 struct {
	Size uint64 `xml:",chardata"`
	Unit string `xml:"unit,attr,omitempty"`
}

type DomainControllerPCIModel struct {
	Name string `xml:"name,attr"`
}

type DomainControllerPCITarget struct {
	ChassisNr *uint
	Chassis   *uint
	Port      *uint
	BusNr     *uint
	Index     *uint
	NUMANode  *uint
}

type DomainControllerPCI struct {
	Model  *DomainControllerPCIModel  `xml:"model"`
	Target *DomainControllerPCITarget `xml:"target"`
	Hole64 *DomainControllerPCIHole64 `xml:"pcihole64"`
}

type DomainControllerUSBMaster struct {
	StartPort uint `xml:"startport,attr"`
}

type DomainControllerUSB struct {
	Port   *uint                      `xml:"ports,attr"`
	Master *DomainControllerUSBMaster `xml:"master"`
}

type DomainControllerVirtIOSerial struct {
	Ports   *uint `xml:"ports,attr"`
	Vectors *uint `xml:"vectors,attr"`
}

type DomainControllerDriver struct {
	Queues     *uint  `xml:"queues,attr"`
	CmdPerLUN  *uint  `xml:"cmd_per_lun,attr"`
	MaxSectors *uint  `xml:"max_sectors,attr"`
	IOEventFD  string `xml:"ioeventfd,attr,omitempty"`
	IOThread   uint   `xml:"iothread,attr,omitempty"`
	IOMMU      string `xml:"iommu,attr,omitempty"`
	ATS        string `xml:"ats,attr,omitempty"`
}

type DomainController struct {
	XMLName      xml.Name                      `xml:"controller"`
	Type         string                        `xml:"type,attr"`
	Index        *uint                         `xml:"index,attr"`
	Model        string                        `xml:"model,attr,omitempty"`
	Driver       *DomainControllerDriver       `xml:"driver"`
	Address      *DomainAddress                `xml:"address"`
	PCI          *DomainControllerPCI          `xml:"-"`
	USB          *DomainControllerUSB          `xml:"-"`
	VirtIOSerial *DomainControllerVirtIOSerial `xml:"-"`
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
	Format   string `xml:"format,attr,omitempty"`
	Name     string `xml:"name,attr,omitempty"`
	WRPolicy string `xml:"wrpolicy,attr,omitempty"`
}

type DomainFilesystemSource struct {
	Mount    *DomainFilesystemSourceMount    `xml:"-"`
	Block    *DomainFilesystemSourceBlock    `xml:"-"`
	File     *DomainFilesystemSourceFile     `xml:"-"`
	Template *DomainFilesystemSourceTemplate `xml:"-"`
	RAM      *DomainFilesystemSourceRAM      `xml:"-"`
	Bind     *DomainFilesystemSourceBind     `xml:"-"`
	Volume   *DomainFilesystemSourceVolume   `xml:"-"`
}

type DomainFilesystemSourceMount struct {
	Dir string `xml:"dir,attr"`
}

type DomainFilesystemSourceBlock struct {
	Dev string `xml:"dev,attr"`
}

type DomainFilesystemSourceFile struct {
	File string `xml:"file,attr"`
}

type DomainFilesystemSourceTemplate struct {
	Name string `xml:"name,attr"`
}

type DomainFilesystemSourceRAM struct {
	Usage uint   `xml:"usage,attr"`
	Units string `xml:"units,attr,omitempty"`
}

type DomainFilesystemSourceBind struct {
	Dir string `xml:"dir,attr"`
}

type DomainFilesystemSourceVolume struct {
	Pool   string `xml:"pool,attr"`
	Volume string `xml:"volume,attr"`
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
	AccessMode     string                          `xml:"accessmode,attr,omitempty"`
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
	Bus    *uint  `xml:"bus,attr"`
	Port   string `xml:"port,attr,omitempty"`
	Device *uint  `xml:"device,attr"`
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
	Random *DomainRNGBackendRandom `xml:"-"`
	EGD    *DomainRNGBackendEGD    `xml:"-"`
}

type DomainRNGBackendEGD struct {
	Type     string                  `xml:"type,attr,omitempty"`
	Sources  []DomainInterfaceSource `xml:"source"`
	Protocol *DomainRNGProtocol      `xml:"protocol"`
}

type DomainRNGBackendRandom struct {
	Device string `xml:",chardata"`
}

type DomainRNG struct {
	XMLName xml.Name          `xml:"rng"`
	Model   string            `xml:"model,attr"`
	Rate    *DomainRNGRate    `xml:"rate"`
	Backend *DomainRNGBackend `xml:"backend"`
	Address *DomainAddress    `xml:"address"`
}

type DomainHostdevSubsysUSB struct {
	Source *DomainHostdevSubsysUSBSource `xml:"source"`
}

type DomainHostdevSubsysUSBSource struct {
	Address *DomainAddressUSB `xml:"address"`
}

type DomainHostdevSubsysSCSI struct {
	SGIO      string                         `xml:"sgio,attr,omitempty"`
	RawIO     string                         `xml:"rawio,attr,omitempty"`
	Source    *DomainHostdevSubsysSCSISource `xml:"source"`
	ReadOnly  *DomainDiskReadOnly            `xml:"readonly"`
	Shareable *DomainDiskShareable           `xml:"shareable"`
}

type DomainHostdevSubsysSCSISource struct {
	Host  *DomainHostdevSubsysSCSISourceHost  `xml:"-"`
	ISCSI *DomainHostdevSubsysSCSISourceISCSI `xml:"-"`
}

type DomainHostdevSubsysSCSIAdapter struct {
	Name string `xml:"name,attr"`
}

type DomainHostdevSubsysSCSISourceHost struct {
	Adapter *DomainHostdevSubsysSCSIAdapter `xml:"adapter"`
	Address *DomainAddressDrive             `xml:"address"`
}

type DomainHostdevSubsysSCSISourceISCSI struct {
	Name string                 `xml:"name,attr"`
	Host []DomainDiskSourceHost `xml:"host"`
	Auth *DomainDiskAuth        `xml:"auth"`
}

type DomainHostdevSubsysSCSIHost struct {
	Source *DomainHostdevSubsysSCSIHostSource `xml:"source"`
}

type DomainHostdevSubsysSCSIHostSource struct {
	Protocol string `xml:"protocol,attr,omitempty"`
	WWPN     string `xml:"wwpn,attr,omitempty"`
}

type DomainHostdevSubsysPCISource struct {
	Address *DomainAddressPCI `xml:"address"`
}

type DomainHostdevSubsysPCIDriver struct {
	Name string `xml:"name,attr,omitempty"`
}

type DomainHostdevSubsysPCI struct {
	Driver *DomainHostdevSubsysPCIDriver `xml:"driver"`
	Source *DomainHostdevSubsysPCISource `xml:"source"`
}

type DomainAddressMDev struct {
	UUID string `xml:"uuid,attr"`
}

type DomainHostdevSubsysMDevSource struct {
	Address *DomainAddressMDev `xml:"address"`
}

type DomainHostdevSubsysMDev struct {
	Model  string                         `xml:"model,attr,omitempty"`
	Source *DomainHostdevSubsysMDevSource `xml:"source"`
}

type DomainHostdevCapsStorage struct {
	Source *DomainHostdevCapsStorageSource `xml:"source"`
}

type DomainHostdevCapsStorageSource struct {
	Block string `xml:"block"`
}

type DomainHostdevCapsMisc struct {
	Source *DomainHostdevCapsMiscSource `xml:"source"`
}

type DomainHostdevCapsMiscSource struct {
	Char string `xml:"char"`
}

type DomainIP struct {
	Address string `xml:"address,attr,omitempty"`
	Family  string `xml:"family,attr,omitempty"`
	Prefix  *uint  `xml:"prefix,attr"`
}

type DomainRoute struct {
	Family  string `xml:"family,attr,omitempty"`
	Address string `xml:"address,attr,omitempty"`
	Gateway string `xml:"gateway,attr,omitempty"`
}

type DomainHostdevCapsNet struct {
	Source *DomainHostdevCapsNetSource `xml:"source"`
	IP     []DomainIP                  `xml:"ip"`
	Route  []DomainRoute               `xml:"route"`
}

type DomainHostdevCapsNetSource struct {
	Interface string `xml:"interface"`
}

type DomainHostdev struct {
	Managed        string                       `xml:"managed,attr,omitempty"`
	SubsysUSB      *DomainHostdevSubsysUSB      `xml:"-"`
	SubsysSCSI     *DomainHostdevSubsysSCSI     `xml:"-"`
	SubsysSCSIHost *DomainHostdevSubsysSCSIHost `xml:"-"`
	SubsysPCI      *DomainHostdevSubsysPCI      `xml:"-"`
	SubsysMDev     *DomainHostdevSubsysMDev     `xml:"-"`
	CapsStorage    *DomainHostdevCapsStorage    `xml:"-"`
	CapsMisc       *DomainHostdevCapsMisc       `xml:"-"`
	CapsNet        *DomainHostdevCapsNet        `xml:"-"`
	Boot           *DomainDeviceBoot            `xml:"boot"`
	Address        *DomainAddress               `xml:"address"`
}

type DomainMemorydevSource struct {
	NodeMask string                         `xml:"nodemask,omitempty"`
	PageSize *DomainMemorydevSourcePagesize `xml:"pagesize"`
	Path     string                         `xml:"path,omitempty"`
}

type DomainMemorydevSourcePagesize struct {
	Value uint64 `xml:",chardata"`
	Unit  string `xml:"unit,attr,omitempty"`
}

type DomainMemorydevTargetNode struct {
	Value uint `xml:",chardata"`
}

type DomainMemorydevTargetSize struct {
	Value uint   `xml:",chardata"`
	Unit  string `xml:"unit,attr,omitempty"`
}

type DomainMemorydevTargetLabel struct {
	Size *DomainMemorydevTargetSize `xml:"size"`
}

type DomainMemorydevTarget struct {
	Size  *DomainMemorydevTargetSize  `xml:"size"`
	Node  *DomainMemorydevTargetNode  `xml:"node"`
	Label *DomainMemorydevTargetLabel `xml:"label"`
}

type DomainMemorydev struct {
	XMLName xml.Name               `xml:"memory"`
	Model   string                 `xml:"model,attr"`
	Access  string                 `xml:"access,attr,omitempty"`
	Source  *DomainMemorydevSource `xml:"source"`
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
	Value    uint   `xml:",chardata"`
	Unit     string `xml:"unit,attr,omitempty"`
	DumpCore string `xml:"dumpCore,attr,omitempty"`
}

type DomainCurrentMemory struct {
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
	UseSerial     string `xml:"useserial,attr,omitempty"`
	RebootTimeout *uint  `xml:"rebootTimeout,attr"`
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

type DomainOSInitEnv struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

type DomainOS struct {
	Type        *DomainOSType      `xml:"type"`
	Init        string             `xml:"init,omitempty"`
	InitArgs    []string           `xml:"initarg"`
	InitEnv     []DomainOSInitEnv  `xml:"initenv"`
	InitDir     string             `xml:"initdir,omitempty"`
	InitUser    string             `xml:"inituser,omitempty"`
	InitGroup   string             `xml:"initgroup,omitempty"`
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
	VendorID string `xml:"vendor_id,attr,omitempty"`
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

type DomainCPUCache struct {
	Level uint   `xml:"level,attr,omitempty"`
	Mode  string `xml:"mode,attr"`
}

type DomainCPU struct {
	Match    string             `xml:"match,attr,omitempty"`
	Mode     string             `xml:"mode,attr,omitempty"`
	Check    string             `xml:"check,attr,omitempty"`
	Model    *DomainCPUModel    `xml:"model"`
	Vendor   string             `xml:"vendor,omitempty"`
	Topology *DomainCPUTopology `xml:"topology"`
	Cache    *DomainCPUCache    `xml:"cache"`
	Features []DomainCPUFeature `xml:"feature"`
	Numa     *DomainNuma        `xml:"numa,omitempty"`
}

type DomainNuma struct {
	Cell []DomainCell `xml:"cell"`
}

type DomainCell struct {
	ID        *uint  `xml:"id,attr"`
	CPUs      string `xml:"cpus,attr"`
	Memory    string `xml:"memory,attr"`
	Unit      string `xml:"unit,attr,omitempty"`
	MemAccess string `xml:"memAccess,attr,omitempty"`
}

type DomainClock struct {
	Offset     string        `xml:"offset,attr,omitempty"`
	Basis      string        `xml:"basis,attr,omitempty"`
	Adjustment int           `xml:"adjustment,attr,omitempty"`
	TimeZone   string        `xml:"timezone,attr,omitempty"`
	Timer      []DomainTimer `xml:"timer"`
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
	EOI string `xml:"eoi,attr,omitempty"`
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

type DomainFeatureIOAPIC struct {
	Driver string `xml:"driver,attr,omitempty"`
}

type DomainFeatureHPT struct {
	Resizing string `xml:"resizing,attr,omitempty"`
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
	IOAPIC     *DomainFeatureIOAPIC `xml:"ioapic"`
	HPT        *DomainFeatureHPT    `xml:"hpt"`
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
	Mode      string `xml:"mode,attr,omitempty"`
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
	CurrentMemory   *DomainCurrentMemory `xml:"currentMemory"`
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

type domainController DomainController

type domainControllerPCI struct {
	DomainControllerPCI
	domainController
}

type domainControllerUSB struct {
	DomainControllerUSB
	domainController
}

type domainControllerVirtIOSerial struct {
	DomainControllerVirtIOSerial
	domainController
}

func (a *DomainControllerPCITarget) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshallUintAttr(&start, "chassisNr", a.ChassisNr, "%d")
	marshallUintAttr(&start, "chassis", a.Chassis, "%d")
	marshallUintAttr(&start, "port", a.Port, "%d")
	marshallUintAttr(&start, "busNr", a.BusNr, "%d")
	marshallUintAttr(&start, "index", a.Index, "%d")
	e.EncodeToken(start)
	if a.NUMANode != nil {
		node := xml.StartElement{
			Name: xml.Name{Local: "node"},
		}
		e.EncodeToken(node)
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *a.NUMANode)))
		e.EncodeToken(node.End())
	}
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainControllerPCITarget) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "chassisNr" {
			if err := unmarshallUintAttr(attr.Value, &a.ChassisNr, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "chassis" {
			if err := unmarshallUintAttr(attr.Value, &a.Chassis, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "port" {
			if err := unmarshallUintAttr(attr.Value, &a.Port, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "busNr" {
			if err := unmarshallUintAttr(attr.Value, &a.BusNr, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "index" {
			if err := unmarshallUintAttr(attr.Value, &a.Index, 10); err != nil {
				return err
			}
		}
	}
	for {
		tok, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			if tok.Name.Local == "node" {
				data, err := d.Token()
				if err != nil {
					return err
				}
				switch data := data.(type) {
				case xml.CharData:
					val, err := strconv.ParseUint(string(data), 10, 64)
					if err != nil {
						return err
					}
					vali := uint(val)
					a.NUMANode = &vali
				}
			}
		}
	}
	return nil
}

func (a *DomainController) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "controller"
	if a.Type == "pci" {
		pci := domainControllerPCI{}
		pci.domainController = domainController(*a)
		if a.PCI != nil {
			pci.DomainControllerPCI = *a.PCI
		}
		return e.EncodeElement(pci, start)
	} else if a.Type == "usb" {
		usb := domainControllerUSB{}
		usb.domainController = domainController(*a)
		if a.USB != nil {
			usb.DomainControllerUSB = *a.USB
		}
		return e.EncodeElement(usb, start)
	} else if a.Type == "virtio-serial" {
		vioserial := domainControllerVirtIOSerial{}
		vioserial.domainController = domainController(*a)
		if a.VirtIOSerial != nil {
			vioserial.DomainControllerVirtIOSerial = *a.VirtIOSerial
		}
		return e.EncodeElement(vioserial, start)
	} else {
		gen := domainController(*a)
		return e.EncodeElement(gen, start)
	}
}

func getAttr(attrs []xml.Attr, name string) (string, bool) {
	for _, attr := range attrs {
		if attr.Name.Local == name {
			return attr.Value, true
		}
	}
	return "", false
}

func (a *DomainController) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing 'type' attribute on domain controller")
	}
	if typ == "pci" {
		var pci domainControllerPCI
		err := d.DecodeElement(&pci, &start)
		if err != nil {
			return err
		}
		*a = DomainController(pci.domainController)
		a.PCI = &pci.DomainControllerPCI
		return nil
	} else if typ == "usb" {
		var usb domainControllerUSB
		err := d.DecodeElement(&usb, &start)
		if err != nil {
			return err
		}
		*a = DomainController(usb.domainController)
		a.USB = &usb.DomainControllerUSB
		return nil
	} else if typ == "virtio-serial" {
		var vioserial domainControllerVirtIOSerial
		err := d.DecodeElement(&vioserial, &start)
		if err != nil {
			return err
		}
		*a = DomainController(vioserial.domainController)
		a.VirtIOSerial = &vioserial.DomainControllerVirtIOSerial
		return nil
	} else {
		var gen domainController
		err := d.DecodeElement(&gen, &start)
		if err != nil {
			return err
		}
		*a = DomainController(gen)
		return nil
	}
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

func (a *DomainFilesystemSource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if a.Mount != nil {
		return e.EncodeElement(a.Mount, start)
	} else if a.Block != nil {
		return e.EncodeElement(a.Block, start)
	} else if a.File != nil {
		return e.EncodeElement(a.File, start)
	} else if a.Template != nil {
		return e.EncodeElement(a.Template, start)
	} else if a.RAM != nil {
		return e.EncodeElement(a.RAM, start)
	} else if a.Bind != nil {
		return e.EncodeElement(a.Bind, start)
	} else if a.Volume != nil {
		return e.EncodeElement(a.Volume, start)
	}
	return nil
}

func (a *DomainFilesystemSource) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if a.Mount != nil {
		return d.DecodeElement(a.Mount, &start)
	} else if a.Block != nil {
		return d.DecodeElement(a.Block, &start)
	} else if a.File != nil {
		return d.DecodeElement(a.File, &start)
	} else if a.Template != nil {
		return d.DecodeElement(a.Template, &start)
	} else if a.RAM != nil {
		return d.DecodeElement(a.RAM, &start)
	} else if a.Bind != nil {
		return d.DecodeElement(a.Bind, &start)
	} else if a.Volume != nil {
		return d.DecodeElement(a.Volume, &start)
	}
	return nil
}

type domainFilesystem DomainFilesystem

func (a *DomainFilesystem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "filesystem"
	if a.Source != nil {
		if a.Source.Mount != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "mount",
			})
		} else if a.Source.Block != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "block",
			})
		} else if a.Source.File != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "file",
			})
		} else if a.Source.Template != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "template",
			})
		} else if a.Source.RAM != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "ram",
			})
		} else if a.Source.Bind != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "bind",
			})
		} else if a.Source.Volume != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "volume",
			})
		}
	}
	fs := domainFilesystem(*a)
	return e.EncodeElement(fs, start)
}

func (a *DomainFilesystem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		typ = "mount"
	}
	a.Source = &DomainFilesystemSource{}
	if typ == "mount" {
		a.Source.Mount = &DomainFilesystemSourceMount{}
	} else if typ == "block" {
		a.Source.Block = &DomainFilesystemSourceBlock{}
	} else if typ == "file" {
		a.Source.File = &DomainFilesystemSourceFile{}
	} else if typ == "template" {
		a.Source.Template = &DomainFilesystemSourceTemplate{}
	} else if typ == "ram" {
		a.Source.RAM = &DomainFilesystemSourceRAM{}
	} else if typ == "bind" {
		a.Source.Bind = &DomainFilesystemSourceBind{}
	} else if typ == "volume" {
		a.Source.Volume = &DomainFilesystemSourceVolume{}
	}
	fs := domainFilesystem(*a)
	err := d.DecodeElement(&fs, &start)
	if err != nil {
		return err
	}
	*a = DomainFilesystem(fs)
	return nil
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

func (a *DomainRNGBackend) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if a.Random != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "model"}, "random",
		})
		return e.EncodeElement(a.Random, start)
	} else if a.EGD != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "model"}, "egd",
		})
		return e.EncodeElement(a.EGD, start)
	}
	return nil
}

func (a *DomainRNGBackend) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	model, ok := getAttr(start.Attr, "model")
	if !ok {
		return nil
	}
	if model == "random" {
		a.Random = &DomainRNGBackendRandom{}
		err := d.DecodeElement(a.Random, &start)
		if err != nil {
			return err
		}
	} else if model == "egd" {
		a.EGD = &DomainRNGBackendEGD{}
		err := d.DecodeElement(a.EGD, &start)
		if err != nil {
			return err
		}
	}
	d.Skip()
	return nil
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

func (a *DomainHostdevSubsysSCSISource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if a.Host != nil {
		return e.EncodeElement(a.Host, start)
	} else if a.ISCSI != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "protocol"}, "iscsi",
		})
		return e.EncodeElement(a.ISCSI, start)
	}
	return nil
}

func (a *DomainHostdevSubsysSCSISource) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	proto, ok := getAttr(start.Attr, "protocol")
	if !ok {
		a.Host = &DomainHostdevSubsysSCSISourceHost{}
		err := d.DecodeElement(a.Host, &start)
		if err != nil {
			return err
		}
	}
	if proto == "iscsi" {
		a.ISCSI = &DomainHostdevSubsysSCSISourceISCSI{}
		err := d.DecodeElement(a.ISCSI, &start)
		if err != nil {
			return err
		}
	}
	d.Skip()
	return nil
}

type domainHostdev DomainHostdev

type domainHostdevSubsysSCSI struct {
	DomainHostdevSubsysSCSI
	domainHostdev
}

type domainHostdevSubsysSCSIHost struct {
	DomainHostdevSubsysSCSIHost
	domainHostdev
}

type domainHostdevSubsysUSB struct {
	DomainHostdevSubsysUSB
	domainHostdev
}

type domainHostdevSubsysPCI struct {
	DomainHostdevSubsysPCI
	domainHostdev
}

type domainHostdevSubsysMDev struct {
	DomainHostdevSubsysMDev
	domainHostdev
}

type domainHostdevCapsStorage struct {
	DomainHostdevCapsStorage
	domainHostdev
}

type domainHostdevCapsMisc struct {
	DomainHostdevCapsMisc
	domainHostdev
}

type domainHostdevCapsNet struct {
	DomainHostdevCapsNet
	domainHostdev
}

func (a *DomainHostdev) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "hostdev"
	if a.SubsysSCSI != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "mode"}, "subsystem",
		})
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "scsi",
		})
		scsi := domainHostdevSubsysSCSI{}
		scsi.domainHostdev = domainHostdev(*a)
		scsi.DomainHostdevSubsysSCSI = *a.SubsysSCSI
		return e.EncodeElement(scsi, start)
	} else if a.SubsysSCSIHost != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "mode"}, "subsystem",
		})
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "scsi_host",
		})
		scsi_host := domainHostdevSubsysSCSIHost{}
		scsi_host.domainHostdev = domainHostdev(*a)
		scsi_host.DomainHostdevSubsysSCSIHost = *a.SubsysSCSIHost
		return e.EncodeElement(scsi_host, start)
	} else if a.SubsysUSB != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "mode"}, "subsystem",
		})
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "usb",
		})
		usb := domainHostdevSubsysUSB{}
		usb.domainHostdev = domainHostdev(*a)
		usb.DomainHostdevSubsysUSB = *a.SubsysUSB
		return e.EncodeElement(usb, start)
	} else if a.SubsysPCI != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "mode"}, "subsystem",
		})
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "pci",
		})
		pci := domainHostdevSubsysPCI{}
		pci.domainHostdev = domainHostdev(*a)
		pci.DomainHostdevSubsysPCI = *a.SubsysPCI
		return e.EncodeElement(pci, start)
	} else if a.SubsysMDev != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "mode"}, "subsystem",
		})
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "mdev",
		})
		mdev := domainHostdevSubsysMDev{}
		mdev.domainHostdev = domainHostdev(*a)
		mdev.DomainHostdevSubsysMDev = *a.SubsysMDev
		return e.EncodeElement(mdev, start)
	} else if a.CapsStorage != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "mode"}, "capabilities",
		})
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "storage",
		})
		storage := domainHostdevCapsStorage{}
		storage.domainHostdev = domainHostdev(*a)
		storage.DomainHostdevCapsStorage = *a.CapsStorage
		return e.EncodeElement(storage, start)
	} else if a.CapsMisc != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "mode"}, "capabilities",
		})
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "misc",
		})
		misc := domainHostdevCapsMisc{}
		misc.domainHostdev = domainHostdev(*a)
		misc.DomainHostdevCapsMisc = *a.CapsMisc
		return e.EncodeElement(misc, start)
	} else if a.CapsNet != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "mode"}, "capabilities",
		})
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "net",
		})
		net := domainHostdevCapsNet{}
		net.domainHostdev = domainHostdev(*a)
		net.DomainHostdevCapsNet = *a.CapsNet
		return e.EncodeElement(net, start)
	} else {
		gen := domainHostdev(*a)
		return e.EncodeElement(gen, start)
	}
}

func (a *DomainHostdev) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	mode, ok := getAttr(start.Attr, "mode")
	if !ok {
		return fmt.Errorf("Missing 'mode' attribute on domain hostdev")
	}
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing 'type' attribute on domain controller")
	}
	if mode == "subsystem" {
		if typ == "scsi" {
			var scsi domainHostdevSubsysSCSI
			err := d.DecodeElement(&scsi, &start)
			if err != nil {
				return err
			}
			*a = DomainHostdev(scsi.domainHostdev)
			a.SubsysSCSI = &scsi.DomainHostdevSubsysSCSI
			return nil
		} else if typ == "scsi_host" {
			var scsi_host domainHostdevSubsysSCSIHost
			err := d.DecodeElement(&scsi_host, &start)
			if err != nil {
				return err
			}
			*a = DomainHostdev(scsi_host.domainHostdev)
			a.SubsysSCSIHost = &scsi_host.DomainHostdevSubsysSCSIHost
			return nil
		} else if typ == "usb" {
			var usb domainHostdevSubsysUSB
			err := d.DecodeElement(&usb, &start)
			if err != nil {
				return err
			}
			*a = DomainHostdev(usb.domainHostdev)
			a.SubsysUSB = &usb.DomainHostdevSubsysUSB
			return nil
		} else if typ == "pci" {
			var pci domainHostdevSubsysPCI
			err := d.DecodeElement(&pci, &start)
			if err != nil {
				return err
			}
			*a = DomainHostdev(pci.domainHostdev)
			a.SubsysPCI = &pci.DomainHostdevSubsysPCI
			return nil
		} else if typ == "mdev" {
			var mdev domainHostdevSubsysMDev
			err := d.DecodeElement(&mdev, &start)
			if err != nil {
				return err
			}
			*a = DomainHostdev(mdev.domainHostdev)
			a.SubsysMDev = &mdev.DomainHostdevSubsysMDev
			return nil
		}
	} else if mode == "capabilities" {
		if typ == "storage" {
			var storage domainHostdevCapsStorage
			err := d.DecodeElement(&storage, &start)
			if err != nil {
				return err
			}
			*a = DomainHostdev(storage.domainHostdev)
			a.CapsStorage = &storage.DomainHostdevCapsStorage
			return nil
		} else if typ == "misc" {
			var misc domainHostdevCapsMisc
			err := d.DecodeElement(&misc, &start)
			if err != nil {
				return err
			}
			*a = DomainHostdev(misc.domainHostdev)
			a.CapsMisc = &misc.DomainHostdevCapsMisc
			return nil
		} else if typ == "net" {
			var net domainHostdevCapsNet
			err := d.DecodeElement(&net, &start)
			if err != nil {
				return err
			}
			*a = DomainHostdev(net.domainHostdev)
			a.CapsNet = &net.DomainHostdevCapsNet
			return nil
		}
	}
	var gen domainHostdev
	err := d.DecodeElement(&gen, &start)
	if err != nil {
		return err
	}
	*a = DomainHostdev(gen)
	return nil
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
	marshallUintAttr(&start, "bus", a.Bus, "%d")
	if a.Port != "" {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "port"}, a.Port,
		})
	}
	marshallUintAttr(&start, "device", a.Device, "%d")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressDrive) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshallUintAttr(&start, "controller", a.Controller, "%d")
	marshallUintAttr(&start, "bus", a.Bus, "%d")
	marshallUintAttr(&start, "target", a.Target, "%d")
	marshallUintAttr(&start, "unit", a.Unit, "%d")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressDIMM) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshallUintAttr(&start, "slot", a.Slot, "%d")
	marshallUint64Attr(&start, "base", a.Base, "0x%x")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressISA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshallUintAttr(&start, "iobase", a.IOBase, "0x%x")
	marshallUintAttr(&start, "irq", a.IRQ, "0x%x")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressVirtioMMIO) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressCCW) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshallUintAttr(&start, "cssid", a.Cssid, "0x%x")
	marshallUintAttr(&start, "ssid", a.Ssid, "0x%x")
	marshallUintAttr(&start, "devno", a.DevNo, "0x%04x")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressVirtioSerial) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshallUintAttr(&start, "controller", a.Controller, "%d")
	marshallUintAttr(&start, "bus", a.Bus, "%d")
	marshallUintAttr(&start, "port", a.Port, "%d")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressSpaprVIO) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshallUint64Attr(&start, "reg", a.Reg, "0x%x")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddress) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if a.USB != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "usb",
		})
		return a.USB.MarshalXML(e, start)
	} else if a.PCI != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "pci",
		})
		return a.PCI.MarshalXML(e, start)
	} else if a.Drive != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "drive",
		})
		return a.Drive.MarshalXML(e, start)
	} else if a.DIMM != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "dimm",
		})
		return a.DIMM.MarshalXML(e, start)
	} else if a.ISA != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "isa",
		})
		return a.ISA.MarshalXML(e, start)
	} else if a.VirtioMMIO != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "virtio-mmio",
		})
		return a.VirtioMMIO.MarshalXML(e, start)
	} else if a.CCW != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "ccw",
		})
		return a.CCW.MarshalXML(e, start)
	} else if a.VirtioSerial != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "virtio-serial",
		})
		return a.VirtioSerial.MarshalXML(e, start)
	} else if a.SpaprVIO != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "spapr-vio",
		})
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
		} else if attr.Name.Local == "device" {
			if err := unmarshallUintAttr(attr.Value, &a.Device, 10); err != nil {
				return err
			}
		}
	}
	d.Skip()
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
	d.Skip()
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
	d.Skip()
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
	d.Skip()
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
	d.Skip()
	return nil
}

func (a *DomainAddressVirtioMMIO) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	d.Skip()
	return nil
}

func (a *DomainAddressCCW) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "cssid" {
			if err := unmarshallUintAttr(attr.Value, &a.Cssid, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "ssid" {
			if err := unmarshallUintAttr(attr.Value, &a.Ssid, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "devno" {
			if err := unmarshallUintAttr(attr.Value, &a.DevNo, 0); err != nil {
				return err
			}
		}
	}
	d.Skip()
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
	d.Skip()
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
	d.Skip()
	return nil
}

func (a *DomainAddress) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var typ string
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			typ = attr.Value
			break
		}
	}
	if typ == "" {
		d.Skip()
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
