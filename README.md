# crock - Rock CLI Clock 🪨

<img src="https://img.shields.io/github/go-mod/go-version/mirumirumi/crock"> <img src="https://img.shields.io/github/v/release/mirumirumi/crock"> <img src="https://img.shields.io/github/workflow/status/mirumirumi/crock/release">

![crock](/assets/main-visual.png)

The crock is an easy-to-read clock that runs on the CLI.  
Do you want to check the current time from a long distance? 👀  
Then crock is perfect for you!

![crocking](/assets/crocking.gif)

## Features

- A big big analog & digital clock 🕰️
- Specify the color of output 🎨
- Oh my gosh, you can even check the date! That's too awesome!! 📅

## Usage

Mac:

```bash
brew tap mirumirumi/crock
brew install crock
```

Linux:

```bash
wget https://raw.githubusercontent.com/mirumirumi/crock/main/crock.sh -P /tmp/
sh /tmp/crock.sh
```

Windows:

```powershell
Invoke-WebRequest https://github.com/mirumirumi/crock/releases/latest/download/crock.exe -OutFile $env:temp\crock.exe  # it would be `C:\Users\[user]\AppData\Local\Temp`
# Move it anywhere you want!
```

then,

```bash
crock
```

start crocking!

## Contribute

PRs are so welcome!  
(Since the author is new to Golang, so there will be lots of room for improvement in the code.)

## Credit and Special Thanks

- ASCII Art: [TEXT-IMAGE.com](https://www.text-image.com/about.html)
- Number font: [りぃ](http://aoirii.babyblue.jp/font/riipopkk/index.html)

## LICENSE

MIT
