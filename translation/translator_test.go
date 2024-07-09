package translation_test

import (
	"testing"

	"github.com/domicmeia/gcp_practice/translation"
)

func TestTranslate(t *testing.T) {
	res := translation.Translate("hello", "english")
	if res != "hello" {
		t.Errorf(`expected "hello" but received "%s"`, res)
	}

	res = translation.Translate("hello", "german")
	if res != "hallo" {
		t.Errorf(`expected "hallo" but received "%s"`, res)
	}

	res = translation.Translate("hello", "finnish")
	if res != "hei" {
		t.Errorf(`expected "hei" but received "%s"`, res)
	}

	res = translation.Translate("hello", "dutch")
	if res != "" {
		t.Errorf(`expected "" but received "%s"`, res)
	}

}
