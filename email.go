package main

import "github.com/SlyMarbo/gmail"

func main() {
	email := gmail.Compose("Email subject", "Email body")
	email.From = "username@gmail.com"
	email.Password = "password"

	// Defaults to "text/plain; charset=utf-8" if unset.
	email.ContentType = "text/html; charset=utf-8"

	// Normally you'll only need one of these, but I thought I'd show both.
	email.AddRecipient("recipient@example.com")
	email.AddRecipients("another@example.com", "more@example.com")

	err := email.Send()
	if err != nil {
		// handle error.
	}
}