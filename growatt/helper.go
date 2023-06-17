package growatt

func checkHTMLResponse(body []byte) error {
	if len(body) > 0 && htmlTagsRegExp.Find(body) != nil {
		return ErrResponseMustNotBeHTML
	}

	return nil
}
