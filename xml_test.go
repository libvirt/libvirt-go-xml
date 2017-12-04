// +build xmlroundtrip

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
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"
)

var xmldirs = []string{
	"testdata/libvirt/tests/bhyveargv2xmldata",
	"testdata/libvirt/tests/bhyvexml2argvdata",
	"testdata/libvirt/tests/bhyvexml2xmloutdata",
	"testdata/libvirt/tests/capabilityschemadata",
	"testdata/libvirt/tests/cputestdata",
	"testdata/libvirt/tests/domaincapsschemadata",
	"testdata/libvirt/tests/domainconfdata",
	"testdata/libvirt/tests/domainschemadata",
	"testdata/libvirt/tests/domainsnapshotxml2xmlin",
	"testdata/libvirt/tests/domainsnapshotxml2xmlout",
	"testdata/libvirt/tests/genericxml2xmlindata",
	"testdata/libvirt/tests/genericxml2xmloutdata",
	"testdata/libvirt/tests/interfaceschemadata",
	"testdata/libvirt/tests/libxlxml2domconfigdata",
	"testdata/libvirt/tests/lxcconf2xmldata",
	"testdata/libvirt/tests/lxcxml2xmldata",
	"testdata/libvirt/tests/lxcxml2xmloutdata",
	"testdata/libvirt/tests/networkxml2confdata",
	"testdata/libvirt/tests/networkxml2firewalldata",
	"testdata/libvirt/tests/networkxml2xmlin",
	"testdata/libvirt/tests/networkxml2xmlout",
	"testdata/libvirt/tests/networkxml2xmlupdatein",
	"testdata/libvirt/tests/networkxml2xmlupdateout",
	"testdata/libvirt/tests/nodedevschemadata",
	"testdata/libvirt/tests/nwfilterxml2firewalldata",
	"testdata/libvirt/tests/nwfilterxml2xmlin",
	"testdata/libvirt/tests/nwfilterxml2xmlout",
	"testdata/libvirt/tests/qemuagentdata",
	"testdata/libvirt/tests/qemuargv2xmldata",
	"testdata/libvirt/tests/qemucapabilitiesdata",
	"testdata/libvirt/tests/qemucaps2xmldata",
	"testdata/libvirt/tests/qemuhotplugtestcpus",
	"testdata/libvirt/tests/qemuhotplugtestdevices",
	"testdata/libvirt/tests/qemuhotplugtestdomains",
	"testdata/libvirt/tests/qemumemlockdata",
	"testdata/libvirt/tests/qemuxml2argvdata",
	"testdata/libvirt/tests/qemuxml2xmloutdata",
	"testdata/libvirt/tests/secretxml2xmlin",
	"testdata/libvirt/tests/securityselinuxlabeldata",
	"testdata/libvirt/tests/sexpr2xmldata",
	"testdata/libvirt/tests/storagepoolschemadata",
	"testdata/libvirt/tests/storagepoolxml2xmlin",
	"testdata/libvirt/tests/storagepoolxml2xmlout",
	"testdata/libvirt/tests/storagevolschemadata",
	"testdata/libvirt/tests/storagevolxml2xmlin",
	"testdata/libvirt/tests/storagevolxml2xmlout",
	"testdata/libvirt/tests/vircaps2xmldata",
	"testdata/libvirt/tests/virstorageutildata",
	"testdata/libvirt/tests/vmx2xmldata",
	"testdata/libvirt/tests/xencapsdata",
	"testdata/libvirt/tests/xlconfigdata",
	"testdata/libvirt/tests/xmconfigdata",
	"testdata/libvirt/tests/xml2sexprdata",
	"testdata/libvirt/tests/xml2vmxdata",
}

var consoletype = "/domain[0]/devices[0]/console[0]/@type"

var blacklist = map[string]bool{
	// intentionally invalid xml
	"testdata/libvirt/tests/genericxml2xmlindata/generic-chardev-unix-redirdev-missing-path.xml":  true,
	"testdata/libvirt/tests/genericxml2xmlindata/generic-chardev-unix-rng-missing-path.xml":       true,
	"testdata/libvirt/tests/qemuxml2argvdata/qemuxml2argv-virtio-rng-egd-crash.xml":               true,
	"testdata/libvirt/tests/genericxml2xmlindata/generic-chardev-unix-smartcard-missing-path.xml": true,
	"testdata/libvirt/tests/genericxml2xmlindata/generic-chardev-tcp-multiple-source.xml":         true,
	// udp source in different order
	"testdata/libvirt/tests/genericxml2xmlindata/generic-chardev-udp.xml":                 true,
	"testdata/libvirt/tests/genericxml2xmlindata/generic-chardev-udp-multiple-source.xml": true,
	"testdata/libvirt/tests/xml2sexprdata/xml2sexpr-fv-serial-udp.xml":                    true,
}

