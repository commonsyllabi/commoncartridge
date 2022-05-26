package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/commonsyllabi/viewer/pkg/commoncartridge"
)

var (
	debug       = flag.Bool("d", false, "debug output")
	metadata    = flag.Bool("m", false, "shows metadata as serialized json")
	as_json     = flag.Bool("j", false, "dumps a serialized json representation")
	items       = flag.Bool("I", false, "lists all items, with their associated resources in the cartridge")
	resources   = flag.Bool("r", false, "lists all resources in the cartridge")
	weblinks    = flag.Bool("weblinks", false, "lists all weblinks in the cartridge")
	assignments = flag.Bool("assignments", false, "lists all assignments in the cartridge")
	topics      = flag.Bool("topics", false, "lists all topics in the cartridge")
	qtis        = flag.Bool("qtis", false, "lists all quizzes in the cartridge")
	ltis        = flag.Bool("ltis", false, "lists all basic LTI links in the cartridge")
	find        = flag.String("f", "", "finds the resource with the related id")
	file        = flag.String("F", "", "finds the file (i.e. webcontent) with the related id and returns the file as a fs.File")
)

func main() {
	flag.Parse()

	if *debug {
		fmt.Println("cosyl v0.1")
	}

	if flag.NArg() == 0 {
		log.Fatal("provide the path of the cartridge to be opened!")
	}

	inputFile := flag.Args()[0]

	cc, err := commoncartridge.Load(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	if *debug {
		fmt.Println("successfully loaded cartridge")
	}

	if *metadata {
		meta, err := cc.Metadata()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(meta)
	}

	if *items {
		items, err := cc.Items()
		if err != nil {
			log.Fatal(err)
		}

		data, _ := json.Marshal(items)
		fmt.Println(string(data))
	}

	if *resources {
		resources, err := cc.Resources()
		if err != nil {
			log.Fatal(err)
		}

		for _, r := range resources {
			if r.Item.Identifierref == "" {
				r.Item.Title = "none"
			}

			fmt.Printf("%+v\n", r)
		}
	}

	if *weblinks {
		weblinks, err := cc.Weblinks()
		if err != nil {
			log.Fatal(err)
		}

		for _, wl := range weblinks {
			fmt.Printf("xml: %s title: %s url: %s\n", wl.XMLName.Local, wl.Title, wl.URL.Href)
		}
	}

	if *assignments {
		assignments, err := cc.Assignments()
		if err != nil {
			log.Fatal(err)
		}

		for _, a := range assignments {
			fmt.Printf("xml: %s title: %s\n", a.XMLName.Local, a.Title)
		}
	}

	if *topics {
		topics, err := cc.Topics()
		if err != nil {
			log.Fatal(err)
		}

		for _, t := range topics {
			fmt.Printf("xml: %s title: %s attachements: %d\n", t.XMLName.Local, t.Title, len(t.Attachments.Attachment))
		}
	}

	if *qtis {
		qtis, err := cc.QTIs()
		if err != nil {
			log.Fatal(err)
		}

		for _, qti := range qtis {
			fmt.Printf("xml: %s title: %s items: %d\n", qti.XMLName.Local, qti.Assessment.Title, len(qti.Assessment.Section.Item))
		}
	}

	if *ltis {
		ltis, err := cc.LTIs()
		if err != nil {
			log.Fatal(err)
		}

		for _, lti := range ltis {
			fmt.Printf("xml: %s title: %s description: %s url: %s\n", lti.XMLName.Local, lti.Title, lti.Description, lti.SecureLaunchURL)
		}
	}

	if *as_json {
		obj, err := cc.MarshalJSON()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(string(obj))
	}

	if *find != "" {
		res, err := cc.Find(*find)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", res)
	}

	if *file != "" {
		file, err := cc.FindFile(*file)
		if err != nil {
			log.Fatal(err)
		}

		info, err := file.Stat()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("found: %s\n", info.Name())

		dst, err := os.Create(info.Name())
		if err != nil {
			log.Fatal(err)
		}

		_, err = io.Copy(dst, file)
		if err != nil {
			log.Fatal(err)
		}

	}
}
