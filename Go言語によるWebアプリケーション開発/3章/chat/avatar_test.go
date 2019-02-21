package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	gomniauthtest "github.com/stretchr/gomniauth/test"
)

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	testUser := &gomniauthtest.TestUser{}
	testUser.On("AvatarURL").Return("", ErrNoAvatarURL)
	testChatUser := &chatUser{User: testUser}
	url, err := authAvatar.GetAvatarURL(testChatUser)
	if err != ErrNoAvatarURL {
		t.Error("値が存在しない場合,AuthAvatar.GetAvatarURLはErrNoAvatatURLを返すべき")
	}

	testUrl := "http://url-to-avatar/"
	testUser = &gomniauthtest.TestUser{}
	testChatUser.User = testUser
	testUser.On("AvatarURL").Return(testUrl, nil)
	url, err = authAvatar.GetAvatarURL(testChatUser)
	if err != nil {
		t.Error("値が存在する場合AuthAvatar.GetAvatarURLはエラーを返すべきではない")
	} else {
		if url != testUrl {
			t.Error("AuthAvatar.GetAvatarURLは正しいURLを返すべき")
		}
	}
}

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvatar GravatarAvatar
	user := &chatUser{uniqueID: "abc"}
	url, err := gravatarAvatar.GetAvatarURL(user)
	if err != nil {
		t.Error("gravatarAvatar.GetAvatarURLはエラーを返すべきではない")
	}
	if url != "//www.gravatar.com/avatar/abc" {
		t.Errorf("gravatarAvatar.GetAvatarURLが%sという値を返しました", url)
	}
}

func TestFileSystemAvatar(t *testing.T) {
	filename := filepath.Join("avatars", "abc.jpg")
	ioutil.WriteFile(filename, []byte{}, 0777)
	defer func() { os.Remove(filename) }()

	var fileSystemAvatar FileSystemAtatar
	user := &chatUser{uniqueID: "abc"}
	url, err := fileSystemAvatar.GetAvatarURL(user)
	if err != nil {
		t.Error("fileSystemAvatar.GetAvatarURLはエラーを返すべきではない")
	}
	if url != "/avatars/abc.jpg" {
		t.Errorf("fileSystemAvatar.GetAvatarURLが%sを返した", url)
	}
}
