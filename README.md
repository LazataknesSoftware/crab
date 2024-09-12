<DIV ALIGN="CENTER"><H1>About this tool</H1></DIV>

**Crab** (🦀) is abbreviation from "**CR**eate **AB**breviations". If you have keywords for your project, but you can't invent name for your project, then **Crab** will help you.

---

<DIV ALIGN="CENTER"><H1>Requirements</H1></DIV>

* Windows 10 1607 or newer
* Android: [Termux](https://github.com/termux/termux-app)
* <S>Linux</S>
* <S>MacOS</S>

---

<DIV ALIGN="CENTER"><H1>Where to get it?</H1></DIV>

Go to Releases and select `crab.exe` if you are on Windows or `crab` for Android (Termux)

---

<DIV ALIGN="CENTER"><H1>Flags of this tool</H1></DIV>

## `-D` flag
Consider following command line:

`$ crab "Lazataknes Software"`

Its output will be `lazataknsw`. Did not you like `software` was shorten to `sw`? There is `-D` goes to help you!

Let's try it!

`$ crab -D "Lazataknes Software"`

Its output will be `lazatasof`.

## `-P` flag
Consider following command line:

`$ crab "Mike at home"`

Its output will be `mikathom`.

Did `at` spoil all? Don't worry! You can run `crab` with `-P` flag! Let's try it!

`$ crab -P "Mike at home"`

Its output will be `mikeho`.

## `-s` flag
Consider following command line:

`$ crab "What is laptop?"`

Its output will be `whaslaptop`.

Do you want to remix it? No problems! `-s` flag will help you.

Try it:

`$ crab -s "What is laptop?"`

Its output will be: `swhl`.

---

## `-r` flag
Consider following command line:

`$ crab "computer game"`

Its output will be: `compg`.

Did you tire to enter `crab "computer game"` every time? Then `-r` flag will help you!

Try it:

`$ crab -r=5 "computer game"`

Its output will be:
```
compgame
compg
compgam
compga
compga
```

## `-u` flag
Consider following command line:

`$ crab "small company which makes games"`

Its output will be: `smcomwhichmg`.

This abbreviation is too long, isn't it? There is `-u` flag which shortens word maximally.

Try it:

`$ crab -u "small company which makes games"`

Its output will be: `scowmgam`.

## `-x` flag
Consider following command line:

`$ crab -u "small company which makes games"`

Do you remember the past abbreviation? It is short, but if you want to more, you can use `-x` flag, which excludes some keywords.

Try it:

`$ crab -x -u "small company which makes games"`

Its output will be: `scowhm`.

<DIV ALIGN="CENTER"><H1>Dictionary</H1></DIV>

All abbreviations like `sw - software`, `tv - television`, `0 - nothing`, etc. `crab` takes from `dict.txt`. 

⚠️ **Caution:** If you will remove `dict.txt`, then `crab` won't run (however, you can bypass it by `-D` flag).

## Adding or removing words
### Adding
If you want to add word, then append abbreviation like `abbr - abbreviation` into end of `dict.txt`.

✔️ Examples:

* `mem - memory`
* `gpu - videocard`
* `ver - version`
* `kb - keyboard`
* `snd - sound`

➡️ **Notes**: 
> 1. The abbreviations which values contain more than word, must be put first in `dict.txt`, otherwise `crab` won't run!
> 2. Empty lines and spaces are forbidden in `dict.txt`

❎ Please, do **NOT** write like:
* `V - volume, version`
* `X - extended, execute`

### Removing
Just remove disliked abbreviation! It is simple, isn't it?

### Behavior when `dict.txt` has values-duplicates (e.g. `V - volume, vol - volume`)
The first key `(V)` will have mory priority than the second.

<DIV ALIGN="CENTER">
  <H1>When you will have questions, then...</H1>
  <span>...make an issue!</span>
</DIV>
