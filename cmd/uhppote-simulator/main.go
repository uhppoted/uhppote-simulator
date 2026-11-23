package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"codeberg.org/uhppoted/uhppoted-core/uhppote"
	"codeberg.org/uhppoted/uhppoted-simulator/commands"
	"codeberg.org/uhppoted/uhppoted-simulator/log"
	"codeberg.org/uhppoted/uhppoted-simulator/simulator"
	"codeberg.org/uhppoted/uhppoted-simulator/simulator/UT0311L04"
)

var options = struct {
	dir   string
	bind  string
	rest  string
	cards uint
	debug bool
}{
	dir:   "./devices",
	bind:  ":60000",
	rest:  ":8000",
	cards: 64,
	debug: false,
}

func main() {
	flag.StringVar(&options.dir, "devices", "devices", "Specifies the simulation device directory")
	flag.StringVar(&options.bind, "bind", ":60000", "Specifies the bind address for the simulator")
	flag.StringVar(&options.rest, "rest", ":8000", "Specifies the bind address for the REST interface")
	flag.UintVar(&options.cards, "cards", 64, "Size of controller cards lists (defaults to 64)")
	flag.BoolVar(&options.debug, "debug", false, "Displays simulator activity")
	flag.Parse()

	if len(flag.Args()) > 0 && flag.Arg(0) == "help" {
		help()
		return
	}

	if len(flag.Args()) > 0 && flag.Arg(0) == "version" {
		version()
		return
	}

	ctx := simulator.Context{
		Directory:   options.dir,
		BindAddress: options.bind,
		RestAddress: options.rest,
		DeviceList:  simulator.NewDeviceList(load(options.dir, int(options.cards))),
	}

	commands.Simulate(&ctx, options.debug)
}

func load(dir string, cardsListSize int) []simulator.Simulator {
	log.Infof("loading devices from '%s'", dir)

	devices := map[uint32]simulator.Simulator{}

	list := []struct {
		glob       string
		compressed bool
	}{
		{"*.json", false},
		{"*.json.gz", true},
	}

	for _, g := range list {
		files, err := filepath.Glob(filepath.Join(dir, g.glob))
		if err == nil {
			for _, f := range files {
				s, err := UT0311L04.Load(f, g.compressed, cardsListSize)
				if err != nil {
					fmt.Printf("   ... error loading device from file '%s': %v\n", f, err)
				} else {
					deviceID := s.DeviceID()
					if devices[deviceID] == nil {
						devices[deviceID] = s
						log.Infof("loaded  device  from '%s'\n", f)
					} else {
						log.Warnf("duplicate serial number %v in device file '%s' - using device loaded from '%s'\n", deviceID, f, devices[deviceID].FilePath())
					}
				}
			}
		}
	}

	l := make([]simulator.Simulator, 0)

	for _, s := range devices {
		l = append(l, s)
	}

	return l
}

func version() {
	fmt.Printf("%v\n", uhppote.VERSION)
}

func help() {
	fmt.Println()
	fmt.Println("  Usage: uhppote-simulator [options] <command>")
	fmt.Println()
	fmt.Println("  By default, the application will run in 'simulation' mode if a command is not specified.")
	fmt.Println()
	fmt.Println("  Commands:")
	fmt.Println()
	fmt.Println("    version          Displays the simulator version")
	fmt.Println("    help             Displays this message")
	fmt.Println("                     For help on a specific command use 'uhppote-cli help <command>'")
	fmt.Println()
	fmt.Println("  Options:")
	fmt.Println()
	fmt.Println("    --devices <directory>  Sets the directory to which to load and save simulator device files")
	fmt.Println("    --debug                Displays internal information for diagnosing errors")
	fmt.Println()
}
