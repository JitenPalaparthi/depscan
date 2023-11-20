package implement

import "errors"

var (
	ErrNilConfig               = errors.New("nil config")
	ErrEmptyPath               = errors.New("empty path")
	ErrInvalidOutfile          = errors.New("invalid outfile")
	ErrInvalidOutfileFormat    = errors.New("invalid outfile format.It must be json | yaml | yml")
	ErrNewImplement            = errors.New("use New function to create Implement object")
	ErrNoDataToGenerateOutfile = errors.New("no data to generate output")
	ErrPathDoesNotExist        = errors.New("path does not exist")
	ErrUnsupportedNPMVersion   = errors.New("unsupported npm version.Currently it supports only lockfileVersion-1,2 and 3 only.Old formats are not supported")
	ErrUnsupportedFile         = errors.New("unsupported file")
	ErrNoPackageFileFound      = errors.New("no package dependency file found")
	ErrInvalidNoOfFiles        = errors.New("invalid number of files")
)
