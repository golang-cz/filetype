package main

func (m *fileMatcher) RegisterJPEG() {
	// JPEG, http://fileformats.archiveteam.org/wiki/JPEG
	//matcher.Register("FF D8 FF", "generic jpg") // generic
	m.Register("FF D8 FF DB", "jpg")
	m.Register("FF D8 FF E0 00 10 4A 46 49 46 00 01", "jpg")
	m.Register("FF D8 FF EE", "generic jpg")

	// JPEG/EXIF
	//m.Register("FF D8 FF E1 ?? ?? 45 78 69 66 00 00", "jpg")
	m.Register("FF D8 FF E1", "JPEG EXIF")

	// JPEG/JFIF
	m.Register("FF D8 FF E0", "JPEG JFIF")

	// JPEG/SPIFF
	m.Register("FF D8 FF E8", "JPEG SPIFF")
}

func (m *fileMatcher) RegisterJPEG2000() {
	// JPEG 2000 (generic prefix)
	m.Register("00 00 00 0C 6A 50 20 20 0D 0A 87 0A", "jp2")
	// JPEG 2000, JP2, http://fileformats.archiveteam.org/wiki/JP2
	m.Register("00 00 00 0C 6A 50 20 20 0D 0A 87 0A {1} {2} {3} {4} 66 74 79 70 6A 70 32 20", "jp2")
	// http://fileformats.archiveteam.org/wiki/JPX
	m.Register("00 00 00 0C 6A 50 20 20 0D 0A 87 0A {1} {2} {3} {4} 66 74 79 70 6A 70 78 20", "jpx")
	// http://fileformats.archiveteam.org/wiki/JPM
	m.Register("00 00 00 0C 6A 50 20 20 0D 0A 87 0A 00 00 00 14 66 74 79 70 6A 70 6D 20", "jpm")
	// http://fileformats.archiveteam.org/wiki/MJ2, Motion JPEG 2000, video+audio
	m.Register("00 00 00 0C 6A 50 20 20 0D 0A 87 0A {1} {2} {3} {4} 66 74 79 70 6D 6A 70 32", "mj2")
	// http://fileformats.archiveteam.org/wiki/JPEG_2000_codestream
	m.Register("FF 4F FF 51", "jpc")
}
