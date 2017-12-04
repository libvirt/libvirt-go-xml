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
 * Copyright (C) 2017 Red Hat, Inc.
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

type NodeDevice struct {
	XMLName    xml.Name             `xml:"device"`
	Name       string               `xml:"name"`
	Path       string               `xml:"path,omitempty"`
	DevNodes   []NodeDeviceDevNode  `xml:"devnode"`
	Parent     string               `xml:"parent,omitempty"`
	Driver     *NodeDeviceDriver    `xml:"driver"`
	Capability NodeDeviceCapability `xml:"capability"`
}

type NodeDeviceDevNode struct {
	Type string `xml:"type,attr,omitempty"`
	Path string `xml:",chardata"`
}

type NodeDeviceDriver struct {
	Name string `xml:"name"`
}

type NodeDeviceCapability struct {
	System    *NodeDeviceSystemCapability
	PCI       *NodeDevicePCICapability
	USB       *NodeDeviceUSBCapability
	USBDevice *NodeDeviceUSBDeviceCapability
	Net       *NodeDeviceNetCapability
	SCSIHost  *NodeDeviceSCSIHostCapability
	SCSI      *NodeDeviceSCSICapability
	Storage   *NodeDeviceStorageCapability
	DRM       *NodeDeviceDRMCapability
	CCW       *NodeDeviceCCWCapability
	MDev      *NodeDeviceMDevCapability
}

