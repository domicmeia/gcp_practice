package translation_test

import (
	"testing"

	"github.com/domicmeia/gcp_practice/translation"
)

func TestTranslate(t *testing.T) {

	tt := []struct {
		Word        string
		Language    string
		Translation string
	}{
		{
			Word: "hello", Language: "english", Translation: "hello",
		},
		{
			Word: "hello", Language: "german", Translation: "hallo",
		},
		{
			Word: "hello", Language: "finnish", Translation: "hei",
		},
		{
			Word: "hello", Language: "dutch", Translation: "",
		},
	}
	underTest := translation.NewStaticService()

	for _, test := range tt {
		res := underTest.Translate(test.Word, test.Language)

		if res != test.Translation {
			t.Errorf(
				`expected "%s" to be "%s" from "%s" but received "%s"`,
				test.Word, test.Language, test.Translation, res)
		}
	}

}
