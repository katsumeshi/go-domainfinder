package thesaurus

// for other api if you extends
type Thesaurus interface {
	Synonyms(term string) ([]string, error)
}
