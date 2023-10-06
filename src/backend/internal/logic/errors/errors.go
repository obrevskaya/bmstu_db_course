package errors

import "errors"

var (
	ErrGet          = errors.New("Don't get user")
	ErrAccess       = errors.New("Don't necessary rights")
	ErrAlreadyExist = errors.New("Already Exists")
	ErrLenPassword  = errors.New("Password less than 6 symbols")
	ErrMail         = errors.New("Email incorrect")
	ErrPoints       = errors.New("Sum of points is negative")
	ErrInsert       = errors.New("Error of insert")
	ErrDelete       = errors.New("Error of delete")
	ErrUpdate       = errors.New("Error of update")
	ErrGetDB        = errors.New("Error of get in data")
	ErrAdd          = errors.New("Error of add")
	ErrCntWine      = errors.New("Count of wine is not enough")
	ErrPrice        = errors.New("Sum of price is negative")
	ErrNotFound     = errors.New("Record not found")
)
