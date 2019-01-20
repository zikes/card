<h1 align="center">
<pre>
╔══════╗
║ CARD ║
╚══════╝</pre>
</h1>

```go
fmt.Print(card.Section{
  Header:       "Jason Hutchinson",
  HeaderMargin: 1,
  Items: []card.Item{
    {Label: "Email", Data: "zikes@zikes.me"},
    {Label: "Web", Data: "https://zikes.me"},
    {Label: "Twitter", Data: "https://twitter.com/zikes"},
    {Label: "GitHub", Data: "https://github.com/zikes"},
  },
})

// Jason Hutchinson
// ----------------
//
//   Email: zikes@zikes.me
//     Web: https://zikes.me
// Twitter: https://twitter.com/zikes
//  GitHub: https://github.com/zikes
```

`card` is a library which will create sections of content aligned by label 
length, with optional headers and configurable spacing after each header.

# Section

```go
type Section struct {
  Header       string
  HeaderMargin int
  ListIndent   int
  Items        []Item
}
```

A Section is an optional header and list of `Item`s.

### Header

The string to use as the header. A line of hyphens will be generated to be
placed below the header to match its length.

### HeaderMargin

The number of blank lines to place after the header. Default 0.

### ListIndent

The number of spaces to precede the list of `Item`s when printed. Default 0.

### Items

A slice of `card.Item`s. If an `Item` with a blank `Label` and `Data` field
are included, a blank line will be inserted in that spot. This is useful
for creating gaps in the list while retaining label alignment.

# Item

```go
type Item struct {
  Label string
  Data  string
}
```

When printed via the parent `Section`'s `String` method, an `Item`'s `Label`
will be automatically bolded and padded to be aligned according to the colon
which is automatically inserted after it.