type NodeDeviceIDName struct {
	ID   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type NodeDevicePCIExpress struct {
	Links []NodeDevicePCIExpressLink `xml:"link"`
}

type NodeDevicePCIExpressLink struct {
	Validity string  `xml:"validity,attr,omitempty"`
	Speed    float64 `xml:"speed,attr,omitempty"`
	Port     int     `xml:"port,attr,omitempty"`
	Width    int     `xml:"width,attr,omitempty"`
}

type NodeDeviceIOMMUGroup struct {
	Number  int                   `xml:"number,attr"`
	Address *NodeDevicePCIAddress `xml:"address"`
}

type NodeDeviceNUMA struct {
	Node int `xml:"node,attr"`
}

type NodeDevicePCICapability struct {
	Domain       *uint                        `xml:"domain,omitempty"`
	Bus          *uint                        `xml:"bus,omitempty"`
	Slot         *uint                        `xml:"slot,omitempty"`
	Function     *uint                        `xml:"function,omitempty"`
	Product      NodeDeviceIDName             `xml:"product,omitempty"`
	Vendor       NodeDeviceIDName             `xml:"vendor,omitempty"`
	IOMMUGroup   *NodeDeviceIOMMUGroup        `xml:"iommuGroup"`
	NUMA         *NodeDeviceNUMA              `xml:"numa"`
	PCIExpress   *NodeDevicePCIExpress        `xml:"pci-express"`
	Capabilities []NodeDevicePCISubCapability `xml:"capability,omitempty"`
}

type NodeDevicePCIAddress struct {
	Domain   string `xml:"domain,attr"`
	Bus      string `xml:"bus,attr"`
	Slot     string `xml:"slot,attr"`
	Function string `xml:"function,attr"`
}

type NodeDevicePCISubCapability struct {
	VirtFunctions *NodeDevicePCIVirtFunctionsCapability
	PhysFunction  *NodeDevicePCIPhysFunctionCapability
	MDevTypes     *NodeDevicePCIMDevTypesCapability
	Bridge        *NodeDevicePCIBridgeCapability
}

type NodeDevicePCIVirtFunctionsCapability struct {
	Address  []NodeDevicePCIAddress `xml:"address,omitempty"`
	MaxCount int                    `xml:"maxCount,attr,omitempty"`
}

type NodeDevicePCIPhysFunctionCapability struct {
	Address NodeDevicePCIAddress `xml:"address,omitempty"`
}

type NodeDevicePCIMDevTypesCapability struct {
	Types []NodeDevicePCIMDevType `xml:"type"`
}

type NodeDevicePCIMDevType struct {
	ID                 string `xml:"id,attr"`
	Name               string `xml:"name"`
	DeviceAPI          string `xml:"deviceAPI"`
	AvailableInstances uint   `xml:"availableInstances"`
}

type NodeDevicePCIBridgeCapability struct {
}

type NodeDeviceSystemHardware struct {
	Vendor  string `xml:"vendor"`
	Version string `xml:"version"`
	Serial  string `xml:"serial"`
	UUID    string `xml:"uuid"`
}

type NodeDeviceSystemFirmware struct {
	Vendor      string `xml:"vendor"`
	Version     string `xml:"version"`
	ReleaseData string `xml:"release_date"`
}

type NodeDeviceSystemCapability struct {
	Product  string                    `xml:"product,omitempty"`
	Hardware *NodeDeviceSystemHardware `xml:"hardware"`
	Firmware *NodeDeviceSystemFirmware `xml:"firmware"`
}

type NodeDeviceUSBDeviceCapability struct {
	Bus     int              `xml:"bus"`
	Device  int              `xml:"device"`
	Product NodeDeviceIDName `xml:"product,omitempty"`
	Vendor  NodeDeviceIDName `xml:"vendor,omitempty"`
}

type NodeDeviceUSBCapability struct {
	Number      int    `xml:"number"`
	Class       int    `xml:"class"`
	Subclass    int    `xml:"subclass"`
	Protocol    int    `xml:"protocol"`
	Description string `xml:"description,omitempty"`
}

type NodeDeviceNetOffloadFeatures struct {
	Name string `xml:"name,attr"`
}

type NodeDeviceNetLink struct {
	State string `xml:"state,attr"`
	Speed string `xml:"speed,attr,omitempty"`
}

type NodeDeviceNetSubCapability struct {
	Type string `xml:"type,attr"`
}

type NodeDeviceNetCapability struct {
	Interface  string                         `xml:"interface"`
	Address    string                         `xml:"address"`
	Link       *NodeDeviceNetLink             `xml:"link"`
	Features   []NodeDeviceNetOffloadFeatures `xml:"feature,omitempty"`
	Capability *NodeDeviceNetSubCapability    `xml:"capability"`
}

type NodeDeviceSCSIVportsOPS struct {
	Vports    int `xml:"vports,omitempty"`
	MaxVports int `xml:"maxvports,,omitempty"`
}

type NodeDeviceSCSIFCHost struct {
	WWNN      string `xml:"wwnn,omitempty"`
	WWPN      string `xml:"wwpn,omitempty"`
	FabricWWN string `xml:"fabric_wwn,omitempty"`
}

type NodeDeviceSCSIHostSubCapability struct {
	VportsOPS *NodeDeviceSCSIVportsOPS `xml:"vports_ops"`
	FCHost    *NodeDeviceSCSIFCHost    `xml:"fc_host"`
}

type NodeDeviceSCSIHostCapability struct {
	Host       int                              `xml:"host"`
	UniqueID   int                              `xml:"unique_id"`
	Capability *NodeDeviceSCSIHostSubCapability `xml:"capability"`
}

type NodeDeviceSCSICapability struct {
	Host   int    `xml:"host"`
	Bus    int    `xml:"bus"`
	Target int    `xml:"target"`
	Lun    int    `xml:"lun"`
	Type   string `xml:"type"`
}

type NodeDeviceStorageSubCapability struct {
	Type           string `xml:"type,attr"`
	MediaAvailable *uint  `xml:"media_available,omitempty"`
	MediaSize      *uint  `xml:"media_size,omitempty"`
	MediaLabel     string `xml:"media_label,omitempty"`
}

type NodeDeviceStorageCapability struct {
	Block        string                          `xml:"block,omitempty"`
	Bus          string                          `xml:"bus,omitempty"`
	DriverType   string                          `xml:"drive_type,omitempty"`
	Model        string                          `xml:"model,omitempty"`
	Vendor       string                          `xml:"vendor,omitempty"`
	Serial       string                          `xml:"serial,omitempty"`
	Size         *uint                           `xml:"size,omitempty"`
	Capatibility *NodeDeviceStorageSubCapability `xml:"capability,omitempty"`
}

type NodeDeviceDRMCapability struct {
	Type string `xml:"type"`
}

type NodeDeviceCCWCapability struct {
	CSSID *uint `xml:"cssid"`
	SSID  *uint `xml:"ssid"`
	DevNo *uint `xml:"devno"`
}

type NodeDeviceMDevCapability struct {
	Type       *NodeDeviceMDevCapabilityType `xml:"type"`
	IOMMUGroup *NodeDeviceIOMMUGroup         `xml:"iommuGroup"`
}

type NodeDeviceMDevCapabilityType struct {
	ID string `xml:"id,attr"`
}

func (c *NodeDeviceCCWCapability) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if c.CSSID != nil {
		cssid := xml.StartElement{
			Name: xml.Name{Local: "cssid"},
		}
		e.EncodeToken(cssid)
		e.EncodeToken(xml.CharData(fmt.Sprintf("0x%x", *c.CSSID)))
		e.EncodeToken(cssid.End())
	}
	if c.SSID != nil {
		ssid := xml.StartElement{
			Name: xml.Name{Local: "ssid"},
		}
		e.EncodeToken(ssid)
		e.EncodeToken(xml.CharData(fmt.Sprintf("0x%x", *c.SSID)))
		e.EncodeToken(ssid.End())
	}
	if c.DevNo != nil {
		devno := xml.StartElement{
			Name: xml.Name{Local: "devno"},
		}
		e.EncodeToken(devno)
		e.EncodeToken(xml.CharData(fmt.Sprintf("0x%04x", *c.DevNo)))
		e.EncodeToken(devno.End())
	}
	e.EncodeToken(start.End())
	return nil
}

