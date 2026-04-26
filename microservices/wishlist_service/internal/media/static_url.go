package media

func StaticUrl(host, filename string) string {
	return "http://" + host + "/api/v1/static/" + filename
}
