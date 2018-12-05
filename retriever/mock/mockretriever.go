package mock

type Retriever struct {
	Content string
}

func (m *Retriever) Get(url string) string {
	return m.Content
}

func (m *Retriever) Post(url string, contents map[string]string) string {
	m.Content = contents["contents"]
	return "ok"
}
