splatstats-uploader-go
=====================

[![Go Report Card](https://goreportcard.com/badge/github.com/cass-dlcm/splatstats-uploader-go)](https://goreportcard.com/report/github.com/cass-dlcm/splatstats-uploader-go)
[![Generate release-artifacts](https://github.com/cass-dlcm/splatstats-uploader-go/actions/workflows/release.yml/badge.svg)](https://github.com/cass-dlcm/splatstats-uploader-go/actions/workflows/release.yml)
[![DeepSource](https://deepsource.io/gh/cass-dlcm/splatstats-uploader-go.svg/?label=active+issues&show_trend=true)](https://deepsource.io/gh/cass-dlcm/splatstats-uploader-go/?ref=repository-badge)

splatstats-uploader-go is a program that uploads battle data from the SplatNet 2 app ("Nintendo Switch Online") to [SplatStats](https://splatstats.cass-dlcm.dev/), a site for recording, visualizing, and aggregating statistics from *Splatoon 2*.

## Usage

```
$ ./splatstatsuploader [-M N] [-f] [-s] [--salmon]
```

The `-M` flag runs the program in monitoring mode, uploading new battle results as you play matches. The script checks for new battles every `N` seconds.

The `-f` flag sends battle data to SplatStats from a folder of SplatNet 2 JSON files.

The `-s` flag saves battle data downloaded from SplatNet 2 to a folder.

The `--salmon` flag uploads Salmon Run shifts.

### Example usage

Running `./splatstatsuploader -M 900` from the command line launches the script in monitoring mode, checking for and uploading battles every 15 minutes.

## Features

- [ ] Full automation of SplatNet cookie generation/acquisition via user log-in
- [ ] Complete battle stats
  - [x] Game mode, stage, weapon
  - [x] Result, final count/percent, turf inked
  - [x] Kills, deaths, assists, specials
  - [x] Rank, level & star emblems (&#9733;), X Rank & Power, weapon freshness
  - [x] Battle start & end times
  - [ ] Ranked power level & League Power
  - [ ] Splatfest support: Title, EXP, Power, Clout, Synergy Bonus, team nickname, win streak
  - [x] Species (Inkling or Octoling)
- [ ] Complete Salmon Run stats

## Setup instructions

1. Download the program from [the releases page](https://github.com/cass-dlcm/splatstatsuploader-go/releases).

2. Run the program as shown in the [usage](#usage) section above.

3. Running the program for the first time will prompt you to enter your SplatStats username and password. If you're using the app in a language other than English, you may enter your [language code](https://github.com/frozenpandaman/splatnet2statink/wiki/languages) (locale) as well.

**NOTE: Read the "Cookie generation" section below before proceeding. [→](#cookie-generation)**

4. <!--You will then be asked to navigate to a specific URL on Nintendo.com, log in, and follow simple instructions to obtain your `session_token`; this will be used to generate an `iksm_session` cookie. If you are opting against automatic cookie generation,--> You will then be asked to navigate to a specific URL on Nintendo.com, log in, and follow simple instructions to obtain your `session_token`. However, this functionality does not work at this time, so instead enter "skip" for this step, at which point you will be asked to manually input your `iksm_session` cookie <!--instead--> (see the [mitmproxy instructions](https://github.com/frozenpandaman/splatnet2statink/wiki/mitmproxy-instructions)).

    This cookie (used to access your SplatNet battle results) along with your SplatStats API key and language will automatically be saved into `config.txt` for you. You're now ready to upload battles!

### Accessing SplatNet 2 from your browser

If you wish to access SplatNet 2 from your computer rather than via the phone app, navigate to [https://app.splatoon2.nintendo.net/home](https://app.splatoon2.nintendo.net/home) (it should show a forbidden error). Use a cookie editor – such as [EditThisCookie](https://chrome.google.com/webstore/detail/editthiscookie/fngmhnnpilhplaeedifhccceomclgfbg?hl=en) for Chrome – to change `iksm_session` to the value you obtained previously (automatically or via [mitmproxy](https://github.com/frozenpandaman/splatnet2statink/wiki/mitmproxy-instructions), stored as  `cookie` in `config.txt`), and refresh the page. If you only want to access SplatNet and don't have a stat.ink API key, simply enter "skip" for this step during setup.

*Splatoon 2* stage rotation information (including Salmon Run) and current SplatNet gear are viewable at [splatoon2.ink](https://splatoon2.ink/).

---

## Cookie generation

For splatstats-uploader to work, a cookie known as `iksm_session` is required to access SplatNet. This cookie may be obtained automatically, using the program, in the future once that functionality is added, or manually via the app. <!--Please read the following sections carefully to decide whether or not you want to use automatic cookie generation.-->

<!-- ### Automatic

Automatic cookie generation involves making a *secure request to two non-Nintendo servers with minimal, non-identifying information*. We aim to be 100% transparent about this and provide in-depth information on security and privacy below. Users who feel uncomfortable with this may opt to manually acquire their cookie instead.

The v1.1.0 update to the Nintendo Switch Online app, released in September 2017, introduced the requirement of a [message authentication code](https://en.wikipedia.org/wiki/Message_authentication_code) (known as `f`), thereby complicating the ability to generate cookies within the script. After figuring out the [key](https://en.wikipedia.org/wiki/Key_\(cryptography\)) previously used to generate `f` tokens, the calculation method was changed in September 2018's v1.4.1 update, heavily obfuscating the new process. As a workaround, an Android server was set up to emulate the app, specifically to generate `f` tokens.

Generation now requires a [hash value](https://en.wikipedia.org/wiki/Hash_function) to further verify the authenticity of the request. The algorithm to calculate this, originally done within the app, is sensitive; to prevent sharing it publicly (i.e. distributing it in the script's source code), @frozenpandaman created a small [API](https://en.wikipedia.org/wiki/Application_programming_interface) which generates a hash value given a valid input. This can be passed to the Android server to generate the corresponding `f` token, which is then used to retrieve an `iksm_session` cookie.

**Privacy statement:** No identifying information is ever sent to the API server. Usernames and passwords are far removed from where the API comes into play and are never readable by anyone but you. Returned hash values are never logged or stored and do not contain meaningful information. It is not possible to use either sent or stored data to identify which account/user performed a request, to view any identifying information about a user, or to gain access to an account.

See the **[API documentation wiki page](https://github.com/frozenpandaman/splatnet2statink/wiki/api-docs)** for more information. -->

### Manual

Users who decide against automatic cookie generation via their computer may instead generate/retrieve `iksm_session` cookies from the SplatNet app.

In this case, users must obtain their cookie from their phone by intercepting their device's web traffic and inputting it into splatstats-uploader when prompted (or manually adding it to `config.txt`). Follow the [mitmproxy instructions](https://github.com/frozenpandaman/splatnet2statink/wiki/mitmproxy-instructions) to obtain and configure your cookie manually. To opt to manually acquire your cookie, enter "skip" when prompted to enter the "Select this account" URL.

## License

[GPLv3](https://www.gnu.org/licenses/gpl-3.0.html)

adapted from [splatnet2statink](https://github.com/frozenpandaman/splatnet2statink)
