package viperconf

// Configuration is  struct of setting all Configuration  */
type Configuration struct {
	Sourcefilerepository SourceFileRepositoryConfiguration
	Logger               LoggerConfiguration
}

// AppConfiguration is struct of App info
type SourceFileRepositoryConfiguration struct {
	RootDirectoryPath     string
	DirectoryPrepare      string
	DirectoryOutput       string
	DirectoryInput        string
	ConfigWordToSpace     string
	ConfigWordToRemove    string
	ConfigWordInsertSpace string
}

// LoggerConfiguration is struct of Logger
type LoggerConfiguration struct {
	Logpath      string
	Mainlogpath  string
	Initlogpath  string
	Debuglogpath string
	Maxsize      int
	Maxbackups   int
	Maxage       int
	Compress     bool
}
