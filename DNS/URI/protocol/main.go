package main

func getMailtoLinkForEmail(email string) string {
	const protocol = "mailto:"
	return protocol + email
}
