package common


func MakeStash(stashPath string) {
	
	exePath, err := os.Executable()

	if err != nil {
		panic(err)
	}

	exeDir := filepath.Dir(exePath)
	absolutePath := filepath.Join(exeDir, stashPath)

	return absolutePath
}