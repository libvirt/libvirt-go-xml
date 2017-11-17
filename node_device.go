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
)

type NodeDevice struct {
	XMLName    xml.Name             `xml:"device"`
	Name       string               `xml:"name"`
	Path       string               `xml:"path,omitempty"`
	Parent     string               `xml:"parent,omitempty"`
	Driver     string               `xml:"driver>name,omitempty"`
	Capability NodeDeviceCapability `xml:"capability"`
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
	Number int `xml:"number,attr"`
}

type NodeDeviceNUMA struct {
	Node int `xml:"node,attr"`
}

type NodeDevicePCICapability struct {
	Domain       int                          `xml:"domain,omitempty"`
	Bus          int                          `xml:"bus,omitempty"`
	Slot         int                          `xml:"slot,omitempty"`
	Function     int                          `xml:"function,omitempty"`
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
	Type     string                 `xml:"type,attr"`
	Address  []NodeDevicePCIAddress `xml:"address,omitempty"`
	MaxCount int                    `xml:"maxCount,attr,omitempty"`
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
	Product  string                    `xml:"product"`
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
	Name string `xml:"number"`
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
	Type           string `xml:"match,attr"`
	MediaAvailable int    `xml:"media_available,omitempty"`
	MediaSize      int    `xml:"media_size,omitempty"`
	MediaLable     int    `xml:"media_label,omitempty"`
}

type NodeDeviceStorageCapability struct {
	Block        string                          `xml:"block"`
	Bus          string                          `xml:"bus"`
	DriverType   string                          `xml:"drive_type"`
	Model        string                          `xml:"model"`
	Vendor       string                          `xml:"vendor"`
	Serial       string                          `xml:"serial"`
	Size         int                             `xml:"size"`
	Capatibility *NodeDeviceStorageSubCapability `xml:"capability,omitempty"`
}

type NodeDeviceDRMCapability struct {
	Type string `xml:"type"`
}

func (c *NodeDeviceCapability) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			switch attr.Value {
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
			}
		}
	}
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
