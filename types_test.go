package tgbot

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFile_Link(t *testing.T) {
	f := &File{
		FileId: "1",
		FileSize: 1,
		FilePath: "test",
	}

	assert.Equal(t, "https://api.telegram.org/file/bottest/test", f.Link("test"))
}