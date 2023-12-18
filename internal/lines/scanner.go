package lines

type LineScanner interface {
	Scan() (string, error)
}
