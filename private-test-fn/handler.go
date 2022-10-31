package function

// Handle a serverless request
func Handle(req []byte) string {
	return "Pulling from private repos is working just fine for ghcr.io\n"
}
