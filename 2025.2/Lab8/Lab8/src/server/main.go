package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

type Server struct {
	hashMap    map[int][]string
	clientData map[string][]int
}

func NewServer() *Server {
	return &Server{
		hashMap:    make(map[int][]string),
		clientData: make(map[string][]int),
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer func() {
		clientIP := conn.RemoteAddr().String()
		s.cleanupClientData(clientIP)
		conn.Close()
	}()

	for {
		var requestType string
		decoder := gob.NewDecoder(conn)
		if err := decoder.Decode(&requestType); err != nil {
			if err.Error() == "EOF" {
				log.Println("Client disconnected")
				return
			}
			log.Println("Error decoding request type:", err)
			return
		}

		if requestType == "store" {
			s.handleStoreRequest(conn, decoder)
		} else if requestType == "create" {
			s.handleCreateRequest(conn, decoder)
		} else if requestType == "delete" {
			s.handleDeleteRequest(conn, decoder)
		} else if requestType == "query" {
			s.handleQueryRequest(conn, decoder)
		} else {
			log.Println("Unknown request type:", requestType)
		}
	}
}

func (s *Server) handleStoreRequest(conn net.Conn, decoder *gob.Decoder) {
	var clientHashes []int
	if err := decoder.Decode(&clientHashes); err != nil {
		log.Println("Error decoding data:", err)
		return
	}

	clientIP := conn.RemoteAddr().String()
	for _, hash := range clientHashes {
		if _, exists := s.hashMap[hash]; !exists {
			s.hashMap[hash] = []string{}
		}
		s.hashMap[hash] = append(s.hashMap[hash], clientIP)
		s.clientData[clientIP] = append(s.clientData[clientIP], hash)
	}

	printHashMap(s.hashMap)
}

func (s *Server) handleCreateRequest(conn net.Conn, decoder *gob.Decoder) {
	var fileHash int
	if err := decoder.Decode(&fileHash); err != nil {
		log.Println("Error decoding file hash:", err)
		return
	}

	clientIP := conn.RemoteAddr().String()

	if _, exists := s.hashMap[fileHash]; !exists {
		s.hashMap[fileHash] = []string{}
	}
	s.hashMap[fileHash] = append(s.hashMap[fileHash], clientIP)
	s.clientData[clientIP] = append(s.clientData[clientIP], fileHash)

	fmt.Printf("File created by %s: Hash %d\n", clientIP, fileHash)
}

func (s *Server) handleDeleteRequest(conn net.Conn, decoder *gob.Decoder) {
	var fileHash int
	if err := decoder.Decode(&fileHash); err != nil {
		log.Println("Error decoding file hash:", err)
		return
	}

	clientIP := conn.RemoteAddr().String()

	if ips, exists := s.hashMap[fileHash]; exists {
		for i, ip := range ips {
			if ip == clientIP {
				s.hashMap[fileHash] = append(ips[:i], ips[i+1:]...)
				break
			}
		}
		if len(s.hashMap[fileHash]) == 0 {
			delete(s.hashMap, fileHash)
		}
	}

	s.clientData[clientIP] = removeFromSlice(s.clientData[clientIP], fileHash)

	fmt.Printf("File deleted by %s: Hash %d\n", clientIP, fileHash)
}

func (s *Server) handleQueryRequest(conn net.Conn, decoder *gob.Decoder) {
	var hash int
	if err := decoder.Decode(&hash); err != nil {
		log.Println("Error decoding hash:", err)
		return
	}

	ips := s.hashMap[hash]

	encoder := gob.NewEncoder(conn)
	encoder.Encode(ips)
}

func (s *Server) cleanupClientData(clientIP string) {
	hashes, exists := s.clientData[clientIP]
	if !exists {
		return
	}

	for _, hash := range hashes {
		ips := s.hashMap[hash]
		for i, ip := range ips {
			if ip == clientIP {
				s.hashMap[hash] = append(ips[:i], ips[i+1:]...)
				break
			}
		}

		if len(s.hashMap[hash]) == 0 {
			delete(s.hashMap, hash)
		}
	}

	delete(s.clientData, clientIP)

	log.Printf("Cleaned up data for client: %s\n", clientIP)
}

func printHashMap(hashMap map[int][]string) {
	fmt.Println("Hash Map:")
	for hash, ips := range hashMap {
		fmt.Printf("Hash: %d\n", hash)
		fmt.Println("  IPs:")
		for _, ip := range ips {
			fmt.Printf("    %s\n", ip)
		}
	}
}

func removeFromSlice(slice []int, val int) []int {
	for i, v := range slice {
		if v == val {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func main() {
	server := NewServer()
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	fmt.Println("Server is listening on port 8080...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		go server.handleConnection(conn)
	}
}
