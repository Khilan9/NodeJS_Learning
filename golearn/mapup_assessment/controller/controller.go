package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"mongoserver/model"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.5.6"

// const connectionString = "mongodb+srv://learncodeonline:hitesh@cluster0.humov.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const dbName = "mapupdb"
const colName = "locations"

// MOST IMPORTANT
var collection *mongo.Collection

// connect with monogoDB

var Collection *mongo.Collection

func init() {

	//client option
	fmt.Print("Trying to connect mongo db")
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is ready")
}

func StoreLocation(w http.ResponseWriter, r *http.Request) {
	start := time.Now() // start timer

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var location model.Location
	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	inserted, err := collection.InsertOne(context.Background(), location)
	if err != nil {
		http.Error(w, "Failed to insert location", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	// Convert insertedID to hex string
	id := ""
	if oid, ok := inserted.InsertedID.(primitive.ObjectID); ok {
		id = oid.Hex()
	}

	// Calculate elapsed time in nanoseconds
	elapsed := time.Since(start).Nanoseconds()

	// Prepare response
	response := map[string]interface{}{
		"id":      id,
		"time_ns": elapsed,
	}

	// Send JSON response
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("Failed to write JSON response:", err)
	}
}

func GetLocationByCategory(w http.ResponseWriter, r *http.Request) {
	start := time.Now() // Start timer

	w.Header().Set("Content-Type", "application/json")

	// Get category from URL
	params := mux.Vars(r)
	category := params["category"]

	// MongoDB filter
	filter := bson.M{"category": category}

	// Find documents
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		http.Error(w, "Error fetching locations", http.StatusInternalServerError)
		log.Println("Mongo find error:", err)
		return
	}
	defer cursor.Close(context.Background())

	// Prepare locations array
	var locations []map[string]interface{}
	for cursor.Next(context.Background()) {
		var loc model.Location
		if err := cursor.Decode(&loc); err != nil {
			log.Println("Decode error:", err)
			continue
		}

		locationMap := map[string]interface{}{
			"id":        loc.ID.Hex(),
			"name":      loc.Name,
			"address":   loc.Address,
			"latitude":  loc.Latitude,
			"longitude": loc.Longitude,
			"category":  loc.Category,
		}
		locations = append(locations, locationMap)
	}

	// Measure elapsed time
	elapsed := time.Since(start).Nanoseconds()

	// Final response including time_ns
	response := map[string]interface{}{
		"locations": locations,
		"time_ns":   elapsed,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("Failed to encode JSON:", err)
	}
}

type SearchRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Category  string  `json:"category"`
	RadiusKm  float64 `json:"radius_km"`
}

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371 // Earth radius in km
	dLat := (lat2 - lat1) * math.Pi / 180
	dLon := (lon2 - lon1) * math.Pi / 180
	lat1 = lat1 * math.Pi / 180
	lat2 = lat2 * math.Pi / 180

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(lat1)*math.Cos(lat2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}
func SearchLocations(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Content-Type", "application/json")

	var req SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Case-insensitive category match
	filter := bson.M{
		"category": bson.M{
			"$regex":   fmt.Sprintf("^%s$", req.Category),
			"$options": "i",
		},
	}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		http.Error(w, "MongoDB error", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	var results []map[string]interface{}
	for cursor.Next(context.Background()) {
		var loc model.Location
		if err := cursor.Decode(&loc); err != nil {
			continue
		}

		dist := haversine(req.Latitude, req.Longitude, loc.Latitude, loc.Longitude)
		if dist <= req.RadiusKm {
			results = append(results, map[string]interface{}{
				"id":       loc.ID.Hex(),
				"name":     loc.Name,
				"address":  loc.Address,
				"distance": dist,
				"category": loc.Category,
			})
		}
	}

	elapsed := time.Since(start).Nanoseconds()

	json.NewEncoder(w).Encode(map[string]interface{}{
		"locations": results,
		"time_ns":   elapsed,
	})
}

func FetchCost(w http.ResponseWriter, r *http.Request) {

}

func Homepage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by LearnCodeOnline</h1>"))
}
