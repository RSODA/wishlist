package media

func StaticURL(host, filename string) string {
	return "http://" + host + "/api/v1/static/" + filename
}
