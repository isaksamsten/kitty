# kitty - the fast, feature-rich, cross-platform, GPU based terminal

This fork makes Kitty more usable on macOS without title bar enabled.

**NOTE** THESE CHANGES WILL NOTE BE MERGED INTO MAIN.

## Changes

-   Enable `drag-window-move` when dragging the tab-bar
-   Center single tab
-   Enable window buttons if `hide_window_decorations = titlebar-only`

## Build

First, if you are using `conda` run `conda deactivate` to ensure
that you are using the system python to build.

```shell
brew bundle
LDFLAGS=-L/opt/homebrew/lib python3 setup.py  --extra-include-dirs /opt/homebrew/Cellar/librsync/{VERSION}/include
pip3 install -r docs/requirements.txt
make docs
LDFLAGS=-L/opt/homebrew/lib python3 setup.py kitty.app  --extra-include-dirs /opt/homebrew/Cellar/librsync/{VERSION}/include
```

Move `kitty.app` to `/Applications`

## Configuration

```
hide_window_decorations titlebar-only
tab_bar_edge top
tab_bar_margin_height 6 6
tab_bar_min_tabs 1
```

Also, for best result, ensure that the active, inactive, tab_bar and the margin
has the same background color:

```
active_tab_background #181725
active_tab_foreground #e0def4
inactive_tab_background #181725
inactive_tab_foreground #817c9c
tab_bar_background #181725
tab_bar_margin_color #181725
```

## Screenshots

![Single tab](https://i.imgur.com/CbILExU.png)

![Two tabs](https://i.imgur.com/F6GS3ij.png)

[Video showcase](https://i.imgur.com/TZNW9uj.mp4)
