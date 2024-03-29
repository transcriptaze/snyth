# NOTES

The current implementation is very much a proof of concept/work in progress and there is an absolute **ton**
of optimization that needs to happen but there is just one of me. Also quite a lot of it was coded around 3AM
so there is weird stuff that only makes sense at that time of the day.

## Browsers

_snyth_ (the _snyth_ ?) has only really been extensively tested on Chrome on an **ancient** and **decrepit**
Macbook Pro. For other browsers and operating systems, YMMV:

| Browser | MacOS | Windows | Linux | Notes                                                                               |
|---------|-------|---------|-------|-------------------------------------------------------------------------------------|
| Chrome  | ✓     | ✓       |       |                                                                                     |
| Opera   | ✓     |         |       |                                                                                     |
| Firefox | ✓     | ✓       |       | Requires Firefox [113.0+](https://www.mozilla.org/en-US/firefox/113.0/releasenotes) |
| Safari  | ✓     | -       |       | Doesn't support [customized web components [2]](#2) and [[3]](#3)                   |
| Edge    | -     | ✓       |       |                                                                                     |

It isn't remotely responsive and almost certainly isn't usable  on a mobile phone or tablet.

_Notes:_
1. Firefox 113.0 fixes a long standing issue [[1]](#2) with _AudioWorklets_ and ES6 modules.
2. Safari works but is rough around the edges (for one thing the canvas resolution seems a bit off).


## MIDI files

At this stage, the MIDI file interpretation is very basic - it does the job (mostly) but otherwise is somewhat
less than ummm, stellar. A lot of MIDI files are not going to sound very good straight out of the box - for one
thing they're often arranged for multiple instruments and for another they often have an extravagant amount of
notes.

The most success I've had so far has been with MIDI files of good fingerstyle guitar arrangements, e.g.:

- [Luca Stricagnoli's](https://www.youtube.com/@LucaStricagnoli) arrangement of [Rondo alla Turca](https://www.youtube.com/watch?v=I3ut6H6-gvE)
- [Eiro Nareth's](https://www.youtube.com/watch?v=SaZiUBfXKEs) arrangement of [Time](https://youtu.be/DRVajO5q-xo) (from the Inception movie).

Other sources:

- [Online Sequencer](https://onlinesequencer.net) MIDI arrangments vary enormously in quality but it's possible
  to find ones that work reasonably well with the _snyth_ and if you dig deep enough and long enough some are 
  even really good. 
- In my experience, [BitMIDI](https://bitmidi.com) is more miss than hit but occasionally you'll stumble across
  something useable.

### CORS

It would be nice to be able to load MIDI files directly from a URL but _snyth_ makes a lot of use of things like
_AudioWorklets_ and _SharedArrayBuffers_ and global current times, all of which are restricted unless the CORS
header is set to _same-origin_. And since it is hosted on Github pages it just isn't possible, so until this 
thing gets its own server you'll have to download the files and load them locally.

## References

<a id="1">1.</a> [Bug #1636121](https://bugzilla.mozilla.org/show_bug.cgi?id=1636121). Fixed in [113.0+](https://www.mozilla.org/en-US/firefox/113.0/releasenotes)  
<a id="2">2.</a> [What web component features are not supported by Safari desktop and Safari iOS?](https://stackoverflow.com/questions/72090155/what-web-component-features-are-not-supported-by-safari-desktop-and-safari-ios)  
<a id="3">3.</a> [Re: Custom form elements ...](https://lists.w3.org/Archives/Public/public-webapps/2013OctDec/1051.html)  
