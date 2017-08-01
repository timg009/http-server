package main

import (
	"path/filepath"
	"strings"
)

// extensionList holds a key-value store with the most common
// file extensions and their corresponding associations.
// There's also file names which are common across platform
// to identify certain file types.
var extensionList = map[string]string{
	// File extensions from https://www.computerhope.com/issues/ch001789.htm
	".aif":       "AIF audio",
	".cda":       "CD audio track",
	".mid":       "MIDI audio",
	".midi":      "MIDI audio",
	".mp3":       "MP3 audio",
	".mpa":       "MPEG-2 audio",
	".ogg":       "Ogg Vorbis audio",
	".wav":       "WAV",
	".wma":       "WMA audio",
	".wpl":       "Windows Media Player playlist",
	".7z":        "7-Zip compressed",
	".arj":       "ARJ compressed",
	".deb":       "Debian software package",
	".pkg":       "Package",
	".rar":       "RAR",
	".rpm":       "Red Hat Package Manager",
	".gz":        "Gzip compressed",
	".z":         "Z compressed",
	".zip":       "Zip compressed",
	".dmg":       "macOS X disk image",
	".iso":       "ISO disc image",
	".toast":     "Toast disc image",
	".vcd":       "Virtual CD",
	".csv":       "Comma separated value",
	".dat":       "Data",
	".db":        "Database",
	".dbf":       "Database",
	".log":       "Log",
	".mdb":       "Microsoft Access database",
	".sav":       "Save",
	".sql":       "SQL database",
	".tar":       "Linux / Unix tarball archive",
	".xml":       "XML",
	".apk":       "Android package",
	".bat":       "Batch",
	".bin":       "Binary",
	".cgi":       "Perl script",
	".pl":        "Perl script",
	".com":       "MS-DOS command",
	".exe":       "Executable",
	".gadget":    "Windows gadget",
	".jar":       "Java Archive",
	".wsf":       "Windows Script",
	".fnt":       "Windows font",
	".fon":       "Generic font",
	".otf":       "Open type font",
	".ttf":       "TrueType font",
	".ai":        "Adobe Illustrator",
	".bmp":       "Bitmap image",
	".gif":       "GIF image",
	".jpg":       "JPEG image",
	".jpeg":      "JPEG image",
	".png":       "PNG image",
	".ps":        "PostScript",
	".psd":       "PSD image",
	".svg":       "Scalable Vector Graphics",
	".tif":       "TIFF image",
	".tiff":      "TIFF image",
	".gitignore": "Git ignore",
	".rspec":     "RSpec",
	".asp":       "Active Server Page",
	".aspx":      "Active Server Page",
	".cer":       "Internet security certificate",
	".cfm":       "ColdFusion Markup",
	".css":       "Cascading Style Sheet",
	".htm":       "HTML",
	".html":      "HTML",
	".js":        "JavaScript",
	".jsp":       "Java Server Page",
	".part":      "Partially downloaded",
	".php":       "PHP",
	".py":        "Python",
	".go":        "Go",
	".rb":        "Ruby",
	".rs":        "Rust",
	".lock":      "Lock",
	".yml":       "YAML",
	".yaml":      "YAML",
	".toml":      "TOML",
	".json":      "JSON",
	".rss":       "RSS",
	".xhtml":     "XHTML",
	".md":        "Markdown",
	".markdown":  "Markdown",
	".key":       "Key",
	".odp":       "OpenOffice Impress presentation",
	".pps":       "PowerPoint slide show",
	".ppt":       "PowerPoint presentation",
	".pptx":      "PowerPoint Open XML presentation",
	".c":         "C, C++ source code",
	".class":     "Java class",
	".cpp":       "C++ source code",
	".cs":        "Visual C# source code",
	".h":         "C, C++, and Objective-C header",
	".java":      "Java Source code",
	".sh":        "Bash shell script",
	".swift":     "Swift source code",
	".vb":        "Visual Basic",
	".ods":       "OpenOffice Calc spreadsheet",
	".xlr":       "Microsoft Works spreadsheet",
	".xls":       "Microsoft Excel",
	".xlsx":      "Microsoft Excel Open XML spreadsheet",
	".bak":       "Backup",
	".cab":       "Windows Cabinet",
	".cfg":       "Configuration",
	".cpl":       "Windows Control panel",
	".cur":       "Windows cursor",
	".dll":       "DLL",
	".dmp":       "Dump",
	".drv":       "Device driver",
	".icns":      "macOS X icon resource",
	".ico":       "Icon",
	".ini":       "Initialization",
	".lnk":       "Windows shortcut",
	".msi":       "Windows installer package",
	".sys":       "Windows system",
	".tmp":       "Temporary",
	".3g2":       "3GPP2 multimedia",
	".3gp":       "3GPP multimedia",
	".avi":       "AVI",
	".flv":       "Adobe Flash",
	".h264":      "H.264 video",
	".m4v":       "Apple MP4 video",
	".mkv":       "Matroska Multimedia Container",
	".mov":       "Apple QuickTime movie",
	".mp4":       "MPEG4 video",
	".mpeg":      "MPEG video",
	".mpg":       "MPEG video",
	".rm":        "RealMedia",
	".swf":       "Shockwave flash",
	".vob":       "DVD Video Object",
	".wmv":       "Windows Media Video",
	".docx":      "Microsoft Word",
	".doc":       "Microsoft Word",
	".odt":       "OpenOffice Writer document",
	".pdf":       "PDF",
	".rtf":       "Real Text",
	".tex":       "A LaTeX document",
	".txt":       "Plain text",
	".wps":       "Microsoft Works",
	".wks":       "Microsoft Works",
	".wpd":       "WordPerfect document",
}

// fileNameList is a map from filename to the file type
// associated with it
var fileNameList = map[string]string{
	"Dockerfile":      "Dockerfile",
	"LICENSE":         "License",
	"CONTRIBUTE":      "Contributor README",
	"README":          "README",
	"README.md":       "README Markdown",
	"README.markdown": "README Markdown",
	"Makefile":        "GNU Make",
	"Gemfile":         "Ruby Gem",
	"Rakefile":        "Ruby Rake",
	"config.ru":       "Ruby Config",
	"Vagrant":         "Vagrant VM",
}

// detectByName tries to find the filetype based on the
// file name using the map above
func detectByName(name string) string {
	// Get the content type based off the full file name
	if content, found := fileNameList[name]; found {
		return content
	}

	// Get the content type based off the file extension
	if content, found := extensionList[filepath.Ext(name)]; found {
		return content
	}

	// Get the content type based off the file name without extension
	if content, found := fileNameList[strings.TrimSuffix(name, filepath.Ext(name))]; found {
		return content
	}

	return ""
}
