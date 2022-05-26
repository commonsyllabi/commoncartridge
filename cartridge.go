package commoncartridge

import (
	"io/fs"

	"github.com/commonsyllabi/commoncartridge/types"
)

type Cartridge interface {
	// MarshalJSON returns a serialized JSON representation
	MarshalJSON() ([]byte, error)

	//-- ParseManifest finds the imsmanifest.xml in the ZipReader and marshals it into a struct
	Manifest() (types.Manifest, error)

	// Title returns the title of the loaded cartridge
	Title() string

	// Metadata returns the metadata fields of the cartridge in a structured fashion
	Metadata() (string, error)

	// Items returns a slice of structs which include the Item, the Resources and the children Items it might have
	Items() ([]FullItem, error)

	// Resources returns a slice of structs which include the resource and, if found, the item in which the resource appears.
	Resources() ([]FullResource, error)
	Weblinks() ([]types.WebLink, error)
	Assignments() ([]types.Assignment, error)
	LTIs() ([]types.CartridgeBasicltiLink, error)
	QTIs() ([]types.Questestinterop, error)
	Topics() ([]types.Topic, error)

	// Find takes an identifier and returns the corresponding resource.
	Find(string) (interface{}, error)

	// FindFile takes an identifier and returns the fs.File that the corresponding node refers to.
	FindFile(string) (fs.File, error)
}
