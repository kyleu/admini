#!/usr/bin/osascript
tell application "iTerm2"
    tell current session of current tab of current window
        write text "cd ~/go/src/github.com/kyleu/admini"
        write text "clear"
        write text "bin/dev.sh"
        split vertically with default profile
    end tell
    tell second session of current tab of current window
        write text "cd ~/go/src/github.com/kyleu/admini/client"
        write text "clear"
        write text "../tools/bin/build-client-watch.sh"
        split horizontally with default profile
    end tell
    tell third session of current tab of current window
        write text "cd ~/go/src/github.com/kyleu/admini"
        write text "clear"
        write text "tools/bin/build-svg.sh"
    end tell
end tell
