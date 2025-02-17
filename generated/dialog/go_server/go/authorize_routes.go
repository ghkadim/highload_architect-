package openapi

var AuthorizeRoutes = []struct {
	Path   string
	Method string
}{

	{
		Path:   "/dialog/{user_id}/send",
		Method: "POST",
	},

	{
		Path:   "/dialog/{user_id}/list",
		Method: "GET",
	},

	{
		Path:   "/dialog/message/{message_id}/read",
		Method: "PUT",
	},
}
