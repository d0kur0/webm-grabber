package fourChannel

type ResponsePosts struct {
	Posts []struct {
		Filename      string `json:"filename"`
		FileExtension string `json:"ext"`
		FileId        int64  `json:"tim"`
	}
}

type ResponseThreads []struct {
	Threads []struct {
		Id int64 `json:"no"`
	}
}
