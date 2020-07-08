# libvirt-go-xml

[![Build Status](https://gitlab.com/libvirt/libvirt-go-xml/badges/master/pipeline.svg)](https://gitlab.com/libvirt/libvirt-go-xml/pipelines)
[![API Documentation](https://img.shields.io/static/v1?label=godev&message=reference&color=00add8)](https://pkg.go.dev/libvirt.org/libvirt-go-xml)

Go API for manipulating libvirt XML documents

This package provides a Go API that defines a set of structs, annotated for use
with "encoding/xml", that can represent libvirt XML documents. There is no
dependancy on the libvirt library itself, so this can be used regardless of
the way in which the application talks to libvirt.

## Documentation

* [API documentation for the bindings](https://pkg.go.dev/libvirt.org/libvirt-go-xml)
* [Libvirt XML schema documentation](https://libvirt.org/format.html):
  * [capabilities](https://libvirt.org/formatcaps.html)
  * [domain](https://libvirt.org/formatdomain.html)
  * [domain capabilities](https://libvirt.org/formatdomaincaps.html)
  * [domain snapshot](https://libvirt.org/formatsnapshot.html)
  * [network](https://libvirt.org/formatnetwork.html)
  * [node device](https://libvirt.org/formatnode.html)
  * [nwfilter](https://libvirt.org/formatnwfilter.html)
  * [secret](https://libvirt.org/formatsecret.html)
  * [storage](https://libvirt.org/formatstorage.html)
  * [storage encryption](https://libvirt.org/formatstorageencryption.html)

## Contributing

The libvirt project aims to add support for new XML elements to
libvirt-go-xml as soon as they are added to the main libvirt C
library. If you are submitting changes to the libvirt C library
that introduce new XML elements, please submit a libvirt-go-xml
change at the same time.

Bug fixes and other improvements to the libvirt-go-xml library are
welcome at any time. The preferred submission method is to use
git send-email to submit patches to the libvir-list@redhat.com
mailing list. eg. to send a single patch

```
  # git send-email --to libvir-list@redhat.com --subject-prefix "PATCH go-xml" \
       --smtp-server=$HOSTNAME -1
```

Or to send all patches on the current branch, against master

```
  $ git send-email --to libvir-list@redhat.com --subject-prefix "PATCH go-xml" \
       --smtp-server=$HOSTNAME --no-chain-reply-to --cover-letter --annotate \
       master..
```

Note the master GIT repository is at

```
   http://libvirt.org/git/?p=libvirt-go.git;a=summary
```

The following automatic read-only mirrors are available as a
convenience to allow contributors to "fork" the repository:

```
  https://gitlab.com/libvirt/libvirt-go
  https://github.com/libvirt/libvirt-go
```

While you can send pull-requests to these mirrors, they will be
re-submitted via emai to the mailing list for review before
being merged, unless they are trivial/obvious bug fixes.

