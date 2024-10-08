package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"slices"
	"strings"
)

func CheckSyntax(line *[]string) {

	lines := []string{}
	for _, checkline := range *line {
		count := len(regexp.MustCompile(" ").FindAllString(strings.Trim(checkline, " "), -1))
		if regexp.MustCompile("(^( {1,}))").MatchString(checkline) || checkline == "" {
			lines = append(lines, fmt.Sprint(0))
		} else if count == 2 {
			lines = append(lines, fmt.Sprint(2))
		} else if count > 2 {
			lines = append(lines, fmt.Sprint(3))
		}
	}
	expected := slices.Clone(lines)
	slices.Sort(expected)
	slices.Reverse(expected)
	if slices.Contains(lines, "0") {
		fmt.Println(`Empty lines in dict.txt are not allowed.
HINT: Remove empty lines.`)
		os.Exit(1)
	} else if !(regexp.MustCompile(strings.Join(expected, "-")).MatchString(strings.Join(lines, "-"))) {
		fmt.Println(`Invalid syntax of dict.txt
HINT: The abbreviations which contain more than word must be placed before the abbreviations which contain 1 word.`)
		os.Exit(1)
	}
}
func main() {
	var free bool
	var random bool
	var repeats uint
	var self bool
	var ultrashort bool
	var exclude bool
	flag.BoolVar(&free, "P", false, "No prefixes (in, of, by, the, etc.)")
	flag.BoolVar(&random, "s", false, "Shuffle shortened words")
	flag.BoolVar(&self, "D", false, "Ignore dictionary")
	flag.BoolVar(&exclude, "x", false, "Exclude some words")
	flag.BoolVar(&ultrashort, "u", false, "Maximum shortings")
	flag.UintVar(&repeats, "r", 1, "Number of repeats")
	flag.Parse()
	done := strings.ToLower(regexp.MustCompile(" +").ReplaceAllString(flag.Arg(0), " "))
	if len(flag.Args()) == 0 || strings.Trim(done, " ") == "" {
		fmt.Println("Crab - CReate ABbreviations\n(C) 2024, Alexander Gavrilovets")
		flag.Usage()
		fmt.Println("Documentation is here: https://github.com/LazataknesSoftware/crab")
		os.Exit(1)
	}
	var text []string
	otherw := []string{}
	const PREFIXES = "\\b (to|from|in|is|of|at|about|for|who|which|what|when|where|whom|the|and|or|may|might|by|[.,!?])\\b|(\\-|\\+|\\b [Aa]\\b)"
	var no_I_Q *regexp.Regexp

	if free {
		temp_prefixes, _ := regexp.Compile(PREFIXES)
		done = temp_prefixes.ReplaceAllString(done, "")
	}
	no_I_Q = regexp.MustCompile(`[^\w ]`)
	done = strings.Trim(no_I_Q.ReplaceAllString(done, ""), " ")

	FutureTemplate := []string{}

	if exclude {
		for _, xr := range strings.Split(done, " ") {
		IsRemoved := rand.Intn(100) + 1
		if IsRemoved >= 50 {
			FutureTemplate = append(FutureTemplate, xr)
		}
		
	}
	done = strings.Join(FutureTemplate, " ")
}

	backup := done
	var i uint = 0

	dictionary, file_removed := os.ReadFile("./dict.txt")
	dictionary_txt := strings.Replace(string(dictionary), "\r\n", "\n", -1)
	text = strings.Split(string(dictionary_txt), "\n")
	if file_removed != nil {
		fmt.Println("Oops, dictionary (dict.txt) removed!")
		os.Exit(1)
	}
	CheckSyntax(&text)

	for i = 0; i < repeats; i++ {
		done = backup
		if !self {
			backup = done
			featured := []string{}
			featured_one := []string{}
			for _, line := range text {
				featured = append(featured, strings.Split(line, " - ")[1])
				featured_one = append(featured_one, strings.Split(line, " - ")[0])
			}
			for _, iter := range text {
				r, _ := regexp.Compile("\\b(" + strings.Split(iter, " - ")[1] + ")\\b")
				done = r.ReplaceAllString(done, strings.Split(iter, " - ")[0])
			}

			otherw = []string{}
			for _, find_anonymous := range strings.Split(done, " ") {
				is_good, _ := regexp.MatchString(PREFIXES, find_anonymous)
				if (!slices.Contains(featured, find_anonymous)) && !is_good && !slices.Contains(featured_one, find_anonymous) {
					otherw = append(otherw, find_anonymous)
				}
			}
		}
		if self {
			backup = done
			otherw = strings.Split(done, " ")
		}
		for _, fi := range otherw {
			rnd := rand.Intn(len(fi)) + 1
			if ultrashort && rnd < 2 {
				rnd = 1
			} else if ultrashort && rnd < 3 {
				rnd = 2
			} else if ultrashort {
				rnd = rand.Intn(3) + 1
			}
			done = strings.ReplaceAll(done, fi, fi[0:rnd])
		}

		if !random {
			fmt.Println(strings.ToLower(strings.ReplaceAll(done, " ", "")))
		} else {
			redone := strings.Split(done, " ")
			rand.Shuffle(len(redone), func(i, r int) {
				redone[i], redone[r] = redone[r], redone[i]
			})
			fmt.Println(strings.ToLower(strings.Join(redone, "")))
		}
	}
}
