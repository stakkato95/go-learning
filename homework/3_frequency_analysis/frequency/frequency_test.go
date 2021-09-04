package frequency_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stakkato95/frequency/frequency"
	"github.com/stretchr/testify/assert"
)

//removing stopwords, removing punctuation and counting words
//https://demos.datasciencedojo.com/demo/stopwords/
//https://www.browserling.com/tools/remove-punctuation
//http://www.writewords.org.uk/word_count.asp
var cases = []struct {
	Text     string
	Expected map[string]bool
}{
	{
		"My choice title Boring Go  properly written Go boring It weird write boring topic I explain Go small feature set step modern programming lan guages Wellwritten Go programs tend straightforward sometimes repetitive Theres inheritance generics  aspectoriented programming function overloading certainly operator overloading Theres pattern matching named parameters exceptions To horror many pointers Gos concurrency model unlike languages its based ideas 1970s algorithm used garbage collector In short Go feels like throwback And thats point Boring trivial Using Go correctly requires understanding features intended fit together While write Go code looks like Java Python youre going unhappy result wonder fuss  Thats comes  It walks features Go explaining best write idiomatic code grow When comes building things  boring great No wants person drive bridge built untested techniques engineer cool The modern depends software depends bridges perhaps  Yet many programming languages add features without thinking impact maintainability codebase Go intended building programs  programs modified dozens developers dozens years Go boring thats fantastic I hope teaches build exciting projects boring code",
		map[string]bool{
			"go":          true,
			"boring":      true,
			"write":       true,
			"thats":       true,
			"programs":    true,
			"programming": true,
			"features":    true,
			"code":        true,
			"theres":      true,
			"overloading": true,
		},
	},
	{
		"This targeted developers looking pick second  fifth lan guage The focus people Go This ranges dont anything Go cute mascot already worked Go tutorial written Go code The focus Learning Go isnt write programs Go its write Go idiomatically More experienced Go developers advice best newer features language The important reader wants learn write Go code looks like Go Experience assumed tools developer trade version control preferably Git IDEs Readers familiar basic computer science concepts like concurrency abstraction explains Go Some code examples downloadable GitHub dozens tried online The Go Playground While internet connection isnt required helpful reviewing executable examples Since Go often used build HTTP servers examples assume familiarity basic HTTP concepts While Gos features languages Go makes different trade offs programs written different structure Learning Go starts look ing set Go development environment covers variables types control structures functions If tempted skip material resist urge  It often details Go code idiomatic Some seems obvious glance actually subtly surprising depth",
		map[string]bool{
			"go":       true,
			"the":      true,
			"code":     true,
			"write":    true,
			"examples": true,
			"written":  true,
			"while":    true,
			"trade":    true,
			"this":     true,
			"some":     true,
		},
	},
}

func TestFrequency(t *testing.T) {
	for i, caze := range cases {
		t.Run(fmt.Sprintf("case %d.", i+1), func(t *testing.T) {
			actual := frequency.Frequency(caze.Text)
			assert.True(t, reflect.DeepEqual(caze.Expected, arrayToSet(actual)))
		})
	}
}

func TestFrequencyWithEmptyInput(t *testing.T) {
	assert.Nil(t, frequency.Frequency(""))
}

func arrayToSet(array []string) map[string]bool {
	out := map[string]bool{}

	for _, item := range array {
		out[item] = true
	}

	return out
}
