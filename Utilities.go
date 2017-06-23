package commandframework_discordgo

import "strings"

func FormatString(what string, toReplace map[string]string) string {
    ret := what
    for key, value := range toReplace {
        ret = strings.Replace(ret, "{"+strings.ToUpper(key)+"}", value, -1)
    }
    return ret
}
