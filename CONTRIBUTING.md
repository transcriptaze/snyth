# Contributing

This is an umbrella repository and all the actual code lives in submodules:

- [snyth-js](https://github.com/transcriptaze/snyth-js) for the HTML, CSS and Javascript source that go to 
  make up the web app. The actual web-app is a web-packed version of this repo.
- [snyth-pd](https://github.com/transcriptaze/snyth-pd) is a basic implementation in _PureData_. At the moment
  there isn't much beyond a basic working demo.
- [snyth-supercollider](https://github.com/transcriptaze/snyth-supercollider) is a _Supercollider_ UGen for the 
  basic _sn_ function and a couple of demo scripts.


### MIDI

While it would be great to have a selection of really good demo MIDI files it's probably legally murky unless 
the arrangements are clearly copyright free. Having said which, if you do have a MIDI file that you would 
like to share:

1. Please create a pull request against the MIDI branch of this repository.
2. The pull request should:
   - Update the MIDI files list with a link to where you are hosting the MIDI file
   - Include the MIDI file
   - Include a _snyth.json_ file
3. After merging the pull request the MIDI files themselves will be stashed somewhere safe (but in good company) 
   until the legal situation is clearer.

Arrangements of classical music are mostly copyright free but be aware that some are copyrighted by the arranger.


