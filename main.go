package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/parnurzeal/gorequest"
	"github.com/visionmedia/go-spin"
)

type Stores struct {
	UpdatedTime string
	Stores      []Store
}
type Store struct {
	StoreNumber, StoreName string
	StoreEnabled           bool
}

const Gold16 = "MG492J/A"
const Gold64 = "MG4J2J/A"
const Gold128 = "MG4E2J/A"
const Silver16 = "MG482J/A"
const Silver64 = "MG4H2J/A"
const Silver128 = "MG4C2J/A"
const Gray16 = "MG472J/A"
const Gray64 = "MG4F2J/A"
const Gray128 = "MG4A2J/A"
const GoldPlus16 = "MGAA2J/A"
const GoldPlus64 = "MGAK2J/A"
const GoldPlus128 = "MGAF2J/A"
const SilverPlus16 = "MGA92J/A"
const SilverPlus64 = "MGAJ2J/A"
const SilverPlus128 = "MGAE2J/A"
const GrayPlus16 = "MGA82J/A"
const GrayPlus64 = "MGAH2J/A"
const GrayPlus128 = "MGAC2J/A"

func main() {
	spinner := spin.New()
	const layout = "15:04:05"
	request := gorequest.New()
	for {
		// get availability json file
		resp, body, errs := request.Get("https://reserve.cdn-apple.com/JP/ja_JP/reserve/iPhone/availability.json").End()
		//resp, body, errs := request.Get("http://localhost:8080/availability.json").End()
		if errs != nil {
			fmt.Println("ERROR: ", errs)
			continue
		}
		if resp.StatusCode != 200 {
			fmt.Println("WARNING: not 200", resp)
			continue
		}
		var data map[string]interface{}
		if err := json.Unmarshal([]byte(body), &data); err != nil {
			fmt.Println("ERROR: Cannot decode " + body)
			continue
		}

		fmt.Printf("\033[2J\033[0;0H\r %s site not yet avaialble %s", time.Now().Format(layout), spinner.Next())
		if len(data) != 0 {
			err := ioutil.WriteFile("/tmp/json."+time.Now().Format("021504")+".out", []byte(body), 0644)
			if err != nil {
				fmt.Println("ERROR: cannot write file", err)
			}
			// getting stores json file
			//resp, body, errs := request.Get("https://reserve.cdn-apple.com/JP/ja_JP/reserve/iPhone/stores.json").End()
			resp, body, errs := request.Get("http://localhost:8080/stores.json").End()
			if errs != nil {
				fmt.Println("ERROR: Cannot get stores data", errs)
			}
			if resp.StatusCode != 200 {
				fmt.Println("WARNING: stores.json not 200", resp)
			}
			stores := &Stores{}
			if err := json.Unmarshal([]byte(body), &stores); err != nil {
				fmt.Println("ERROR: Cannot decode " + body)
			}
			fmt.Printf("\033[2J\033[0;0H")
			fmt.Printf("%-15s %8s %8s %8s %8s %8s %8s %8s %8s %8s %8s %8s %8s %8s %8s %8s %8s %8s %8s\n", "Store",
				"Silv16", "Silv64", "Silv128",
				"Gold16", "Gold64", "Gold128",
				"Gray16", "Gray64", "Gray128",
				"Silv-P16", "Silv-P64", "Silv-P128",
				"Gold-P16", "Gold-P64", "Gold-P128",
				"Gray-P16", "Gray-P64", "Gray-P128",
			)
			for _, s := range stores.Stores {
				fmt.Printf("%s", s.StoreName)
				fmt.Printf("\r\033[15C%8t %8t %8t %8t %8t %8t %8t %8t %8t %8t %8t %8t %8t %8t %8t %8t %8t %8t\n",
					data[s.StoreNumber].(map[string]interface{})[Silver16],
					data[s.StoreNumber].(map[string]interface{})[Silver64],
					data[s.StoreNumber].(map[string]interface{})[Silver128],
					data[s.StoreNumber].(map[string]interface{})[Gold16],
					data[s.StoreNumber].(map[string]interface{})[Gold64],
					data[s.StoreNumber].(map[string]interface{})[Gold128],
					data[s.StoreNumber].(map[string]interface{})[Gray16],
					data[s.StoreNumber].(map[string]interface{})[Gray64],
					data[s.StoreNumber].(map[string]interface{})[Gray128],
					data[s.StoreNumber].(map[string]interface{})[SilverPlus16],
					data[s.StoreNumber].(map[string]interface{})[SilverPlus64],
					data[s.StoreNumber].(map[string]interface{})[SilverPlus128],
					data[s.StoreNumber].(map[string]interface{})[GoldPlus16],
					data[s.StoreNumber].(map[string]interface{})[GoldPlus64],
					data[s.StoreNumber].(map[string]interface{})[GoldPlus128],
					data[s.StoreNumber].(map[string]interface{})[GrayPlus16],
					data[s.StoreNumber].(map[string]interface{})[GrayPlus64],
					data[s.StoreNumber].(map[string]interface{})[GrayPlus128],
				)
			}
			fmt.Print("\x07")
		}
		time.Sleep(1 * time.Second)
	}
}
