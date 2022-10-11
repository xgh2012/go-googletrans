package translator

import (
	"testing"
)

// TestTranslator_Translate calls translate.translate.
func TestTranslator_Translate(t *testing.T) {
	origin := "你好，世界！"
	dest := "Hello World!"
	c := Config{
		Proxy:       "http://127.0.0.1:1080",
		UserAgent:   []string{"Custom Agent"},
		ServiceUrls: []string{"translate.google.com.hk"},
	}
	trans := New(c)
	result, srcLangBack, err := trans.Translate(origin, "auto", "en")

	if result.Text != dest || err != nil {
		t.Fatalf(`%q, %v, srcLangBack %v, Want match for %q, nil`, result.Text, srcLangBack, err, dest)
	}
}
