package types

type ChapterType string

const (
	Codeblock    ChapterType = "codeblock"
	Example      ChapterType = "example"
	Grid         ChapterType = "grid"
	Image        ChapterType = "image"
	List         ChapterType = "list"
	ListNoBullet ChapterType = "list-no-bullet"
	Notice       ChapterType = "notice"
	Quote        ChapterType = "quote"
	ReadAloud    ChapterType = "read-aloud"
	Rule         ChapterType = "rule"
	Sidebar      ChapterType = "sidebar"
	Table        ChapterType = "table"
	Text         ChapterType = "text"
)
