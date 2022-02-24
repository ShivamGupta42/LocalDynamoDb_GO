package DynamoGO

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"time"
)

// Use struct tags much like the standard JSON library,
// you can embed anonymous structs too!
type widget struct {
	UserID int       // Hash key, a.k.a. partition key
	Time   time.Time // Range key, a.k.a. sort key

	Msg       string              `dynamo:"Message"`    // Change name in the database
	Count     int                 `dynamo:",omitempty"` // Omits if zero value
	Children  []widget            // Lists
	Friends   []string            `dynamo:",set"` // Sets
	Set       map[string]struct{} `dynamo:",set"` // Map sets, too!
	SecretKey string              `dynamo:"-"`    // Ignored
}

func main() {
	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &aws.Config{Region: aws.String("us-west-2")})
	table := db.Table("Widgets")

	// put item
	w := widget{UserID: 613, Time: time.Now(), Msg: "hello"}
	table.Put(w).Run()

	// get the same item
	var result widget
	table.Get("UserID", w.UserID).
		Range("Time", dynamo.Equal, w.Time).
		One(&result)

	// get all items
	var results []widget
	table.Scan().All(&results)

	// use placeholders in filter expressions (see Expressions section below)
	var filtered []widget
	table.Scan().Filter("'Count' > ?", 10).All(&filtered)
}
