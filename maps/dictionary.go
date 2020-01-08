package maps

type Dictionary map[string]string

func (d Dictionary) Search(keyword string) (string, error) {
	definition, ok := d[keyword]
	if !ok {
		return "", errKeywordNotFind
	}

	return definition, nil
}

func (d Dictionary) Add(key, val string) {
	d[key] = val
}
