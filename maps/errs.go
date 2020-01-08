package maps

type DictionaryErr string

const (
	errKeywordNotFind      = DictionaryErr("could not find the keyword you were looking for")
	errWordAlreadyExist    = DictionaryErr("the keyword you insert is already exist")
	errKeywordDoesNotExist = DictionaryErr("the keyword does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}
