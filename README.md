
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

- [x]  Button (2024/04/14)
- [x]  Dropdown (2024/04/14)
- [x]  Modal (2024/04/14)
- [x]  Kbd (2024/04/14)
- [x]  Breadcrumbs (2024/04/14)
- [x]  Checkbox (2024/04/14)
- [x]  Radio (2024/04/14)
- [x]  Accordion (2024/04/14)
- [x]  Badge (2024/04/14)

Comming soon:

- [ ]  Card
- [ ]  Table
- [ ]  Navbar
- [ ]  Pagination
- [ ]  Tab
- [ ]  Toast
- [ ]  File Input
- [ ]  Range
- [ ]  Select
- [ ]  Text input
- [ ]  Textarea
- [ ]  Code

And next (not my priority):

- [ ]  Carousel
- [ ]  Chat bubble
- [ ]  Timeline
- [ ]  Bottom navigation
- [ ]  Alert
- [ ]  Progress
- [ ]  Radial Progress
- [ ]  Tooltip
- [ ]  Toggle
- [ ]  Divider
- [ ]  Drawer
- [ ]  Indicator
- [ ]  Collapse
- [ ]  Stat
- [ ]  Avatar
- [ ]  Steps
- [ ]  Menu

Not implementable:

- Countdown (sorry, templ doesn´t allows dynamic style attribute for set `--value` css variable)

## Acknowledgements

 - [Templ](https://github.com/a-h/templ)
 - [Daisy UI](https://daisyui.com)
 - [Tailwind CSS](https://tailwindcss.com)
 - [Golang](https://go.dev)


## Authors

- [@hbourgeot](https://www.github.com/hbourgeot)
