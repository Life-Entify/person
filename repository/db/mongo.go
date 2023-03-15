package person

import (
	"context"
	"encoding/json"
	"reflect"
	"strings"

	common "github.com/life-entify/person/common"
	config "github.com/life-entify/person/config"
	errors "github.com/life-entify/person/errors"
	"github.com/life-entify/person/v1"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	COLLECTION = "persons"
)

type MongoDB struct {
	uri      string
	database string
}

func NewMongoDB(config config.IConfig) *MongoDB {
	return &MongoDB{
		uri:      config.GetDBUrl(),
		database: config.GetDBName(),
	}
}
func (db *MongoDB) Connect() (*mongo.Client, *mongo.Collection) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.uri))
	if err != nil {
		panic(err)
	}
	collection := client.Database(db.database).Collection(COLLECTION, &options.CollectionOptions{})
	return client, collection
}
func MongoDisconnect(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
func (db *MongoDB) GetNextPersonId(ctx context.Context) (int64, error) {
	filter := bson.D{}
	client, collection := db.Connect()
	defer MongoDisconnect(client)
	opts := options.Find().SetSort(bson.D{{Key: "person_id", Value: -1}}).SetLimit(1)
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return 0, err
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	if len(results) > 0 {
		var resultPerson person.Person
		common.JSONToStruct(results[0], &resultPerson)
		return resultPerson.PersonId + 1, nil
	}
	return 1, nil
}
func (db *MongoDB) FindPersons(ctx context.Context, filterObj *person.Person, page *Pagination) ([]*person.Person, error) {
	client, collection := db.Connect()
	defer MongoDisconnect(client)
	var filter = bson.M{}
	if filterObj != nil {
		filterItems := []bson.M{}
		if filterObj.XId != "" {
			id, _ := primitive.ObjectIDFromHex(filterObj.XId)
			idFilter := bson.M{"_id": id}
			filterItems = append(filterItems, idFilter)
		}
		if !reflect.ValueOf(filterObj.Profile).IsZero() {

			if filterObj.Profile.LastName != "" {
				lastNameFilter := bson.M{"profile.last_name": bson.M{"$regex": primitive.Regex{Pattern: filterObj.Profile.LastName, Options: "i"}}}
				filterItems = append(filterItems, lastNameFilter)
			}
			if filterObj.Profile.FirstName != "" {
				firstNameFilter := bson.M{"profile.first_name": bson.M{"$regex": primitive.Regex{Pattern: filterObj.Profile.FirstName, Options: "i"}}}
				filterItems = append(filterItems, firstNameFilter)
			}
			if filterObj.Profile.MiddleName != "" {
				middleNameFilter := bson.M{"profile.middle_name": bson.M{"$regex": primitive.Regex{Pattern: filterObj.Profile.MiddleName, Options: "i"}}}
				filterItems = append(filterItems, middleNameFilter)
			}
			if filterObj.Profile.PhoneNumber != "" {
				//TODO:// show error when it has (+)
				phoneFilter := bson.M{"profile.phone_number": bson.M{"$regex": primitive.Regex{Pattern: filterObj.Profile.PhoneNumber, Options: "i"}}}
				filterItems = append(filterItems, phoneFilter)
			}
			if filterObj.Profile.NationalIdentity != "" {
				ninFilter := bson.M{"profile.national_identity": bson.M{"$regex": primitive.Regex{Pattern: filterObj.Profile.NationalIdentity, Options: "i"}}}
				filterItems = append(filterItems, ninFilter)
			}
		}
		if len(filterItems) > 0 {
			filter = bson.M{"$or": filterItems}
		}
	}
	option := options.Find().SetSort(bson.D{{Key: "_id", Value: -1}})
	if !reflect.ValueOf(page).IsZero() {
		if page.Skip != 0 {
			option.SetSkip(page.Skip)
		}
		if page.Limit != 0 {
			option.SetLimit(page.Limit)
		}
	}
	cursor, err := collection.Find(ctx, filter, option)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	defer cursor.Close(ctx)
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, errors.Errorf(err.Error())
	}
	var resultPersons []*person.Person
	for _, pt := range results {
		var resultPerson person.Person
		err = common.JSONToStruct(pt, &resultPerson)
		if err != nil {
			return nil, errors.Errorf(err.Error())
		}
		resultPersons = append(resultPersons, &resultPerson)
	}
	return resultPersons, nil
}
func (db *MongoDB) FindPersonByPhone(ctx context.Context, phone string) ([]byte, error) {
	client, collection := db.Connect()
	defer MongoDisconnect(client)
	var result bson.M
	err := collection.FindOne(ctx, bson.D{{Key: "profile.phone_number", Value: phone}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	jsonData, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
func (db *MongoDB) FindPersonById(ctx context.Context, id primitive.ObjectID) (*person.Person, error) {
	client, collection := db.Connect()
	defer MongoDisconnect(client)
	var (
		result    bson.M
		newPerson person.Person
	)
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	err = common.JSONToStruct(result, &newPerson)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	return &newPerson, nil
}
func (db *MongoDB) FindPersonsByPersonID(ctx context.Context, ids []int64) ([]*person.Person, error) {
	client, collection := db.Connect()
	defer MongoDisconnect(client)
	cursor, err := collection.Find(ctx, bson.D{{Key: "person_id", Value: bson.D{{Key: "$in", Value: ids}}}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, errors.Errorf(err.Error())
	}
	var resultPersons []*person.Person
	for _, pt := range results {
		var resultPerson person.Person
		common.JSONToStruct(pt, &resultPerson)
		resultPersons = append(resultPersons, &resultPerson)
	}
	return resultPersons, nil
}
func (db *MongoDB) FindPersonsByID(ctx context.Context, ids []primitive.ObjectID) ([]*person.Person, error) {
	client, collection := db.Connect()
	defer MongoDisconnect(client)
	cursor, err := collection.Find(ctx, bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: ids}}}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, errors.Errorf(err.Error())
	}
	var resultPersons []*person.Person
	for _, pt := range results {
		var resultPerson person.Person
		common.JSONToStruct(pt, &resultPerson)
		resultPersons = append(resultPersons, &resultPerson)
	}
	return resultPersons, nil
}
func (db *MongoDB) FindPersonProfile(ctx context.Context, _id string) ([]byte, error) {
	client, coll := db.Connect()
	defer MongoDisconnect(client)
	objectId, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return nil, err
	}
	var result bson.M
	opts := options.FindOne().SetProjection(bson.D{{Key: "profile", Value: 1}, {Key: "_id", Value: 1}, {Key: "person_id", Value: 1}})
	err = coll.FindOne(ctx, bson.D{{Key: "_id", Value: objectId}}, opts).Decode(result)
	if err != nil {
		return nil, err
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
func (db *MongoDB) CreatePerson(ctx context.Context, profile *person.Profile, nextOfKins []*person.NextOfKin, checkExistence bool) (*person.Person, error) {
	client, coll := db.Connect()
	defer MongoDisconnect(client)
	personId, err := db.GetNextPersonId(ctx)
	if err != nil {
		return nil, errors.Errorf(err.Error())
	}
	if checkExistence {
		sPersonByte, err := db.FindPersonByPhone(context.TODO(), profile.PhoneNumber)
		if err != nil {
			return nil, errors.Errorf(err.Error())
		}
		if len(sPersonByte) > 0 {
			return nil, errors.Errorf("user with this phone number already exists")
		}
	}

	person := person.Person{
		Profile:  profile,
		PersonId: personId,
	}
	if nextOfKins != nil {
		person.NextOfKins = nextOfKins
	}
	personByte, err := json.Marshal(&person)
	if err != nil {
		return nil, errors.Errorf(err)
	}
	var jsonPerson interface{}
	err = json.Unmarshal(personByte, &jsonPerson)
	if err != nil {
		return nil, errors.Errorf(err)
	}
	value, err := coll.InsertOne(ctx, jsonPerson)
	if err != nil {
		return nil, errors.Errorf(err)
	}
	if oid, ok := value.InsertedID.(primitive.ObjectID); ok {
		person.XId = oid.Hex()
	}
	return &person, nil
}
func (db *MongoDB) UpdatePerson(ctx context.Context, _id primitive.ObjectID, profile *person.Profile) (*mongo.UpdateResult, error) {
	client, coll := db.Connect()
	defer MongoDisconnect(client)
	filter := bson.D{primitive.E{Key: "_id", Value: _id}}
	upsert := true
	opts := options.UpdateOptions{
		Upsert: &upsert,
	}
	arrayFiler := options.ArrayFilters{}
	errors.Errorf(profile)
	if !reflect.ValueOf(profile.Addresses).IsZero() && len(profile.Addresses) > 0 {
		arrayFiler.Filters = bson.A{bson.M{"x._id": profile.Addresses[0].XId}}
		opts.ArrayFilters = &arrayFiler
	}

	var setUpdate = bson.M{}
	var onDateInsert = bson.M{}

	values := reflect.ValueOf(profile).Elem()
	fields := values.Type()
	for i := 0; i < values.NumField(); i++ {
		f := values.Field(i)
		if fields.Field(i).Name == "Addresses" {
			addArray := f
			for j := 0; j < addArray.Len(); j++ {
				addUnit := addArray.Index(j).Elem()
				addUnitFields := addUnit.Type()
				for k := 0; k < addUnit.NumField(); k++ {
					addValue := addUnit.Field(k)
					if addValue.CanInterface() {
						addField := addUnitFields.Field(k)
						addFieldValue := addValue.Interface()
						tag := strings.Split(addField.Tag.Get("json"), ",")[0]
						// check if the value is empty
						if reflect.Zero(addValue.Type()).Interface() != addFieldValue {
							setUpdate["profile.addresses.$[x]."+tag] = addFieldValue
							onDateInsert[tag] = addFieldValue
						}
					}
				}
			}
		} else {
			if f.CanInterface() {
				if reflect.Zero(f.Type()).Interface() != f.Interface() {
					tag := strings.Split(fields.Field(i).Tag.Get("json"), ",")[0]
					setUpdate["profile."+tag] = f.Interface()
				}
			}
		}
	}
	updateData := bson.M{
		"$set": setUpdate,
	}
	value, err := coll.UpdateOne(ctx, filter, updateData, &opts)
	if err != nil {
		if strings.Contains(err.Error(), "The path 'profile.addresses' must exist") {
			_, err1 := coll.UpdateOne(ctx, filter,
				bson.M{"$set": bson.M{"profile.addresses": []interface{}{onDateInsert}}},
			)
			if err1 != nil {
				return nil, errors.Errorf("%s => caused by $s", err1.Error(), err.Error())
			}
			_, err = coll.UpdateOne(ctx, filter, updateData, &opts)
			if err != nil {
				return nil, errors.Errorf("%s => caused by $s", err.Error(), err.Error())
			}
		}
		return nil, errors.Errorf(err.Error())
	}
	return value, nil
}
func (db *MongoDB) DeletePersons(ctx context.Context, _ids []primitive.ObjectID) (*mongo.DeleteResult, error) {
	client, coll := db.Connect()
	defer MongoDisconnect(client)
	filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: _ids}}}}
	result, err := coll.DeleteMany(ctx, filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}
