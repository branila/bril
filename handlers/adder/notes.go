package adder

func getNotes(note string) []string {
	if note == "" {
		return nil
	}

	return []string{note}
}
