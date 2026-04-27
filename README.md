<a href="#"></a>
<DIV ALIGN="CENTER"><H1>Об этой программе</H1></DIV>

**Crab** (🦀) - это английская аббревиатура из "**CR**eate **AB**breviations". Если у вас есть ключевые слова для проекта, но вы не можете придумать ему название, то **Crab** вам поможет.

---

<DIV ALIGN="CENTER"><H1>Поддерживаемые операционные системы</H1></DIV>

* Windows 10 1607 или новее
* Android: [Termux](https://github.com/termux/termux-app)
* <S>Linux</S>
* <S>MacOS</S>

---

<DIV ALIGN="CENTER"><H1>Где я могу это загрузить?</H1></DIV>

Перейдите в раздел "Releases" и выберите `crab.exe`, если у вас Windows, или `crab`, если вы на Android и там установлен Termux.

---

<DIV ALIGN="CENTER"><H1>Опции этой программы</H1></DIV>

## Опция `-D`
Посмотрите на следующую команду:

`$ crab "Lazataknes Software"`

Результатом будет `lazataknsw`. Не понравилось, как `software` сократилось в `sw`? Вам нужна опция `-D`!

Попробуем!

`$ crab -D "Lazataknes Software"`

Результатом будет `lazatasof`.

## Опция `-P`
Посмотрите на следующую команду:

`$ crab "Mike at home"`

Результатом будет `mikathom`.

`at` всё испортило? Секунду! вы можете запустить `crab` с опцией `-P`! Попробуем!

`$ crab -P "Mike at home"`

Результатом будет `mikeho`.

## `-s` flag
Посмотрите на следующую команду:

`$ crab "What is laptop?"`

Результатом будет `whaslaptop`.

Хотите перемешать сокращения? Без проблем! Опция `-s` Вам поможет.

Попробуйте:

`$ crab -s "What is laptop?"`

Результатом будет `swhl`.

---

## `-r` flag
Посмотрите на следующую команду:

`$ crab "computer game"`

Результатом будет `compg`.

Не хочется вводить `crab "computer game"` каждый раз? Тогда опция `-r` спешит на помощь!

Попробуйте:

`$ crab -r=5 "computer game"`

Результатом будет
```
compgame
compg
compgam
compga
compga
```

## `-u` flag
Посмотрите на следующую команду:

`$ crab "small company which makes games"`

Результатом будет `smcomwhichmg`.

Эта аббревиатура получилась слишком длинной, не так ли? Для решения таких проблем есть опция `-u`, которая максимально сокращает слова.

Попробуйте:

`$ crab -u "small company which makes games"`

Результатом будет `scowmgam`.

## `-x` flag
Посмотрите на следующую команду:

`$ crab -u "small company which makes games"`

Помните предыдущую аббревиатуру? Она короткая, но если вы хотите сократить ее еще больше, вы можете использовать опцию `-x`, которая исключается некоторые ключевые слова.

Попробуйте:

`$ crab -x -u "small company which makes games"`

Результатом будет `scowhm`.

<DIV ALIGN="CENTER"><H1>Словарь</H1></DIV>

Все аббревиатуры по типу `sw - software`, `tv - television`, `0 - nothing` и т.д. `crab` получает из `dict.txt`. 

⚠️ **Осторожно:** Если вы удалите `dict.txt`, то `crab` не запуститься (однако вы можете это обойти с помощью опции `-D`).

## Добавление или удаление сокращений
### Добавление
Если вы хотите добавить слово, то добавьте выражение вида `abbr - abbreviation` в конец `dict.txt`.

✔️ Примеры:

* `mem - memory`
* `gpu - videocard`
* `ver - version`
* `kb - keyboard`
* `snd - sound`

➡️ **Замечания**: 
> 1. Аббревиатуры, состоящие более чем из одного слова, должны быть первыми в `dict.txt`, иначе `crab` не запустится!
> 2. Пустые линии и пробелы в `dict.txt` запрещены.

❎ **НЕ** пишите так:
* `V - volume, version`
* `X - extended, execute`

### Удаление
Просто удалите ненужную аббревиатуру.

---
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
Just remove disliked abbreviation.