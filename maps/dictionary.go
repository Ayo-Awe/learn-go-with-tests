package maps

type Dictionary map[string]string

func (d *Dictionary) Search(word string) string {
	dict := *d
	return dict[word]
}
