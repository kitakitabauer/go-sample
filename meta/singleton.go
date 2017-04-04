package meta

type MetaWordser interface {
	PositiveWords() []string
	NegativeWords() []string
}

// metaWords is implemented by MetaWordser
type metaWords struct {
	positiveWords []string
	negativeWords []string
}

func (m *metaWords) PositiveWords() []string {
	return m.positiveWords
}

func (m *metaWords) NetativeWords() []string {
	return m.negativeWords
}

var sharedInstance = &metaWords{}

// GetInstance is singleton of meta words
func GetInstance() MetaWordser {
	return sharedInstance
}


