func readData(rw *bufio.ReadWriter) {

	for {
		str, err := rw.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		if str == "" {
			return
		}
		if str != "\n" {

			chain := make([]Block, 0)
			if err := json.Unmarshal([]byte(str), &chain); err != nil {
				log.Fatal(err)
			}

			mutex.Lock()
			if len(chain) > len(Blockchain) {
				Blockchain = chain
				bytes, err := json.MarshalIndent(Blockchain, "", "  ")
				if err != nil {

					log.Fatal(err)
				}
				// Green console color: 	\x1b[32m
				// Reset console color: 	\x1b[0m
				fmt.Printf("\x1b[32m%s\x1b[0m> ", string(bytes))
			}
			mutex.Unlock()
		}
	}
}