func (c *NodeDeviceCCWCapability) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			cdata, err := d.Token()
			if err != nil {
				return err
			}

			if tok.Name.Local != "cssid" &&
				tok.Name.Local != "ssid" &&
				tok.Name.Local != "devno" {
				continue
			}

			chardata, ok := cdata.(xml.CharData)
			if !ok {
				return fmt.Errorf("Expected text for CCW '%s'", tok.Name.Local)
			}

			valstr := strings.TrimPrefix(string(chardata), "0x")
			val, err := strconv.ParseUint(valstr, 16, 64)
			if err != nil {
				return err
			}

			vali := uint(val)
			if tok.Name.Local == "cssid" {
				c.CSSID = &vali
			} else if tok.Name.Local == "ssid" {
				c.SSID = &vali
			} else if tok.Name.Local == "devno" {
				c.DevNo = &vali
			}
		}
	}
	return nil
}

func (c *NodeDevicePCISubCapability) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing node device capability type")
	}

	switch typ {
	case "virt_functions":
		var virtFuncCaps NodeDevicePCIVirtFunctionsCapability
		if err := d.DecodeElement(&virtFuncCaps, &start); err != nil {
			return err
		}
		c.VirtFunctions = &virtFuncCaps
	case "phys_function":
		var physFuncCaps NodeDevicePCIPhysFunctionCapability
		if err := d.DecodeElement(&physFuncCaps, &start); err != nil {
			return err
		}
		c.PhysFunction = &physFuncCaps
	case "mdev_types":
		var mdevTypeCaps NodeDevicePCIMDevTypesCapability
		if err := d.DecodeElement(&mdevTypeCaps, &start); err != nil {
			return err
		}
		c.MDevTypes = &mdevTypeCaps
	case "pci-bridge":
		var bridgeCaps NodeDevicePCIBridgeCapability
		if err := d.DecodeElement(&bridgeCaps, &start); err != nil {
			return err
		}
		c.Bridge = &bridgeCaps
	}
	d.Skip()
	return nil
}

