package maps

type Dictionary map[string]string

func (d Dictionary) Search(keyword string) (string, DictionaryErr) {
	definition, ok := d[keyword]
	if !ok {
		return "", errKeywordNotFind
	}

	return definition, ""
}

func (d Dictionary) Add(keyword, definition string) DictionaryErr {
	_, err := d.Search(keyword)

	switch err {
	case "":
		return errWordAlreadyExist
	case errKeywordNotFind:
		d[keyword] = definition
	default:
		return ""
	}

	return ""
}

func (d Dictionary) Update(keyword, definition string) DictionaryErr {
	_, err := d.Search(keyword)

	switch err {
	case "":
		d[keyword] = definition
		return ""
	case errKeywordNotFind:
		return errKeywordDoesNotExist
	default:
		return err
	}
}

func (d Dictionary) Delete(keyword string) {
	delete(d, keyword)
}
