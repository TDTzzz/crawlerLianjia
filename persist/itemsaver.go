package persist

import (
	"context"
	"github.com/TDTzzz/crawlerLianjia/model"
	"github.com/olivere/elastic/v7"
	"log"
)

func ItemSaver(index string) (chan model.HouseDetail, error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	out := make(chan model.HouseDetail)

	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("got item"+"#%d: %v", itemCount, item)
			itemCount++
			err := Save(client, index, item)
			if err != nil {
				log.Println("Item Saver: err "+"saving item %v : %v", item, err)
			}
		}
	}()
	return out, nil
}

func Save(client *elastic.Client, index string, item model.HouseDetail) error {
	indexService := client.Index().Index(index).BodyJson(item)
	_, err := indexService.Do(context.Background())
	return err
}
