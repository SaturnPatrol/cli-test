/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/spf13/cobra"
)

// kafkaCmd represents the kafka command
var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		NewPsi().write()
	},
}

func init() {
	rootCmd.AddCommand(kafkaCmd)
}

type PromStand struct {
	servers []string
	tls     int
}

func NewProm() PromStand {
	return PromStand{servers: []string{"localhost:9092"}, tls: 2}
}

type PsiStand struct {
	servers []string
	tls     int
}

func NewPsi() *PsiStand {
	return &PsiStand{servers: []string{"192.168.1.43:9092"}, tls: 1}
}

func (s *PromStand) write() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	p, _ := sarama.NewSyncProducer(s.servers, config)
	p.SendMessage(&sarama.ProducerMessage{
		Topic: "firstTopic",
		Value: sarama.StringEncoder("foo"),
	})
}

func (s *PsiStand) write() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	p, err := sarama.NewSyncProducer(s.servers, config)
	if err != nil {
		println(err)
	}
	part, off, _ := p.SendMessage(&sarama.ProducerMessage{
		Topic: "firstTopic",
		Value: sarama.StringEncoder("foqweqweqweo"),
		// Partition: 2,
	})

	fmt.Printf("Message produced to partition %v, offset %v\n", part, off)
}

type Writer interface {
	write()
}
