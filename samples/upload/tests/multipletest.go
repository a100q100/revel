package tests

import (
	"net/url"

	"github.com/revel/revel/samples/upload/app/routes"

	"github.com/revel/revel"
)

type MultipleTest struct {
	revel.TestSuite
}

func (t *MultipleTest) TestThatMultipleFilesUploadWorks() {
	// Make sure it is not allowed to submit less than 2 files.
	t.PostFile(routes.Multiple.HandleUpload(), url.Values{}, url.Values{
		"file": {
			"github.com/revel/revel/samples/upload/public/img/favicon.png",
		},
	})
	t.AssertOk()
	t.AssertContains("You cannot submit less than 2 files")

	// Make sure upload of 2 files works.
	t.PostFile(routes.Multiple.HandleUpload(), url.Values{}, url.Values{
		"file[]": {
			"github.com/revel/revel/samples/upload/public/img/favicon.png",
			"github.com/revel/revel/samples/upload/public/img/glyphicons-halflings.png",
		},
	})
	revel.WARN.Println(string(t.ResponseBody))
	t.AssertOk()
	t.AssertContains("Successfully uploaded")
	t.AssertContains("favicon.png")
	t.AssertContains("glyphicons-halflings.png")
	t.AssertContains("image/png")
}
