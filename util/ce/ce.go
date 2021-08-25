package ce

import "errors"

type ChainError interface {
	AE(error) ChainError
	EF() error
}

type chainError struct {
	errs []error
}

func (c *chainError) AE(err error) ChainError {
	c.errs = append(c.errs, err)
	return c
}

func (c *chainError) EF() error {
	var message string
	var err error
	if len(c.errs) != 0 {
		for _, er := range c.errs {
			if er != nil {
				if len(message) == 0 {
					message += er.Error()
				} else {
					message += " => " + er.Error()
				}
			}
		}
		err = errors.New(message)
		return err
	}
	return nil
}

func NewChainError() ChainError {
	return &chainError{}
}
