package os

import (
	"testing"
)

/*
The test in this package is just to test that the methods in the os package have been called.
Whether an error is returned when calling is not very important.
*/

func TestExecutable(t *testing.T) {
	o := &OS{}
	r, e := o.Executable()

	if e != nil {
		t.Logf("%+v\n", e)
		return
	}

	t.Logf("%s\n", r)
}

func TestUserHomeDir(t *testing.T) {
	o := &OS{}
	r, e := o.UserHomeDir()

	if e != nil {
		t.Logf("%+v\n", e)
		return
	}

	t.Logf("%s\n", r)
}

func TestUserConfigDir(t *testing.T) {
	o := &OS{}
	r, e := o.UserConfigDir()

	if e != nil {
		t.Logf("%+v\n", e)
		return
	}

	t.Logf("%s\n", r)

}

func TestUserCacheDir(t *testing.T) {
	o := &OS{}
	r, e := o.UserCacheDir()

	if e != nil {
		t.Logf("%+v\n", e)
		return
	}

	t.Logf("%s\n", r)

}

func TestGetegid(t *testing.T) {
	o := &OS{}
	r := o.Getegid()
	t.Logf("%d\n", r)
}

func TestGeteuid(t *testing.T) {
	o := &OS{}
	r := o.Geteuid()
	t.Logf("%d\n", r)
}

func TestGetgid(t *testing.T) {
	o := &OS{}
	r := o.Getgid()
	t.Logf("%d\n", r)
}

func TestGetgroups(t *testing.T) {
	o := &OS{}
	r, e := o.Getgroups()

	if e != nil {
		t.Logf("%+v\n", e)
		return
	}

	t.Logf("%v\n", r)
}

func TestGetpagesize(t *testing.T) {
	o := &OS{}
	r := o.Getpagesize()

	t.Logf("%d\n", r)
}

func TestGetpid(t *testing.T) {
	o := &OS{}
	r := o.Getpid()

	t.Logf("%d\n", r)
}

func TestGetppid(t *testing.T) {
	o := &OS{}
	r := o.Getppid()

	t.Logf("%d\n", r)
}

func TestGetuid(t *testing.T) {
	o := &OS{}
	r := o.Getuid()

	t.Logf("%d\n", r)
}

func TestGetwd(t *testing.T) {
	o := &OS{}
	r, e := o.Getwd()

	if e != nil {
		t.Logf("%+v\n", e)
		return
	}

	t.Logf("%s\n", r)
}
