package extractor

func TestData() *RunResult {
	return &RunResult{
		Repository: &Repository{
			StarsCount: 100,
		},
	}
}
