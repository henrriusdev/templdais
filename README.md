
# TemplDais

[Templ Components](https://github.com/a-h/templ) maded with [Tailwind](https://tailwindcss.com) and [DaisyUI](https://daisyui.com), for easy use.


## Badges

![GitHub Tag](https://img.shields.io/github/v/tag/hbourgeot/templdais)
![GitHub Repo stars](https://img.shields.io/github/stars/hbourgeot/templdais)
![GitHub commit activity](https://img.shields.io/github/commit-activity/w/hbourgeot/templdais)



## Features

- Components ready to use.
- Includes Types as `Links`, `Attributes` and more.
- Color and size customizable.


## Installation

Install the UI with `go get`

```bash
  go get github.com/hbourgeot/templdais
```
    
## Demo

```gohtml
// /path/to/your/file.templ

package main

templ Page() {
    <div class="w-full h-full flex justify-center items-center">
    @Button("primary", "button", "sm", "circle"){
        This works!
    }
}
```

Here, the arguments are:
1. The value `primary` is for the **brand color**.
2. The value `button` is for the **HTML Tag**.
3. `sm` is the size.
4. `circle` is the shape of the button, can be `square` or `circle`.



## Roadmap

[DaisyUI](https://daisyui.com) Components ported to TemplDais:

### Added 2024/04/14

- [x]  Button
- [x]  Dropdown
- [x]  Modal
- [x]  Kbd
- [x]  Breadcrumbs
- [x]  Checkbox
- [x]  Radio
- [x]  Accordion
- [x]  Badge
- [x]  Card (responsive card for now)

### Added 2024/04/15

- [x]  Table
- [x]  Navbar
- [x]  Pagination
- [x]  Toast
- [x]  File Input
- [x]  Range
- [x]  Select

Comming soon:

- [ ]  Text input
- [ ]  Textarea
- [ ]  Code
- [ ]  Alert
- [ ]  Timeline
- [ ]  Divider
- [ ]  Indicator

**PD**: A Documentation web of *TemplDais* is incoming soon (in development)
**PD 2**: I've plained to use or create a Icons Kit for *TemplDais*, for now, you must provide the url of the icons (svg)

And next (not my priority):

- [ ]  Carousel
- [ ]  Chat bubble
- [ ]  Bottom navigation
- [ ]  Progress
- [ ]  Radial Progress
- [ ]  Tooltip
- [ ]  Toggle
- [ ]  Drawer
- [ ]  Collapse
- [ ]  Stat
- [ ]  Avatar
- [ ]  Steps
- [ ]  Menu
- [ ]  Advanced pagination

Not implementable (yet):

- Countdown (sorry, templ doesn´t allows dynamic style attribute for set `--value` css variable)
- Tab (it needs a tabs and tabpanel, maybe in a long future I'll work on this)

## Acknowledgements

 - [Templ](https://github.com/a-h/templ)
 - [Daisy UI](https://daisyui.com)
 - [Tailwind CSS](https://tailwindcss.com)
 - [Golang](https://go.dev)


## Authors

- [@hbourgeot](https://www.github.com/hbourgeot)