var extraActualNodes = map[string][]string{

	"testdata/libvirt/tests/xml2sexprdata/xml2sexpr-pv-vcpus.xml":              []string{consoletype},
	"testdata/libvirt/tests/xml2sexprdata/xml2sexpr-pv.xml":                    []string{consoletype},
	"testdata/libvirt/tests/xml2sexprdata/xml2sexpr-pv-bootloader.xml":         []string{consoletype},
	"testdata/libvirt/tests/xml2sexprdata/xml2sexpr-pv-bootloader-cmdline.xml": []string{consoletype},
	"testdata/libvirt/tests/xml2sexprdata/xml2sexpr-pci-devs.xml":              []string{consoletype},
	"testdata/libvirt/tests/xml2sexprdata/xml2sexpr-net-routed.xml":            []string{consoletype},
	"testdata/libvirt/tests/xml2sexprdata/xml2sexpr-net-e1000.xml":             []string{consoletype},
	"testdata/libvirt/tests/xml2sexprdata/xml2sexpr-net-bridged.xml":           []string{consoletype},
	"testdata/libvirt/tests/xml2sexprdata/xml2sexpr-fv-kernel.xml":             []string{consoletype},
	"testdata/libvirt/tests/xml2sexprdata/xml2sexpr-escape.xml":                []string{consoletype},
	"testdata/libvirt/tests/xml2sexprdata/xml2sexpr-disk-file.xml":             []string{consoletype},
	"testdata/libvirt/tests/xml2sexprdata/xml2sexpr-disk-drv-loop.xml":         []string{consoletype},
	"testdata/libvirt/tests/xml2sexprdata/xml2sexpr-disk-drv-blktap2.xml":      []string{consoletype},
	"testdata/libvirt/tests/xml2sexprdata/xml2sexpr-disk-drv-blktap2-raw.xml":  []string{consoletype},
	"testdata/libvirt/tests/xml2sexprdata/xml2sexpr-disk-drv-blktap.xml":       []string{consoletype},
	"testdata/libvirt/tests/xml2sexprdata/xml2sexpr-disk-drv-blktap-raw.xml":   []string{consoletype},
	"testdata/libvirt/tests/xml2sexprdata/xml2sexpr-disk-drv-blktap-qcow.xml":  []string{consoletype},
	"testdata/libvirt/tests/xml2sexprdata/xml2sexpr-disk-drv-blkback.xml":      []string{consoletype},
	"testdata/libvirt/tests/xml2sexprdata/xml2sexpr-disk-block.xml":            []string{consoletype},
	"testdata/libvirt/tests/xml2sexprdata/xml2sexpr-disk-block-shareable.xml":  []string{consoletype},
	"testdata/libvirt/tests/xml2sexprdata/xml2sexpr-bridge-ipaddr.xml":         []string{consoletype},
	"testdata/libvirt/tests/xml2sexprdata/xml2sexpr-boot-grub.xml":             []string{consoletype},
	"testdata/libvirt/tests/xml2sexprdata/xml2sexpr-no-source-cdrom.xml": []string{
		"/domain[0]/devices[0]/disk[1]/@type",
	},

	"testdata/libvirt/tests/qemuxml2argvdata/qemuxml2argv-fs9p-ccw.xml": []string{
		"/domain[0]/devices[0]/filesystem[1]/@type",
		"/domain[0]/devices[0]/filesystem[2]/@type",
	},
	"testdata/libvirt/tests/qemuxml2argvdata/qemuxml2argv-fs9p.xml": []string{
		"/domain[0]/devices[0]/filesystem[1]/@type",
		"/domain[0]/devices[0]/filesystem[2]/@type",
	},
	"testdata/libvirt/tests/qemuxml2argvdata/qemuxml2argv-disk-drive-discard.xml": []string{
		"/domain[0]/devices[0]/disk[0]/@type",
	},
	"testdata/libvirt/tests/genericxml2xmlindata/generic-chardev-udp.xml": []string{
		"/domain[0]/devices[0]/channel[0]/source[0]/@mode",
	},
	"testdata/libvirt/tests/qemuxml2argvdata/qemuxml2argv-disk-mirror-old.xml": []string{
		"/domain[0]/devices[0]/disk[0]/mirror[0]/@type",
		"/domain[0]/devices[0]/disk[0]/mirror[0]/source[0]",
		"/domain[0]/devices[0]/disk[2]/mirror[0]/@type",
		"/domain[0]/devices[0]/disk[2]/mirror[0]/format[0]",
		"/domain[0]/devices[0]/disk[2]/mirror[0]/source[0]",
	},

	"testdata/libvirt/tests/networkxml2xmlin/openvswitch-net.xml": []string{
		"/network[0]/virtualport[0]/parameters[0]",
	},
	"testdata/libvirt/tests/networkxml2xmlout/openvswitch-net.xml": []string{
		"/network[0]/virtualport[0]/parameters[0]",
	},
	"testdata/libvirt/tests/networkxml2xmlupdateout/openvswitch-net-modified.xml": []string{
		"/network[0]/virtualport[0]/parameters[0]",
	},
	"testdata/libvirt/tests/networkxml2xmlupdateout/openvswitch-net-more-portgroups.xml": []string{
		"/network[0]/virtualport[0]/parameters[0]",
	},
	"testdata/libvirt/tests/networkxml2xmlupdateout/openvswitch-net-without-alice.xml": []string{
		"/network[0]/virtualport[0]/parameters[0]",
	},
}

