package main

import (
	"flag"
	"fmt"
	"go-led-strip/controller"
	"go-led-strip/triggers"
	"go-led-strip/effects"
	"go-led-strip/ledstrips"
	"go-led-strip/services"
	"go-led-strip/services/config"
	"golang.org/x/image/colornames"
)

func main() {
	stripConfigPath := "configs/ledStrip.yaml"
	effectsConfigPath := flag.String("effects", "effects.yaml", "")
	flag.Parse()
	Initialize(stripConfigPath, "configs/"+*effectsConfigPath)
}

func Initialize(ledStripConfigPath, effectsConfigPath string) {
	fmt.Println("initializing...")
	stripConfig := config.GetLedStripConfig(ledStripConfigPath)
	strip := initializeLedStrip(stripConfig)
	strip.Fill(colornames.Black)
	effectsConfig := config.GetEffectsConfig(effectsConfigPath)
	ledEffects := initializeEffects(effectsConfig)
	service := services.NewLedStripService(services.NewStripFrameService(stripConfig.LedCount, ledEffects), strip)
	service.Test()
	service.InitTicker()
	service.Start()
	controller.NewLedStripController(service)
	triggers.NewSwitchTrigger(service)
}

func initializeLedStrip(config *config.LedStripConfig) ledstrips.LedStrip {
	switch config.StripType {
	case ledstrips.APA102:
		fmt.Println("Found APA102 led strip with", config.LedCount, "leds")
		return ledstrips.NewApa102LedStripFromConfig(config)
	}
	return nil
}

func initializeEffects(config *config.EffectsConfig) []effects.LedEffect {
	results := make([]effects.LedEffect, len(config.Effects))
	for i := range results {
		switch config.Effects[i].Type {
		case effects.LENGTH_BREATHING:
			results[i] = effects.NewLengthBreathingFromConfig(config.Effects[i])
		case effects.MOVING_SEGMENTS:
			results[i] = effects.NewMovingSegmentsFromConfig(config.Effects[i])
		}
		results[i].Start()
	}
	fmt.Println("Found", len(results), "effects defined")
	return results
}
