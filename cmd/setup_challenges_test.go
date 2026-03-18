package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_noopProvider(t *testing.T) {
	p := &noopProvider{}

	testCases := []struct {
		desc    string
		domain  string
		token   string
		keyAuth string
	}{
		{desc: "normal values", domain: "example.com", token: "abc123", keyAuth: "keyAuth"},
		{desc: "empty values", domain: "", token: "", keyAuth: ""},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			err := p.Present(test.domain, test.token, test.keyAuth)
			assert.NoError(t, err)

			err = p.CleanUp(test.domain, test.token, test.keyAuth)
			assert.NoError(t, err)
		})
	}
}
