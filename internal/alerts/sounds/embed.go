package sounds

import "embed"

//go:embed chime.wav bell.wav ping.wav pop.wav ding.wav
var SoundFiles embed.FS
