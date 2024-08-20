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

func main() {
        var free bool
        var random bool
        var self bool
        flag.BoolVar(&free, "P", false, "No prefixes (in, of, by, etc.)")
        flag.BoolVar(&random, "s", false, "Shuffle shortened words")
        flag.BoolVar(&self, "D", false, "Ignore dictionary")
        flag.Parse()
        done := strings.ToLower(flag.Arg(0))
        if len(flag.Args()) == 0 || strings.Trim(done, " ") == "" {
                fmt.Println("Crab - CReate ABbreviations\n(C) 2024, Alexander Gavrilovets\n\n")
                flag.Usage()
                os.Exit(1)
        }
        var text []string
        otherw := []string{}
        const PREFIXES = "\\b (to|from|in|is|of|at|about|and|or|may|might|by|[.,!?]) \\b|(\\-|\\+)"
        var no_I_Q *regexp.Regexp

        if free {
                temp_prefixes, _ := regexp.Compile(PREFIXES)
                done = temp_prefixes.ReplaceAllString(done, " ")
        }

        if !self {
                dictionary, file_removed := os.ReadFile("./dict.txt")
                dictionary_txt := strings.Replace(string(dictionary), "\r\n", "\n", -1)
                text_temp := strings.Split(string(dictionary_txt), "\n")
                text = slices.Delete(text_temp, len(text_temp)-1, len(text_temp))
                if file_removed != nil {
                        fmt.Println("Oops, dictionary (dict.txt) removed!")
                        os.Exit(1)
                }

                no_I_Q, _ = regexp.Compile("[!?]")
                done = no_I_Q.ReplaceAllString(done, "")
                featured := []string{}
                featured_one := []string{}
                for _, line := range text {
                        featured = append(featured, strings.Split(line, " - ")[1])
                        featured_one = append(featured_one, strings.Split(line, " - ")[0])
                }
                for _, iter := range text {
                        r, _ := regexp.Compile("\\b" + strings.Split(iter, " - ")[1] + "\\b")
                        done = r.ReplaceAllString(done, strings.Split(iter, " - ")[0])
                }

                for _, find_anonymous := range strings.Split(done, " ") {
                        is_good, _ := regexp.MatchString(PREFIXES, find_anonymous)
                        if (!slices.Contains(featured, find_anonymous)) && !is_good && !slices.Contains(featured_one, find_anonymous) {
                                otherw = append(otherw, find_anonymous)
                        }
                }
        }
        if self {
                otherw = strings.Split(done, " ")
        }
        for _, fi := range otherw {
                done = strings.ReplaceAll(done, fi, fi[0:(rand.Intn(len(fi))+1)])
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
