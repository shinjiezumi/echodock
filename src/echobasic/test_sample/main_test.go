package test_sample

import (
	"testing"
)

func TestDecode(t *testing.T) {
	post, err := decode("./post.json")
	if err != nil {
		t.Error(err)
	}
	if post.Id != 1 {
		t.Error("wrong id, was expecting 1 but got", post.Id)
	}

	if post.Content != "HelloWorld!" {
		t.Error("wrong conetnt, was expecting 'HelloWorld!' but got", post.Content)
	}
}

func TestDecode_InvalidFileName(t *testing.T) {
	post, err := decode("./post.jsonaaa")
	if post.Id != 0 || post.Content != "" || err == nil {
		t.Error(err)
	}
}

func TestEncode(t *testing.T) {
	t.Skip("Skipping encoding for now")
}
