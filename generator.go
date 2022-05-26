package commoncartridge

//go:generate echo "Generating Manifest, Item, Resource, ..."
//go:generate bash -c "zek -P types -t manifest -r 'Item Resource' -o ./types/autogen_manifest.go ./types/examples/manifest.xml"
//go:generate bash -c "zek -P types -t item -r 'Item' -o ./types/autogen_item.go ./types/examples/item.xml"
//go:generate bash -c "zek -P types -t resource -o ./types/autogen_resource.go ./types/examples/resource.xml"

//go:generate echo "Generating Topic, LTI, QTI, WebLink, ..."
//go:generate bash -c "zek -P types -t topic -o ./types/autogen_topic.go ./types/examples/topic.xml"
//go:generate bash -c "zek -P types -t lti -o ./types/autogen_lti.go ./types/examples/lti.xml"
//go:generate bash -c "zek -P types -t qti -o ./types/autogen_qti.go ./types/examples/qti.xml"
//go:generate bash -c "zek -P types -t weblink -o ./types/autogen_weblink.go ./types/examples/weblink.xml"
//go:generate bash -c "zek -P types -t assignment -o ./types/autogen_assignment.go ./types/examples/assignment.xml"

//go:generate echo "...done!"
