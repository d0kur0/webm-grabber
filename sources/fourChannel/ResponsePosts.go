package fourChannel

type ResponsePosts struct {
	Posts []struct {
		Filename      string `json:"filename"`
		FileExtension string `json:"ext"`
		FileId        int    `json:"tim"`
	}
}
