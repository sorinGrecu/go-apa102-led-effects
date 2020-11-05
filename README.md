# go-apa102-led-effects
Controlling APA102 and some led animations through Go
First project done in Go, learning playground. Still a lot to learn.

How to: 
Install Go on a Raspberry Pi, connect an APA102 ledstrip, edit the configs/effects.yaml and configs/ledStrip.yaml and have fun.

What it does:
  - sets up the driver for APA102
  - plugs in some effects
  - reads configuration for effects and led strip from yaml files
  - exposes endpoints to control the led strip (on/off, all I needed to be able to control it with Alexa)
  
  
Future plans:
  - going to port more animations from my Java projects into this one
  - expose endpoints to control the animation properties (speed, colours, spacing, etc)
  - a lot more
  
