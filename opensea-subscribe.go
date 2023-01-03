package main

import (
	"fmt"
	"github.com/foundVanting/opensea-stream-go/entity"
	"github.com/foundVanting/opensea-stream-go/opensea"
	"github.com/foundVanting/opensea-stream-go/types"
	"github.com/mitchellh/mapstructure"
	"github.com/nshafer/phx"
)

func main() {
	client := opensea.NewStreamClient(types.TESTNET, "", phx.LogInfo, func(err error) {
		fmt.Println("NewStreamClient err:", err)
	})
	client.Connect()

	go client.OnItemListed("untitled-collection-14425637", func(response any) {
		var itemListedEvent entity.ItemListedEvent
		err := mapstructure.Decode(response, &itemListedEvent)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
		}
		fmt.Printf("%+v\n", itemListedEvent)
	})
	go client.OnItemSold("untitled-collection-14425637", func(response any) {
		var itemSoldEvent entity.ItemSoldEvent
		err := mapstructure.Decode(response, &itemSoldEvent)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
		}
		fmt.Printf("%+v\n", itemSoldEvent)
	})
	go client.OnItemCancelled("untitled-collection-14425637", func(response any) {
		var itemListedEvent entity.ItemListedEvent
		err := mapstructure.Decode(response, &itemListedEvent)
		if err != nil {
			fmt.Println("mapstructure.Decode err:", err)
		}
		fmt.Printf("%+v\n", itemListedEvent)
	})
	select {}
}
