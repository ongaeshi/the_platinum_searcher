package option

type Option struct {
	NoColor          bool     `long:"nocolor" description:"Don't print color codes in results (Disabled by default)"`
	NoGroup          bool     `long:"nogroup" description:"Don't print file name at header (Disabled by default)"`
	FilesWithMatches bool     `short:"l" long:"files-with-matches" description:"Only print filenames that contain matches"`
	VcsIgnore        []string `long:"vcs-ignore" description:"VCS ignore files (Default: .gitignore, .hgignore)"`
	Ignore           []string `long:"ignore" description:"Ignore files/directories matching pattern"`
	IgnoreCase       bool     `short:"i" long:"ignore-case" description:"Match case insensitively"`
	SmartCase        bool     `short:"S" long:"smart-case" description:"Match case insensitively unless PATTERN contains uppercase characters"`
	FilesWithRegexp  string   `short:"g" description:"Print filenames matching PATTERN"`
	FileSearchRegexp string   `short:"G" long:"file-search-regexp" description:"PATTERN Limit search to filenames matching PATTERN"`
	Depth            int      `long:"depth" default:"25" default-mask:"-" description:"Search up to NUM derectories deep (Default: 25)"`
	Follow           bool     `short:"f" long:"follow" description:"Follow symlinks"`
	After            int      `short:"A" long:"after" description:"Print lines after match"`
	Before           int      `short:"B" long:"before" description:"Print lines before match"`
	Context          int      `short:"C" long:"context" description:"Print lines before and after match"`
	OutputEncode     []string `short:"o" long:"output-encode" description:"Specify output encoding (none, jis, sjis, euc)"`
	Proc             int      // Number of goroutine. Not user option.
	Version          bool     `long:"version" description:"Show version"`
}

func (self *Option) VcsIgnores() []string {
	if len(self.VcsIgnore) == 0 {
		self.VcsIgnore = []string{".gitignore", ".hgignore"}
	}
	return self.VcsIgnore
}
