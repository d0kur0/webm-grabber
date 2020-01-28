package twoChannel

type ResponsePosts struct {
	Threads []struct {
		Posts []struct {
			Files []struct {
				Name    string `json:"fullname"`
				Path    string `json:"path"`
				Preview string `json:"thumbnail"`
			}
		}
	}
}