func (c *NodeDevicePCISubCapability) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if c.VirtFunctions != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "virt_functions",
		})
		return e.EncodeElement(c.VirtFunctions, start)
	} else if c.PhysFunction != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "phys_function",
		})
		return e.EncodeElement(c.PhysFunction, start)
	} else if c.MDevTypes != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "mdev_types",
		})
		return e.EncodeElement(c.MDevTypes, start)
	} else if c.Bridge != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "pci-bridge",
		})
		return e.EncodeElement(c.Bridge, start)
	}
	return nil
}

func (c *NodeDeviceCapability) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing node device capability type")
	}

	switch typ {
	case "pci":
		var pciCaps NodeDevicePCICapability
		if err := d.DecodeElement(&pciCaps, &start); err != nil {
			return err
		}
		c.PCI = &pciCaps
	case "system":
		var systemCaps NodeDeviceSystemCapability
		if err := d.DecodeElement(&systemCaps, &start); err != nil {
			return err
		}
		c.System = &systemCaps
	case "usb_device":
		var usbdevCaps NodeDeviceUSBDeviceCapability
		if err := d.DecodeElement(&usbdevCaps, &start); err != nil {
			return err
		}
		c.USBDevice = &usbdevCaps
	case "usb":
		var usbCaps NodeDeviceUSBCapability
		if err := d.DecodeElement(&usbCaps, &start); err != nil {
			return err
		}
		c.USB = &usbCaps
	case "net":
		var netCaps NodeDeviceNetCapability
		if err := d.DecodeElement(&netCaps, &start); err != nil {
			return err
		}
		c.Net = &netCaps
	case "scsi_host":
		var scsiHostCaps NodeDeviceSCSIHostCapability
		if err := d.DecodeElement(&scsiHostCaps, &start); err != nil {
			return err
		}
		c.SCSIHost = &scsiHostCaps
	case "scsi":
		var scsiCaps NodeDeviceSCSICapability
		if err := d.DecodeElement(&scsiCaps, &start); err != nil {
			return err
		}
		c.SCSI = &scsiCaps
	case "storage":
		var storageCaps NodeDeviceStorageCapability
		if err := d.DecodeElement(&storageCaps, &start); err != nil {
			return err
		}
		c.Storage = &storageCaps
	case "drm":
		var drmCaps NodeDeviceDRMCapability
		if err := d.DecodeElement(&drmCaps, &start); err != nil {
			return err
		}
		c.DRM = &drmCaps
	case "ccw":
		var ccwCaps NodeDeviceCCWCapability
		if err := d.DecodeElement(&ccwCaps, &start); err != nil {
			return err
		}
		c.CCW = &ccwCaps
	case "mdev":
		var mdevCaps NodeDeviceMDevCapability
		if err := d.DecodeElement(&mdevCaps, &start); err != nil {
			return err
		}
		c.MDev = &mdevCaps
	}
	d.Skip()
	return nil
}

func (c *NodeDeviceCapability) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if c.PCI != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "pci",
		})
		return e.EncodeElement(c.PCI, start)
	} else if c.System != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "system",
		})
		return e.EncodeElement(c.System, start)
	} else if c.USB != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "usb",
		})
		return e.EncodeElement(c.USB, start)
	} else if c.USBDevice != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "usb_device",
		})
		return e.EncodeElement(c.USBDevice, start)
	} else if c.Net != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "net",
		})
		return e.EncodeElement(c.Net, start)
	} else if c.SCSI != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "scsi",
		})
		return e.EncodeElement(c.SCSI, start)
	} else if c.SCSIHost != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "scsi_host",
		})
		return e.EncodeElement(c.SCSIHost, start)
	} else if c.Storage != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "storage",
		})
		return e.EncodeElement(c.Storage, start)
	} else if c.DRM != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "drm",
		})
		return e.EncodeElement(c.DRM, start)
	} else if c.CCW != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "ccw",
		})
		return e.EncodeElement(c.CCW, start)
	} else if c.MDev != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "mdev",
		})
		return e.EncodeElement(c.MDev, start)
	}
	return nil
}

func (c *NodeDevice) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), c)
}

func (c *NodeDevice) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(c, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}
