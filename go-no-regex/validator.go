package main

import (
	"errors"
	"net"
	"net/mail"
	"net/url"
	"strings"
	"unicode"

	"golang.org/x/net/html"
)

var (
	ErrInvalidEmail    = errors.New("invalid email")
	ErrDomainNotFound  = errors.New("domain not found")
	ErrInvalidPassword = errors.New("invalid password")
	ErrNoAnchorFound   = errors.New("anchor not found")
	ErrNoImplemented   = errors.New("not yet implemented")
)

func validateMail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return ErrInvalidEmail
	}
	parts := strings.Split(email, "@")
	_, err = net.LookupMX(parts[1])
	if err != nil {
		return ErrDomainNotFound
	}
	return nil
}

func validatePassword(pass string) error {
	isMoreThan8 := len(pass) > 8
	var isLower, isUpper, isSym bool
	for _, r := range pass {
		if !isLower && unicode.IsLower(r) {
			isLower = true
		}
		if !isUpper && unicode.IsUpper(r) {
			isUpper = true
		}
		if !isSym && (unicode.IsSymbol(r) || unicode.IsPunct(r)) {
			isSym = true
		}
	}
	isValid := isMoreThan8 && isLower && isUpper && isSym
	if !isValid {
		return ErrInvalidPassword
	}
	return nil
}

func getURLFromFristAnchor(htmlString string) (*url.URL, error) {
	root, err := html.Parse(strings.NewReader(htmlString))
	if err != nil {
		return nil, ErrNoAnchorFound
	}
	var f func(n *html.Node) (*url.URL, error)
	f = func(n *html.Node) (*url.URL, error) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					u, err := url.Parse(a.Val)
					if err != nil {
						return nil, ErrNoAnchorFound
					}
					return u, nil
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			u, err := f(c)
			if err == nil {
				return u, err
			}
		}
		return nil, ErrNoAnchorFound
	}

	return f(root)
}
