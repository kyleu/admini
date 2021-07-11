package download

import (
	"fmt"

	"github.com/kyleu/admini/app/util"
)

func DownloadLinks(version string) Links {
	if availableLinks == nil {
		ret := Links{}
		add := func(url string, mode string, os string, arch string) {
			ret = append(ret, &Link{URL: url, Mode: mode, OS: os, Arch: arch})
		}
		addDefault := func(mode string, os string, arch string) {
			var u string
			switch mode {
			case modeServer:
				u = fmt.Sprintf("%s_%s_%s_%s.zip", util.AppKey, version, os, arch)
			case modeDesktop:
				u = fmt.Sprintf("%s_desktop_%s_%s_%s.zip", util.AppKey, version, os, arch)
			}
			add(u, mode, os, arch)
		}
		addArms := func(mode string, os string) {
			for _, arm := range []string{archARMV5, archARMV6, archARMV7} {
				addDefault(mode, os, arm)
			}
		}
		addWeird := func(mode string, os string) {
			for _, weird := range []string{archMips64Hard, archMips64LEHard, archMips64LESoft, archMips64Soft, archMipsHard, archMipsLEHard, archMipsLESoft, archMipsSoft} {
				addDefault(mode, os, weird)
			}
		}
		addDefault(modeServer, osAIX, archPPC64)

		addDefault(modeServer, osDragonfly, archAMD64)

		addDefault(modeServer, osFreeBSD, archARM64)
		addArms(modeServer, osFreeBSD)
		addDefault(modeServer, osFreeBSD, archI386)
		addDefault(modeServer, osFreeBSD, archAMD64)

		addDefault(modeServer, osIllumos, archAMD64)

		addDefault(modeServer, osJS, archWASM)

		addDefault(modeServer, osLinux, archARM64)
		addArms(modeServer, osLinux)
		addDefault(modeServer, osLinux, archI386)
		addDefault(modeServer, osLinux, archAMD64)
		addDefault(modeServer, osLinux, archPPC64)
		addDefault(modeServer, osLinux, archRISCV64)
		addDefault(modeServer, osLinux, archS390X)
		addDefault(modeServer, osLinux, archAMD64)
		addDefault(modeDesktop, osLinux, archAMD64)
		addWeird(modeServer, osLinux)

		addDefault(modeServer, osMac, archARM64)
		addDefault(modeServer, osMac, archAMD64)
		addDefault(modeDesktop, osMac, archAMD64)

		addDefault(modeServer, osMobile, archAndroid)
		addDefault(modeServer, osMobile, archIOS)

		addDefault(modeServer, osNetBSD, archARMV7)
		addDefault(modeServer, osNetBSD, archI386)
		addDefault(modeServer, osNetBSD, archAMD64)

		addDefault(modeServer, osOpenBSD, archARM64)
		addArms(modeServer, osOpenBSD)
		addDefault(modeServer, osOpenBSD, archI386)
		addDefault(modeServer, osOpenBSD, archAMD64)

		addArms(modeServer, osPlan9)
		addDefault(modeServer, osPlan9, archI386)
		addDefault(modeServer, osPlan9, archAMD64)

		addDefault(modeServer, osSolaris, archAMD64)

		addArms(modeServer, osWindows)
		addDefault(modeServer, osWindows, archI386)
		addDefault(modeServer, osWindows, archAMD64)
		addDefault(modeDesktop, osWindows, archAMD64)

		availableLinks = ret
	}
	return availableLinks
}
