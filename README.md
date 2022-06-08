# Common Cartridge

Package for parsing IMSCC-compliant files. A [CommonCartridge](https://www.imsglobal.org/activity/common-cartridge) is a file used to export and import data from learning management systems such as Moodle, Canvas, Brightspace, Sakai, etc. In their own words:

> Common Cartridge (CC) is a set of open standards developed by the IMS member community that enable interoperability between content and systems. Common Cartridge basically solves two problems. The first is to provide a standard way to represent digital course materials for use in online learning systems so that such content can be developed in one format and used across a wide variety of learning systems. The second is to enable new publishing models for online course materials and digital books that are modular, web-distributed, interactive, and customizable.

This package allows you to inspect and access specific parts of the cartridge, either through a command-line interface provided, or through a web interface accessible at [viewer.commonsyllabi.org](https://viewer.commonsyllabi.org).

Find the full documentation for the package at [pkg.go.dev/github.com/commonsyllabi/commoncartridge](https://pkg.go.dev/github.com/commonsyllabi/commoncartridge).

## Installation

```
go get github.com/commonsyllabi/commoncartridge
```

## Usage

### CLI

You can use the command-line interface by passing it a `.zip` or `.imscc` file with a `imsmanifest.xml`, for instance, to access the metadata fields of the test file located in the `test_files` folder:

```
cosyl -m test_01.imscc
```

To list all commands:

```
cosyl --help
```
### Package

You can use this package directly in your Go app by importing it in this way:

```
package main

import "github.com/commonsyllabi/commoncartridge"

func main(){
    cc, err := commoncartridge.Load("test_01.imscc")
    if err != nil {
        log.Fatal(err)
    }

    // prints the JSON representation of the cartridge
    obj, err := cc.MarshalJSON()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Print(string(obj))
}

```

## Note on generating IMSCC structs

Due to the naming complications of the official XSD files and the exorbitant costs of IMSCC resources in terms of test files and validator software, the IMSCC structs are generated from the sample `.xml` files in `types/examples`, using [zek](https://github.com/miku/zek). You can regenerate the structs by running `go generate ./...` from the root folder.

__NOTE__: the current version of zek does not allow for the generation of non-nested structs ([#14](https://github.com/miku/zek/issues/14)). Please see [this fork](https://github.com/periode/zek) for a working version.

## Alternatives

- [github.com/AndcultureCode/Common-Cartridge](https://github.com/AndcultureCode/Common-Cartridge/), written in C#, seemingly unmaintained.
- [github.com/instructure/common-cartridge-viewer](https://github.com/instructure/common-cartridge-viewer), written in JS, maintained by Instructure.
- [github.com/vhl/common_cartridge_parser](https://github.com/vhl/common_cartridge_parser), written in Ruby, seemingly unmaintained.

## Credits

This work has been funded by the [Prototype Fund](https://prototypefund.de) and the [Bundesministerium für Bildung und Forschung](https://www.bmbf.de/bmbf/de/home/home_node.html), and has been developed by [Pierre Depaz](https://github.com/periode), [Tobias Schmidt](https://github.com/grobie) and [Pat Shiu](https://github.com/patshiu).