var extraExpectNodes = map[string][]string{
	"testdata/libvirt/tests/genericxml2xmlindata/generic-chardev-unix.xml": []string{
		"/domain[0]/devices[0]/channel[1]/source[0]",
	},
	"testdata/libvirt/tests/qemuxml2argvdata/qemuxml2argv-usb-redir-filter.xml": []string{
		"/domain[0]/devices[0]/redirfilter[0]/usbdev[1]/@vendor",
		"/domain[0]/devices[0]/redirfilter[0]/usbdev[1]/@product",
		"/domain[0]/devices[0]/redirfilter[0]/usbdev[1]/@class",
		"/domain[0]/devices[0]/redirfilter[0]/usbdev[1]/@version",
	},
	"testdata/libvirt/tests/domainschemadata/domain-parallels-ct-simple.xml": []string{
		"/domain[0]/description[0]",
	},
}

func testRoundTrip(t *testing.T, xml string, filename string) {
	if strings.HasSuffix(filename, "-invalid.xml") {
		return
	}

	var doc Document
	if strings.HasPrefix(xml, "<domain ") {
		doc = &Domain{}
	} else if strings.HasPrefix(xml, "<capabilities") {
		doc = &Caps{}
	} else if strings.HasPrefix(xml, "<network") {
		doc = &Network{}
	} else if strings.HasPrefix(xml, "<secret") {
		doc = &Secret{}
	} else {
		return
	}
	err := doc.Unmarshal(xml)
	if err != nil {
		t.Fatal(fmt.Errorf("Cannot parse file %s: %s\n", filename, err))
	}

	newxml, err := doc.Marshal()
	if err != nil {
		t.Fatal(err)
	}

	extraExpectNodes, _ := extraExpectNodes[filename]
	extraActualNodes, _ := extraActualNodes[filename]
	err = testCompareXML(filename, xml, newxml, extraExpectNodes, extraActualNodes)
	if err != nil {
		t.Fatal(err)
	}
}

func syncGit(t *testing.T) {
	_, err := os.Stat("testdata/libvirt/tests")
	if err != nil {
		if os.IsNotExist(err) {
			err := exec.Command("git", "clone", "--depth", "1", "git://libvirt.org/libvirt.git", "testdata/libvirt").Run()
			if err != nil {
				t.Fatal(err)
			}
		} else {
			t.Fatal(err)
		}
	} else {
		here, err := os.Getwd()
		if err != nil {
			t.Fatal(err)
		}
		err = os.Chdir("testdata/libvirt")
		if err != nil {
			t.Fatal(err)
		}
		defer func() {
			os.Chdir(here)
		}()
		err = exec.Command("git", "pull").Run()
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestRoundTrip(t *testing.T) {
	syncGit(t)
	for _, xmldir := range xmldirs {
		xmlfiles, err := ioutil.ReadDir(xmldir)
		if err != nil {
			t.Fatal(err)
		}

		for _, xmlfile := range xmlfiles {
			if !xmlfile.IsDir() && strings.HasSuffix(xmlfile.Name(), ".xml") {
				fname := xmldir + "/" + xmlfile.Name()
				_, ok := blacklist[fname]
				if ok {
					continue
				}
				xml, err := ioutil.ReadFile(fname)
				if err != nil {
					t.Fatal(err)
				}
				testRoundTrip(t, string(xml), fname)
			}
		}
	}
}
