package model

type Actions []Action

type BE struct {
	Title                   string   `json:"title"`
	Content                 string   `json:"content"`
	Id                      int      `json:"id"`
	OriginalLocale          string   `json:"originalLocale"`
	CreationTimeSeconds     int      `json:"creationTimeSeconds"`
	AuthorHandle            string   `json:"authorHandle"`
	Locale                  string   `json:"locale"`
	ModificationTimeSeconds int      `json:"modificationTimeSeconds"`
	AllowViewHistory        bool     `json:"allowViewHistory"`
	Tags                    []string `json:"tags"`
	Rating                  int      `json:"rating"`
}

type C struct {
	Text                string `json:"text"`
	Rating              int    `json:"rating"`
	Id                  int    `json:"id"`
	CreationTimeSeconds int    `json:"creationTimeSeconds"`
	CommentatorHandle   string `json:"commentatorHandle"`
	Locale              string `json:"locale"`
	ParentCommentId     int    `json:"parentCommentId"`
}

type Action struct {
	TimeSeconds int64   `json:"timeSeconds"`
	BlogEntry   *BE `json:"blogEntry"`
	Comment     *C   `json:"comment"`
}

type UserLogIn struct {
	Username string `json:"id1"`
	Id2 string `json:"id2"`
}

type UserSignIn struct {
	Id1 string `json:"id1"`
	Id2 string `json:"id2"`
}
