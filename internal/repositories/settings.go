package repositories

import (
	"github.com/everestafrica/everest-api/internal/database"
	"github.com/everestafrica/everest-api/internal/models"
	"gorm.io/gorm"
)

type ISettingsRepository interface {
	CreateCustomCategory(category *models.CustomCategory) error
	DeleteCustomCategory(id string) error
	FindAllCustomCategories(userId string) (*[]models.CustomCategory, error)
	CreatePriceAlert(alert *models.PriceAlert) error
	DeletePriceAlert(id string) error
	FindAllPriceAlerts(userId string) (*[]models.PriceAlert, error)
}

type settingsRepo struct {
	db *gorm.DB
}

// NewSettingsRepo will instantiate Settings Repository
func NewSettingsRepo() ISettingsRepository {
	return &settingsRepo{
		db: database.DB(),
	}
}

func (r *settingsRepo) CreateCustomCategory(category *models.CustomCategory) error {
	return r.db.Create(&category).Error
}

func (r *settingsRepo) FindAllCustomCategories(userId string) (*[]models.CustomCategory, error) {
	var categories []models.CustomCategory
	if err := r.db.Where("user_id = ? ", userId).Find(&categories).Error; err != nil {
		return nil, err
	}
	return &categories, nil
}

func (r *settingsRepo) DeleteCustomCategory(id string) error {
	var category models.CustomCategory
	if err := r.db.Where("id = ? ", id).Delete(&category).Error; err != nil {
		return err
	}
	return nil
}

func (r *settingsRepo) CreatePriceAlert(alert *models.PriceAlert) error {
	return r.db.Create(&alert).Error
}
func (r *settingsRepo) DeletePriceAlert(id string) error {
	var alert models.PriceAlert
	if err := r.db.Where("id = ? ", id).Delete(&alert).Error; err != nil {
		return err
	}
	return nil
}

func (r *settingsRepo) FindAllPriceAlerts(userId string) (*[]models.PriceAlert, error) {
	var alerts []models.PriceAlert
	if err := r.db.Where("user_id = ? ", userId).Find(&alerts).Error; err != nil {
		return nil, err
	}
	return &alerts, nil
}

func (r *settingsRepo) UpdateAlert(alert models.PriceAlert) error {
	return r.db.Save(alert).Error
}

/*

func main() {
    // Set up cron job to run the monitorStock function every minute
    c := cron.New()
    c.AddFunc("@every 1m", monitorStock)
    c.Start()

    // Run the program indefinitely
    select {}
}

func monitorStock() {
    stockSymbol := "AAPL"
    alertPrice := 110.00

    // Fetch the stock price from an external API
    price, err := getStockPrice(stockSymbol)
    if err != nil {
        log.Printf("Failed to fetch stock price: %v", err)
        return
    }

    // Check if the stock price is above the alert price and send alert if necessary
    if price >= alertPrice {
        sendAlert(stockSymbol, price, alertPrice)
    }
}

func getStockPrice(symbol string) (float64, error) {
    // Call an external API to fetch the stock price
    // You can use a package like net/http or a third-party library like go-resty to do this
}

func sendAlert(symbol string, price float64, alertPrice float64) {
    // Send an alert notification to the user
    // You can use email, SMS, or push notifications for this purpose
}


func main() {
    // Set up Kafka producer
    brokers := []string{"localhost:9092"}
    producer, err := createKafkaProducer(brokers)
    if err != nil {
        log.Fatalf("Failed to create Kafka producer: %v", err)
    }

    // Set up Kafka consumer
    consumer, err := createKafkaConsumer(brokers, "stock-price-topic")
    if err != nil {
        log.Fatalf("Failed to create Kafka consumer: %v", err)
    }

    // Run the program indefinitely
    for {
        select {
        case msg := <-consumer.Messages():
            // Parse the message and get the stock price
            stockPrice, err := parseStockPriceMessage(msg.Value)
            if err != nil {
                log.Printf("Failed to parse stock price message: %v", err)
                continue
            }

            // Check if the stock price is above

// example2

broker := []string{"localhost:9092"}
    topic := "test"

    // Create a new Kafka producer
    config := sarama.NewConfig()
    config.Producer.RequiredAcks = sarama.WaitForAll
    config.Producer.Retry.Max = 5
    producer, err := sarama.NewSyncProducer(broker, config)
    if err != nil {
        log.Fatalln("Failed to start Sarama producer:", err)
    }
    defer func() {
        if err := producer.Close(); err != nil {
            log.Fatalln("Failed to close Sarama producer:", err)
        }
    }()

    // Produce a message to the Kafka topic
    message := &sarama.ProducerMessage{
        Topic: topic,
        Value: sarama.StringEncoder("Hello, Kafka!"),
    }
    partition, offset, err := producer.SendMessage(message)
    if err != nil {
        log.Fatalln("Failed to send message to Kafka:", err)
    }
    fmt.Printf("Produced message to partition %d, offset %d\n", partition, offset)
}

// example3

	broker := "localhost:9092"
	topic := "test"

	// Create a new Kafka producer
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
	})
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	// Produce a message to the Kafka topic
	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte("Hello, Kafka!"),
	}
	err = producer.Produce(message, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Produced message to topic %s\n", topic)
}
*/
