# win_reload_proxy_settings

Ever wanted to change settings in:
Computer\HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Internet Settings

but those settings never got actually reloaded into your open browsers, probably because you are running an old and vulnerable version of windows - a disaster waiting to unfold but totally unrelated to this repository? (hey Mr.Robot, I'm here!)

wininet.dll's InternetSetOptionA for the rescue... But in golang :(

$INTERNET_OPTION_SETTINGS_CHANGED   = 39
$INTERNET_OPTION_REFRESH            = 37
