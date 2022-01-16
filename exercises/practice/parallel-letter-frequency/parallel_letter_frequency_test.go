package letter

import (
	"reflect"
	"testing"
)

// In the separate file parallel_letter_frequency.go, you are given a function,
// Frequency(), to sequentially count letter frequencies in a single text.
//
// Perform this exercise on parallelism using Go concurrency features.
// Make concurrent calls to Frequency and combine results to obtain the answer.

var (
	ralph_ellison = `I am an invisible man. No, I am not a spook like those who haunted Edgar Allan Poe;
nor am I one of your Hollywood-movie ectoplasms. I am a man of substance, of flesh and bone,
fiber and liquids—and I might even be said to possess a mind. I am invisible, understand,
simply because people refuse to see me. Like the bodiless heads you see sometimes in circus sideshows,
it is as though I have been surrounded by mirrors of hard, distorting glass.
When they approach me they see only my surroundings, themselves, or figments of their imagination—indeed,
everything and anything except me.`

	dostoevsky = `I am a sick man. ... I am a spiteful man. I am an unattractive man. I believe my liver is diseased.
However, I know nothing at all about my disease, and do not know for certain what ails me.
I don't consult a doctor for it, and never have, though I have a respect for medicine and doctors.
Besides, I am extremely superstitious, sufficiently so to respect medicine, anyway
(I am well-educated enough not to be superstitious, but I am superstitious). No, I refuse to consult a doctor from spite.
That you probably will not understand. Well, I understand it, though. Of course,
I can't explain who it is precisely that I am mortifying in this case by my spite:
I am perfectly well aware that I cannot "pay out" the doctors by not consulting them;
I know better than anyone that by all this I am only injuring myself and no one else. But still, if I don't consult a doctor it is from spite. My liver is bad, well then let it hurt even worse!
`

	mark_twain = `You don't know about me without you have read a book by the name of The Adventures of Tom Sawyer;
but that ain't no matter. That book was made by Mr. Mark Twain, and he told the truth, mainly.
There was things which he stretched, but mainly he told the truth. That is nothing. I never seen anybody but lied one time or another,
without it was Aunt Polly, or the widow, or maybe Mary. Aunt Polly - Tom's Aunt Polly, she is - and Mary,
and the Widow Douglas is all told about in that book, which is mostly a true book, with some stretchers, as I said before.`

	gary_provost = `This sentence has five words. Here are five more words. Five-word sentences are fine.
But several together become monotonous. Listen to what is happening. The writing is getting boring.
The sound of it drones. It’s like a stuck record. The ear demands some variety. Now listen. I vary the sentence length,
and I create music. Music. The writing sings. It has a pleasant rhythm, a lilt, a harmony. I use short sentences.
And I use sentences of medium length. And sometimes, when I am certain the reader is rested, I will engage him with a sentence of considerable length,
a sentence that burns with energy and builds with all the impetus of a crescendo, the roll of the drums, the crash of the cymbals–sounds that say listen
to this, it is important.`

)


func OriginalFrequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func TestConcurrentFrequency(t *testing.T) {
	seq := OriginalFrequency(ralph_ellison + dostoevsky + mark_twain + gary_provost)
	con := ConcurrentFrequency([]string{ralph_ellison, dostoevsky, mark_twain, gary_provost})
	if !reflect.DeepEqual(con, seq) {
		t.Fatal("ConcurrentFrequency wrong result")
	}
}

func TestSequentialFrequency(t *testing.T) {
	oSeq := OriginalFrequency(ralph_ellison + dostoevsky + mark_twain + gary_provost)
	seq := Frequency(ralph_ellison + dostoevsky + mark_twain + gary_provost)
	if !reflect.DeepEqual(oSeq, seq) {
		t.Fatal("Frequency wrong result")
	}
}

func BenchmarkSequentialFrequency(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		Frequency(ralph_ellison + dostoevsky + mark_twain + gary_provost)
	}
}

func BenchmarkConcurrentFrequency(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		ConcurrentFrequency([]string{ralph_ellison, dostoevsky, mark_twain, gary_provost})
	}
}
