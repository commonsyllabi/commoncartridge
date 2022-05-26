package commoncartridge

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/commonsyllabi/commoncartridge/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const singleTestFile = "./test_files/test_01.imscc"
const allTestFilesDir = "./test_files/dump"

func TestLoadEmpty(t *testing.T) {
	_, err := Load("")
	assert.NotNil(t, err)
}

func TestExactCartridge(t *testing.T) {
	var cc Cartridge = load(t, singleTestFile)
	manifest, err := cc.Manifest()
	assert.Nil(t, err)
	assert.Equal(t, manifest.Metadata.Lom.General.Title.String.Text, "Loaded Course")
	assert.Equal(t, manifest.Metadata.Lom.General.Description.String.Text, "Sample Description")
	assert.Equal(t, manifest.Metadata.Lom.General.Keyword.String.Text, "Test, Attempt")
	assert.Equal(t, manifest.Metadata.Lom.General.Language, "en-US")
	assert.Equal(t, manifest.Metadata.Lom.LifeCycle.Contribute.Date.DateTime, "2014-09-08")
	assert.Equal(t, manifest.Metadata.Lom.Rights.CopyrightAndOtherRestrictions.Value, "yes")
	assert.Equal(t, manifest.Metadata.Lom.Rights.Description.String, "Private (Copyrighted) - http://en.wikipedia.org/wiki/Copyright")
	assert.Equal(t, manifest.Organizations.Organization.Item.Identifier, "LearningModules")

	public := manifest.Organizations.Organization.Item.Item[0]
	assert.Equal(t, len(public.Item), 11)

	locked := manifest.Organizations.Organization.Item.Item[1]
	assert.Equal(t, len(locked.Item), 1)
	assert.Equal(t, len(manifest.Resources.Resource), 120)
}

func TestLoadCorrect(t *testing.T) {
	cc := load(t, singleTestFile)
	assert.NotEmpty(t, cc, IMSCC{})
}

func TestLoadAll(t *testing.T) {
	cwd, _ := os.Getwd()

	_, err := os.Stat(filepath.Join(cwd, allTestFilesDir))
	if os.IsNotExist(err) {
		t.Skip("dump folder does not exist, skipping")
	}

	files, err := ioutil.ReadDir(filepath.Join(cwd, allTestFilesDir))
	require.Nil(t, err)

	t.Logf("Test loading %d cartridges\n", len(files))

	for i, file := range files {
		if file.IsDir() {
			continue
		}

		cc, err := Load(filepath.Join(allTestFilesDir, file.Name()))
		require.Nil(t, err)
		assert.NotEmpty(t, cc, IMSCC{})

		assert.NotEqual(t, cc.Title(), "")
		t.Logf("Parsed %d/%d - %s\n", i+1, len(files), cc.Title())
	}
}

func TestMetadata(t *testing.T) {
	cc := load(t, singleTestFile)
	meta, err := cc.Metadata()
	require.Nil(t, err)
	assert.IsType(t, "", meta)
}

func TestItems(t *testing.T) {
	cc := load(t, singleTestFile)
	items, err := cc.Items()
	require.Nil(t, err)

	assert.IsType(t, []FullItem{}, items)
	assert.IsType(t, types.Item{}, items[0].Item)
	assert.Equal(t, len(items), 2)
}

func TestResources(t *testing.T) {
	cc := load(t, singleTestFile)
	resources, err := cc.Resources()
	require.Nil(t, err)

	assert.IsType(t, []FullResource{}, resources)
	assert.IsType(t, types.Resource{}, resources[0].Resource)
	assert.Equal(t, len(resources), 120)
}

func TestAssignments(t *testing.T) {
	cc := load(t, singleTestFile)
	assignments, err := cc.Assignments()
	require.Nil(t, err)

	assert.IsType(t, []types.Assignment{}, assignments)
	assert.Contains(t, assignments[0].XMLName.Local, "assignment")
}

func TestLTIs(t *testing.T) {
	cc := load(t, singleTestFile)
	ltis, err := cc.LTIs()
	require.Nil(t, err)

	assert.IsType(t, []types.CartridgeBasicltiLink{}, ltis)
	assert.Contains(t, ltis[0].XMLName.Local, "cartridge_basiclti_link")
}

func TestQTIs(t *testing.T) {
	cc := load(t, singleTestFile)
	qtis, err := cc.QTIs()
	require.Nil(t, err)

	assert.IsType(t, []types.Questestinterop{}, qtis)
	assert.Contains(t, qtis[0].XMLName.Local, "questestinterop")
}

func TestTopics(t *testing.T) {
	cc := load(t, singleTestFile)
	topics, err := cc.Topics()
	require.Nil(t, err)

	assert.IsType(t, []types.Topic{}, topics)
	assert.Contains(t, topics[0].XMLName.Local, "topic")
}

func TestWeblinks(t *testing.T) {
	cc := load(t, singleTestFile)
	weblinks, err := cc.Weblinks()

	require.Nil(t, err)

	assert.Contains(t, weblinks[0].XMLName.Local, "webLink")
}

func TestFind(t *testing.T) {
	cc := load(t, singleTestFile)

	found, err := cc.Find("ic1b5d76bd9a4bd37eb78cf0bcb5b84da")
	require.Nil(t, err)
	assert.IsType(t, types.Resource{}, found)

	found, err = cc.Find("i528c2ce0186a758d13a9bd193bd88611")
	require.Nil(t, err)
	assert.IsType(t, types.Topic{}, found)

	found, err = cc.Find("ibb3ca45e774c0c487daeb9352e7a4553")
	require.Nil(t, err)
	assert.IsType(t, types.WebLink{}, found)

	found, err = cc.Find("ie801a403cd25e9a771ab7e3a2d6bea3a")
	require.Nil(t, err)
	assert.IsType(t, types.Assignment{}, found)

	found, err = cc.Find("iad7e264143b9f2ec9dbc71a9d166f6f2")
	require.Nil(t, err)
	assert.IsType(t, types.Questestinterop{}, found)

	found, err = cc.Find("iae0220efe8693f664806e9bfe43b6e30")
	require.Nil(t, err)
	assert.IsType(t, types.CartridgeBasicltiLink{}, found)

	found, err = cc.Find("i3755487a331b36c76cec8bbbcdb7cc66")
	require.Nil(t, err)
	assert.IsType(t, types.Resource{}, found)
}

func TestFindFile(t *testing.T) {
	cc := load(t, singleTestFile)
	_, err := cc.FindFile("i3755487a331b36c76cec8bbbcdb7cc66")
	require.Nil(t, err)
}

func TestMarshalJSON(t *testing.T) {
	cc := load(t, singleTestFile)
	obj, err := cc.MarshalJSON()
	require.Nil(t, err)

	assert.NotEqual(t, len(obj), 0)
}

func load(t *testing.T, p string) Cartridge {
	cc, err := Load(p)
	require.Nil(t, err)
	return cc
}
