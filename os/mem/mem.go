package mem

// OS defination
type OS struct {
	ExecutableFunc    func() (string, error)
	UserHomeDirFunc   func() (string, error)
	UserConfigDirFunc func() (string, error)
	UserCacheDirFunc  func() (string, error)
	GetegidFunc       func() int
	GeteuidFunc       func() int
	GetgidFunc        func() int
	GetgroupsFunc     func() ([]int, error)
	GetpagesizeFunc   func() int
	GetpidFunc        func() int
	GetppidFunc       func() int
	GetuidFunc        func() int
	GetwdFunc         func() (string, error)
}

// New defination
func New() *OS {
	strErrF := func() (string, error) {
		return "Placeholder", nil
	}

	intF := func() int {
		return 6
	}

	intsErrF := func() ([]int, error) {
		return []int{6, 6, 6}, nil
	}

	return &OS{
		ExecutableFunc:    strErrF,
		UserCacheDirFunc:  strErrF,
		UserHomeDirFunc:   strErrF,
		UserConfigDirFunc: strErrF,
		GetegidFunc:       intF,
		GeteuidFunc:       intF,
		GetgidFunc:        intF,
		GetgroupsFunc:     intsErrF,
		GetpagesizeFunc:   intF,
		GetpidFunc:        intF,
		GetppidFunc:       intF,
		GetuidFunc:        intF,
		GetwdFunc:         strErrF,
	}
}

// Executable defination
func (o *OS) Executable() (string, error) {
	return o.ExecutableFunc()
}

// UserHomeDir defination
func (o *OS) UserHomeDir() (string, error) {
	return o.UserHomeDirFunc()
}

// UserConfigDir defination
func (o *OS) UserConfigDir() (string, error) {
	return o.UserCacheDirFunc()
}

// UserCacheDir defination
func (o *OS) UserCacheDir() (string, error) {
	return o.UserCacheDirFunc()
}

// Getegid defination
func (o *OS) Getegid() int {
	return o.GetegidFunc()
}

// Geteuid defination
func (o *OS) Geteuid() int {
	return o.GeteuidFunc()
}

// Getgid defination
func (o *OS) Getgid() int {
	return o.GetgidFunc()
}

// Getgroups defination
func (o *OS) Getgroups() ([]int, error) {
	return o.GetgroupsFunc()
}

// Getpagesize defination
func (o *OS) Getpagesize() int {
	return o.GetpagesizeFunc()
}

// Getpid defination
func (o *OS) Getpid() int {
	return o.GetpidFunc()
}

// Getppid defination
func (o *OS) Getppid() int {
	return o.GetppidFunc()
}

// Getuid defination
func (o *OS) Getuid() int {
	return o.GetuidFunc()
}

// Getwd defination
func (o *OS) Getwd() (dir string, err error) {
	return o.GetwdFunc()
}
